package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/uptrace/bun"
)

type Strategy struct {
	bun.BaseModel `bun:"table:strategies"`

	ID            int64    `bun:"id,pk,autoincrement"`
	Domain        uint32   `bun:"domain,notnull"`
	Name          string   `bun:"name,notnull"`
	TargetAddress string   `bun:"target_address,notnull"`
	Selector      string   `bun:"selector,notnull"`
	Proof         []string `bun:"proof,array,notnull"`
}

func (s *Strategy) Target() common.Address {
	return common.HexToAddress(s.TargetAddress)
}

func (s *Strategy) ProofBytes() [][32]byte {
	proof := make([][32]byte, 0, len(s.Proof))
	for _, v := range s.Proof {
		proof = append(proof, common.HexToHash(v))
	}
	return proof
}

func (s *Strategy) SelectorBytes() []byte {
	return common.Hex2Bytes(s.Selector)
}
