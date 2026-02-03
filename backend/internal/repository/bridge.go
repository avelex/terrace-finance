package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/uptrace/bun"
)

type BridgeRepository struct {
	db *bun.DB
}

func NewBridgeRepository(db *bun.DB) *BridgeRepository {
	return &BridgeRepository{db: db}
}

func (r *BridgeRepository) CreateBridgeOp(ctx context.Context, send models.SendFunds) error {
	bridge := send.BridgeOp()

	_, err := r.db.NewInsert().Model(&bridge).On("CONFLICT (id) DO NOTHING").Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to create bridge op: %w", err)
	}

	return nil
}

func (r *BridgeRepository) CompleteBridgeOp(ctx context.Context, received models.ReceivedFunds) error {
	_, err := r.db.NewUpdate().Model(new(models.BridgeOp)).
		Set("to_chain_id = ?", received.ChainID.String()).
		Set("received_amount = ?", received.Amount.String()).
		Set("received_tx_hash = ?", received.TxHash.String()).
		Set("received_at = ?", time.Unix(received.Timestamp, 0)).
		Where("id = ?", received.ID.String()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to complete bridge op: %w", err)
	}

	return nil
}

func (r *BridgeRepository) GetPendingOps(ctx context.Context) ([]models.BridgeOp, error) {
	var ops []models.BridgeOp
	err := r.db.NewSelect().Model(&ops).
		Where("cctp_message IS NULL").
		Where("received_tx_hash IS NULL").
		Order("sent_at ASC").
		Limit(25).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get pending ops: %w", err)
	}

	return ops, nil
}

func (r *BridgeRepository) UpdateMessageAndAttestation(ctx context.Context, id, message, attestation string) error {
	_, err := r.db.NewUpdate().Model(new(models.BridgeOp)).
		Set("cctp_message = ?", message).
		Set("cctp_attestation = ?", attestation).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to update message and attestation: %w", err)
	}

	return nil
}
