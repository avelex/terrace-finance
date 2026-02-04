import { createPublicClient, createWalletClient, custom, http } from 'viem';
import { mainnet } from 'viem/chains';

// Public client for reading blockchain data
export const publicClient = createPublicClient({
  chain: mainnet,
  transport: http(),
});

// Create wallet client from browser provider (MetaMask)
export async function createWalletClientFromProvider() {
  if (typeof window === 'undefined' || !window.ethereum) {
    throw new Error('No ethereum provider found. Please install MetaMask.');
  }

  const walletClient = createWalletClient({
    chain: mainnet,
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
