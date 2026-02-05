'use client';

import useSWR from 'swr';
import { UserDeposit } from '@/lib/types';
import { API_BASE_URL } from '@/lib/constants';

interface UseDepositsResult {
  deposits: UserDeposit[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

const fetcher = async (url: string): Promise<UserDeposit[]> => {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Failed to fetch deposits: ${response.status}`);
  }
  return response.json();
};

export function useDeposits(address: `0x${string}` | null): UseDepositsResult {
  const { data, error, isLoading, mutate } = useSWR<UserDeposit[]>(
    address ? `${API_BASE_URL}/api/wallet/${address}/deposit_stake/list` : null,
    fetcher,
    {
      revalidateOnFocus: false,
      dedupingInterval: 5000,
    }
  );

  return {
    deposits: data ?? [],
    isLoading,
    error: error?.message ?? null,
    refetch: async () => { await mutate(); },
  };
}
