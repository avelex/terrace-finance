'use client';

import { useState, memo } from 'react';
import { useBridgeOperations } from '@/hooks/useBridgeOperations';
import { DOMAIN_NAMES, CircleDomain } from '@/lib/types';
import { formatUsdcValue, formatDate, shortenAddress } from '@/lib/format';
import { getAddressExplorerUrl, getTxExplorerUrl } from '@/lib/constants';
import { StatusBadge } from './StatusBadge';

const DEFAULT_LIMIT = 10;

function getBridgeStatus(receivedTxHash: string): { label: string; variant: 'completed' | 'pending' } {
    return receivedTxHash
        ? { label: 'Completed', variant: 'completed' }
        : { label: 'Pending', variant: 'pending' };
}

function getDomainName(domain: number): string {
    return DOMAIN_NAMES[domain as CircleDomain] || `Domain ${domain}`;
}

export const OperationsTable = memo(function OperationsTable() {
    const [page, setPage] = useState(0);
    const { operations, total, isLoading, error, refetch } = useBridgeOperations(page, DEFAULT_LIMIT);

    const totalPages = Math.ceil(total / DEFAULT_LIMIT);
    const hasPrev = page > 0;
    const hasNext = page < totalPages - 1;

    if (isLoading) {
        return (
            <div className="list-card operations-card">
                <div className="list-header">
                    Bridge Operations
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
            <div className="list-card operations-card">
                <div className="list-header">
                    Bridge Operations
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder list-error">Error loading operations</div>
            </div>
        );
    }

    if (operations.length === 0) {
        return (
            <div className="list-card operations-card">
                <div className="list-header">
                    Bridge Operations
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder">No operations yet</div>
            </div>
        );
    }

    return (
        <div className="list-card operations-card">
            <div className="list-header">
                Bridge Operations
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
                            <th>From</th>
                            <th>To</th>
                            <th>Amount</th>
                            <th>Status</th>
                            <th>Sent At</th>
                            <th>Received At</th>
                        </tr>
                    </thead>
                    <tbody>
                        {operations.map((op, index) => {
                            const status = getBridgeStatus(op.receivedTxHash);
                            return (
                                <tr key={op.id || `op-${index}`}>
                                    <td>
                                        <a
                                            href={getAddressExplorerUrl(op.fromDomain, op.fromTerrace)}
                                            target="_blank"
                                            rel="noopener noreferrer"
                                            className="explorer-link"
                                        >
                                            {getDomainName(op.fromDomain)}
                                        </a>
                                    </td>
                                    <td>
                                        <a
                                            href={getAddressExplorerUrl(op.toDomain, op.toTerrace)}
                                            target="_blank"
                                            rel="noopener noreferrer"
                                            className="explorer-link"
                                        >
                                            {getDomainName(op.toDomain)}
                                        </a>
                                    </td>
                                    <td>{formatUsdcValue(BigInt(op.sendAmount))} USDC</td>
                                    <td>
                                        <StatusBadge label={status.label} variant={status.variant} />
                                    </td>
                                    <td>
                                        {op.sentTxHash ? (
                                            <a
                                                href={getTxExplorerUrl(op.fromDomain, op.sentTxHash)}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="explorer-link"
                                            >
                                                {formatDate(op.sentAt)}
                                            </a>
                                        ) : formatDate(op.sentAt)}
                                    </td>
                                    <td>
                                        {op.receivedTxHash ? (
                                            <a
                                                href={getTxExplorerUrl(op.toDomain, op.receivedTxHash)}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="explorer-link"
                                            >
                                                {formatDate(op.receivedAt)}
                                            </a>
                                        ) : formatDate(op.receivedAt)}
                                    </td>
                                </tr>
                            );
                        })}
                    </tbody>
                </table>
            </div>
            <div className="pagination">
                <button
                    className="pagination-button"
                    onClick={() => setPage(p => p - 1)}
                    disabled={!hasPrev}
                >
                    Previous
                </button>
                <span className="pagination-info">
                    Page {page + 1} of {totalPages || 1}
                </span>
                <button
                    className="pagination-button"
                    onClick={() => setPage(p => p + 1)}
                    disabled={!hasNext}
                >
                    Next
                </button>
            </div>
        </div>
    );
});
