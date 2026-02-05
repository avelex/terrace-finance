'use client';

import { useWallet } from '@/contexts/WalletContext';
import { shortenAddress } from '@/lib/format';

export function Header() {
    const { address, isConnected, isConnecting, connect, disconnect } = useWallet();

    return (
        <header className="header">
            <div className="header-logo">Terrace Finance</div>

            {isConnected ? (
                <button
                    className="connect-button connected"
                    onClick={disconnect}
                >
                    {shortenAddress(address!)}
                </button>
            ) : (
                <button
                    className="connect-button"
                    onClick={connect}
                    disabled={isConnecting}
                >
                    {isConnecting ? 'Connecting...' : 'Connect Wallet'}
                </button>
            )}
        </header>
    );
}
