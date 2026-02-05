'use client';

import useSWR from 'swr';
import { PaginatedBridgeOps } from '@/lib/types';
import { API_BASE_URL } from '@/lib/constants';

interface UseBridgeOperationsResult {
    operations: PaginatedBridgeOps['data'];
    total: number;
    page: number;
    limit: number;
    isLoading: boolean;
    error: string | null;
    refetch: () => Promise<void>;
}

const fetcher = async (url: string): Promise<PaginatedBridgeOps> => {
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`Failed to fetch operations: ${response.status}`);
    }
    return response.json();
};

export function useBridgeOperations(page: number = 0, limit: number = 10): UseBridgeOperationsResult {
    const { data, error, isLoading, mutate } = useSWR<PaginatedBridgeOps>(
        `${API_BASE_URL}/api/bridge/operations?limit=${limit}&page=${page}`,
        fetcher,
        {
            revalidateOnFocus: false,
            dedupingInterval: 5000,
        }
    );

    return {
        operations: data?.data ?? [],
        total: data?.total ?? 0,
        page: data?.page ?? page,
        limit: data?.limit ?? limit,
        isLoading,
        error: error?.message ?? null,
        refetch: async () => { await mutate(); },
    };
}
