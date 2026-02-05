'use client';

import { useState, useEffect, useCallback } from 'react';
import { UserUnifiedPermit } from '@/lib/types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || '';

interface UsePermitsResult {
  permits: UserUnifiedPermit[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export function usePermits(address: `0x${string}` | null): UsePermitsResult {
  const [permits, setPermits] = useState<UserUnifiedPermit[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchPermits = useCallback(async () => {
    if (!address) {
      setPermits([]);
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch(`${API_BASE_URL}/api/wallet/${address}/unify/list`);
      
      if (!response.ok) {
        throw new Error(`Failed to fetch permits: ${response.status}`);
      }

      const data: UserUnifiedPermit[] = await response.json();
      setPermits(data ?? []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
      setPermits([]);
    } finally {
      setIsLoading(false);
    }
  }, [address]);

  useEffect(() => {
    fetchPermits();
  }, [fetchPermits]);

  return {
    permits,
    isLoading,
    error,
    refetch: fetchPermits,
  };
}
