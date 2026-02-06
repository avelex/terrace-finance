'use client';

import { useState, memo } from 'react';
import { DOMAIN_NAMES, CircleDomain, DISPLAYED_DOMAINS } from '@/lib/types';
import { API_BASE_URL } from '@/lib/constants';

function getDomainName(domain: number): string {
    return DOMAIN_NAMES[domain as CircleDomain] || `Domain ${domain}`;
}

export const BridgeCard = memo(function BridgeCard() {
    const [srcDomain, setSrcDomain] = useState<number>(DISPLAYED_DOMAINS[0]);
    const [dstDomain, setDstDomain] = useState<number>(DISPLAYED_DOMAINS[1]);
    const [isLoading, setIsLoading] = useState(false);
    const [result, setResult] = useState<{ txHash?: string; error?: string } | null>(null);

    const handleBridge = async () => {
        if (srcDomain === dstDomain) {
            setResult({ error: 'Source and destination must be different' });
            return;
        }

        setIsLoading(true);
        setResult(null);

        try {
            const response = await fetch(`${API_BASE_URL}/api/bridge/send/${srcDomain}/${dstDomain}`);
            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.message || `Bridge failed: ${response.status}`);
            }

            setResult({ txHash: data.txHash });
        } catch (err) {
            setResult({ error: err instanceof Error ? err.message : 'Bridge failed' });
        } finally {
            setIsLoading(false);
        }
    };

    const swapDomains = () => {
        setSrcDomain(dstDomain);
        setDstDomain(srcDomain);
        setResult(null);
    };

    return (
        <div className="list-card bridge-card">
            <div className="list-header">Bridge</div>

            <div className="bridge-form">
                <div className="bridge-row">
                    <span className="bridge-row-label">From</span>
                    <select
                        className="bridge-select"
                        value={srcDomain}
                        onChange={(e) => {
                            setSrcDomain(Number(e.target.value));
                            setResult(null);
                        }}
                        disabled={isLoading}
                    >
                        {DISPLAYED_DOMAINS.map((domain) => (
                            <option key={domain} value={domain}>
                                {getDomainName(domain)}
                            </option>
                        ))}
                    </select>
                </div>

                <button
                    className="bridge-swap-button"
                    onClick={swapDomains}
                    disabled={isLoading}
                    type="button"
                    title="Swap"
                >
                    ⇅
                </button>

                <div className="bridge-row">
                    <span className="bridge-row-label">To</span>
                    <select
                        className="bridge-select"
                        value={dstDomain}
                        onChange={(e) => {
                            setDstDomain(Number(e.target.value));
                            setResult(null);
                        }}
                        disabled={isLoading}
                    >
                        {DISPLAYED_DOMAINS.map((domain) => (
                            <option key={domain} value={domain}>
                                {getDomainName(domain)}
                            </option>
                        ))}
                    </select>
                </div>

                <button
                    className="bridge-button"
                    onClick={handleBridge}
                    disabled={isLoading || srcDomain === dstDomain}
                >
                    {isLoading ? 'Bridging...' : 'Bridge'}
                </button>

                {result?.error && (
                    <div className="bridge-error">{result.error}</div>
                )}
                {result?.txHash && (
                    <div className="bridge-success">
                        Tx: {result.txHash.slice(0, 10)}...{result.txHash.slice(-8)}
                    </div>
                )}
            </div>
        </div>
    );
});
