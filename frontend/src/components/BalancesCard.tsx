'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
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
    const { balances, isLoading, error } = useBalances(address);

    if (!isConnected) {
        return (
            <div className="balances-card">
                <div className="balances-header">Balances</div>
                <div className="balances-placeholder">
                    Connect wallet to view balances
                </div>
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
    );
}
