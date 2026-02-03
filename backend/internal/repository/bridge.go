package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/ethereum/go-ethereum/common"
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
		Set("received_amount = ?", received.Amount.String()).
		Set("received_tx_hash = ?", received.TxHash.String()).
		Set("received_at = ?", time.Unix(received.Timestamp, 0)).
		Where("id = ?", common.BytesToHash(received.ID[:]).String()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to complete bridge op: %w", err)
	}

	return nil
}

func (r *BridgeRepository) UpdateReceivedTxHash(ctx context.Context, id string, txHash common.Hash) error {
	_, err := r.db.NewUpdate().Model(new(models.BridgeOp)).
		Set("received_tx_hash = ?", txHash.String()).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to update received tx hash: %w", err)
	}

	return nil
}

func (r *BridgeRepository) GetSendingOps(ctx context.Context) ([]models.BridgeOp, error) {
	var ops []models.BridgeOp
	err := r.db.NewSelect().Model(&ops).
		Where("cctp_message IS NULL").
		Where("received_tx_hash IS NULL").
		Order("sent_at ASC").
		Limit(10).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get sending ops: %w", err)
	}

	return ops, nil
}

func (r *BridgeRepository) GetPendingOps(ctx context.Context) ([]models.BridgeOp, error) {
	var ops []models.BridgeOp
	err := r.db.NewSelect().Model(&ops).
		Where("cctp_message IS NOT NULL").
		Where("received_tx_hash IS NULL").
		Order("sent_at ASC").
		Limit(10).
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
