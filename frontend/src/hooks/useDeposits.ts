'use client';

import { useState, useEffect, useCallback } from 'react';
import { UserDeposit } from '@/lib/types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || '';

interface UseDepositsResult {
  deposits: UserDeposit[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export function useDeposits(address: `0x${string}` | null): UseDepositsResult {
  const [deposits, setDeposits] = useState<UserDeposit[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchDeposits = useCallback(async () => {
    if (!address) {
      setDeposits([]);
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch(`${API_BASE_URL}/api/wallet/${address}/deposit_stake/list`);
      
      if (!response.ok) {
        throw new Error(`Failed to fetch deposits: ${response.status}`);
      }

      const data: UserDeposit[] = await response.json();
      setDeposits(data ?? []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
      setDeposits([]);
    } finally {
      setIsLoading(false);
    }
  }, [address]);

  useEffect(() => {
    fetchDeposits();
  }, [fetchDeposits]);

  return {
    deposits,
    isLoading,
    error,
    refetch: fetchDeposits,
  };
}
