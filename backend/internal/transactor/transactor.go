package transactor

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/avelex/terrace-finance/backend/internal/models"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
)

type BridgeRepository interface {
	GetPendingOps(ctx context.Context) ([]models.BridgeOp, error)
	UpdateReceivedTxHash(ctx context.Context, id string, txHash common.Hash) error
}

type UserRepository interface {
	GetPendingPermitDeposits(ctx context.Context) ([]models.UserUnifiedPermits, error)
	UpdatePermitDepositTxHash(ctx context.Context, id string, owner common.Address, domain uint32, txHash common.Hash) error
	GetPendingHubDeposits(ctx context.Context) ([]models.UserDeposit, error)
	UpdateDepositTxHash(ctx context.Context, id string, txHash common.Hash) error
}

type Transactor struct {
	hubDomain enum.CircleDomain
	domains   map[enum.CircleDomain]*DomainTransactor
	brideRepo BridgeRepository
	userRepo  UserRepository
}

func NewTransactor(hubDomain enum.CircleDomain, domains map[enum.CircleDomain]*DomainTransactor, br BridgeRepository, ur UserRepository) *Transactor {
	return &Transactor{
		hubDomain: hubDomain,
		domains:   domains,
		brideRepo: br,
		userRepo:  ur,
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

func (t *Transactor) TerraceBalance(ctx context.Context, domain enum.CircleDomain, token common.Address) (*big.Int, error) {
	dt, exists := t.domains[domain]
	if !exists {
		return nil, fmt.Errorf("domain %d not found", domain)
	}

	return dt.TerraceBalanceOf(ctx, token)
}

func (t *Transactor) TerraceAddress(domain enum.CircleDomain) common.Address {
	return t.domains[domain].vaultAddress
}

func (t *Transactor) SendAllFunds(ctx context.Context, srcDomain, dstDomain enum.CircleDomain, maxFee *big.Int) (string, error) {
	dt, exists := t.domains[srcDomain]
	if !exists {
		return "", fmt.Errorf("domain %d not found", srcDomain)
	}

	if srcDomain == t.hubDomain {
		receipt, err := dt.SendAllFundsToTerrace(dstDomain, maxFee, enum.FAST_TRANSFER)
		if err != nil {
			return "", fmt.Errorf("send all funds to terrace: %w", err)
		}

		return receipt.TxHash.String(), nil
	} else if dstDomain == t.hubDomain {
		receipt, err := dt.SendAllFundsToHub(maxFee, enum.FAST_TRANSFER)
		if err != nil {
			return "", fmt.Errorf("send all funds to hub: %w", err)
		}

		return receipt.TxHash.String(), nil
	} else {
		return "", fmt.Errorf("invalid sending path")
	}
}

func (t *Transactor) BatchExecute(ctx context.Context, domain enum.CircleDomain, targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (string, error) {
	dt, exists := t.domains[domain]
	if !exists {
		return "", fmt.Errorf("domain %d not found", domain)
	}

	receipt, err := dt.BatchExecute(targets, selectors, data, proofs)
	if err != nil {
		return "", fmt.Errorf("batch execute: %w", err)
	}

	return receipt.TxHash.String(), nil
}

func (t *Transactor) process(ctx context.Context) error {
	if err := t.processPendingBridgeOps(ctx); err != nil {
		return fmt.Errorf("process pending bridge ops: %w", err)
	}

	if err := t.processGatewayDeposits(ctx); err != nil {
		return fmt.Errorf("process gateway deposits: %w", err)
	}

	if err := t.processGatewayMints(ctx); err != nil {
		return fmt.Errorf("process gateway mints: %w", err)
	}

	return nil
}

func (t *Transactor) processPendingBridgeOps(ctx context.Context) error {
	ops, err := t.brideRepo.GetPendingOps(ctx)
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

		go func(domain enum.CircleDomain, message []byte, attestation []byte) {
			defer wg.Done()

			receipt, err := t.domains[domain].ReceiveMessage(message, attestation)
			if err != nil {
				log.Error().Err(err).Msg("receive message failed")
				return
			}

			if err := t.brideRepo.UpdateReceivedTxHash(ctx, op.ID, receipt.TxHash); err != nil {
				log.Error().Err(err).Msg("update received tx hash failed")
				return
			}
		}(enum.CircleDomain(op.ToDomain), msg, att)
	}

	wg.Wait()

	return nil
}

func (t *Transactor) processGatewayDeposits(ctx context.Context) error {
	ops, err := t.userRepo.GetPendingPermitDeposits(ctx)
	if err != nil {
		return fmt.Errorf("get pending permit deposits: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(ops))

	for _, op := range ops {
		go func(permit models.UserUnifiedPermits) {
			defer wg.Done()

			domain := enum.CircleDomain(permit.Domain)
			token := common.HexToAddress(permit.Token)
			owner := common.HexToAddress(permit.Owner)
			deadline := new(big.Int).SetInt64(permit.Deadline)

			amount, ok := new(big.Int).SetString(permit.Value, 10)
			if !ok {
				log.Error().Msg("unable to parse amount")
				return
			}

			sig, err := hexutil.Decode(permit.Signature)
			if err != nil {
				log.Error().Err(err).Msg("unable to decode signature")
				return
			}

			receipt, err := t.domains[domain].DepositWithPermit(token, owner, amount, deadline, sig)
			if err != nil {
				log.Error().Err(err).Msg("deposit with permit failed")
				return
			}

			if err := t.userRepo.UpdatePermitDepositTxHash(ctx, permit.DepositID.String(), owner, uint32(domain), receipt.TxHash); err != nil {
				log.Error().Err(err).Msg("update permit deposit tx hash failed")
				return
			}
		}(op)
	}

	wg.Wait()

	return nil
}

func (t *Transactor) processGatewayMints(ctx context.Context) error {
	mints, err := t.userRepo.GetPendingHubDeposits(ctx)
	if err != nil {
		return fmt.Errorf("get pending hub deposits: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(mints))

	for _, deposit := range mints {
		attestation, err := hexutil.Decode(deposit.Attestation)
		if err != nil {
			log.Error().Err(err).Msg("unable to decode attestation")
			return err
		}

		sig, err := hexutil.Decode(deposit.Signature)
		if err != nil {
			log.Error().Err(err).Msg("unable to decode signature")
			return err
		}

		receipt, err := t.domains[t.hubDomain].DepositAndStake(attestation, sig)
		if err != nil {
			log.Error().Err(err).Msg("deposit and stake failed")
			return err
		}

		if err := t.userRepo.UpdateDepositTxHash(ctx, deposit.ID.String(), receipt.TxHash); err != nil {
			log.Error().Err(err).Msg("update deposit tx hash failed")
			return err
		}
	}

	return nil
}
