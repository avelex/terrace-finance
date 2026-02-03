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
	cctp_client "github.com/avelex/terrace-finance/backend/internal/cctp/client"
	"github.com/avelex/terrace-finance/backend/internal/cctp/processor"
	"github.com/avelex/terrace-finance/backend/internal/event_handler"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/repository"
	"github.com/avelex/terrace-finance/backend/internal/strategy"
	"github.com/avelex/terrace-finance/backend/internal/transactor"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	eventscale "github.com/eventscale/eventscale/pkg/sdk-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/labstack/echo/v4"
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

	domains, err := connectDomainTransactors(ctx, conf)
	if err != nil {
		return fmt.Errorf("connect domain transactors: %w", err)
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

	ectx, err := eventscale.Connect(ctx, conf.EventscaleURL)
	if err != nil {
		return fmt.Errorf("connect to eventscale: %w", err)
	}

	cctpClient := cctp_client.NewClient()

	repo := repository.NewBridgeRepository(bunDB)
	aaveRepo := repository.NewAaveRepository(bunDB)

	manager := strategy.NewManager(aaveRepo)

	eventHandler := event_handler.NewEventHandler(repo, manager)
	processor := processor.New(cctpClient, repo)
	transactor := transactor.NewTransactor(enum.ARC_DOMAIN, domains, repo, cctpClient)

	apiHandler := api.NewHandler(transactor)

	echoRouter := echo.New()
	apiGroup := echoRouter.Group("/api")

	apiHandler.Register(apiGroup)

	terraceSub, err := eventscale.SubscribeEventBlock(ectx, eventHandler.Handle,
		eventscale.WithEventBlockConsumerName("terrace-operator"),
	)
	if err != nil {
		return fmt.Errorf("subscribe to eventscale: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		log.Info().Msg("subscribed to terrace contracts")

		defer wg.Done()
		terraceSub.Start(ctx)
	}()

	go func() {
		log.Info().Msg("started CCTP processor")

		defer wg.Done()
		processor.Start(ctx)
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
	wg.Wait()

	return nil
}

func connectDomainTransactors(ctx context.Context, cfg config.Config) (map[uint32]*transactor.DomainTransactor, error) {
	pk, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("load private key: %w", err)
	}

	transactors := make(map[uint32]*transactor.DomainTransactor)

	for _, network := range cfg.Networks {
		client, err := ethclient.DialContext(ctx, network.URL)
		if err != nil {
			return nil, fmt.Errorf("connect to %s: %w", network.Name, err)
		}

		transactor, err := transactor.NewDomainTransactor(pk, client, network.Domain, network.Contract)
		if err != nil {
			return nil, fmt.Errorf("create transactor for %s: %w", network.Name, err)
		}

		transactors[network.Domain] = transactor
	}

	return transactors, nil

}
