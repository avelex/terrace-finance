'use client';

import { useState } from 'react';
import { useWallet } from '@/contexts/WalletContext';

export function DepositCard() {
    const { isConnected } = useWallet();
    const [amount, setAmount] = useState('');

    const handleDeposit = () => {
        if (!isConnected || !amount) return;
        // TODO: Implement deposit logic
        console.log('Depositing:', amount, 'USDC');
    };

    return (
        <div className="deposit-card">
            <div className="apy-display">
                ~10% APY
            </div>

            <div className="input-row">
                <span className="token-label">USDC</span>
                <input
                    type="number"
                    className="amount-input"
                    placeholder="0"
                    value={amount}
                    onChange={(e) => setAmount(e.target.value)}
                    min="0"
                    step="any"
                />
            </div>

            <button
                className="deposit-button"
                onClick={handleDeposit}
                disabled={!isConnected || !amount}
            >
                Deposit
            </button>
        </div>
    );
}
