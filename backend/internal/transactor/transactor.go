package transactor

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	cctp_client "github.com/avelex/terrace-finance/backend/internal/cctp/client"
	"github.com/avelex/terrace-finance/backend/internal/models"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type Repository interface {
	GetPendingOps(ctx context.Context) ([]models.BridgeOp, error)
	UpdateReceivedTxHash(ctx context.Context, id string, txHash common.Hash) error
}

type Transactor struct {
	hubDomain uint32
	domains   map[uint32]*DomainTransactor
	repo      Repository
	cctp      *cctp_client.Client
}

func NewTransactor(hubDomain uint32, domains map[uint32]*DomainTransactor, repo Repository, cctp *cctp_client.Client) *Transactor {
	return &Transactor{
		hubDomain: hubDomain,
		domains:   domains,
		repo:      repo,
		cctp:      cctp,
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

func (t *Transactor) SendAllFunds(ctx context.Context, srcDomain, dstDomain uint32) (string, error) {
	dt, exists := t.domains[srcDomain]
	if !exists {
		return "", fmt.Errorf("domain %d not found", srcDomain)
	}

	fees, err := t.cctp.Fees(ctx, srcDomain, dstDomain)
	if err != nil {
		return "", fmt.Errorf("get fees: %w", err)
	}

	fastMinFee := fees.FastTransfer()
	maxFee := calculateMaxFee(fastMinFee)

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
		return "", fmt.Errorf("unknown send path")
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

// fee in bps, for example, 1 = 0.01% = 0.0001
func calculateMaxFee(fee decimal.Decimal) *big.Int {
	if fee.IsZero() {
		return big.NewInt(0)
	}

	// Convert to subunits and add 20% buffer
	feeSubunits := fee.Div(decimal.NewFromInt(1_00_00)).Mul(decimal.NewFromInt(1_000_000))
	maxFee := feeSubunits.Mul(decimal.NewFromFloat(1.2))
	return maxFee.BigInt()
}
