package wallet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/avelex/terrace-finance/backend/config"
	"github.com/avelex/terrace-finance/backend/internal/abi"
	circle_gateway "github.com/avelex/terrace-finance/backend/internal/circle/gateway/client"
	"github.com/avelex/terrace-finance/backend/internal/models"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/models/request"
	"github.com/avelex/terrace-finance/backend/internal/models/response"
	"github.com/avelex/terrace-finance/backend/internal/repository"
	"github.com/avelex/terrace-finance/backend/internal/services/network"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type Service struct {
	protocol config.Protocol
	gateway  *circle_gateway.Client
	ns       *network.Service
	repo     *repository.UserRepository
}

func NewService(cfg config.Protocol, gateway *circle_gateway.Client, ns *network.Service, repo *repository.UserRepository) *Service {
	return &Service{
		protocol: cfg,
		gateway:  gateway,
		ns:       ns,
		repo:     repo,
	}
}

func (s *Service) ProtocolBalances(ctx context.Context, user common.Address) (*response.ProtocolBalances, error) {
	supportedDomains := enum.SupportedDomains()
	if len(supportedDomains) == 0 {
		return nil, fmt.Errorf("no supported domains")
	}

	unifiedBalances, err := s.gateway.Balances(ctx, user, supportedDomains)
	if err != nil {
		log.Error().Err(err).Msg("gateway balances")

		if !errors.Is(err, circle_gateway.ErrUnifiedBalanceNotFound) {
			return nil, fmt.Errorf("gateway balances: %w", err)
		}

		unifiedBalances = make(map[enum.CircleDomain]decimal.Decimal)
	}

	usdcBalances := make(map[enum.CircleDomain]decimal.Decimal)
	for _, domain := range supportedDomains {
		network := enum.NetworkByDomain(domain)
		client, err := s.ns.Client(network)
		if err != nil {
			return nil, fmt.Errorf("client for %s: %w", network, err)
		}

		usdcCaller, err := abi.NewIERC20Caller(enum.USDC_MAPPING[network], client)
		if err != nil {
			return nil, fmt.Errorf("usdc caller for %s: %w", network, err)
		}

		usdcBalance, err := usdcCaller.BalanceOf(&bind.CallOpts{Context: ctx}, user)
		if err != nil {
			return nil, fmt.Errorf("usdc balance in %s: %w", network, err)
		}

		usdcBalances[domain] = decimal.NewFromBigInt(usdcBalance, 0)
	}

	hubClient, err := s.ns.Client(enum.NetworkByDomain(s.protocol.Hub))
	if err != nil {
		return nil, fmt.Errorf("hub client: %w", err)
	}

	stableCaller, err := abi.NewIERC20Caller(s.protocol.Stablecoin, hubClient)
	if err != nil {
		return nil, fmt.Errorf("stablecoin caller: %w", err)
	}

	stableBalance, err := stableCaller.BalanceOf(&bind.CallOpts{Context: ctx}, user)
	if err != nil {
		return nil, fmt.Errorf("stablecoin balance: %w", err)
	}

	stakingCaller, err := abi.NewIERC20Caller(s.protocol.Staking, hubClient)
	if err != nil {
		return nil, fmt.Errorf("staking caller: %w", err)
	}

	stakingBalance, err := stakingCaller.BalanceOf(&bind.CallOpts{Context: ctx}, user)
	if err != nil {
		return nil, fmt.Errorf("staking balance: %w", err)
	}

	return &response.ProtocolBalances{
		UnifiedUSDC:       unifiedBalances,
		USDC:              usdcBalances,
		StablecoinBalance: decimal.NewFromBigInt(stableBalance, 0),
		StakingBalance:    decimal.NewFromBigInt(stakingBalance, 0),
	}, nil
}

func (s *Service) UserDeposits(ctx context.Context, user common.Address) ([]models.UserDeposit, error) {
	return s.repo.GetUserDeposits(ctx, user)
}

func (s *Service) UserUnifyDeposits(ctx context.Context, user common.Address) ([]models.UserUnifiedPermits, error) {
	return s.repo.GetUserUnifiedPermits(ctx, user)
}

