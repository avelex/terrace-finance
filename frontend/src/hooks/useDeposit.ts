'use client';

import { useState, useCallback } from 'react';
import { pad, maxUint256, Hex } from 'viem';
import { TokenBalancesByDomain } from '@/lib/types';
import { createWalletClientFromProvider } from '@/lib/wallet';
import {
  API_BASE_URL,
  GATEWAY_TRANSFER_API_URL,
  DESTINATION_DOMAIN,
  DESTINATION_RECIPIENT,
  DESTINATION_CALLER,
  GATEWAY_MINTER_ADDRESS,
  GATEWAY_WALLET_ADDRESS,
  USDC_ADDRESSES,
  DOMAIN_TO_CHAIN,
  GAS_FEES,
  MAX_FEE,
} from '@/lib/constants';

// EIP-712 Domain for Gateway burn intents
const domain = {
  name: 'GatewayWallet',
  version: '1',
};

// EIP-712 Types
const TransferSpec = [
  { name: 'version', type: 'uint32' },
  { name: 'sourceDomain', type: 'uint32' },
  { name: 'destinationDomain', type: 'uint32' },
  { name: 'sourceContract', type: 'bytes32' },
  { name: 'destinationContract', type: 'bytes32' },
  { name: 'sourceToken', type: 'bytes32' },
  { name: 'destinationToken', type: 'bytes32' },
  { name: 'sourceDepositor', type: 'bytes32' },
  { name: 'destinationRecipient', type: 'bytes32' },
  { name: 'sourceSigner', type: 'bytes32' },
  { name: 'destinationCaller', type: 'bytes32' },
  { name: 'value', type: 'uint256' },
  { name: 'salt', type: 'bytes32' },
  { name: 'hookData', type: 'bytes' },
] as const;

const BurnIntent = [
  { name: 'maxBlockHeight', type: 'uint256' },
  { name: 'maxFee', type: 'uint256' },
  { name: 'spec', type: 'TransferSpec' },
] as const;

// Convert address to bytes32
function addressToBytes32(address: string): Hex {
  return pad(address.toLowerCase() as Hex, { size: 32 });
}

// Generate random bytes32 salt
function randomSalt(): Hex {
  const bytes = new Uint8Array(32);
  crypto.getRandomValues(bytes);
  return ('0x' + Array.from(bytes).map(b => b.toString(16).padStart(2, '0')).join('')) as Hex;
}

/**
 * Calculate Circle Gateway fee for transfer
 * Formula: maxFee ≥ gas fee + (transfer amount * 0.00005)
 * @see https://developers.circle.com/gateway/references/fees
 */
function calculateGatewayFee(
  amount: bigint,
  sourceDomain: number,
  destinationDomain: number
): bigint {
  const gasFee = GAS_FEES[sourceDomain] ?? BigInt(10000);

  // Transfer fee: 0.005% (only for crosschain transfers)
  const isCrosschain = sourceDomain !== destinationDomain;
  const transferFee = isCrosschain
    ? amount * BigInt(5) / BigInt(100000)
    : BigInt(0);

  return gasFee + transferFee;
}

// Create burn intent for a source domain
function createBurnIntent(
  sourceDomain: number,
  amount: bigint,
  depositorAddress: `0x${string}`
) {
  const sourceToken = USDC_ADDRESSES[sourceDomain];
  const destToken = USDC_ADDRESSES[DESTINATION_DOMAIN];

  return {
    maxBlockHeight: maxUint256,
    maxFee: MAX_FEE,
    spec: {
      version: 1,
      sourceDomain: sourceDomain,
      destinationDomain: DESTINATION_DOMAIN,
      sourceContract: GATEWAY_WALLET_ADDRESS,
      destinationContract: GATEWAY_MINTER_ADDRESS,
      sourceToken: sourceToken,
      destinationToken: destToken,
      sourceDepositor: depositorAddress,
      destinationRecipient: DESTINATION_RECIPIENT,
      sourceSigner: depositorAddress,
      destinationCaller: DESTINATION_CALLER,
      value: amount,
      salt: randomSalt(),
      hookData: '0x' as Hex,
    },
  };
}

// Create EIP-712 typed data for signing
function burnIntentTypedData(burnIntent: ReturnType<typeof createBurnIntent>) {
  return {
    types: {
      TransferSpec,
      BurnIntent,
    },
    domain,
    primaryType: 'BurnIntent' as const,
    message: {
      ...burnIntent,
      spec: {
        ...burnIntent.spec,
        sourceContract: addressToBytes32(burnIntent.spec.sourceContract),
        destinationContract: addressToBytes32(burnIntent.spec.destinationContract),
        sourceToken: addressToBytes32(burnIntent.spec.sourceToken),
        destinationToken: addressToBytes32(burnIntent.spec.destinationToken),
        sourceDepositor: addressToBytes32(burnIntent.spec.sourceDepositor),
        destinationRecipient: addressToBytes32(burnIntent.spec.destinationRecipient),
        sourceSigner: addressToBytes32(burnIntent.spec.sourceSigner),
        destinationCaller: addressToBytes32(burnIntent.spec.destinationCaller),
      },
    },
  };
}

