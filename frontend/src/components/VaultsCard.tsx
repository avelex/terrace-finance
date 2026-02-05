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
    const { vaults, isLoading, error } = useVaultsInfo();

    if (isLoading) {
        return (
            <div className="list-card vaults-card">
                <div className="list-header">Vaults</div>
                <div className="list-placeholder">Loading...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="list-card vaults-card">
                <div className="list-header">Vaults</div>
                <div className="list-placeholder list-error">Error loading vaults</div>
            </div>
        );
    }

    if (vaults.length === 0) {
        return (
            <div className="list-card vaults-card">
                <div className="list-header">Vaults</div>
                <div className="list-placeholder">No vaults found</div>
            </div>
        );
    }

    return (
        <div className="list-card vaults-card">
            <div className="list-header">Vaults</div>
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
                                <td>{formatUsdcValue(vault.balance)} USDC</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
});