func (s *Service) UnifyUSDC(ctx context.Context, user common.Address, domains []enum.CircleDomain) (*response.PermitPayload, error) {
	permitsPayload := make(map[enum.CircleDomain]string)
	permits := make([]models.UserUnifiedPermits, 0, len(domains))
	depositID := uuid.New()

	for _, domain := range domains {
		network := enum.NetworkByDomain(domain)

		client, err := s.ns.Client(network)
		if err != nil {
			return nil, fmt.Errorf("client for %s: %w", network, err)
		}

		usdcAddress := enum.USDC_MAPPING[network]

		usdcCaller, err := abi.NewUSDCCaller(usdcAddress, client)
		if err != nil {
			return nil, fmt.Errorf("usdc caller for %s: %w", network, err)
		}

		usdcBalance, err := usdcCaller.BalanceOf(&bind.CallOpts{Context: ctx}, user)
		if err != nil {
			return nil, fmt.Errorf("usdc balance in %s: %w", network, err)
		}

		if domain == enum.ARC_DOMAIN {
			usdcBalance = new(big.Int).Sub(usdcBalance, big.NewInt(1))
		}

		permitNonce, err := usdcCaller.Nonces(&bind.CallOpts{Context: ctx}, user)
		if err != nil {
			return nil, fmt.Errorf("usdc permit nonce in %s: %w", network, err)
		}

		usdcName, err := usdcCaller.Name(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("usdc name in %s: %w", network, err)
		}

		usdcVersion, err := usdcCaller.Version(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("usdc version in %s: %w", network, err)
		}

		chainID, err := client.ChainID(ctx)
		if err != nil {
			return nil, fmt.Errorf("chain id for %s: %w", network, err)
		}

		deadline := time.Now().Unix() + 3600
		gatewayWallet := enum.GATEWAY_WALLET_MAPPING[network]

		typedData := NewEip712TypedData(usdcName, usdcVersion, usdcAddress, user, gatewayWallet, usdcBalance, chainID.Int64(), deadline, permitNonce)
		bytes, err := json.Marshal(&typedData)
		if err != nil {
			return nil, fmt.Errorf("marshal typed data: %w", err)
		}

		permitsPayload[domain] = string(bytes)
		permits = append(permits, models.UserUnifiedPermits{
			ID:            depositID,
			Owner:         user.String(),
			Token:         usdcAddress.String(),
			Value:         usdcBalance.String(),
			Deadline:      deadline,
			Domain:        uint32(domain),
			GatewayWallet: gatewayWallet.String(),
		})
	}

	if err := s.repo.InitiateUserDepositFlow(ctx, permits); err != nil {
		return nil, fmt.Errorf("unable to initiate user deposit flow: %w", err)
	}

	return &response.PermitPayload{
		ID:      depositID.String(),
		Domains: permitsPayload,
	}, nil
}

func (s *Service) SaveUnifyPermitsSignatures(ctx context.Context, user common.Address, signedPermits request.SignedPermits) error {
	permits := make([]models.UserUnifiedPermits, 0, len(signedPermits.Domains))

	for domain, sig := range signedPermits.Domains {
		permits = append(permits, models.UserUnifiedPermits{
			ID:        signedPermits.ID,
			Owner:     user.String(),
			Domain:    uint32(domain),
			Signature: sig,
		})
	}

	if err := s.repo.UpdateUserUnifiedPermitsSignatures(ctx, permits); err != nil {
		return fmt.Errorf("failed to save unify permits: %w", err)
	}

	return nil
}

func (s *Service) InitDepositAndStake(ctx context.Context, user common.Address, deposit request.DepositAndStake) error {
	userDeposit := models.UserDeposit{
		ID:              uuid.New(),
		Address:         user.String(),
		Value:           deposit.Amount.String(),
		DestDomain:      uint32(s.protocol.Hub),
		DestGatewayMint: enum.GATEWAY_MINT_MAPPING[enum.NetworkByDomain(s.protocol.Hub)].String(),
		Attestation:     deposit.Attestation,
		Signature:       deposit.Signature,
	}

	if err := s.repo.CreateUserDeposit(ctx, userDeposit); err != nil {
		return fmt.Errorf("failed to create user deposit: %w", err)
	}

	return nil
}

func NewEip712TypedData(name string, version string, usdc, depositor, spender common.Address, balance *big.Int, chainID, deadline int64, nonce *big.Int) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:              name,
			Version:           version,
			ChainId:           math.NewHexOrDecimal256(chainID),
			VerifyingContract: usdc.String(),
		},
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Permit": []apitypes.Type{
				{Name: "owner", Type: "address"},
				{Name: "spender", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
			},
		},
		PrimaryType: "Permit",
		Message: apitypes.TypedDataMessage{
			"owner":    depositor.String(),
			"spender":  spender.String(),
			"value":    balance.String(),
			"nonce":    nonce.String(),
			"deadline": fmt.Sprintf("%d", deadline),
		},
	}
}
