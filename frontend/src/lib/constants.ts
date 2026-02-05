import { Chain } from 'viem';
import { arcTestnet, baseSepolia, sepolia, avalancheFuji } from 'viem/chains';
import { CircleDomain } from './types';

/**
 * API base URL for backend requests
 */
export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || '';

/**
 * Circle Gateway constants
 */
export const GATEWAY_TRANSFER_API_URL = 'https://gateway-api-testnet.circle.com/v1/transfer';
export const DESTINATION_DOMAIN = 26; // Arc testnet
export const DESTINATION_RECIPIENT = '0x851addF749C87F1159F30399F2ca160F8Ccdb6B1';
export const DESTINATION_CALLER = '0x851addF749C87F1159F30399F2ca160F8Ccdb6B1';
export const GATEWAY_MINTER_ADDRESS = '0x0022222ABE238Cc2C7Bb1f21003F0a260052475B';
export const GATEWAY_WALLET_ADDRESS = '0x0077777d7EBA4688BDeF3E311b846F25870A19B9';

/**
 * Maximum fee for gateway operations (1.0 USDC)
 */
export const MAX_FEE = BigInt(1_000000);

/**
 * USDC contract addresses by Circle domain
 */
export const USDC_ADDRESSES: Record<number, `0x${string}`> = {
    [CircleDomain.ARC]: '0x3600000000000000000000000000000000000000',
    [CircleDomain.ETHEREUM]: '0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238',
    [CircleDomain.BASE]: '0x036CbD53842c5426634e7929541eC2318f3dCF7e',
    [CircleDomain.AVAX]: '0x5425890298aed601595a70AB815c96711a31Bc65',
};

/**
 * Map Circle domain to viem Chain config
 */
export const DOMAIN_TO_CHAIN: Record<number, Chain> = {
    [CircleDomain.ETHEREUM]: sepolia,
    [CircleDomain.AVAX]: avalancheFuji,
    [CircleDomain.BASE]: baseSepolia,
    [CircleDomain.ARC]: arcTestnet,
};

/**
 * Map chain ID to viem Chain config
 */
export const CHAIN_ID_TO_CHAIN: Record<number, Chain> = {
    11155111: sepolia,
    43113: avalancheFuji,
    84532: baseSepolia,
    5042002: arcTestnet,
};

/**
 * Gas fees by domain (in USDC subunits, 6 decimals)
 * These are approximate values based on network costs
 */
export const GAS_FEES: Record<number, bigint> = {
    [CircleDomain.ETHEREUM]: BigInt(2000000), // ~$2.00
    [CircleDomain.AVAX]: BigInt(20000),       // ~$0.02
    [CircleDomain.BASE]: BigInt(10000),       // ~$0.01
    [CircleDomain.ARC]: BigInt(10000),        // ~$0.01
};

/**
 * Block explorer URLs by domain (testnets)
 */
export const EXPLORER_URLS: Record<number, string> = {
    [CircleDomain.ETHEREUM]: 'https://eth-sepolia.blockscout.com',
    [CircleDomain.AVAX]: 'https://testnet.snowtrace.io',
    [CircleDomain.BASE]: 'https://base-sepolia.blockscout.com',
    [CircleDomain.ARC]: 'https://testnet.arcscan.app',
};

/**
 * Get explorer URL for an address
 */
export function getAddressExplorerUrl(domain: number, address: string): string {
    const baseUrl = EXPLORER_URLS[domain];
    return baseUrl ? `${baseUrl}/address/${address}` : '#';
}

/**
 * Get explorer URL for a transaction
 */
export function getTxExplorerUrl(domain: number, txHash: string): string {
    const baseUrl = EXPLORER_URLS[domain];
    return baseUrl ? `${baseUrl}/tx/${txHash}` : '#';
}
