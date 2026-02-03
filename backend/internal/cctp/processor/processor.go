package processor

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/extclients/cctp"
	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/rs/zerolog/log"
)

type Repository interface {
	GetPendingOps(ctx context.Context) ([]models.BridgeOp, error)
	UpdateMessageAndAttestation(ctx context.Context, id, message, attestation string) error
}

type Processor struct {
	client *cctp.Client
	repo   Repository
}

func New(client *cctp.Client, repo Repository) *Processor {
	return &Processor{
		client: client,
		repo:   repo,
	}
}

func (p *Processor) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := p.process(ctx); err != nil {
				log.Error().Err(err).Msg("process failed")
			}
		}
	}
}

func (p *Processor) process(ctx context.Context) error {
	ops, err := p.repo.GetPendingOps(ctx)
	if err != nil {
		return fmt.Errorf("get pending ops: %w", err)
	}

	for _, v := range ops {
		srcDomain := strconv.FormatUint(uint64(v.FromDomain), 10)

		msg, att, err := p.client.MessageAndAttestation(ctx, srcDomain, v.SentTxHash)
		if err != nil {
			log.Error().Err(err).Msg("message and attestation")
			continue
		}

		if err := p.repo.UpdateMessageAndAttestation(ctx, v.ID, msg, att); err != nil {
			log.Error().Err(err).Msg("update message and attestation")
			continue
		}
	}

	return nil
}
