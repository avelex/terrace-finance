package transactor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/avelex/terrace-finance/backend/internal/abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type DomainTransactor struct {
	logger zerolog.Logger

	lock   *sync.Mutex
	client *ethclient.Client

	opts       *bind.TransactOpts
	transactor *abi.TerraceTransactor
}

func NewDomainTransactor(key *ecdsa.PrivateKey, client *ethclient.Client, domain uint32, contract common.Address) (*DomainTransactor, error) {
	transactor, err := abi.NewTerraceTransactor(contract, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
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

	return &DomainTransactor{
		logger:     log.With().Uint32("domain", domain).Logger(),
		lock:       &sync.Mutex{},
		client:     client,
		opts:       opts,
		transactor: transactor,
	}, nil
}

func (dt *DomainTransactor) ReceiveMessage(message, attestation []byte) (*types.Receipt, error) {
	dt.lock.Lock()
	defer dt.lock.Unlock()

	tx, err := dt.transactor.ReceiveMessage(dt.opts, message, attestation)
	if err != nil {
		return nil, fmt.Errorf("failed to receive message: %w", err)
	}

	txLog := dt.logger.With().
		Str("tx_hash", tx.Hash().String()).
		Str("method", "ReceiveMessage").
		Logger()

	txLog.Info().Msg("transaction sent, waiting for confirmation")

	receipt, err := bind.WaitMined(context.Background(), dt.client, tx)
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
