'use client';

import { useWallet } from '@/contexts/WalletContext';
import { useBalances } from '@/hooks/useBalances';
import { useUnify } from '@/hooks/useUnify';
import { DOMAIN_NAMES, DISPLAYED_DOMAINS, TokenBalancesByDomain } from '@/lib/types';
import { formatBalance, formatUsdcValue } from '@/lib/format';

interface BalanceListProps {
    title: string;
    balances: TokenBalancesByDomain | undefined;
}

function BalanceList({ title, balances }: BalanceListProps) {
    return (
        <div className="modal-balance-list">
            <div className="modal-balance-title">{title}:</div>
            <ul className="modal-balance-items">
                {DISPLAYED_DOMAINS.map((domain) => (
                    <li key={domain} className="modal-balance-item">
                        - {DOMAIN_NAMES[domain]}: {formatUsdcValue(BigInt(balances?.[domain] ?? 0))}
                    </li>
                ))}
            </ul>
        </div>
    );
}

interface BalancesModalProps {
    onClose: () => void;
}

export function BalancesModal({ onClose }: BalancesModalProps) {
    const { address } = useWallet();
    const { balances, refetch } = useBalances(address);
    const { isUnifying, currentStep: unifyStep, error: unifyError, unify } = useUnify();

    const hasUsdcBalance = balances?.usdc &&
        Object.values(balances.usdc).some(balance => balance > 0);

    const handleUnify = async () => {
        if (!address || !balances?.usdc) return;

        const id = await unify(address, balances.usdc);
        if (id) {
            refetch();
        }
    };

    const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
        if (e.target === e.currentTarget) {
            onClose();
        }
    };

    return (
        <div className="modal-overlay" onClick={handleOverlayClick}>
            <div className="modal-content">
                <button className="modal-close" onClick={onClose}>×</button>

                <div className="modal-header">Balances</div>

                <div className="modal-balances-columns">
                    <BalanceList title="USDC" balances={balances?.usdc} />
                    <BalanceList title="Unified USDC" balances={balances?.unifiedUsdc} />
                </div>

                <button
                    className="modal-unify-button"
                    onClick={handleUnify}
                    disabled={!hasUsdcBalance || isUnifying}
                >
                    {isUnifying ? (unifyStep || 'Processing...') : 'Unify USDC'}
                </button>

                {unifyError && (
                    <div className="modal-error">{unifyError}</div>
                )}
            </div>
        </div>
    );
}
