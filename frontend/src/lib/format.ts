/**
 * Shared formatting utilities for the frontend
 */

/**
 * Format a number as a balance with locale formatting
 * @param value - The numeric value to format
 * @param decimals - Optional, max decimal places (default: 2)
 */
export function formatBalance(value: number, decimals: number = 2): string {
    return value.toLocaleString('en-US', {
        minimumFractionDigits: 0,
        maximumFractionDigits: decimals,
    });
}

/**
 * Format a USDC value (6 decimals) for display
 * @param value - The raw value as string, number, or BigInt (in smallest units)
 */
export function formatUsdcValue(value: number | bigint): string {
    let num: number;
    if (typeof value === 'bigint') {
        num = Number(value) / 1e6;
    } else {
        num = value / 1e6;
    }
    return num.toLocaleString('en-US', {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2,
    });
}

/**
 * Format an ISO date string for display
 * @param dateStr - ISO date string
 */
export function formatDate(dateStr: string): string {
    if (!dateStr || dateStr === '0001-01-01T00:00:00Z') return '-';
    return new Date(dateStr).toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
    });
}

/**
 * Shorten an Ethereum address for display
 * @param address - Full address
 * @param chars - Number of characters to show on each side (default: 4)
 */
export function shortenAddress(address: string, chars: number = 4): string {
    return `${address.slice(0, chars + 2)}...${address.slice(-chars)}`;
}

/**
 * Format a staking balance value (18 decimals) for display
 * @param value - The raw value (in smallest units with 18 decimals)
 */
export function formatStakingBalance(value: number | bigint): string {
    let num: number;
    if (typeof value === 'bigint') {
        // For BigInt with 18 decimals, convert to number
        num = Number(value) / 1e18;
    } else {
        num = value / 1e18;
    }
    return num.toLocaleString('en-US', {
        minimumFractionDigits: 0,
        maximumFractionDigits: 2,
    });
}

/**
 * Convert BigInt (6 decimals) to human-readable number string
 */
export function bigintToHumanReadable(value: bigint, decimals: number = 6): string {
    const divisor = BigInt(10 ** decimals);
    const integerPart = value / divisor;
    const fractionalPart = value % divisor;
    const fractionalStr = fractionalPart.toString().padStart(decimals, '0');
    // Trim trailing zeros but keep at least 2 decimal places
    const trimmed = fractionalStr.replace(/0+$/, '').padEnd(2, '0');
    return `${integerPart}.${trimmed}`;
}

