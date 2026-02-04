package transactor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/avelex/terrace-finance/backend/internal/abi"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type DomainTransactor struct {
	logger zerolog.Logger
	domain uint32

	lock   *sync.Mutex
	client *ethclient.Client
	opts   *bind.TransactOpts

	vaultAddress         common.Address
	gatewayWalletAddress common.Address
	gatewayMintAddress   common.Address

	vault         *abi.TerraceTransactor
	gatewayWallet *abi.GatewayWalletTransactor
	gatewayMint   *abi.GatewayMintTransactor
}

func NewDomainTransactor(key *ecdsa.PrivateKey, client *ethclient.Client, domain uint32, vault common.Address, gatewayWallet common.Address, gatewayMint common.Address) (*DomainTransactor, error) {
	vaultTransactor, err := abi.NewTerraceTransactor(vault, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault transactor: %w", err)
	}

	gatewayWalletTransactor, err := abi.NewGatewayWalletTransactor(gatewayWallet, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gateway wallet transactor: %w", err)
	}

	gatewayMintTransactor, err := abi.NewGatewayMintTransactor(gatewayMint, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gateway mint transactor: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, fmt.Errorf("create tx opts: %w", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), opts.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasLimit = 600_000

	return &DomainTransactor{
		domain:               domain,
		logger:               log.With().Uint32("domain", domain).Logger(),
		lock:                 &sync.Mutex{},
		client:               client,
		vaultAddress:         vault,
		gatewayWalletAddress: gatewayWallet,
		gatewayMintAddress:   gatewayMint,
		opts:                 opts,
		vault:                vaultTransactor,
		gatewayWallet:        gatewayWalletTransactor,
		gatewayMint:          gatewayMintTransactor,
	}, nil
}

func (dt *DomainTransactor) ReceiveMessage(message, attestation []byte) (*types.Receipt, error) {
	receipt, err := dt.sendTransaction(context.Background(), "ReceiveMessage", func() (*types.Transaction, error) {
		return dt.vault.ReceiveMessage(dt.opts, message, attestation)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to receive message: %w", err)
	}

	return receipt, nil
}

func (dt *DomainTransactor) SendAllFundsToHub(maxFee *big.Int, minFinalityThreshold uint32) (*types.Receipt, error) {
	receipt, err := dt.sendTransaction(context.Background(), "SendAllToHub", func() (*types.Transaction, error) {
		return dt.vault.SendAllToHub(dt.opts, maxFee, minFinalityThreshold)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send all funds to hub: %w", err)
	}

	return receipt, nil
}

func (dt *DomainTransactor) SendAllFundsToTerrace(domain enum.CircleDomain, maxFee *big.Int, minFinalityThreshold uint32) (*types.Receipt, error) {
	receipt, err := dt.sendTransaction(context.Background(), "SendAllToTerrace", func() (*types.Transaction, error) {
		return dt.vault.SendAllToTerrace(dt.opts, uint32(domain), maxFee, minFinalityThreshold)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send all funds to terrace: %w", err)
	}

	return receipt, nil
}

func (dt *DomainTransactor) BatchExecute(targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Receipt, error) {
	receipt, err := dt.sendTransaction(context.Background(), "BatchExecute", func() (*types.Transaction, error) {
		return dt.vault.BatchExecute(dt.opts, targets, selectors, data, proofs)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to batch execute: %w", err)
	}

	return receipt, nil
}

func (dt *DomainTransactor) TerraceBalanceOf(ctx context.Context, token common.Address) (*big.Int, error) {
	caller, err := abi.NewIERC20Caller(token, dt.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create ERC20 caller: %w", err)
	}

	opts := &bind.CallOpts{Context: ctx}

	return caller.BalanceOf(opts, dt.vaultAddress)
}

func (dt *DomainTransactor) DepositWithPermit(token, owner common.Address, amount, deadline *big.Int, sig []byte) (*types.Receipt, error) {
	receipt, err := dt.sendTransaction(context.Background(), "DepositWithPermit", func() (*types.Transaction, error) {
		return dt.gatewayWallet.DepositWithPermit(dt.opts, token, owner, amount, deadline, sig)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deposit with permit: %w", err)
	}

	return receipt, nil
}

func (dt *DomainTransactor) GatewayMint(attestation, sig []byte) (*types.Receipt, error) {
	receipt, err := dt.sendTransaction(context.Background(), "GatewayMint", func() (*types.Transaction, error) {
		return dt.gatewayMint.GatewayMint(dt.opts, attestation, sig)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to gateway mint: %w", err)
	}

	return receipt, nil
}

func (dt *DomainTransactor) sendTransaction(ctx context.Context, method string, txFunc func() (*types.Transaction, error)) (*types.Receipt, error) {
	dt.lock.Lock()
	defer dt.lock.Unlock()

	methodLog := dt.logger.With().Str("method", method).Logger()
	methodLog.Info().Msg("sending transaction")

	tx, err := txFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	txLog := methodLog.With().
		Str("tx_hash", tx.Hash().String()).
		Logger()

	txLog.Info().Msg("transaction sent, waiting for confirmation")

	receipt, err := bind.WaitMined(ctx, dt.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for receipt: %w", err)
	}

	dt.opts.Nonce.Add(dt.opts.Nonce, big.NewInt(1))

	if receipt.Status == 0 {
		txLog.Info().Msg("transaction reverted")

		return nil, fmt.Errorf("transaction reverted")
	}

	txLog.Info().Msg("transaction success")

	return receipt, nil
}
