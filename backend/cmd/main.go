package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"

	"github.com/avelex/terrace-finance/backend/config"
	"github.com/avelex/terrace-finance/backend/db"
	"github.com/avelex/terrace-finance/backend/internal/api"
	cctp_client "github.com/avelex/terrace-finance/backend/internal/circle/cctp/client"
	cctp_processor "github.com/avelex/terrace-finance/backend/internal/circle/cctp/processor"
	gateway_client "github.com/avelex/terrace-finance/backend/internal/circle/gateway/client"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/repository"
	"github.com/avelex/terrace-finance/backend/internal/services/network"
	"github.com/avelex/terrace-finance/backend/internal/services/strategy"
	"github.com/avelex/terrace-finance/backend/internal/services/wallet"
	"github.com/avelex/terrace-finance/backend/internal/transactor"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	context.AfterFunc(ctx, func() {
		log.Info().Msg("shutting down")
	})

	if err := run(ctx); err != nil {
		log.Fatal().Err(err).Msg("run failed")
	}
}

func run(ctx context.Context) error {
	conf, err := config.Load("config.yaml")
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	pgPool, err := pgxpool.New(ctx, conf.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("connect to db")
	}

	defer pgPool.Close()

	sqldb := stdlib.OpenDBFromPool(pgPool)
	bunDB := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Migrate(sqldb); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	// ectx, err := eventscale.Connect(ctx, conf.EventscaleURL)
	// if err != nil {
	// 	return fmt.Errorf("connect to eventscale: %w", err)
	// }

	cctpClient := cctp_client.NewClient()
	gatewayClient := gateway_client.NewClient()

	repo := repository.NewBridgeRepository(bunDB)
	aaveRepo := repository.NewAaveRepository(bunDB)
	strategyRepo := repository.NewStrategyRepository(bunDB)
	userRepo := repository.NewUserRepository(bunDB)

	networkService, err := network.NewService(ctx, conf)
	if err != nil {
		return fmt.Errorf("create network service: %w", err)
	}

	domains, err := connectDomainTransactors(conf, networkService)
	if err != nil {
		return fmt.Errorf("connect domain transactors: %w", err)
	}

	cctpProcessor := cctp_processor.New(cctpClient, repo)
	transactor := transactor.NewTransactor(conf.Protocol.Hub, domains, repo, userRepo)

	walletService := wallet.NewService(conf.Protocol, gatewayClient, networkService, userRepo)
	strategyService := strategy.NewService(transactor, aaveRepo, strategyRepo, cctpClient)
	//eventHandler := event_handler.NewEventHandler(repo, strategyService)

	apiHandler := api.NewHandler(strategyService, walletService)

	echoRouter := echo.New()
	apiGroup := echoRouter.Group("/api")
	apiGroup.Use(middleware.CORS())

	apiHandler.Register(apiGroup)

	// terraceSub, err := eventscale.SubscribeEventBlock(ectx, eventHandler.Handle,
	// 	eventscale.WithEventBlockConsumerName("terrace-operator"),
	// )
	// if err != nil {
	// 	return fmt.Errorf("subscribe to eventscale: %w", err)
	// }

	wg := &sync.WaitGroup{}
	wg.Add(3)

	// go func() {
	// 	log.Info().Msg("subscribed to terrace contracts")

	// 	defer wg.Done()
	// 	terraceSub.Start(ctx)
	// }()

	go func() {
		log.Info().Msg("started CCTP processor")

		defer wg.Done()
		cctpProcessor.Start(ctx)
	}()

	go func() {
		log.Info().Msg("started transactor")

		defer wg.Done()
		transactor.Start(ctx)
	}()

	go func() {
		log.Info().Msg("started API server")

		defer wg.Done()
		echoRouter.Start(":8080")
	}()

	<-ctx.Done()

	if err := echoRouter.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("shutdown API server: %w", err)
	}

	wg.Wait()

	return nil
}

func connectDomainTransactors(cfg config.Config, svc *network.Service) (map[enum.CircleDomain]*transactor.DomainTransactor, error) {
	pk, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("load private key: %w", err)
	}

	transactors := make(map[enum.CircleDomain]*transactor.DomainTransactor)

	for _, network := range cfg.Networks {
		client, err := svc.Client(network.Name)
		if err != nil {
			return nil, fmt.Errorf("connect to %s: %w", network.Name, err)
		}

		gatewayWallet := enum.GATEWAY_WALLET_MAPPING[network.Name]
		gatewayMint := common.Address{}

		if network.Domain == uint32(cfg.Protocol.Hub) {
			gatewayMint = cfg.Protocol.Stablecoin
		}

		transactor, err := transactor.NewDomainTransactor(pk, client, network.Domain, network.Vault, gatewayWallet, gatewayMint)
		if err != nil {
			return nil, fmt.Errorf("create transactor for %s: %w", network.Name, err)
		}

		transactors[enum.CircleDomain(network.Domain)] = transactor
	}

	return transactors, nil

}
