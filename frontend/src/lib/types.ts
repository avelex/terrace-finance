// Circle CCTP Domain IDs
export enum CircleDomain {
  ETHEREUM = 0,
  AVAX = 1,
  BASE = 6,
  ARC = 26,
}

// Human-readable chain names
export const DOMAIN_NAMES: Record<CircleDomain, string> = {
  [CircleDomain.ETHEREUM]: 'Ethereum',
  [CircleDomain.AVAX]: 'Avalanche',
  [CircleDomain.BASE]: 'Base',
  [CircleDomain.ARC]: 'Arc',
};

// Domains to display in the UI (ordered as in design)
export const DISPLAYED_DOMAINS: CircleDomain[] = [
  CircleDomain.ARC,
  CircleDomain.ETHEREUM,
  CircleDomain.BASE,
  CircleDomain.AVAX,
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

// EIP-712 TypedData structure (from viem)
export interface TypedDataDomain {
  name?: string;
  version?: string;
  chainId?: number | string; // can be hex string like "0xa4b1"
  verifyingContract?: `0x${string}`;
  salt?: `0x${string}`;
}

export interface TypedData {
  domain: TypedDataDomain;
  types: Record<string, { name: string; type: string }[]>;
  primaryType: string;
  message: Record<string, unknown>;
}

// POST /api/wallet/:address/unify - Request
export interface UnifyRequest {
  domains: number[];
}

// POST /api/wallet/:address/unify - Response
// Note: Backend sends typedData as JSON string, parsed in useUnify hook
export interface UnifyResponse {
  id: string;
  domains: Record<number, TypedData | string>;
}

// POST /api/wallet/:address/unify/permits - Request
export interface PermitsRequest {
  id: string;
  domains: Record<number, string>; // domain -> signature (0x...)
}

// POST /api/wallet/:address/deposit_stake - Request
export interface DepositStakeRequest {
  id: string;
  attestation: string;
  signature: string;
}

// GET /api/wallet/:address/deposit_stake/list - Response item
export interface UserDeposit {
  id: string;
  address: string;
  value: string;
  destDomain: number;
  destGatewayMint: string;
  attestation: string;
  signature: string;
  txHash: string;
  reason: string;
  createdAt: string;
  depositedAt: string;
}

// GET /api/bridge/operations - Response item
export interface BridgeOperation {
  id: string;
  fromTerrace: string;
  fromDomain: number;
  toTerrace: string;
  toDomain: number;
  sendAmount: string;
  receivedAmount: string;
  sentTxHash: string;
  receivedTxHash: string;
  sentAt: string;
  receivedAt: string;
}

// GET /api/bridge/operations - Paginated response
export interface PaginatedBridgeOps {
  data: BridgeOperation[];
  total: number;
  page: number;
  limit: number;
}

// GET /api/wallet/:address/unify/list - Response item
export interface UserUnifiedPermit {
  id: string;
  owner: string;
  token: string;
  value: string;
  deadline: number;
  signature: string;
  domain: number;
  gatewayWallet: string;
  txHash: string;
  createdAt: string;
  depositedAt: string;
}

