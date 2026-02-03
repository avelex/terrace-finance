package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"

	"github.com/avelex/terrace-finance/backend/config"
	"github.com/avelex/terrace-finance/backend/db"
	"github.com/avelex/terrace-finance/backend/internal/cctp/processor"
	"github.com/avelex/terrace-finance/backend/internal/event_handler"
	"github.com/avelex/terrace-finance/backend/internal/extclients/cctp"
	"github.com/avelex/terrace-finance/backend/internal/repository"

	eventscale "github.com/eventscale/eventscale/pkg/sdk-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
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
	conf, err := config.Load()
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

	ectx, err := eventscale.Connect(ctx, conf.EventscaleURL)
	if err != nil {
		return fmt.Errorf("connect to eventscale: %w", err)
	}

	bridgeRepo := repository.NewBridgeRepository(bunDB)
	handler := event_handler.NewHandler(bridgeRepo)
	processor := processor.New(cctp.NewClient(), bridgeRepo)

	sub, err := eventscale.SubscribeEventBlock(ectx, handler.Handle,
		eventscale.WithEventBlockConsumerName("terrace-operator"),
	)
	if err != nil {
		return fmt.Errorf("subscribe to eventscale: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		log.Info().Msg("subscribed to eventscale")

		defer wg.Done()
		sub.Start(ctx)
	}()

	go func() {
		log.Info().Msg("started CCTP processor")

		defer wg.Done()
		processor.Start(ctx)
	}()

	<-ctx.Done()
	wg.Wait()

	return nil
}
