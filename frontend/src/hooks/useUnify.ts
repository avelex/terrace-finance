'use client';

import { useState, useCallback } from 'react';
import { Chain } from 'viem';
import { arcTestnet, baseSepolia, sepolia, avalancheFuji } from 'viem/chains';
import { UnifyRequest, UnifyResponse, PermitsRequest, TypedData, TokenBalancesByDomain } from '@/lib/types';
import { createWalletClientFromProvider, createPublicClientForChain } from '@/lib/wallet';

// Map chainId to chain config for adding chains to wallet
const CHAIN_MAP: Record<number, Chain> = {
  11155111: sepolia,
  43113: avalancheFuji,
  84532: baseSepolia,
  5042002: arcTestnet,
};

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || '';

interface UseUnifyResult {
  isUnifying: boolean;
  currentStep: string | null;
  error: string | null;
  unify: (address: `0x${string}`, usdcBalances: TokenBalancesByDomain) => Promise<boolean>;
}

export function useUnify(): UseUnifyResult {
  const [isUnifying, setIsUnifying] = useState(false);
  const [currentStep, setCurrentStep] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  const unify = useCallback(async (
    address: `0x${string}`,
    usdcBalances: TokenBalancesByDomain
  ): Promise<boolean> => {
    setIsUnifying(true);
    setError(null);
    setCurrentStep('Initializing...');

    try {
      // Step 1: Get domains with balance > 0
      const domainsWithBalance = Object.entries(usdcBalances)
        .filter(([, balance]) => balance > 0)
        .map(([domain]) => Number(domain));

      if (domainsWithBalance.length === 0) {
        setError('No USDC balance to unify');
        return false;
      }

      setCurrentStep(`Requesting permits for ${domainsWithBalance.length} chains...`);

      // Step 2: POST /api/wallet/:address/unify
      const unifyRequest: UnifyRequest = { domains: domainsWithBalance };
      const unifyResponse = await fetch(`${API_BASE_URL}/api/wallet/${address}/unify`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(unifyRequest),
      });

      if (!unifyResponse.ok) {
        throw new Error(`Failed to initiate unify: ${unifyResponse.status}`);
      }

      const unifyData: UnifyResponse = await unifyResponse.json();
      const { id, domains: typedDataByDomain } = unifyData;

      // Step 3: Sign each typedData
      const walletClient = await createWalletClientFromProvider();
      const [account] = await walletClient.getAddresses();

      const signatures: Record<number, string> = {};
      const domainEntries = Object.entries(typedDataByDomain);
      
      for (let i = 0; i < domainEntries.length; i++) {
        const [domainStr, typedDataRaw] = domainEntries[i];
        const domain = Number(domainStr);
        
        // Backend sends typedData as JSON string, parse it
        const typedData: TypedData = typeof typedDataRaw === 'string' 
          ? JSON.parse(typedDataRaw) 
          : typedDataRaw;
        
        // Get chainId from typedData domain (can be hex string or number)
        const requiredChainId = typedData.domain.chainId;
        if (requiredChainId) {
          const chainIdHex = typeof requiredChainId === 'string' && requiredChainId.startsWith('0x')
            ? requiredChainId
            : `0x${Number(requiredChainId).toString(16)}`;
          
          setCurrentStep(`Switching to chain ${chainIdHex}...`);
          
          try {
            await window.ethereum?.request({
              method: 'wallet_switchEthereumChain',
              params: [{ chainId: chainIdHex }],
            });
          } catch (switchError: unknown) {
            // Chain not recognized or not added - try to add it
            const err = switchError as { code?: number; message?: string };
            const isUnrecognizedChain = err.code === 4902 || 
              (err.message && err.message.toLowerCase().includes('unrecognized chain'));
            
            if (isUnrecognizedChain) {
              const chainIdNum = parseInt(chainIdHex, 16);
              const chainConfig = CHAIN_MAP[chainIdNum];
              
              if (!chainConfig) {
                throw new Error(`Chain ${chainIdHex} not supported. Please add it manually.`);
              }
              
              // Add the chain to wallet
              setCurrentStep(`Adding ${chainConfig.name} to wallet...`);
              
              await window.ethereum?.request({
                method: 'wallet_addEthereumChain',
                params: [{
                  chainId: chainIdHex,
                  chainName: chainConfig.name,
                  nativeCurrency: chainConfig.nativeCurrency,
                  rpcUrls: chainConfig.rpcUrls.default.http,
                  blockExplorerUrls: chainConfig.blockExplorers?.default 
                    ? [chainConfig.blockExplorers.default.url] 
                    : undefined,
                }],
              });
            } else {
              throw switchError;
            }
          }
          
          // Recreate wallet client after chain switch
          const newWalletClient = await createWalletClientFromProvider();
          
          setCurrentStep(`Signing ${i + 1}/${domainEntries.length}...`);

          // Normalize chainId to number for viem
          const normalizedDomain = {
            ...typedData.domain,
            chainId: typedData.domain.chainId 
              ? (typeof typedData.domain.chainId === 'string' 
                  ? parseInt(typedData.domain.chainId, 16) 
                  : typedData.domain.chainId)
              : undefined,
          };

          const signature = await newWalletClient.signTypedData({
            account,
            domain: normalizedDomain,
            types: typedData.types,
            primaryType: typedData.primaryType,
            message: typedData.message,
          });

          // Verify signature using public client for this chain
          setCurrentStep(`Verifying signature ${i + 1}/${domainEntries.length}...`);
          
          const chainIdNum = normalizedDomain.chainId!;
          const chainPublicClient = createPublicClientForChain(chainIdNum);
          
          const isValid = await chainPublicClient.verifyTypedData({
            address: account,
            domain: normalizedDomain,
            types: typedData.types as never, // Dynamic types from backend
            primaryType: typedData.primaryType,
            message: typedData.message,
            signature,
          });

          if (!isValid) {
            throw new Error(
              `Invalid permit signature for ${account}: ${signature}`,
            );
          }

          signatures[domain] = signature;
        } else {
          setCurrentStep(`Signing ${i + 1}/${domainEntries.length}...`);

          const signature = await walletClient.signTypedData({
            account,
            domain: typedData.domain as { chainId?: number; name?: string; version?: string; verifyingContract?: `0x${string}`; salt?: `0x${string}` },
            types: typedData.types,
            primaryType: typedData.primaryType,
            message: typedData.message,
          });

          signatures[domain] = signature;
        }
      }

      // Step 4: POST /api/wallet/:address/unify/permits
      setCurrentStep('Submitting permits...');

      const permitsRequest: PermitsRequest = {
        id,
        domains: signatures,
      };

      const permitsResponse = await fetch(`${API_BASE_URL}/api/wallet/${address}/unify/permits`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(permitsRequest),
      });

      if (!permitsResponse.ok) {
        throw new Error(`Failed to submit permits: ${permitsResponse.status}`);
      }

      setCurrentStep('Complete!');
      return true;
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Unknown error';
      setError(message);
      console.error('Unify error:', err);
      return false;
    } finally {
      setIsUnifying(false);
      // Clear step after a delay
      setTimeout(() => setCurrentStep(null), 2000);
    }
  }, []);

  return {
    isUnifying,
    currentStep,
    error,
    unify,
  };
}
