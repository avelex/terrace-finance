package enum

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
