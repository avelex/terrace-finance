package client

import (
	"math/big"

	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type balancesResponse struct {
	Balances []balanceSource `json:"balances"`
}

func (b *balancesResponse) Map() map[enum.CircleDomain]decimal.Decimal {
	out := make(map[enum.CircleDomain]decimal.Decimal)
	for _, balance := range b.Balances {
		out[balance.Domain] = balance.Balance
	}
	return out
}

type balancesRequest struct {
	Token   string          `json:"token"`
	Sources []balanceSource `json:"sources"`
}

type balanceSource struct {
	Depositor string            `json:"depositor"`
	Domain    enum.CircleDomain `json:"domain"`
	Balance   decimal.Decimal   `json:"balance,omitempty"`
}

func newBalancesRequest(depositor common.Address, domains []enum.CircleDomain) *balancesRequest {
	sources := make([]balanceSource, len(domains))
	for _, domain := range domains {
		sources = append(sources, balanceSource{
			Depositor: depositor.String(),
			Domain:    domain,
		})
	}
	return &balancesRequest{
		Token:   "USDC",
		Sources: sources,
	}
}

/*
	{
	  "transferId": "3c90c3cc-0d44-4b50-8888-8dd25736052a",
	  "attestation": "<string>",
	  "signature": "<string>",
	  "fees": {
	    "total": "1.178",
	    "token": "USDC",
	    "perIntent": [
	      {
	        "transferSpecHash": "0x1234567890123456789012345678901234567890123456789012345678901234",
	        "domain": 0,
	        "baseFee": "2",
	        "transferFee": "0.125"
	      }
	    ]
	  },
	  "expirationBlock": "<string>"
	}
*/
type TransferRequest struct {
	ID          string `json:"transferId"`
	Attestation string `json:"attestation"`
	Signature   string `json:"signature"`
	Fees        struct {
		Total     decimal.Decimal `json:"total"`
		Token     string          `json:"token"`
		PerIntent []struct {
			TransferSpecHash string            `json:"transferSpecHash"`
			Domain           enum.CircleDomain `json:"domain"`
			BaseFee          decimal.Decimal   `json:"baseFee"`
			TransferFee      decimal.Decimal   `json:"transferFee"`
		} `json:"perIntent"`
	} `json:"fees"`
	ExpirationBlock *big.Int `json:"expirationBlock"`
}
