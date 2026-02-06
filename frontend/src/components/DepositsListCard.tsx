'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useDeposits } from '@/hooks/useDeposits';
import { DOMAIN_NAMES, CircleDomain } from '@/lib/types';
import { formatUsdcValue, formatDate } from '@/lib/format';
import { StatusBadge, getDepositStatus } from './StatusBadge';

export function DepositsListCard() {
    const { address, isConnected } = useWallet();
    const { deposits, isLoading, error, refetch } = useDeposits(address);

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
                <div className="list-header">
                    Deposits
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
                    Deposits
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder list-error">Error loading deposits</div>
            </div>
        );
    }

    if (deposits.length === 0) {
        return (
            <div className="list-card">
                <div className="list-header">
                    Deposits
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder">No deposits yet</div>
            </div>
        );
    }

    return (
        <div className="list-card">
            <div className="list-header">
                Deposits
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
                                    <td>{formatUsdcValue(BigInt(deposit.value))} USDC</td>
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
