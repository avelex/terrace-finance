'use client';

import Link from 'next/link';
import { useWallet } from '@/contexts/WalletContext';
import { shortenAddress } from '@/lib/format';

export function Header() {
    const { address, isConnected, isConnecting, connect, disconnect } = useWallet();

    return (
        <header className="header">
            <div className="header-left">
                <Link href="/" className="header-logo">Terrace Finance</Link>
                <nav className="header-nav">
                    <Link href="/explorer" className="nav-link">Explorer</Link>
                </nav>
            </div>

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
