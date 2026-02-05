import { memo } from 'react';

interface StatusBadgeProps {
    label: string;
    variant: 'completed' | 'signed' | 'pending' | 'failed';
}

/**
 * Reusable status badge component for displaying transaction/permit status
 */
export const StatusBadge = memo(function StatusBadge({ label, variant }: StatusBadgeProps) {
    const className = `status-badge status-${variant}`;
    return <span className={className}>{label}</span>;
});

/**
 * Determine deposit status from txHash and reason
 */
export function getDepositStatus(txHash: string, reason: string): { label: string; variant: StatusBadgeProps['variant'] } {
    if (txHash) return { label: 'Completed', variant: 'completed' };
    if (reason) return { label: 'Failed', variant: 'failed' };
    return { label: 'Pending', variant: 'pending' };
}

/**
 * Determine permit group status
 */
export function getPermitGroupStatus(allExecuted: boolean, allSigned: boolean): { label: string; variant: StatusBadgeProps['variant'] } {
    if (allExecuted) return { label: 'Executed', variant: 'completed' };
    if (allSigned) return { label: 'Signed', variant: 'signed' };
    return { label: 'Pending', variant: 'pending' };
}
