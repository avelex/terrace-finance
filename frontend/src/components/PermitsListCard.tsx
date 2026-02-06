'use client';

import { useMemo } from 'react';
import { useWallet } from '@/contexts/WalletContext';
import { usePermits } from '@/hooks/usePermits';
import { DOMAIN_NAMES, CircleDomain, UserUnifiedPermit } from '@/lib/types';
import { formatUsdcValue, formatDate } from '@/lib/format';
import { getTxExplorerUrl } from '@/lib/constants';
import { StatusBadge, getPermitGroupStatus } from './StatusBadge';

interface GroupedPermit {
    id: string;
    permits: UserUnifiedPermit[];
    totalValue: number;
    domains: number[];
    createdAt: string;
    allExecuted: boolean;
    allSigned: boolean;
}

function groupPermitsById(permits: UserUnifiedPermit[]): GroupedPermit[] {
    const grouped = new Map<string, GroupedPermit>();

    for (const permit of permits) {
        if (!grouped.has(permit.id)) {
            grouped.set(permit.id, {
                id: permit.id,
                permits: [],
                totalValue: 0,
                domains: [],
                createdAt: permit.createdAt,
                allExecuted: true,
                allSigned: true,
            });
        }
        const group = grouped.get(permit.id)!;
        group.permits.push(permit);
        group.totalValue += parseFloat(permit.value);
        group.domains.push(permit.domain);
        if (!permit.txHash) group.allExecuted = false;
        if (!permit.signature) group.allSigned = false;
    }

    return Array.from(grouped.values()).sort(
        (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
    );
}

export function PermitsListCard() {
    const { address, isConnected } = useWallet();
    const { permits, isLoading, error, refetch } = usePermits(address);

    const groupedPermits = useMemo(() => groupPermitsById(permits), [permits]);

    if (!isConnected) {
        return (
            <div className="list-card">
                <div className="list-header">Permits</div>
                <div className="list-placeholder">Connect wallet to view permits</div>
            </div>
        );
    }

    if (isLoading) {
        return (
            <div className="list-card">
                <div className="list-header">
                    Permits
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="list-card">
                <div className="list-header">
                    Permits
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder list-error">Error loading permits</div>
            </div>
        );
    }

    if (groupedPermits.length === 0) {
        return (
            <div className="list-card">
                <div className="list-header">
                    Permits
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder">No permits yet</div>
            </div>
        );
    }

    return (
        <div className="list-card">
            <div className="list-header">
                Permits
                <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                    </svg>
                </button>
            </div>
            <div className="list-table-wrapper">
                <table className="list-table">
                    <thead>
                        <tr>
                            <th>Domains</th>
                            <th>Total Value</th>
                            <th>Status</th>
                            <th>Executed At</th>
                        </tr>
                    </thead>
                    <tbody>
                        {groupedPermits.map((group) => {
                            const status = getPermitGroupStatus(group.allExecuted, group.allSigned);
                            const domainNames = group.domains
                                .map(d => DOMAIN_NAMES[d as CircleDomain] || `D${d}`)
                                .join(', ');
                            return (
                                <tr key={group.id}>
                                    <td>{domainNames}</td>
                                    <td>{formatUsdcValue(BigInt(group.totalValue))} USDC</td>
                                    <td>
                                        <StatusBadge label={status.label} variant={status.variant} />
                                    </td>
                                    <td>
                                        {group.permits[0]?.txHash ? (
                                            <a
                                                href={getTxExplorerUrl(group.permits[0].domain, group.permits[0].txHash)}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="explorer-link"
                                            >
                                                {formatDate(group.createdAt)}
                                            </a>
                                        ) : formatDate(group.createdAt)}
                                    </td>
                                </tr>
                            );
                        })}
                    </tbody>
                </table>
            </div>
        </div>
    );
}
