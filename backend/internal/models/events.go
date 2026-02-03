package models

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type SendFunds struct {
	ID          [32]byte       `json:"id"`
	FromTerrace common.Address `json:"fromTerrace"`
	FromDomain  uint32         `json:"fromDomain"`
	ToTerrace   common.Address `json:"toTerrace"`
	ToDomain    uint32         `json:"toDomain"`
	Amount      *big.Int       `json:"amount"`

	TxHash    common.Hash `json:"-"`
	Timestamp int64       `json:"-"`
	ChainID   *big.Int    `json:"-"`
}

func (s SendFunds) BridgeOp() BridgeOp {
	return BridgeOp{
		ID:          common.BytesToHash(s.ID[:]).String(),
		FromTerrace: s.FromTerrace.String(),
		FromDomain:  s.FromDomain,
		FromChainID: s.ChainID.String(),
		ToTerrace:   s.ToTerrace.String(),
		ToDomain:    s.ToDomain,
		SendAmount:  s.Amount.String(),
		SentTxHash:  s.TxHash.String(),
		SentAt:      time.Unix(s.Timestamp, 0),
	}
}

type ReceivedFunds struct {
	ID          [32]byte       `json:"id"`
	FromTerrace common.Address `json:"fromTerrace"`
	FromDomain  uint32         `json:"fromDomain"`
	ToTerrace   common.Address `json:"toTerrace"`
	ToDomain    uint32         `json:"toDomain"`
	Amount      *big.Int       `json:"amount"`

	TxHash    common.Hash `json:"-"`
	Timestamp int64       `json:"-"`
	ChainID   *big.Int    `json:"-"`
}
