'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useDeposits } from '@/hooks/useDeposits';
import { DOMAIN_NAMES, CircleDomain } from '@/lib/types';

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

function getStatus(txHash: string, reason: string): { label: string; className: string } {
    if (txHash) return { label: 'Completed', className: 'status-completed' };
    if (reason) return { label: 'Failed', className: 'status-failed' };
    return { label: 'Pending', className: 'status-pending' };
}

export function DepositsListCard() {
    const { address, isConnected } = useWallet();
    const { deposits, isLoading, error } = useDeposits(address);

    if (!isConnected) {
        return (
            <div className="list-card">
                <div className="list-header">Deposits</div>
                <div className="list-placeholder">Connect wallet to view deposits</div>
            </div>
        );
    }

    if (isLoading) {
        return (
            <div className="list-card">
                <div className="list-header">Deposits</div>
                <div className="list-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="list-card">
                <div className="list-header">Deposits</div>
                <div className="list-placeholder list-error">Error loading deposits</div>
            </div>
        );
    }

    if (deposits.length === 0) {
        return (
            <div className="list-card">
                <div className="list-header">Deposits</div>
                <div className="list-placeholder">No deposits yet</div>
            </div>
        );
    }

    return (
        <div className="list-card">
            <div className="list-header">Deposits</div>
            <div className="list-table-wrapper">
                <table className="list-table">
                    <thead>
                        <tr>
                            <th>Domain</th>
                            <th>Value</th>
                            <th>Status</th>
                            <th>Created</th>
                        </tr>
                    </thead>
                    <tbody>
                        {deposits.map((deposit) => {
                            const status = getStatus(deposit.txHash, deposit.reason);
                            const domainName = DOMAIN_NAMES[deposit.destDomain as CircleDomain] || `Domain ${deposit.destDomain}`;
                            return (
                                <tr key={deposit.id}>
                                    <td>{domainName}</td>
                                    <td>{formatValue(deposit.value)} USDC</td>
                                    <td>
                                        <span className={`status-badge ${status.className}`}>
                                            {status.label}
                                        </span>
                                    </td>
                                    <td>{formatDate(deposit.createdAt)}</td>
                                </tr>
                            );
                        })}
                    </tbody>
                </table>
            </div>
        </div>
    );
}
