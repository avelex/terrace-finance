package bridge

import (
	"context"
	"fmt"
	"math/big"

	cctp_client "github.com/avelex/terrace-finance/backend/internal/circle/cctp/client"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/models/response"
	"github.com/avelex/terrace-finance/backend/internal/repository"
	"github.com/avelex/terrace-finance/backend/internal/transactor"
	"github.com/shopspring/decimal"
)

type Service struct {
	repo       *repository.BridgeRepository
	transactor *transactor.Transactor
	cctp       *cctp_client.Client
}

func NewService(repo *repository.BridgeRepository, transactor *transactor.Transactor, cctp *cctp_client.Client) *Service {
	return &Service{
		repo:       repo,
		transactor: transactor,
		cctp:       cctp,
	}
}

func (s *Service) BridgeFunds(ctx context.Context, srcDomain, dstDomain enum.CircleDomain) (string, error) {
	fees, err := s.cctp.Fees(ctx, uint32(srcDomain), uint32(dstDomain))
	if err != nil {
		return "", fmt.Errorf("get fees: %w", err)
	}

	fastMinFee := fees.FastTransfer()
	maxFee := calculateMaxFee(fastMinFee)

	txHash, err := s.transactor.SendAllFunds(ctx, srcDomain, dstDomain, maxFee)
	if err != nil {
		return "", fmt.Errorf("send all funds: %w", err)
	}

	return txHash, nil
}

func (s *Service) GetVaultsInfo(ctx context.Context) ([]response.VaultInfo, error) {
	supportedDomains := enum.SupportedDomains()
	vaults := make([]response.VaultInfo, 0, len(supportedDomains))
	for _, domain := range supportedDomains {
		address := s.transactor.TerraceAddress(domain)
		balance, err := s.transactor.TerraceBalance(ctx, domain, enum.USDC_MAPPING[enum.NetworkByDomain(domain)])
		if err != nil {
			return nil, fmt.Errorf("unable to get vault balance: %w", err)
		}
		vaults = append(vaults, response.VaultInfo{
			Domain:  domain,
			Address: address,
			Balance: decimal.NewFromBigInt(balance, 0),
		})
	}

	return vaults, nil
}

func (s *Service) GetBridgeOperations(ctx context.Context, limit, page int) (response.PagableBridgeOps, error) {
	ops, total, err := s.repo.GetBridgeOperations(ctx, limit, page)
	if err != nil {
		return response.PagableBridgeOps{}, fmt.Errorf("unable to get bridge operations: %w", err)
	}

	return response.NewPagableBridgeOps(ops, total, int64(limit), int64(page)), nil
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
