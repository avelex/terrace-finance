package enum

import "github.com/ethereum/go-ethereum/common"

const (
	FAST_TRANSFER     = 1_000
	STANDARD_TRANSFER = 2_000
)

const (
	ETHEREUM_DOMAIN = 0
	AVAX_DOMAIN     = 1
	OPTIMISM_DOMAIN = 2
	ARBITRUM_DOMAIN = 3
	BASE_DOMAIN     = 6
	POLYGON_DOMAIN  = 7
	SONIC_DOMAIN    = 13
	ARC_DOMAIN      = 26
)

var USDC_MAPPING = map[string]common.Address{
	"base":     common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"),
	"arbitrum": common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
}

var AAVE_V3 = map[string]common.Address{
	"base":     common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
	"arbitrum": common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
}

var DomainNetworkMapping = map[uint32]string{
	ETHEREUM_DOMAIN: "ethereum",
	AVAX_DOMAIN:     "avalanche",
	OPTIMISM_DOMAIN: "optimism",
	ARBITRUM_DOMAIN: "arbitrum",
	BASE_DOMAIN:     "base",
	POLYGON_DOMAIN:  "polygon",
	SONIC_DOMAIN:    "sonic",
}

var NetworkDomainMapping = map[string]uint32{
	"ethereum":  ETHEREUM_DOMAIN,
	"avalanche": AVAX_DOMAIN,
	"optimism":  OPTIMISM_DOMAIN,
	"arbitrum":  ARBITRUM_DOMAIN,
	"base":      BASE_DOMAIN,
	"polygon":   POLYGON_DOMAIN,
	"sonic":     SONIC_DOMAIN,
}

func NetworkByDomain(domain uint32) string {
	return DomainNetworkMapping[domain]
}

func DomainByNetwork(network string) uint32 {
	return NetworkDomainMapping[network]
}
