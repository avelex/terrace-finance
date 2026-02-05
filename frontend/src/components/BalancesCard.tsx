'use client';

import { useState } from 'react';
import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
import { useUnify } from '@/hooks/useUnify';
import { useDeposit } from '@/hooks/useDeposit';
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
    const { isUnifying, currentStep: unifyStep, error: unifyError, unify } = useUnify();
    const { isDepositing, currentStep: depositStep, error: depositError, deposit } = useDeposit();


    // Check if there are any USDC balances > 0
    const hasUsdcBalance = balances?.usdc &&
        Object.values(balances.usdc).some(balance => balance > 0);

    // Check if there are any unified USDC balances > 0
    const hasUnifiedBalance = balances?.unifiedUsdc &&
        Object.values(balances.unifiedUsdc).some(balance => balance > 0);

    const handleUnify = async () => {
        if (!address || !balances?.usdc) return;

        const id = await unify(address, balances.usdc);
        if (id) {
            refetch();
        }
    };

    const handleDeposit = async () => {
        if (!address || !balances?.unifiedUsdc) return;

        const success = await deposit(address, balances.unifiedUsdc);
        if (success) {
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
                <div className="action-buttons">
                    <button className="unify-button" disabled>
                        Unify USDC
                    </button>
                    <button className="deposit-button" disabled>
                        Deposit
                    </button>
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

    const currentStep = unifyStep || depositStep;
    const actionError = unifyError || depositError;

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

            <div className="action-buttons">
                <button
                    className="unify-button"
                    onClick={handleUnify}
                    disabled={!hasUsdcBalance || isUnifying || isDepositing}
                >
                    {isUnifying ? (unifyStep || 'Processing...') : 'Unify USDC'}
                </button>

                <button
                    className="deposit-button"
                    onClick={handleDeposit}
                    disabled={!hasUnifiedBalance || isDepositing || isUnifying}
                >
                    {isDepositing ? (depositStep || 'Processing...') : 'Deposit'}
                </button>
            </div>

            {actionError && (
                <div className="action-error">{actionError}</div>
            )}
        </div>
    );
}

