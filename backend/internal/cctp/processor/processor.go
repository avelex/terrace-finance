package processor

import (
	"context"
	"fmt"
	"strconv"
	"time"

	cctp_client "github.com/avelex/terrace-finance/backend/internal/cctp/client"
	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/rs/zerolog/log"
)

type Repository interface {
	GetSendingOps(ctx context.Context) ([]models.BridgeOp, error)
	UpdateMessageAndAttestation(ctx context.Context, id, message, attestation string) error
}

type Processor struct {
	client *cctp_client.Client
	repo   Repository
}

func New(client *cctp_client.Client, repo Repository) *Processor {
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
	ops, err := p.repo.GetSendingOps(ctx)
	if err != nil {
		return fmt.Errorf("get sending ops: %w", err)
	}

	for _, v := range ops {
		srcDomain := strconv.FormatUint(uint64(v.FromDomain), 10)

		msg, att, err := p.client.MessageAndAttestation(ctx, srcDomain, v.SentTxHash)
		if err != nil {
			log.Info().Str("event_id", v.ID).Msg("message and attestation not ready")
			continue
		}

		if err := p.repo.UpdateMessageAndAttestation(ctx, v.ID, msg, att); err != nil {
			log.Error().Err(err).Msg("update message and attestation")
			continue
		}
	}

	return nil
}
