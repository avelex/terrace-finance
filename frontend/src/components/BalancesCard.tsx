'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
import { useUnify } from '@/hooks/useUnify';
import { CircleDomain, DOMAIN_NAMES, DISPLAYED_DOMAINS, TokenBalancesByDomain } from '@/lib/types';

function formatBalance(value: number): string {
    return value.toLocaleString('en-US', {
        minimumFractionDigits: 0,
        maximumFractionDigits: 2
    });
}

interface BalanceListProps {
    title: string;
    balances: TokenBalancesByDomain | undefined;
}

function BalanceList({ title, balances }: BalanceListProps) {
    return (
        <div className="balance-list">
            <div className="balance-list-title">{title}:</div>
            <ul className="balance-list-items">
                {DISPLAYED_DOMAINS.map((domain) => (
                    <li key={domain} className="balance-list-item">
                        - {DOMAIN_NAMES[domain]}: {formatBalance(balances?.[domain] ?? 0)}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export function BalancesCard() {
    const { address, isConnected } = useWallet();
    const { balances, isLoading, error, refetch } = useBalances(address);
    const { isUnifying, currentStep, error: unifyError, unify } = useUnify();

    // Check if there are any USDC balances > 0
    const hasUsdcBalance = balances?.usdc &&
        Object.values(balances.usdc).some(balance => balance > 0);

    const handleUnify = async () => {
        if (!address || !balances?.usdc) return;

        const success = await unify(address, balances.usdc);
        if (success) {
            // Refetch balances after successful unify
            refetch();
        }
    };

    if (!isConnected) {
        return (
            <div className="balances-wrapper">
                <div className="balances-card">
                    <div className="balances-header">Balances</div>
                    <div className="balances-placeholder">
                        Connect wallet to view balances
                    </div>
                </div>
                <button className="unify-button" disabled>
                    Unify USDC
                </button>
            </div>
        );
    }

    if (isLoading) {
        return (
            <div className="balances-card">
                <div className="balances-header">Balances</div>
                <div className="balances-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="balances-card">
                <div className="balances-header">Balances</div>
                <div className="balances-placeholder balances-error">
                    Error loading balances
                </div>
            </div>
        );
    }

    return (
        <div className="balances-wrapper">
            <div className="balances-card">
                <div className="balances-header">Balances</div>

                <div className="balances-content">
                    <div className="balances-columns">
                        <BalanceList title="USDC" balances={balances?.usdc} />
                        <BalanceList title="Unified USDC" balances={balances?.unifiedUsdc} />
                    </div>

                    <div className="balances-totals">
                        <div className="balance-total-item">
                            Staking: {formatBalance(balances?.stakingBalance ?? 0)}
                        </div>
                        <div className="balance-total-item">
                            Stablecoin: {formatBalance(balances?.stablecoinBalance ?? 0)}
                        </div>
                    </div>
                </div>
            </div>

            <button
                className="unify-button"
                onClick={handleUnify}
                disabled={!hasUsdcBalance || isUnifying}
            >
                {isUnifying ? (currentStep || 'Processing...') : 'Unify USDC'}
            </button>

            {unifyError && (
                <div className="unify-error">{unifyError}</div>
            )}
        </div>
    );
}
