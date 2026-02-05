'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useDeposits } from '@/hooks/useDeposits';
import { DOMAIN_NAMES, CircleDomain } from '@/lib/types';
import { formatUsdcValue, formatDate } from '@/lib/format';
import { StatusBadge, getDepositStatus } from './StatusBadge';

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
                            const status = getDepositStatus(deposit.txHash, deposit.reason);
                            const domainName = DOMAIN_NAMES[deposit.destDomain as CircleDomain] || `Domain ${deposit.destDomain}`;
                            return (
                                <tr key={deposit.id}>
                                    <td>{domainName}</td>
                                    <td>{formatUsdcValue(deposit.value)} USDC</td>
                                    <td>
                                        <StatusBadge label={status.label} variant={status.variant} />
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
