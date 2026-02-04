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

type Repository interface {
	GetPendingOps(ctx context.Context) ([]models.BridgeOp, error)
	UpdateReceivedTxHash(ctx context.Context, id string, txHash common.Hash) error
}

type Transactor struct {
	hubDomain enum.CircleDomain
	domains   map[enum.CircleDomain]*DomainTransactor
	repo      Repository
}

func NewTransactor(hubDomain enum.CircleDomain, domains map[enum.CircleDomain]*DomainTransactor, repo Repository) *Transactor {
	return &Transactor{
		hubDomain: hubDomain,
		domains:   domains,
		repo:      repo,
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
	return t.domains[domain].contract
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

		go func(domain enum.CircleDomain, message []byte, attestation []byte) {
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
		}(enum.CircleDomain(op.ToDomain), msg, att)
	}

	wg.Wait()

	return nil
}
