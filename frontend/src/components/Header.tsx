'use client';

import Link from 'next/link';
import { useWallet } from '@/contexts/WalletContext';
import { shortenAddress } from '@/lib/format';
import { StakingBalance } from './StakingBalance';
import { UsdcBalance } from './UsdcBalance';

export function Header() {
    const { address, isConnected, isConnecting, connect, disconnect } = useWallet();

    return (
        <header className="header">
            <div className="header-left">
                <Link href="/" className="header-logo">TERRACE</Link>
            </div>

            <nav className="header-nav">
                <Link href="/explorer" className="nav-link">EXPLORER</Link>
                <Link href="/dashboard" className="nav-link">DASHBOARD</Link>
                <Link href="/about" className="nav-link">ABOUT</Link>
            </nav>

            <div className="header-right">
                <StakingBalance />
                <UsdcBalance />

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
                        {isConnecting ? 'CONNECTING...' : 'CONNECT WALLET'}
                    </button>
                )}
            </div>
        </header>
    );
}

