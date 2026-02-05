package models

import (
	"math/big"
	"time"

	utils "github.com/avelex/terrace-finance/backend/internal/utils/events"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type EventLocation struct {
	TxHash    common.Hash `json:"-"`
	Timestamp int64       `json:"-"`
	Network   string      `json:"-"`
}

func NewEventLocation(txHash common.Hash, timestamp int64, network string) EventLocation {
	return EventLocation{
		TxHash:    txHash,
		Timestamp: timestamp,
		Network:   network,
	}
}

type SendFunds struct {
	EventLocation

	ID          [32]byte       `json:"id"`
	FromTerrace common.Address `json:"fromTerrace"`
	FromDomain  uint32         `json:"fromDomain"`
	ToTerrace   common.Address `json:"toTerrace"`
	ToDomain    uint32         `json:"toDomain"`
	Amount      *big.Int       `json:"amount"`
}

func (s SendFunds) BridgeOp() BridgeOp {
	sentAt := time.Unix(s.Timestamp, 0)
	if sentAt.IsZero() {
		sentAt = time.Now()
	}

	return BridgeOp{
		ID:          common.BytesToHash(s.ID[:]).String(),
		FromTerrace: s.FromTerrace.String(),
		FromDomain:  s.FromDomain,
		ToTerrace:   s.ToTerrace.String(),
		ToDomain:    s.ToDomain,
		SendAmount:  s.Amount.String(),
		SentTxHash:  s.TxHash.String(),
		SentAt:      sentAt,
	}
}

type ReceivedFunds struct {
	EventLocation

	ID          [32]byte       `json:"id"`
	FromTerrace common.Address `json:"fromTerrace"`
	FromDomain  uint32         `json:"fromDomain"`
	ToTerrace   common.Address `json:"toTerrace"`
	ToDomain    uint32         `json:"toDomain"`
	Amount      *big.Int       `json:"amount"`
}

type ReserveDataUpdated struct {
	EventLocation
	Reserve             common.Address `json:"reserve"`
	LiquidityRate       *big.Int       `json:"liquidityRate"`
	StableBorrowRate    *big.Int       `json:"stableBorrowRate"`
	VariableBorrowRate  *big.Int       `json:"variableBorrowRate"`
	LiquidityIndex      *big.Int       `json:"liquidityIndex"`
	VariableBorrowIndex *big.Int       `json:"variableBorrowIndex"`
}

func (r *ReserveDataUpdated) SupplyAPY() decimal.Decimal {
	return utils.ConvertRayToDecimal(r.LiquidityRate).Mul(decimal.NewFromInt(100))
}

func (r *ReserveDataUpdated) BorrowAPY() decimal.Decimal {
	return utils.ConvertRayToDecimal(r.VariableBorrowRate).Mul(decimal.NewFromInt(100))
}

func (r *ReserveDataUpdated) Snapshot() AaveV3SupplySnapshot {
	return AaveV3SupplySnapshot{
		Network:   r.Network,
		APY:       r.SupplyAPY(),
		Timestamp: time.Unix(r.Timestamp, 0),
	}
}
