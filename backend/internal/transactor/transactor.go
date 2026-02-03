package transactor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	GetPendingOps(ctx context.Context) ([]models.BridgeOp, error)
	UpdateReceivedTxHash(ctx context.Context, id string, txHash common.Hash) error
}

type Transactor struct {
	domains map[uint32]*DomainTransactor
	repo    Repository
}

func NewTransactor(domains map[uint32]*DomainTransactor, repo Repository) *Transactor {
	return &Transactor{
		domains: domains,
		repo:    repo,
	}
}

func (t *Transactor) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := t.process(ctx); err != nil {
				log.Error().Err(err).Msg("process tx failed")
			}
		}
	}
}

func (t *Transactor) process(ctx context.Context) error {
	ops, err := t.repo.GetPendingOps(ctx)
	if err != nil {
		return fmt.Errorf("get pending ops: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(ops))

	for _, op := range ops {
		msg, err := hexutil.Decode(op.CCTPMessage)
		if err != nil {
			continue
		}

		att, err := hexutil.Decode(op.CCTPAttestation)
		if err != nil {
			continue
		}

		go func(domain uint32, message []byte, attestation []byte) {
			defer wg.Done()

			receipt, err := t.domains[domain].ReceiveMessage(message, attestation)
			if err != nil {
				log.Error().Err(err).Msg("receive message failed")
				return
			}

			if err := t.repo.UpdateReceivedTxHash(ctx, op.ID, receipt.TxHash); err != nil {
				log.Error().Err(err).Msg("update received tx hash failed")
				return
			}
		}(op.ToDomain, msg, att)
	}

	wg.Wait()

	return nil
}
