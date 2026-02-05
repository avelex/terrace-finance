'use client';

import useSWR from 'swr';
import { VaultInfo } from '@/lib/types';
import { API_BASE_URL } from '@/lib/constants';

interface UseVaultsInfoResult {
    vaults: VaultInfo[];
    isLoading: boolean;
    error: string | null;
    refetch: () => Promise<void>;
}

const fetcher = async (url: string): Promise<VaultInfo[]> => {
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`Failed to fetch vaults: ${response.status}`);
    }
    return response.json();
};

export function useVaultsInfo(): UseVaultsInfoResult {
    const { data, error, isLoading, mutate } = useSWR<VaultInfo[]>(
        `${API_BASE_URL}/api/bridge/vaults`,
        fetcher,
        {
            revalidateOnFocus: false,
            dedupingInterval: 10000,
        }
    );

    // Sort by domain
    const sortedVaults = (data ?? []).sort((a, b) => a.domain - b.domain);

    return {
        vaults: sortedVaults,
        isLoading,
        error: error?.message ?? null,
        refetch: async () => { await mutate(); },
    };
}
