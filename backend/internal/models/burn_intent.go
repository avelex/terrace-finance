package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BurnIntent struct {
	MaxBlockHeight *big.Int     `json:"maxBlockHeight"`
	MaxFee         *big.Int     `json:"maxFee"`
	Spec           TransferSpec `json:"spec"`
	Signature      string       `json:"signature"`
}

type TransferSpec struct {
	Version              int         `json:"version"`
	SourceDomain         uint32      `json:"sourceDomain"`
	DestinationDomain    uint32      `json:"destinationDomain"`
	SourceContract       common.Hash `json:"sourceContract"`
	DestinationContract  common.Hash `json:"destinationContract"`
	SourceToken          common.Hash `json:"sourceToken"`
	DestinationToken     common.Hash `json:"destinationToken"`
	SourceDepositor      common.Hash `json:"sourceDepositor"`
	DestinationRecipient common.Hash `json:"destinationRecipient"`
	SourceSigner         common.Hash `json:"sourceSigner"`
	DestinationCaller    common.Hash `json:"destinationCaller"`
	Value                *big.Int    `json:"value"`
	Salt                 []byte      `json:"salt"`
	HookData             []byte      `json:"hookData"`
}
