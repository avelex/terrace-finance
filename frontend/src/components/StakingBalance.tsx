'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
import { formatBalance, formatStakingBalance } from '@/lib/format';

export function StakingBalance() {
    const { address, isConnected } = useWallet();
    const { balances, isLoading } = useBalances(address);

    // stakingBalance is already in human-readable format from backend
    const formattedBalance = formatStakingBalance(balances?.stakingBalance ?? 0);

    return (
        <div className="balance-badge">
            <span className="balance-badge-label">Staking</span>
            <span className="balance-badge-divider">|</span>
            <span className="balance-badge-value">{formattedBalance}</span>
        </div>
    );
}