interface UseDepositResult {
  isDepositing: boolean;
  currentStep: string | null;
  error: string | null;
  deposit: (address: `0x${string}`, unifiedBalances: TokenBalancesByDomain) => Promise<boolean>;
}

export function useDeposit(): UseDepositResult {
  const [isDepositing, setIsDepositing] = useState(false);
  const [currentStep, setCurrentStep] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const deposit = useCallback(async (
    address: `0x${string}`,
    unifiedBalances: TokenBalancesByDomain,
  ): Promise<boolean> => {
    setIsDepositing(true);
    setError(null);
    setCurrentStep('Initializing deposit...');

    try {
      // Get domains with balance > 0
      const domainsWithBalance = Object.entries(unifiedBalances)
        .filter(([, balance]) => balance > 0)
        .map(([domain, balance]) => ({
          domain: Number(domain),
          balance: BigInt(balance),
        }));

      if (domainsWithBalance.length === 0) {
        setError('No unified USDC balance to deposit');
        return false;
      }

      setCurrentStep(`Creating burn intents for ${domainsWithBalance.length} chains...`);

      const walletClient = await createWalletClientFromProvider();
      const [account] = await walletClient.getAddresses();

      // Create and sign burn intents
      const requests: { burnIntent: ReturnType<typeof burnIntentTypedData>['message']; signature: string }[] = [];

      let totalAmount = BigInt(0);

      for (let i = 0; i < domainsWithBalance.length; i++) {
        const { domain: sourceDomain, balance } = domainsWithBalance[i];

        const chain = DOMAIN_TO_CHAIN[sourceDomain];
        if (!chain) {
          console.warn(`Unknown chain for domain ${sourceDomain}, skipping`);
          continue;
        }

        // Switch to the correct chain
        const chainIdHex = `0x${chain.id.toString(16)}`;
        setCurrentStep(`Switching to ${chain.name}...`);

        try {
          await window.ethereum?.request({
            method: 'wallet_switchEthereumChain',
            params: [{ chainId: chainIdHex }],
          });
        } catch (switchError: unknown) {
          const err = switchError as { code?: number; message?: string };
          if (err.code === 4902 || err.message?.toLowerCase().includes('unrecognized chain')) {
            // Add chain
            await window.ethereum?.request({
              method: 'wallet_addEthereumChain',
              params: [{
                chainId: chainIdHex,
                chainName: chain.name,
                nativeCurrency: chain.nativeCurrency,
                rpcUrls: chain.rpcUrls.default.http,
                blockExplorerUrls: chain.blockExplorers?.default ? [chain.blockExplorers.default.url] : undefined,
              }],
            });
          } else {
            throw switchError;
          }
        }

        // Calculate fee based on source and destination domains
        const fee = calculateGatewayFee(balance, sourceDomain, DESTINATION_DOMAIN);
        const balanceAfterFee = balance - fee;

        // Create burn intent
        setCurrentStep(`Creating burn intent ${i + 1}/${domainsWithBalance.length}...`);
        const intent = createBurnIntent(sourceDomain, balanceAfterFee, account);
        const typedData = burnIntentTypedData(intent);

        setCurrentStep(`Signing burn intent ${i + 1}/${domainsWithBalance.length}...`);
        const newWalletClient = await createWalletClientFromProvider();

        const signature = await newWalletClient.signTypedData({
          account,
          domain: typedData.domain,
          types: typedData.types,
          primaryType: typedData.primaryType,
          message: typedData.message,
        });

        requests.push({ burnIntent: typedData.message, signature })
        totalAmount += balanceAfterFee;
      }

      if (requests.length === 0) {
        setError('No valid burn intents created');
        return false;
      }

      // Send to Circle Gateway API
      setCurrentStep('Sending to Circle Gateway...');

      const response = await fetch(GATEWAY_TRANSFER_API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(requests, (_key, value) =>
          typeof value === 'bigint' ? value.toString() : value
        ),
      });

      if (!response.ok) {
        const text = await response.text();
        throw new Error(`Gateway API error: ${response.status} ${text}`);
      }

      const json = await response.json();
      console.log('Gateway API response:', json);

      const attestation = json?.attestation;
      const operatorSig = json?.signature;

      if (!attestation || !operatorSig) {
        throw new Error('Missing attestation or signature in Gateway response');
      }

      // Send attestation to our backend
      setCurrentStep('Submitting to backend...');

      const backendResponse = await fetch(`${API_BASE_URL}/api/wallet/${address}/deposit_stake`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          amount: totalAmount.toString(),
          attestation,
          signature: operatorSig,
        }),
      });

      if (!backendResponse.ok) {
        throw new Error(`Backend error: ${backendResponse.status}`);
      }

      setCurrentStep('Deposit complete!');
      return true;
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Unknown error';
      setError(message);
      console.error('Deposit error:', err);
      return false;
    } finally {
      setIsDepositing(false);
      setTimeout(() => setCurrentStep(null), 2000);
    }
  }, []);

  return {
    isDepositing,
    currentStep,
    error,
    deposit,
  };
}
