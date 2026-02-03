package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	WAD_DECIMALS = 18
	RAY_DECIMALS = 27
)

func ConvertToDecimal(amount decimal.Decimal, decimals int32) decimal.Decimal {
	return decimal.NewFromBigInt(amount.BigInt(), -decimals)
}

func ConvertBigIntToDecimal(amount *big.Int, decimals int32) decimal.Decimal {
	return decimal.NewFromBigInt(amount, -decimals)
}

func ConvertRayToDecimal(amount *big.Int) decimal.Decimal {
	return ConvertBigIntToDecimal(amount, RAY_DECIMALS)
}

func ConvertWadToDecimal(amount *big.Int) decimal.Decimal {
	return ConvertBigIntToDecimal(amount, WAD_DECIMALS)
}
