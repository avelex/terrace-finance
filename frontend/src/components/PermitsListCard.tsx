'use client';

import { useMemo } from 'react';
import { useWallet } from '@/contexts/WalletContext';
import { usePermits } from '@/hooks/usePermits';
import { DOMAIN_NAMES, CircleDomain, UserUnifiedPermit } from '@/lib/types';

function formatValue(value: string): string {
    const num = parseFloat(value) / 1e6; // USDC has 6 decimals
    return num.toLocaleString('en-US', {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2
    });
}

function formatDate(dateStr: string): string {
    if (!dateStr || dateStr === '0001-01-01T00:00:00Z') return '-';
    return new Date(dateStr).toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
}

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

function getGroupStatus(group: GroupedPermit): { label: string; className: string } {
    if (group.allExecuted) return { label: 'Executed', className: 'status-completed' };
    if (group.allSigned) return { label: 'Signed', className: 'status-signed' };
    return { label: 'Pending', className: 'status-pending' };
}

export function PermitsListCard() {
    const { address, isConnected } = useWallet();
    const { permits, isLoading, error } = usePermits(address);

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
                <div className="list-header">Permits</div>
                <div className="list-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="list-card">
                <div className="list-header">Permits</div>
                <div className="list-placeholder list-error">Error loading permits</div>
            </div>
        );
    }

    if (groupedPermits.length === 0) {
        return (
            <div className="list-card">
                <div className="list-header">Permits</div>
                <div className="list-placeholder">No permits yet</div>
            </div>
        );
    }

    return (
        <div className="list-card">
            <div className="list-header">Permits</div>
            <div className="list-table-wrapper">
                <table className="list-table">
                    <thead>
                        <tr>
                            <th>Domains</th>
                            <th>Total Value</th>
                            <th>Status</th>
                            <th>Created</th>
                        </tr>
                    </thead>
                    <tbody>
                        {groupedPermits.map((group) => {
                            const status = getGroupStatus(group);
                            const domainNames = group.domains
                                .map(d => DOMAIN_NAMES[d as CircleDomain] || `D${d}`)
                                .join(', ');
                            return (
                                <tr key={group.id}>
                                    <td>{domainNames}</td>
                                    <td>{formatValue(group.totalValue.toString())} USDC</td>
                                    <td>
                                        <span className={`status-badge ${status.className}`}>
                                            {status.label}
                                        </span>
                                    </td>
                                    <td>{formatDate(group.createdAt)}</td>
                                </tr>
                            );
                        })}
                    </tbody>
                </table>
            </div>
        </div>
    );
}
