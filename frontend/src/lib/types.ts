// Circle CCTP Domain IDs
export enum CircleDomain {
  ETHEREUM = 0,
  AVAX = 1,
  OPTIMISM = 2,
  ARBITRUM = 3,
  BASE = 6,
  POLYGON = 7,
  ARC = 26,
}

// Human-readable chain names
export const DOMAIN_NAMES: Record<CircleDomain, string> = {
  [CircleDomain.ETHEREUM]: 'Ethereum',
  [CircleDomain.AVAX]: 'Avalanche',
  [CircleDomain.OPTIMISM]: 'Optimism',
  [CircleDomain.ARBITRUM]: 'Arbitrum',
  [CircleDomain.BASE]: 'Base',
  [CircleDomain.POLYGON]: 'Polygon',
  [CircleDomain.ARC]: 'Arc',
};

// Domains to display in the UI (ordered as in design)
export const DISPLAYED_DOMAINS: CircleDomain[] = [
  CircleDomain.ARC,
  CircleDomain.ETHEREUM,
  CircleDomain.ARBITRUM,
  CircleDomain.BASE,
  CircleDomain.AVAX,
  CircleDomain.POLYGON,
];

// Token balances by domain
export interface TokenBalancesByDomain {
  [domain: number]: number;
}

// API Response from GET /api/wallet/:address/balances
export interface BalancesResponse {
  stablecoinBalance: number;
  stakingBalance: number;
  usdc: TokenBalancesByDomain;
  unifiedUsdc: TokenBalancesByDomain;
}
