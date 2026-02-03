package client_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/avelex/terrace-finance/backend/internal/cctp/client"

	"github.com/stretchr/testify/require"
)

func Test_Fees(t *testing.T) {
	r := require.New(t)

	cctpClient := client.NewClient()
	ethereumDomain := "0"
	baseDomain := "6"

	fees, err := cctpClient.Fees(context.Background(), ethereumDomain, baseDomain)
	r.NoError(err)

	fmt.Printf("Fast Fee: %v\n", fees.FastTransfer())
	fmt.Printf("Standard Fee: %v\n", fees.StandardTransfer())
}

func Test_MessageAndAttestation(t *testing.T) {
	r := require.New(t)

	cctpClient := client.NewClient()
	srcDomain := "6"
	txHash := "0x9ce60a2cbf89c6dfe6d8535a6bea086417318859e0acdf93ff8afdd2a925e206"

	message, attestation, err := cctpClient.MessageAndAttestation(context.Background(), srcDomain, txHash)
	r.NoError(err)

	fmt.Printf("Message: %v\n", message)
	fmt.Printf("Attestation: %v\n", attestation)
}
