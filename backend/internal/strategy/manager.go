package strategy

import (
	"context"
	"fmt"
	"math/big"

	cctp_client "github.com/avelex/terrace-finance/backend/internal/cctp/client"
	"github.com/avelex/terrace-finance/backend/internal/models"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/repository"
	"github.com/avelex/terrace-finance/backend/internal/transactor"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type AaveRepository interface {
	InsertAaveV3SupplySnapshot(ctx context.Context, snapshot models.AaveV3SupplySnapshot) error
}

type Manager struct {
	transactor   *transactor.Transactor
	aaveRepo     AaveRepository
	strategyRepo *repository.StrategyRepository
	cctp         *cctp_client.Client
}

func NewManager(transactor *transactor.Transactor, aaveRepo AaveRepository, strategyRepo *repository.StrategyRepository, cctp *cctp_client.Client) *Manager {
	return &Manager{
		transactor:   transactor,
		aaveRepo:     aaveRepo,
		strategyRepo: strategyRepo,
		cctp:         cctp,
	}
}

func (m *Manager) UpdateReserveData(ctx context.Context, reserve models.ReserveDataUpdated) error {
	// skip, if network is not supported
	usdc, exists := enum.USDC_MAPPING[reserve.Network]
	if !exists {
		return nil
	}

	// skip, if reserve is not USDC
	if reserve.Reserve.Cmp(usdc) != 0 {
		return nil
	}

	if err := m.aaveRepo.InsertAaveV3SupplySnapshot(ctx, reserve.Snapshot()); err != nil {
		return fmt.Errorf("insert aave v3 supply snapshot: %w", err)
	}

	return nil
}

func (m *Manager) SendAllFunds(ctx context.Context, srcDomain, dstDomain uint32) (string, error) {
	fees, err := m.cctp.Fees(ctx, srcDomain, dstDomain)
	if err != nil {
		return "", fmt.Errorf("get fees: %w", err)
	}

	fastMinFee := fees.FastTransfer()
	maxFee := calculateMaxFee(fastMinFee)

	txHash, err := m.transactor.SendAllFunds(ctx, srcDomain, dstDomain, maxFee)
	if err != nil {
		return "", fmt.Errorf("send all funds: %w", err)
	}

	return txHash, nil
}

func (m *Manager) SupplyAllFundsToAaveV3(ctx context.Context, domain uint32) (string, error) {
	network := enum.NetworkByDomain(domain)
	usdc := enum.USDC_MAPPING[network]
	aavePool := enum.AAVE_V3[network]

	terraceBalance, err := m.transactor.TerraceBalance(ctx, domain, usdc)
	if err != nil {
		return "", fmt.Errorf("get terrace balance: %w", err)
	}

	supplyStrategy, err := m.strategyRepo.GetByDomainAndName(ctx, domain, "supply")
	if err != nil {
		return "", fmt.Errorf("get supply strategy: %w", err)
	}

	approveStrategy, err := m.strategyRepo.GetByDomainAndName(ctx, domain, "approve")
	if err != nil {
		return "", fmt.Errorf("get approve strategy: %w", err)
	}

	encodedApproveData, err := encodeApprove(aavePool, terraceBalance)
	if err != nil {
		return "", fmt.Errorf("encode approve data: %w", err)
	}

	encodedSupplyData, err := encodeSupply(usdc, terraceBalance, m.transactor.TerraceAddress(domain))
	if err != nil {
		return "", fmt.Errorf("encode supply data: %w", err)
	}

	targets := []common.Address{approveStrategy.Target(), supplyStrategy.Target()}
	selectors := [][]byte{approveStrategy.SelectorBytes(), supplyStrategy.SelectorBytes()}
	data := [][]byte{encodedApproveData, encodedSupplyData}
	proofs := [][][32]byte{approveStrategy.ProofBytes(), supplyStrategy.ProofBytes()}

	txHash, err := m.transactor.BatchExecute(ctx, domain, targets, selectors, data, proofs)
	if err != nil {
		return "", fmt.Errorf("batch execute: %w", err)
	}

	return txHash, nil
}

func (m *Manager) WithdrawAllFundsFromAaveV3(ctx context.Context, domain uint32) (string, error) {
	network := enum.NetworkByDomain(domain)
	usdc := enum.USDC_MAPPING[network]

	encodedWithdrawData, err := encodeWithdraw(usdc, m.transactor.TerraceAddress(domain))
	if err != nil {
		return "", fmt.Errorf("encode withdraw data: %w", err)
	}

	withdrawStrategy, err := m.strategyRepo.GetByDomainAndName(ctx, domain, "withdraw")
	if err != nil {
		return "", fmt.Errorf("get withdraw strategy: %w", err)
	}

	targets := []common.Address{withdrawStrategy.Target()}
	selectors := [][]byte{withdrawStrategy.SelectorBytes()}
	data := [][]byte{encodedWithdrawData}
	proofs := [][][32]byte{withdrawStrategy.ProofBytes()}

	txHash, err := m.transactor.BatchExecute(ctx, domain, targets, selectors, data, proofs)
	if err != nil {
		return "", fmt.Errorf("batch execute: %w", err)
	}

	return txHash, nil
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

func encodeApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	typeUint256, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	typeAddress, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{
		{Type: typeAddress},
		{Type: typeUint256},
	}

	encodedData, err := arguments.Pack(spender, amount)
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}

func encodeSupply(asset common.Address, amount *big.Int, onBehalfOf common.Address) ([]byte, error) {
	typeUint256, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	typeAddress, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}
	typeUint16, err := abi.NewType("uint16", "", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{
		{Type: typeAddress},
		{Type: typeUint256},
		{Type: typeAddress},
		{Type: typeUint16},
	}

	encodedData, err := arguments.Pack(asset, amount, onBehalfOf, uint16(0))
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}

func encodeWithdraw(asset common.Address, to common.Address) ([]byte, error) {
	maxUint256, ok := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	if !ok {
		return nil, fmt.Errorf("failed to set max uint256")
	}

	typeUint256, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	typeAddress, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{
		{Type: typeAddress},
		{Type: typeUint256},
		{Type: typeAddress},
	}

	encodedData, err := arguments.Pack(asset, maxUint256, to)
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}
