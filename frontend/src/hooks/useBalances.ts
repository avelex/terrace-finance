'use client';

import { useState, useEffect, useCallback } from 'react';
import { BalancesResponse } from '@/lib/types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || '';

interface UseBalancesResult {
  balances: BalancesResponse | null;
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export function useBalances(address: `0x${string}` | null): UseBalancesResult {
  const [balances, setBalances] = useState<BalancesResponse | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchBalances = useCallback(async () => {
    if (!address) {
      setBalances(null);
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch(`${API_BASE_URL}/api/wallet/${address}/balances`);
      
      if (!response.ok) {
        throw new Error(`Failed to fetch balances: ${response.status}`);
      }

      const data: BalancesResponse = await response.json();
      setBalances(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
      setBalances(null);
    } finally {
      setIsLoading(false);
    }
  }, [address]);

  useEffect(() => {
    fetchBalances();
  }, [fetchBalances]);

  return {
    balances,
    isLoading,
    error,
    refetch: fetchBalances,
  };
}
