import { createPublicClient, createWalletClient, custom, http, Chain } from 'viem';
import { arcTestnet, baseSepolia, sepolia, avalancheFuji } from 'viem/chains';

// Map chainId to chain config
const CHAIN_MAP: Record<number, Chain> = {
  11155111: sepolia,
  43113: avalancheFuji,
  84532: baseSepolia,
  5042002: arcTestnet,
};

// Public client for reading blockchain data (default mainnet)
export const publicClient = createPublicClient({
  chain: arcTestnet,
  transport: http(),
});

// Create public client for specific chain
export function createPublicClientForChain(chainId: number) {
  const chain = CHAIN_MAP[chainId] || { 
    id: chainId, 
    name: `Chain ${chainId}`,
    nativeCurrency: { name: 'ETH', symbol: 'ETH', decimals: 18 },
    rpcUrls: { default: { http: [] } },
  } as Chain;

  return createPublicClient({
    chain,
    transport: http(),
  });
}

// Create wallet client from browser provider (MetaMask)
export async function createWalletClientFromProvider() {
  if (typeof window === 'undefined' || !window.ethereum) {
    throw new Error('No ethereum provider found. Please install MetaMask.');
  }

  const walletClient = createWalletClient({
    chain: arcTestnet,
    transport: custom(window.ethereum),
  });

  return walletClient;
}

// Request accounts from MetaMask
export async function requestAccounts(): Promise<`0x${string}`[]> {
  if (typeof window === 'undefined' || !window.ethereum) {
    throw new Error('No ethereum provider found. Please install MetaMask.');
  }

  const accounts = await window.ethereum.request({
    method: 'eth_requestAccounts',
  }) as `0x${string}`[];

  return accounts;
}

// Get current accounts without requesting
export async function getAccounts(): Promise<`0x${string}`[]> {
  if (typeof window === 'undefined' || !window.ethereum) {
    return [];
  }

  const accounts = await window.ethereum.request({
    method: 'eth_accounts',
  }) as `0x${string}`[];

  return accounts;
}

// Shorten address for display
export function shortenAddress(address: string, chars = 4): string {
  return `${address.slice(0, chars + 2)}...${address.slice(-chars)}`;
}
