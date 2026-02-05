'use client';

import { useState, memo } from 'react';
import { useBridgeOperations } from '@/hooks/useBridgeOperations';
import { DOMAIN_NAMES, CircleDomain } from '@/lib/types';
import { formatUsdcValue, formatDate } from '@/lib/format';
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
    const { operations, total, isLoading, error } = useBridgeOperations(page, DEFAULT_LIMIT);

    const totalPages = Math.ceil(total / DEFAULT_LIMIT);
    const hasPrev = page > 0;
    const hasNext = page < totalPages - 1;

    if (isLoading) {
        return (
            <div className="list-card operations-card">
                <div className="list-header">Bridge Operations</div>
                <div className="list-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="list-card operations-card">
                <div className="list-header">Bridge Operations</div>
                <div className="list-placeholder list-error">Error loading operations</div>
            </div>
        );
    }

    if (operations.length === 0) {
        return (
            <div className="list-card operations-card">
                <div className="list-header">Bridge Operations</div>
                <div className="list-placeholder">No operations yet</div>
            </div>
        );
    }

    return (
        <div className="list-card operations-card">
            <div className="list-header">Bridge Operations</div>
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
                                    <td>{getDomainName(op.fromDomain)}</td>
                                    <td>{getDomainName(op.toDomain)}</td>
                                    <td>{formatUsdcValue(op.sendAmount)} USDC</td>
                                    <td>
                                        <StatusBadge label={status.label} variant={status.variant} />
                                    </td>
                                    <td>{formatDate(op.sentAt)}</td>
                                    <td>{formatDate(op.receivedAt)}</td>
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
