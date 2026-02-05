'use client';

import useSWR from 'swr';
import { UserUnifiedPermit } from '@/lib/types';
import { API_BASE_URL } from '@/lib/constants';

interface UsePermitsResult {
  permits: UserUnifiedPermit[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

const fetcher = async (url: string): Promise<UserUnifiedPermit[]> => {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Failed to fetch permits: ${response.status}`);
  }
  return response.json();
};

export function usePermits(address: `0x${string}` | null): UsePermitsResult {
  const { data, error, isLoading, mutate } = useSWR<UserUnifiedPermit[]>(
    address ? `${API_BASE_URL}/api/wallet/${address}/unify/list` : null,
    fetcher,
    {
      revalidateOnFocus: false,
      dedupingInterval: 5000,
    }
  );

  return {
    permits: data ?? [],
    isLoading,
    error: error?.message ?? null,
    refetch: async () => { await mutate(); },
  };
}
