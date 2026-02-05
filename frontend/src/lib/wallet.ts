import { createPublicClient, createWalletClient, custom, http, Chain } from 'viem';
import { arcTestnet } from 'viem/chains';
import { CHAIN_ID_TO_CHAIN } from './constants';

// Public client for reading blockchain data (default to Arc testnet)
export const publicClient = createPublicClient({
  chain: arcTestnet,
  transport: http(),
});

// Create public client for specific chain
export function createPublicClientForChain(chainId: number) {
  const chain = CHAIN_ID_TO_CHAIN[chainId] || {
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
