package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InitiateUserDepositFlow(ctx context.Context, deposit models.UserDeposit, permits []models.UserUnifiedPermits) error {
	err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctxTx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(&deposit).Exec(ctxTx)
		if err != nil {
			return fmt.Errorf("unable to create user deposit: %w", err)
		}

		_, err = tx.NewInsert().Model(&permits).Exec(ctxTx)
		if err != nil {
			return fmt.Errorf("unable to create user permits: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to initiate user deposit flow: %w", err)
	}

	return nil
}

func (r *UserRepository) UpdateUserUnifiedPermitsSignatures(ctx context.Context, permits []models.UserUnifiedPermits) error {
	err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctxTx context.Context, tx bun.Tx) error {
		for _, permit := range permits {
			_, err := tx.NewUpdate().
				Model(&models.UserUnifiedPermits{}).
				Set("signature = ?", permit.Signature).
				Where("deposit_id = ?", permit.DepositID).
				Where("owner = ?", permit.Owner).
				Where("domain = ?", permit.Domain).
				Exec(ctxTx)
			if err != nil {
				return fmt.Errorf("unable to update user permit: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to update user unified permits signatures: %w", err)
	}

	return nil
}

func (r *UserRepository) SaveDepositAttestationAndSignature(ctx context.Context, id uuid.UUID, attestation, signature string) error {
	_, err := r.db.NewUpdate().
		Model(&models.UserDeposit{}).
		Set("attestation = ?", attestation).
		Set("signature = ?", signature).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to save deposit attestation and signature: %w", err)
	}

	return nil
}

func (r *UserRepository) GetPendingPermitDeposits(ctx context.Context) ([]models.UserUnifiedPermits, error) {
	var permits []models.UserUnifiedPermits
	err := r.db.NewSelect().
		Model(&permits).
		Where("signature IS NOT NULL").
		Where("deposited_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get pending permit deposits: %w", err)
	}

	return permits, nil
}

func (r *UserRepository) UpdatePermitDepositTxHash(ctx context.Context, id string, owner common.Address, domain uint32, txHash common.Hash) error {
	log.Info().Msgf("update permit deposit tx hash: %s, %s, %d, %s", id, owner, domain, txHash)

	res, err := r.db.NewUpdate().
		Model((*models.UserUnifiedPermits)(nil)).
		Set("tx_hash = ?", txHash.String()).
		Set("deposited_at = ?", time.Now()).
		Where("deposit_id = ?", id).
		Where("owner = ?", owner.String()).
		Where("domain = ?", domain).Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to update permit deposit tx hash: %w", err)
	}

	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		log.Warn().Msg("no rows affected")
	}

	return nil
}

func (r *UserRepository) GetPendingHubDeposits(ctx context.Context) ([]models.UserDeposit, error) {
	var deposits []models.UserDeposit
	err := r.db.NewSelect().
		Model(&deposits).
		Where("attestation IS NOT NULL").
		Where("signature IS NOT NULL").
		Where("deposited_at IS NULL").
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get pending hub deposits: %w", err)
	}

	return deposits, nil
}

func (r *UserRepository) UpdateDepositTxHash(ctx context.Context, id string, txHash common.Hash) error {
	_, err := r.db.NewUpdate().
		Model(&models.UserDeposit{}).
		Where("id = ?", id).
		Set("deposit_tx_hash = ?", txHash.String()).
		Set("deposited_at = ?", time.Now()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to update deposit tx hash: %w", err)
	}

	return nil
}
