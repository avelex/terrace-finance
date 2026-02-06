'use client';

import Image from 'next/image';
import { useState, useMemo } from 'react';
import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
import { useDeposit } from '@/hooks/useDeposit';
import { formatUsdcValue, bigintToHumanReadable } from '@/lib/format';
import { DISPLAYED_DOMAINS, TokenBalancesByDomain } from '@/lib/types';

export function DepositCard() {
    const { address, isConnected } = useWallet();
    const { balances, refetch } = useBalances(address);
    const { isDepositing, currentStep, error: depositError, deposit } = useDeposit();
    const [amount, setAmount] = useState('');

    // Sum of unified USDC balances from all domains (raw values as BigInt)
    const totalUnifiedUsdc = useMemo(() => {
        return Object.values(balances?.unifiedUsdc ?? {}).reduce(
            (sum, domain) => sum + BigInt(domain ?? 0),
            BigInt(0)
        );
    }, [balances?.unifiedUsdc]);

    // Parse user input amount (user enters human-readable, convert to BigInt with 6 decimals)
    const parsedAmount = useMemo(() => {
        const num = parseFloat(amount);
        if (isNaN(num) || num <= 0) return BigInt(0);
        // Convert to 6 decimals as BigInt
        return BigInt(Math.floor(num * 1e6));
    }, [amount]);

    // Validation using BigInt comparisons
    const isAmountValid = parsedAmount > BigInt(0) && parsedAmount <= totalUnifiedUsdc;
    const exceedsMax = parsedAmount > totalUnifiedUsdc && parsedAmount > BigInt(0);

    // Create filtered balances based on user amount (distribute across domains)
    const getDepositBalances = (): TokenBalancesByDomain => {
        const result: TokenBalancesByDomain = {};
        let remaining = parsedAmount;

        // Distribute across domains in order until amount is fulfilled
        for (const domain of DISPLAYED_DOMAINS) {
            const domainBalance = BigInt(balances?.unifiedUsdc?.[domain] ?? 0);
            if (domainBalance <= BigInt(0) || remaining <= BigInt(0)) continue;

            const toDeposit = domainBalance < remaining ? domainBalance : remaining;
            result[domain] = Number(toDeposit);
            remaining -= toDeposit;
        }

        return result;
    };

    const handleDeposit = async () => {
        if (!isConnected || !address || !isAmountValid) return;

        const depositBalances = getDepositBalances();
        const success = await deposit(address, depositBalances);

        if (success) {
            setAmount('');
            refetch();
        }
    };

    const handleMaxClick = () => {
        // Set max amount (convert from BigInt to human-readable string)
        const maxHumanReadable = bigintToHumanReadable(totalUnifiedUsdc, 6);
        setAmount(maxHumanReadable);
    };

    return (
        <div className="deposit-card">
            <div className="input-row">
                <span className="token-label">
                    <Image src="/usdc-logo.png" alt="USDC" width={24} height={24} />
                    USDC
                </span>
                <input
                    type="text"
                    inputMode="decimal"
                    className="amount-input"
                    placeholder="0.00"
                    value={amount}
                    onChange={(e) => {
                        // Allow only numbers and dots, convert commas to dots
                        const value = e.target.value.replace(/,/g, '.');
                        if (value === '' || /^\d*\.?\d*$/.test(value)) {
                            setAmount(value);
                        }
                    }}
                    disabled={isDepositing}
                    pattern="^[0-9]*[.,]?[0-9]*$"
                />
            </div>

            {isConnected && (
                <div className="balance-info">
                    <span>Available: {formatUsdcValue(totalUnifiedUsdc)} USDC</span>
                    <button
                        className="max-button"
                        onClick={handleMaxClick}
                        disabled={isDepositing || totalUnifiedUsdc <= BigInt(0)}
                    >
                        MAX
                    </button>
                </div>
            )}

            {exceedsMax && (
                <div className="deposit-error">
                    Amount exceeds available balance
                </div>
            )}

            {depositError && (
                <div className="deposit-error">
                    {depositError}
                </div>
            )}

            <button
                className="deposit-button"
                onClick={handleDeposit}
                disabled={!isConnected || !isAmountValid || isDepositing}
            >
                {isDepositing ? (currentStep || 'PROCESSING...') : 'DEPOSIT'}
            </button>
        </div>
    );
}



