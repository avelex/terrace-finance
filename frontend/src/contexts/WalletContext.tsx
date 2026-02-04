'use client';

import React, { createContext, useContext, useState, useEffect, useCallback, ReactNode } from 'react';
import { requestAccounts, getAccounts } from '@/lib/wallet';

interface WalletContextType {
    address: `0x${string}` | null;
    isConnected: boolean;
    isConnecting: boolean;
    connect: () => Promise<void>;
    disconnect: () => void;
}

const WalletContext = createContext<WalletContextType | undefined>(undefined);

export function WalletProvider({ children }: { children: ReactNode }) {
    const [address, setAddress] = useState<`0x${string}` | null>(null);
    const [isConnecting, setIsConnecting] = useState(false);

    const isConnected = address !== null;

    // Check for existing connection on mount
    useEffect(() => {
        const checkConnection = async () => {
            try {
                const accounts = await getAccounts();
                if (accounts.length > 0) {
                    setAddress(accounts[0]);
                }
            } catch (error) {
                console.error('Error checking wallet connection:', error);
            }
        };

        checkConnection();
    }, []);

    // Listen for account changes
    useEffect(() => {
        if (typeof window === 'undefined' || !window.ethereum) return;

        const handleAccountsChanged = (accounts: unknown) => {
            const typedAccounts = accounts as `0x${string}`[];
            if (typedAccounts.length === 0) {
                setAddress(null);
            } else {
                setAddress(typedAccounts[0]);
            }
        };

        window.ethereum.on('accountsChanged', handleAccountsChanged);

        return () => {
            window.ethereum?.removeListener('accountsChanged', handleAccountsChanged);
        };
    }, []);

    const connect = useCallback(async () => {
        setIsConnecting(true);
        try {
            const accounts = await requestAccounts();
            if (accounts.length > 0) {
                setAddress(accounts[0]);
            }
        } catch (error) {
            console.error('Error connecting wallet:', error);
        } finally {
            setIsConnecting(false);
        }
    }, []);

    const disconnect = useCallback(() => {
        setAddress(null);
    }, []);

    return (
        <WalletContext.Provider
            value={{
                address,
                isConnected,
                isConnecting,
                connect,
                disconnect,
            }}
        >
            {children}
        </WalletContext.Provider>
    );
}

export function useWallet() {
    const context = useContext(WalletContext);
    if (context === undefined) {
        throw new Error('useWallet must be used within a WalletProvider');
    }
    return context;
}
