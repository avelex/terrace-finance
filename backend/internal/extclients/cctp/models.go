package cctp

import "github.com/shopspring/decimal"

type messagesResponse struct {
	Messages []struct {
		Message     string `json:"message"`
		EventNonce  string `json:"eventNonce"`
		Attestation string `json:"attestation"`
		Version     int    `json:"cctpVersion"`
		Status      string `json:"status"`
	} `json:"messages"`
}

type Fees []struct {
	FinalityThreshold int             `json:"finalityThreshold"`
	MinimumFee        decimal.Decimal `json:"minimumFee"`
}

func (f Fees) StandardTransfer() decimal.Decimal {
	for _, fee := range f {
		if fee.FinalityThreshold > 1_000 {
			return fee.MinimumFee
		}
	}
	return decimal.Zero
}

func (f Fees) FastTransfer() decimal.Decimal {
	for _, fee := range f {
		if fee.FinalityThreshold <= 1_000 {
			return fee.MinimumFee
		}
	}
	return decimal.Zero
}
