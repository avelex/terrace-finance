'use client';

import { useState } from 'react';
import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
import { formatBalance, formatUsdcValue } from '@/lib/format';
import { DISPLAYED_DOMAINS } from '@/lib/types';
import { BalancesModal } from './BalancesModal';

export function UsdcBalance() {
    const { address, isConnected } = useWallet();
    const { balances, isLoading } = useBalances(address);
    const [isModalOpen, setIsModalOpen] = useState(false);

    if (!isConnected || isLoading || !balances) {
        return null;
    }

    // Sum of unified USDC balances from all domains (values are already in human-readable)
    const totalUnifiedUsdc = Object.values(balances.unifiedUsdc).reduce(
        (sum, domain) => sum + (BigInt(domain) ?? 0),
        BigInt(0)
    );

    const formattedBalance = formatUsdcValue(totalUnifiedUsdc);

    return (
        <>
            <button
                className="balance-badge balance-badge-clickable"
                onClick={() => setIsModalOpen(true)}
            >
                <span className="balance-badge-label">USDC</span>
                <span className="balance-badge-divider">|</span>
                <span className="balance-badge-value">{formattedBalance}</span>
            </button>

            {isModalOpen && (
                <BalancesModal onClose={() => setIsModalOpen(false)} />
            )}
        </>
    );
}
