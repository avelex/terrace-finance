'use client';

import useSWR from 'swr';
import { BalancesResponse } from '@/lib/types';
import { API_BASE_URL } from '@/lib/constants';

interface UseBalancesResult {
  balances: BalancesResponse | null;
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

const fetcher = async (url: string): Promise<BalancesResponse> => {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Failed to fetch balances: ${response.status}`);
  }
  return response.json();
};

export function useBalances(address: `0x${string}` | null): UseBalancesResult {
  const { data, error, isLoading, mutate } = useSWR<BalancesResponse>(
    address ? `${API_BASE_URL}/api/wallet/${address}/balances` : null,
    fetcher,
    {
      revalidateOnFocus: false,
      dedupingInterval: 5000,
    }
  );

  return {
    balances: data ?? null,
    isLoading,
    error: error?.message ?? null,
    refetch: async () => { await mutate(); },
  };
}
