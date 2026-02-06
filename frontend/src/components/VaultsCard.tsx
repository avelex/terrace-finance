'use client';

import { memo } from 'react';
import { useVaultsInfo } from '@/hooks/useVaultsInfo';
import { DOMAIN_NAMES, CircleDomain } from '@/lib/types';
import { formatUsdcValue, shortenAddress } from '@/lib/format';
import { getAddressExplorerUrl } from '@/lib/constants';

function getDomainName(domain: number): string {
    return DOMAIN_NAMES[domain as CircleDomain] || `Domain ${domain}`;
}

export const VaultsCard = memo(function VaultsCard() {
    const { vaults, isLoading, error, refetch } = useVaultsInfo();

    if (isLoading) {
        return (
            <div className="list-card vaults-card">
                <div className="list-header">
                    Vaults
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="list-card vaults-card">
                <div className="list-header">
                    Vaults
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder list-error">Error loading vaults</div>
            </div>
        );
    }

    if (vaults.length === 0) {
        return (
            <div className="list-card vaults-card">
                <div className="list-header">
                    Vaults
                    <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                        </svg>
                    </button>
                </div>
                <div className="list-placeholder">No vaults found</div>
            </div>
        );
    }

    return (
        <div className="list-card vaults-card">
            <div className="list-header">
                Vaults
                <button className="refresh-button" onClick={() => refetch()} title="Refresh">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                    </svg>
                </button>
            </div>
            <div className="list-table-wrapper">
                <table className="list-table">
                    <thead>
                        <tr>
                            <th>Network</th>
                            <th>Address</th>
                            <th>Balance</th>
                        </tr>
                    </thead>
                    <tbody>
                        {vaults.map((vault) => (
                            <tr key={vault.domain}>
                                <td>{getDomainName(vault.domain)}</td>
                                <td>
                                    <a
                                        href={getAddressExplorerUrl(vault.domain, vault.address)}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="explorer-link"
                                    >
                                        {shortenAddress(vault.address)}
                                    </a>
                                </td>
                                <td>{formatUsdcValue(BigInt(vault.balance))} USDC</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
});
