// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GatewayWalletMetaData contains all meta data concerning the GatewayWallet contract.
var GatewayWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AccountDenylisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"}],\"name\":\"BurnFeeTooHighAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotDelegateToSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursorOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositValueMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"DepositorIsBlacklisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxBlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"IntentExpiredAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"IntentValueMustBePositiveAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBurnSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"intentContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedContract\",\"type\":\"address\"}],\"name\":\"InvalidIntentSourceContractAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"intentSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actualSigner\",\"type\":\"address\"}],\"name\":\"InvalidIntentSourceSignerAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"actualMagic\",\"type\":\"bytes4\"}],\"name\":\"InvalidTransferPayloadMagic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"actualMagic\",\"type\":\"bytes4\"}],\"name\":\"InvalidTransferSpecMagic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actualVersion\",\"type\":\"uint32\"}],\"name\":\"InvalidTransferSpecVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedBurn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustHaveAtLeastOneBurnIntent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRelevantBurnIntents\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWithdrawingBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllSameToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadDataTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedTotalLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTotalLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadOverallLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"actualSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredOffset\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetElementHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"actualSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredOffset\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetElementTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes4\",\"name\":\"actualMagic\",\"type\":\"bytes4\"}],\"name\":\"TransferPayloadSetInvalidElementMagic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedTotalLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTotalLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetOverallLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferSpecHash\",\"type\":\"bytes32\"}],\"name\":\"TransferSpecHashUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferSpecHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedHookDataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferSpecLength\",\"type\":\"uint256\"}],\"name\":\"TransferSpecInvalidHookData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedTotalLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTotalLength\",\"type\":\"uint256\"}],\"name\":\"TransferSpecOverallLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"UnauthorizedDenylister\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"UnsupportedTokenAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalNotYetAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalValueExceedsAvailableBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalValueMustBePositive\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"BurnSignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"BurnSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Denylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDenylister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"DenylisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferSpecHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationRecipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAvailable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromWithdrawing\",\"type\":\"uint256\"}],\"name\":\"GatewayBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawingBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"UnDenylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldDelay\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingAvailable\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawing\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalBlock\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"addBurnSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"addDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"availableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"depositors\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"denylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denylister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"depositWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validBefore\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositWithAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"depositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"calldataBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"gatewayBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"intent\",\"type\":\"bytes\"}],\"name\":\"getTypedDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauser_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"denylister_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"supportedTokens_\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"domain_\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalDelay_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burnSigner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAuthorizedForBalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isBurnSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isDenylisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferSpecHash\",\"type\":\"bytes32\"}],\"name\":\"isTransferSpecHashUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"removeBurnSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"removeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"unDenylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"updateDenylister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawalDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"withdrawableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"withdrawalBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"withdrawingBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GatewayWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayWalletMetaData.ABI instead.
var GatewayWalletABI = GatewayWalletMetaData.ABI

// GatewayWallet is an auto generated Go binding around an Ethereum contract.
type GatewayWallet struct {
	GatewayWalletCaller     // Read-only binding to the contract
	GatewayWalletTransactor // Write-only binding to the contract
	GatewayWalletFilterer   // Log filterer for contract events
}

// GatewayWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayWalletSession struct {
	Contract     *GatewayWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayWalletCallerSession struct {
	Contract *GatewayWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GatewayWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayWalletTransactorSession struct {
	Contract     *GatewayWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GatewayWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayWalletRaw struct {
	Contract *GatewayWallet // Generic contract binding to access the raw methods on
}

// GatewayWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayWalletCallerRaw struct {
	Contract *GatewayWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayWalletTransactorRaw struct {
	Contract *GatewayWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayWallet creates a new instance of GatewayWallet, bound to a specific deployed contract.
func NewGatewayWallet(address common.Address, backend bind.ContractBackend) (*GatewayWallet, error) {
	contract, err := bindGatewayWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayWallet{GatewayWalletCaller: GatewayWalletCaller{contract: contract}, GatewayWalletTransactor: GatewayWalletTransactor{contract: contract}, GatewayWalletFilterer: GatewayWalletFilterer{contract: contract}}, nil
}

// NewGatewayWalletCaller creates a new read-only instance of GatewayWallet, bound to a specific deployed contract.
func NewGatewayWalletCaller(address common.Address, caller bind.ContractCaller) (*GatewayWalletCaller, error) {
	contract, err := bindGatewayWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletCaller{contract: contract}, nil
}

// NewGatewayWalletTransactor creates a new write-only instance of GatewayWallet, bound to a specific deployed contract.
func NewGatewayWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayWalletTransactor, error) {
	contract, err := bindGatewayWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletTransactor{contract: contract}, nil
}

// NewGatewayWalletFilterer creates a new log filterer instance of GatewayWallet, bound to a specific deployed contract.
func NewGatewayWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayWalletFilterer, error) {
	contract, err := bindGatewayWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletFilterer{contract: contract}, nil
}

// bindGatewayWallet binds a generic wrapper to an already deployed contract.
func bindGatewayWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayWallet *GatewayWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayWallet.Contract.GatewayWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayWallet *GatewayWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayWallet.Contract.GatewayWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayWallet *GatewayWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayWallet.Contract.GatewayWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayWallet *GatewayWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayWallet *GatewayWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayWallet *GatewayWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayWallet.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xb727b579.
//
// Solidity: function EIP712_DOMAIN_TYPE_HASH() view returns(bytes32)
func (_GatewayWallet *GatewayWalletCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "EIP712_DOMAIN_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xb727b579.
//
// Solidity: function EIP712_DOMAIN_TYPE_HASH() view returns(bytes32)
func (_GatewayWallet *GatewayWalletSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _GatewayWallet.Contract.EIP712DOMAINTYPEHASH(&_GatewayWallet.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xb727b579.
//
// Solidity: function EIP712_DOMAIN_TYPE_HASH() view returns(bytes32)
func (_GatewayWallet *GatewayWalletCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _GatewayWallet.Contract.EIP712DOMAINTYPEHASH(&_GatewayWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayWallet *GatewayWalletCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayWallet *GatewayWalletSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayWallet.Contract.UPGRADEINTERFACEVERSION(&_GatewayWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayWallet *GatewayWalletCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayWallet.Contract.UPGRADEINTERFACEVERSION(&_GatewayWallet.CallOpts)
}

// AvailableBalance is a free data retrieval call binding the contract method 0x3ccb64ae.
//
// Solidity: function availableBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCaller) AvailableBalance(opts *bind.CallOpts, token common.Address, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "availableBalance", token, depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableBalance is a free data retrieval call binding the contract method 0x3ccb64ae.
//
// Solidity: function availableBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletSession) AvailableBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.AvailableBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// AvailableBalance is a free data retrieval call binding the contract method 0x3ccb64ae.
//
// Solidity: function availableBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCallerSession) AvailableBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.AvailableBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address depositor, uint256 id) view returns(uint256 balance)
func (_GatewayWallet *GatewayWalletCaller) BalanceOf(opts *bind.CallOpts, depositor common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "balanceOf", depositor, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address depositor, uint256 id) view returns(uint256 balance)
func (_GatewayWallet *GatewayWalletSession) BalanceOf(depositor common.Address, id *big.Int) (*big.Int, error) {
	return _GatewayWallet.Contract.BalanceOf(&_GatewayWallet.CallOpts, depositor, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address depositor, uint256 id) view returns(uint256 balance)
func (_GatewayWallet *GatewayWalletCallerSession) BalanceOf(depositor common.Address, id *big.Int) (*big.Int, error) {
	return _GatewayWallet.Contract.BalanceOf(&_GatewayWallet.CallOpts, depositor, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] depositors, uint256[] ids) view returns(uint256[] balances)
func (_GatewayWallet *GatewayWalletCaller) BalanceOfBatch(opts *bind.CallOpts, depositors []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "balanceOfBatch", depositors, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] depositors, uint256[] ids) view returns(uint256[] balances)
func (_GatewayWallet *GatewayWalletSession) BalanceOfBatch(depositors []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _GatewayWallet.Contract.BalanceOfBatch(&_GatewayWallet.CallOpts, depositors, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] depositors, uint256[] ids) view returns(uint256[] balances)
func (_GatewayWallet *GatewayWalletCallerSession) BalanceOfBatch(depositors []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _GatewayWallet.Contract.BalanceOfBatch(&_GatewayWallet.CallOpts, depositors, ids)
}

// Denylister is a free data retrieval call binding the contract method 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (_GatewayWallet *GatewayWalletCaller) Denylister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "denylister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Denylister is a free data retrieval call binding the contract method 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (_GatewayWallet *GatewayWalletSession) Denylister() (common.Address, error) {
	return _GatewayWallet.Contract.Denylister(&_GatewayWallet.CallOpts)
}

// Denylister is a free data retrieval call binding the contract method 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (_GatewayWallet *GatewayWalletCallerSession) Denylister() (common.Address, error) {
	return _GatewayWallet.Contract.Denylister(&_GatewayWallet.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(uint32)
func (_GatewayWallet *GatewayWalletCaller) Domain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "domain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(uint32)
func (_GatewayWallet *GatewayWalletSession) Domain() (uint32, error) {
	return _GatewayWallet.Contract.Domain(&_GatewayWallet.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(uint32)
func (_GatewayWallet *GatewayWalletCallerSession) Domain() (uint32, error) {
	return _GatewayWallet.Contract.Domain(&_GatewayWallet.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() pure returns(bytes32)
func (_GatewayWallet *GatewayWalletCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() pure returns(bytes32)
func (_GatewayWallet *GatewayWalletSession) DomainSeparator() ([32]byte, error) {
	return _GatewayWallet.Contract.DomainSeparator(&_GatewayWallet.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() pure returns(bytes32)
func (_GatewayWallet *GatewayWalletCallerSession) DomainSeparator() ([32]byte, error) {
	return _GatewayWallet.Contract.DomainSeparator(&_GatewayWallet.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GatewayWallet *GatewayWalletCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GatewayWallet *GatewayWalletSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GatewayWallet.Contract.Eip712Domain(&_GatewayWallet.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GatewayWallet *GatewayWalletCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GatewayWallet.Contract.Eip712Domain(&_GatewayWallet.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_GatewayWallet *GatewayWalletCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_GatewayWallet *GatewayWalletSession) FeeRecipient() (common.Address, error) {
	return _GatewayWallet.Contract.FeeRecipient(&_GatewayWallet.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_GatewayWallet *GatewayWalletCallerSession) FeeRecipient() (common.Address, error) {
	return _GatewayWallet.Contract.FeeRecipient(&_GatewayWallet.CallOpts)
}

// GetTypedDataHash is a free data retrieval call binding the contract method 0xdc4aa513.
//
// Solidity: function getTypedDataHash(bytes intent) view returns(bytes32)
func (_GatewayWallet *GatewayWalletCaller) GetTypedDataHash(opts *bind.CallOpts, intent []byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "getTypedDataHash", intent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTypedDataHash is a free data retrieval call binding the contract method 0xdc4aa513.
//
// Solidity: function getTypedDataHash(bytes intent) view returns(bytes32)
func (_GatewayWallet *GatewayWalletSession) GetTypedDataHash(intent []byte) ([32]byte, error) {
	return _GatewayWallet.Contract.GetTypedDataHash(&_GatewayWallet.CallOpts, intent)
}

// GetTypedDataHash is a free data retrieval call binding the contract method 0xdc4aa513.
//
// Solidity: function getTypedDataHash(bytes intent) view returns(bytes32)
func (_GatewayWallet *GatewayWalletCallerSession) GetTypedDataHash(intent []byte) ([32]byte, error) {
	return _GatewayWallet.Contract.GetTypedDataHash(&_GatewayWallet.CallOpts, intent)
}

// IsAuthorizedForBalance is a free data retrieval call binding the contract method 0x74392bf3.
//
// Solidity: function isAuthorizedForBalance(address token, address depositor, address addr) view returns(bool)
func (_GatewayWallet *GatewayWalletCaller) IsAuthorizedForBalance(opts *bind.CallOpts, token common.Address, depositor common.Address, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "isAuthorizedForBalance", token, depositor, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedForBalance is a free data retrieval call binding the contract method 0x74392bf3.
//
// Solidity: function isAuthorizedForBalance(address token, address depositor, address addr) view returns(bool)
func (_GatewayWallet *GatewayWalletSession) IsAuthorizedForBalance(token common.Address, depositor common.Address, addr common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsAuthorizedForBalance(&_GatewayWallet.CallOpts, token, depositor, addr)
}

// IsAuthorizedForBalance is a free data retrieval call binding the contract method 0x74392bf3.
//
// Solidity: function isAuthorizedForBalance(address token, address depositor, address addr) view returns(bool)
func (_GatewayWallet *GatewayWalletCallerSession) IsAuthorizedForBalance(token common.Address, depositor common.Address, addr common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsAuthorizedForBalance(&_GatewayWallet.CallOpts, token, depositor, addr)
}

// IsBurnSigner is a free data retrieval call binding the contract method 0x3660dbd0.
//
// Solidity: function isBurnSigner(address signer) view returns(bool)
func (_GatewayWallet *GatewayWalletCaller) IsBurnSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "isBurnSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBurnSigner is a free data retrieval call binding the contract method 0x3660dbd0.
//
// Solidity: function isBurnSigner(address signer) view returns(bool)
func (_GatewayWallet *GatewayWalletSession) IsBurnSigner(signer common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsBurnSigner(&_GatewayWallet.CallOpts, signer)
}

// IsBurnSigner is a free data retrieval call binding the contract method 0x3660dbd0.
//
// Solidity: function isBurnSigner(address signer) view returns(bool)
func (_GatewayWallet *GatewayWalletCallerSession) IsBurnSigner(signer common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsBurnSigner(&_GatewayWallet.CallOpts, signer)
}

// IsDenylisted is a free data retrieval call binding the contract method 0xe877a526.
//
// Solidity: function isDenylisted(address addr) view returns(bool)
func (_GatewayWallet *GatewayWalletCaller) IsDenylisted(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "isDenylisted", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDenylisted is a free data retrieval call binding the contract method 0xe877a526.
//
// Solidity: function isDenylisted(address addr) view returns(bool)
func (_GatewayWallet *GatewayWalletSession) IsDenylisted(addr common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsDenylisted(&_GatewayWallet.CallOpts, addr)
}

// IsDenylisted is a free data retrieval call binding the contract method 0xe877a526.
//
// Solidity: function isDenylisted(address addr) view returns(bool)
func (_GatewayWallet *GatewayWalletCallerSession) IsDenylisted(addr common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsDenylisted(&_GatewayWallet.CallOpts, addr)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address token) view returns(bool)
func (_GatewayWallet *GatewayWalletCaller) IsTokenSupported(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "isTokenSupported", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address token) view returns(bool)
func (_GatewayWallet *GatewayWalletSession) IsTokenSupported(token common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsTokenSupported(&_GatewayWallet.CallOpts, token)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address token) view returns(bool)
func (_GatewayWallet *GatewayWalletCallerSession) IsTokenSupported(token common.Address) (bool, error) {
	return _GatewayWallet.Contract.IsTokenSupported(&_GatewayWallet.CallOpts, token)
}

// IsTransferSpecHashUsed is a free data retrieval call binding the contract method 0x17580158.
//
// Solidity: function isTransferSpecHashUsed(bytes32 transferSpecHash) view returns(bool)
func (_GatewayWallet *GatewayWalletCaller) IsTransferSpecHashUsed(opts *bind.CallOpts, transferSpecHash [32]byte) (bool, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "isTransferSpecHashUsed", transferSpecHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferSpecHashUsed is a free data retrieval call binding the contract method 0x17580158.
//
// Solidity: function isTransferSpecHashUsed(bytes32 transferSpecHash) view returns(bool)
func (_GatewayWallet *GatewayWalletSession) IsTransferSpecHashUsed(transferSpecHash [32]byte) (bool, error) {
	return _GatewayWallet.Contract.IsTransferSpecHashUsed(&_GatewayWallet.CallOpts, transferSpecHash)
}

// IsTransferSpecHashUsed is a free data retrieval call binding the contract method 0x17580158.
//
// Solidity: function isTransferSpecHashUsed(bytes32 transferSpecHash) view returns(bool)
func (_GatewayWallet *GatewayWalletCallerSession) IsTransferSpecHashUsed(transferSpecHash [32]byte) (bool, error) {
	return _GatewayWallet.Contract.IsTransferSpecHashUsed(&_GatewayWallet.CallOpts, transferSpecHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayWallet *GatewayWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayWallet *GatewayWalletSession) Owner() (common.Address, error) {
	return _GatewayWallet.Contract.Owner(&_GatewayWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayWallet *GatewayWalletCallerSession) Owner() (common.Address, error) {
	return _GatewayWallet.Contract.Owner(&_GatewayWallet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayWallet *GatewayWalletCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayWallet *GatewayWalletSession) Paused() (bool, error) {
	return _GatewayWallet.Contract.Paused(&_GatewayWallet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayWallet *GatewayWalletCallerSession) Paused() (bool, error) {
	return _GatewayWallet.Contract.Paused(&_GatewayWallet.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_GatewayWallet *GatewayWalletCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_GatewayWallet *GatewayWalletSession) Pauser() (common.Address, error) {
	return _GatewayWallet.Contract.Pauser(&_GatewayWallet.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_GatewayWallet *GatewayWalletCallerSession) Pauser() (common.Address, error) {
	return _GatewayWallet.Contract.Pauser(&_GatewayWallet.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GatewayWallet *GatewayWalletCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GatewayWallet *GatewayWalletSession) PendingOwner() (common.Address, error) {
	return _GatewayWallet.Contract.PendingOwner(&_GatewayWallet.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GatewayWallet *GatewayWalletCallerSession) PendingOwner() (common.Address, error) {
	return _GatewayWallet.Contract.PendingOwner(&_GatewayWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayWallet *GatewayWalletCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayWallet *GatewayWalletSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayWallet.Contract.ProxiableUUID(&_GatewayWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayWallet *GatewayWalletCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayWallet.Contract.ProxiableUUID(&_GatewayWallet.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0x1453b987.
//
// Solidity: function totalBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCaller) TotalBalance(opts *bind.CallOpts, token common.Address, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "totalBalance", token, depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalance is a free data retrieval call binding the contract method 0x1453b987.
//
// Solidity: function totalBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletSession) TotalBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.TotalBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// TotalBalance is a free data retrieval call binding the contract method 0x1453b987.
//
// Solidity: function totalBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCallerSession) TotalBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.TotalBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x3bbe1ecd.
//
// Solidity: function withdrawableBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCaller) WithdrawableBalance(opts *bind.CallOpts, token common.Address, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "withdrawableBalance", token, depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x3bbe1ecd.
//
// Solidity: function withdrawableBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletSession) WithdrawableBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawableBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x3bbe1ecd.
//
// Solidity: function withdrawableBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCallerSession) WithdrawableBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawableBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// WithdrawalBlock is a free data retrieval call binding the contract method 0xfefeec13.
//
// Solidity: function withdrawalBlock(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCaller) WithdrawalBlock(opts *bind.CallOpts, token common.Address, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "withdrawalBlock", token, depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalBlock is a free data retrieval call binding the contract method 0xfefeec13.
//
// Solidity: function withdrawalBlock(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletSession) WithdrawalBlock(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawalBlock(&_GatewayWallet.CallOpts, token, depositor)
}

// WithdrawalBlock is a free data retrieval call binding the contract method 0xfefeec13.
//
// Solidity: function withdrawalBlock(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCallerSession) WithdrawalBlock(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawalBlock(&_GatewayWallet.CallOpts, token, depositor)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_GatewayWallet *GatewayWalletCaller) WithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "withdrawalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_GatewayWallet *GatewayWalletSession) WithdrawalDelay() (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawalDelay(&_GatewayWallet.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xa7ab6961.
//
// Solidity: function withdrawalDelay() view returns(uint256)
func (_GatewayWallet *GatewayWalletCallerSession) WithdrawalDelay() (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawalDelay(&_GatewayWallet.CallOpts)
}

// WithdrawingBalance is a free data retrieval call binding the contract method 0xdf0c6690.
//
// Solidity: function withdrawingBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCaller) WithdrawingBalance(opts *bind.CallOpts, token common.Address, depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GatewayWallet.contract.Call(opts, &out, "withdrawingBalance", token, depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawingBalance is a free data retrieval call binding the contract method 0xdf0c6690.
//
// Solidity: function withdrawingBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletSession) WithdrawingBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawingBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// WithdrawingBalance is a free data retrieval call binding the contract method 0xdf0c6690.
//
// Solidity: function withdrawingBalance(address token, address depositor) view returns(uint256)
func (_GatewayWallet *GatewayWalletCallerSession) WithdrawingBalance(token common.Address, depositor common.Address) (*big.Int, error) {
	return _GatewayWallet.Contract.WithdrawingBalance(&_GatewayWallet.CallOpts, token, depositor)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GatewayWallet *GatewayWalletTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GatewayWallet *GatewayWalletSession) AcceptOwnership() (*types.Transaction, error) {
	return _GatewayWallet.Contract.AcceptOwnership(&_GatewayWallet.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GatewayWallet *GatewayWalletTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GatewayWallet.Contract.AcceptOwnership(&_GatewayWallet.TransactOpts)
}

// AddBurnSigner is a paid mutator transaction binding the contract method 0xb3f69f8d.
//
// Solidity: function addBurnSigner(address signer) returns()
func (_GatewayWallet *GatewayWalletTransactor) AddBurnSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "addBurnSigner", signer)
}

// AddBurnSigner is a paid mutator transaction binding the contract method 0xb3f69f8d.
//
// Solidity: function addBurnSigner(address signer) returns()
func (_GatewayWallet *GatewayWalletSession) AddBurnSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.AddBurnSigner(&_GatewayWallet.TransactOpts, signer)
}

// AddBurnSigner is a paid mutator transaction binding the contract method 0xb3f69f8d.
//
// Solidity: function addBurnSigner(address signer) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) AddBurnSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.AddBurnSigner(&_GatewayWallet.TransactOpts, signer)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xe909ebfa.
//
// Solidity: function addDelegate(address token, address delegate) returns()
func (_GatewayWallet *GatewayWalletTransactor) AddDelegate(opts *bind.TransactOpts, token common.Address, delegate common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "addDelegate", token, delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xe909ebfa.
//
// Solidity: function addDelegate(address token, address delegate) returns()
func (_GatewayWallet *GatewayWalletSession) AddDelegate(token common.Address, delegate common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.AddDelegate(&_GatewayWallet.TransactOpts, token, delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xe909ebfa.
//
// Solidity: function addDelegate(address token, address delegate) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) AddDelegate(token common.Address, delegate common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.AddDelegate(&_GatewayWallet.TransactOpts, token, delegate)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address token) returns()
func (_GatewayWallet *GatewayWalletTransactor) AddSupportedToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "addSupportedToken", token)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address token) returns()
func (_GatewayWallet *GatewayWalletSession) AddSupportedToken(token common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.AddSupportedToken(&_GatewayWallet.TransactOpts, token)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address token) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) AddSupportedToken(token common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.AddSupportedToken(&_GatewayWallet.TransactOpts, token)
}

// Denylist is a paid mutator transaction binding the contract method 0x3371bfff.
//
// Solidity: function denylist(address addr) returns()
func (_GatewayWallet *GatewayWalletTransactor) Denylist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "denylist", addr)
}

// Denylist is a paid mutator transaction binding the contract method 0x3371bfff.
//
// Solidity: function denylist(address addr) returns()
func (_GatewayWallet *GatewayWalletSession) Denylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Denylist(&_GatewayWallet.TransactOpts, addr)
}

// Denylist is a paid mutator transaction binding the contract method 0x3371bfff.
//
// Solidity: function denylist(address addr) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) Denylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Denylist(&_GatewayWallet.TransactOpts, addr)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 value) returns()
func (_GatewayWallet *GatewayWalletTransactor) Deposit(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "deposit", token, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 value) returns()
func (_GatewayWallet *GatewayWalletSession) Deposit(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Deposit(&_GatewayWallet.TransactOpts, token, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 value) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) Deposit(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Deposit(&_GatewayWallet.TransactOpts, token, value)
}

// DepositFor is a paid mutator transaction binding the contract method 0xb3db428b.
//
// Solidity: function depositFor(address token, address depositor, uint256 value) returns()
func (_GatewayWallet *GatewayWalletTransactor) DepositFor(opts *bind.TransactOpts, token common.Address, depositor common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "depositFor", token, depositor, value)
}

// DepositFor is a paid mutator transaction binding the contract method 0xb3db428b.
//
// Solidity: function depositFor(address token, address depositor, uint256 value) returns()
func (_GatewayWallet *GatewayWalletSession) DepositFor(token common.Address, depositor common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositFor(&_GatewayWallet.TransactOpts, token, depositor, value)
}

// DepositFor is a paid mutator transaction binding the contract method 0xb3db428b.
//
// Solidity: function depositFor(address token, address depositor, uint256 value) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) DepositFor(token common.Address, depositor common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositFor(&_GatewayWallet.TransactOpts, token, depositor, value)
}

// DepositWithAuthorization is a paid mutator transaction binding the contract method 0x438d4835.
//
// Solidity: function depositWithAuthorization(address token, address from, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_GatewayWallet *GatewayWalletTransactor) DepositWithAuthorization(opts *bind.TransactOpts, token common.Address, from common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "depositWithAuthorization", token, from, value, validAfter, validBefore, nonce, signature)
}

// DepositWithAuthorization is a paid mutator transaction binding the contract method 0x438d4835.
//
// Solidity: function depositWithAuthorization(address token, address from, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_GatewayWallet *GatewayWalletSession) DepositWithAuthorization(token common.Address, from common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithAuthorization(&_GatewayWallet.TransactOpts, token, from, value, validAfter, validBefore, nonce, signature)
}

// DepositWithAuthorization is a paid mutator transaction binding the contract method 0x438d4835.
//
// Solidity: function depositWithAuthorization(address token, address from, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, bytes signature) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) DepositWithAuthorization(token common.Address, from common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithAuthorization(&_GatewayWallet.TransactOpts, token, from, value, validAfter, validBefore, nonce, signature)
}

// DepositWithAuthorization0 is a paid mutator transaction binding the contract method 0x8a94d4fc.
//
// Solidity: function depositWithAuthorization(address token, address from, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_GatewayWallet *GatewayWalletTransactor) DepositWithAuthorization0(opts *bind.TransactOpts, token common.Address, from common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "depositWithAuthorization0", token, from, value, validAfter, validBefore, nonce, v, r, s)
}

// DepositWithAuthorization0 is a paid mutator transaction binding the contract method 0x8a94d4fc.
//
// Solidity: function depositWithAuthorization(address token, address from, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_GatewayWallet *GatewayWalletSession) DepositWithAuthorization0(token common.Address, from common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithAuthorization0(&_GatewayWallet.TransactOpts, token, from, value, validAfter, validBefore, nonce, v, r, s)
}

// DepositWithAuthorization0 is a paid mutator transaction binding the contract method 0x8a94d4fc.
//
// Solidity: function depositWithAuthorization(address token, address from, uint256 value, uint256 validAfter, uint256 validBefore, bytes32 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) DepositWithAuthorization0(token common.Address, from common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithAuthorization0(&_GatewayWallet.TransactOpts, token, from, value, validAfter, validBefore, nonce, v, r, s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x26a3fb30.
//
// Solidity: function depositWithPermit(address token, address owner, uint256 value, uint256 deadline, bytes signature) returns()
func (_GatewayWallet *GatewayWalletTransactor) DepositWithPermit(opts *bind.TransactOpts, token common.Address, owner common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "depositWithPermit", token, owner, value, deadline, signature)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x26a3fb30.
//
// Solidity: function depositWithPermit(address token, address owner, uint256 value, uint256 deadline, bytes signature) returns()
func (_GatewayWallet *GatewayWalletSession) DepositWithPermit(token common.Address, owner common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithPermit(&_GatewayWallet.TransactOpts, token, owner, value, deadline, signature)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x26a3fb30.
//
// Solidity: function depositWithPermit(address token, address owner, uint256 value, uint256 deadline, bytes signature) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) DepositWithPermit(token common.Address, owner common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithPermit(&_GatewayWallet.TransactOpts, token, owner, value, deadline, signature)
}

// DepositWithPermit0 is a paid mutator transaction binding the contract method 0x8ef59739.
//
// Solidity: function depositWithPermit(address token, address owner, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GatewayWallet *GatewayWalletTransactor) DepositWithPermit0(opts *bind.TransactOpts, token common.Address, owner common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "depositWithPermit0", token, owner, value, deadline, v, r, s)
}

// DepositWithPermit0 is a paid mutator transaction binding the contract method 0x8ef59739.
//
// Solidity: function depositWithPermit(address token, address owner, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GatewayWallet *GatewayWalletSession) DepositWithPermit0(token common.Address, owner common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithPermit0(&_GatewayWallet.TransactOpts, token, owner, value, deadline, v, r, s)
}

// DepositWithPermit0 is a paid mutator transaction binding the contract method 0x8ef59739.
//
// Solidity: function depositWithPermit(address token, address owner, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) DepositWithPermit0(token common.Address, owner common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.DepositWithPermit0(&_GatewayWallet.TransactOpts, token, owner, value, deadline, v, r, s)
}

// GatewayBurn is a paid mutator transaction binding the contract method 0x8ec3dbb9.
//
// Solidity: function gatewayBurn(bytes calldataBytes, bytes signature) returns()
func (_GatewayWallet *GatewayWalletTransactor) GatewayBurn(opts *bind.TransactOpts, calldataBytes []byte, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "gatewayBurn", calldataBytes, signature)
}

// GatewayBurn is a paid mutator transaction binding the contract method 0x8ec3dbb9.
//
// Solidity: function gatewayBurn(bytes calldataBytes, bytes signature) returns()
func (_GatewayWallet *GatewayWalletSession) GatewayBurn(calldataBytes []byte, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.GatewayBurn(&_GatewayWallet.TransactOpts, calldataBytes, signature)
}

// GatewayBurn is a paid mutator transaction binding the contract method 0x8ec3dbb9.
//
// Solidity: function gatewayBurn(bytes calldataBytes, bytes signature) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) GatewayBurn(calldataBytes []byte, signature []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.GatewayBurn(&_GatewayWallet.TransactOpts, calldataBytes, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xf49ebd82.
//
// Solidity: function initialize(address pauser_, address denylister_, address[] supportedTokens_, uint32 domain_, uint256 withdrawalDelay_, address burnSigner_, address feeRecipient_) returns()
func (_GatewayWallet *GatewayWalletTransactor) Initialize(opts *bind.TransactOpts, pauser_ common.Address, denylister_ common.Address, supportedTokens_ []common.Address, domain_ uint32, withdrawalDelay_ *big.Int, burnSigner_ common.Address, feeRecipient_ common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "initialize", pauser_, denylister_, supportedTokens_, domain_, withdrawalDelay_, burnSigner_, feeRecipient_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf49ebd82.
//
// Solidity: function initialize(address pauser_, address denylister_, address[] supportedTokens_, uint32 domain_, uint256 withdrawalDelay_, address burnSigner_, address feeRecipient_) returns()
func (_GatewayWallet *GatewayWalletSession) Initialize(pauser_ common.Address, denylister_ common.Address, supportedTokens_ []common.Address, domain_ uint32, withdrawalDelay_ *big.Int, burnSigner_ common.Address, feeRecipient_ common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Initialize(&_GatewayWallet.TransactOpts, pauser_, denylister_, supportedTokens_, domain_, withdrawalDelay_, burnSigner_, feeRecipient_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf49ebd82.
//
// Solidity: function initialize(address pauser_, address denylister_, address[] supportedTokens_, uint32 domain_, uint256 withdrawalDelay_, address burnSigner_, address feeRecipient_) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) Initialize(pauser_ common.Address, denylister_ common.Address, supportedTokens_ []common.Address, domain_ uint32, withdrawalDelay_ *big.Int, burnSigner_ common.Address, feeRecipient_ common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Initialize(&_GatewayWallet.TransactOpts, pauser_, denylister_, supportedTokens_, domain_, withdrawalDelay_, burnSigner_, feeRecipient_)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc8393ba9.
//
// Solidity: function initiateWithdrawal(address token, uint256 value) returns()
func (_GatewayWallet *GatewayWalletTransactor) InitiateWithdrawal(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "initiateWithdrawal", token, value)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc8393ba9.
//
// Solidity: function initiateWithdrawal(address token, uint256 value) returns()
func (_GatewayWallet *GatewayWalletSession) InitiateWithdrawal(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.InitiateWithdrawal(&_GatewayWallet.TransactOpts, token, value)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc8393ba9.
//
// Solidity: function initiateWithdrawal(address token, uint256 value) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) InitiateWithdrawal(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.InitiateWithdrawal(&_GatewayWallet.TransactOpts, token, value)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayWallet *GatewayWalletTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayWallet *GatewayWalletSession) Pause() (*types.Transaction, error) {
	return _GatewayWallet.Contract.Pause(&_GatewayWallet.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayWallet *GatewayWalletTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayWallet.Contract.Pause(&_GatewayWallet.TransactOpts)
}

// RemoveBurnSigner is a paid mutator transaction binding the contract method 0x7e2acfa3.
//
// Solidity: function removeBurnSigner(address signer) returns()
func (_GatewayWallet *GatewayWalletTransactor) RemoveBurnSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "removeBurnSigner", signer)
}

// RemoveBurnSigner is a paid mutator transaction binding the contract method 0x7e2acfa3.
//
// Solidity: function removeBurnSigner(address signer) returns()
func (_GatewayWallet *GatewayWalletSession) RemoveBurnSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.RemoveBurnSigner(&_GatewayWallet.TransactOpts, signer)
}

// RemoveBurnSigner is a paid mutator transaction binding the contract method 0x7e2acfa3.
//
// Solidity: function removeBurnSigner(address signer) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) RemoveBurnSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.RemoveBurnSigner(&_GatewayWallet.TransactOpts, signer)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x020d308d.
//
// Solidity: function removeDelegate(address token, address delegate) returns()
func (_GatewayWallet *GatewayWalletTransactor) RemoveDelegate(opts *bind.TransactOpts, token common.Address, delegate common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "removeDelegate", token, delegate)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x020d308d.
//
// Solidity: function removeDelegate(address token, address delegate) returns()
func (_GatewayWallet *GatewayWalletSession) RemoveDelegate(token common.Address, delegate common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.RemoveDelegate(&_GatewayWallet.TransactOpts, token, delegate)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x020d308d.
//
// Solidity: function removeDelegate(address token, address delegate) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) RemoveDelegate(token common.Address, delegate common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.RemoveDelegate(&_GatewayWallet.TransactOpts, token, delegate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayWallet *GatewayWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayWallet *GatewayWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayWallet.Contract.RenounceOwnership(&_GatewayWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayWallet *GatewayWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayWallet.Contract.RenounceOwnership(&_GatewayWallet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayWallet *GatewayWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayWallet *GatewayWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.TransferOwnership(&_GatewayWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.TransferOwnership(&_GatewayWallet.TransactOpts, newOwner)
}

// UnDenylist is a paid mutator transaction binding the contract method 0x9cab0c1c.
//
// Solidity: function unDenylist(address addr) returns()
func (_GatewayWallet *GatewayWalletTransactor) UnDenylist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "unDenylist", addr)
}

// UnDenylist is a paid mutator transaction binding the contract method 0x9cab0c1c.
//
// Solidity: function unDenylist(address addr) returns()
func (_GatewayWallet *GatewayWalletSession) UnDenylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UnDenylist(&_GatewayWallet.TransactOpts, addr)
}

// UnDenylist is a paid mutator transaction binding the contract method 0x9cab0c1c.
//
// Solidity: function unDenylist(address addr) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) UnDenylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UnDenylist(&_GatewayWallet.TransactOpts, addr)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayWallet *GatewayWalletTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayWallet *GatewayWalletSession) Unpause() (*types.Transaction, error) {
	return _GatewayWallet.Contract.Unpause(&_GatewayWallet.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayWallet *GatewayWalletTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayWallet.Contract.Unpause(&_GatewayWallet.TransactOpts)
}

// UpdateDenylister is a paid mutator transaction binding the contract method 0xa946de04.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (_GatewayWallet *GatewayWalletTransactor) UpdateDenylister(opts *bind.TransactOpts, newDenylister common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "updateDenylister", newDenylister)
}

// UpdateDenylister is a paid mutator transaction binding the contract method 0xa946de04.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (_GatewayWallet *GatewayWalletSession) UpdateDenylister(newDenylister common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdateDenylister(&_GatewayWallet.TransactOpts, newDenylister)
}

// UpdateDenylister is a paid mutator transaction binding the contract method 0xa946de04.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) UpdateDenylister(newDenylister common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdateDenylister(&_GatewayWallet.TransactOpts, newDenylister)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_GatewayWallet *GatewayWalletTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "updateFeeRecipient", newFeeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_GatewayWallet *GatewayWalletSession) UpdateFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdateFeeRecipient(&_GatewayWallet.TransactOpts, newFeeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address newFeeRecipient) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) UpdateFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdateFeeRecipient(&_GatewayWallet.TransactOpts, newFeeRecipient)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address newPauser) returns()
func (_GatewayWallet *GatewayWalletTransactor) UpdatePauser(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "updatePauser", newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address newPauser) returns()
func (_GatewayWallet *GatewayWalletSession) UpdatePauser(newPauser common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdatePauser(&_GatewayWallet.TransactOpts, newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address newPauser) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) UpdatePauser(newPauser common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdatePauser(&_GatewayWallet.TransactOpts, newPauser)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x52cd67e9.
//
// Solidity: function updateWithdrawalDelay(uint256 newDelay) returns()
func (_GatewayWallet *GatewayWalletTransactor) UpdateWithdrawalDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "updateWithdrawalDelay", newDelay)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x52cd67e9.
//
// Solidity: function updateWithdrawalDelay(uint256 newDelay) returns()
func (_GatewayWallet *GatewayWalletSession) UpdateWithdrawalDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdateWithdrawalDelay(&_GatewayWallet.TransactOpts, newDelay)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x52cd67e9.
//
// Solidity: function updateWithdrawalDelay(uint256 newDelay) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) UpdateWithdrawalDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpdateWithdrawalDelay(&_GatewayWallet.TransactOpts, newDelay)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayWallet *GatewayWalletTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayWallet *GatewayWalletSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpgradeToAndCall(&_GatewayWallet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayWallet *GatewayWalletTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayWallet.Contract.UpgradeToAndCall(&_GatewayWallet.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_GatewayWallet *GatewayWalletTransactor) Withdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GatewayWallet.contract.Transact(opts, "withdraw", token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_GatewayWallet *GatewayWalletSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Withdraw(&_GatewayWallet.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_GatewayWallet *GatewayWalletTransactorSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _GatewayWallet.Contract.Withdraw(&_GatewayWallet.TransactOpts, token)
}

// GatewayWalletBurnSignerAddedIterator is returned from FilterBurnSignerAdded and is used to iterate over the raw logs and unpacked data for BurnSignerAdded events raised by the GatewayWallet contract.
type GatewayWalletBurnSignerAddedIterator struct {
	Event *GatewayWalletBurnSignerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletBurnSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletBurnSignerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletBurnSignerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletBurnSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletBurnSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletBurnSignerAdded represents a BurnSignerAdded event raised by the GatewayWallet contract.
type GatewayWalletBurnSignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnSignerAdded is a free log retrieval operation binding the contract event 0x895cdec14fd06461967f67fd12f3dc81c22b59fe80d1caf6f07b1902a8eaa35e.
//
// Solidity: event BurnSignerAdded(address indexed signer)
func (_GatewayWallet *GatewayWalletFilterer) FilterBurnSignerAdded(opts *bind.FilterOpts, signer []common.Address) (*GatewayWalletBurnSignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "BurnSignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletBurnSignerAddedIterator{contract: _GatewayWallet.contract, event: "BurnSignerAdded", logs: logs, sub: sub}, nil
}

// WatchBurnSignerAdded is a free log subscription operation binding the contract event 0x895cdec14fd06461967f67fd12f3dc81c22b59fe80d1caf6f07b1902a8eaa35e.
//
// Solidity: event BurnSignerAdded(address indexed signer)
func (_GatewayWallet *GatewayWalletFilterer) WatchBurnSignerAdded(opts *bind.WatchOpts, sink chan<- *GatewayWalletBurnSignerAdded, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "BurnSignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletBurnSignerAdded)
				if err := _GatewayWallet.contract.UnpackLog(event, "BurnSignerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurnSignerAdded is a log parse operation binding the contract event 0x895cdec14fd06461967f67fd12f3dc81c22b59fe80d1caf6f07b1902a8eaa35e.
//
// Solidity: event BurnSignerAdded(address indexed signer)
func (_GatewayWallet *GatewayWalletFilterer) ParseBurnSignerAdded(log types.Log) (*GatewayWalletBurnSignerAdded, error) {
	event := new(GatewayWalletBurnSignerAdded)
	if err := _GatewayWallet.contract.UnpackLog(event, "BurnSignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletBurnSignerRemovedIterator is returned from FilterBurnSignerRemoved and is used to iterate over the raw logs and unpacked data for BurnSignerRemoved events raised by the GatewayWallet contract.
type GatewayWalletBurnSignerRemovedIterator struct {
	Event *GatewayWalletBurnSignerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletBurnSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletBurnSignerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletBurnSignerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletBurnSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletBurnSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletBurnSignerRemoved represents a BurnSignerRemoved event raised by the GatewayWallet contract.
type GatewayWalletBurnSignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnSignerRemoved is a free log retrieval operation binding the contract event 0x61bf1e503074577bd48a3a903d7b74860236be7560e7324f6614d093a166ff45.
//
// Solidity: event BurnSignerRemoved(address indexed signer)
func (_GatewayWallet *GatewayWalletFilterer) FilterBurnSignerRemoved(opts *bind.FilterOpts, signer []common.Address) (*GatewayWalletBurnSignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "BurnSignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletBurnSignerRemovedIterator{contract: _GatewayWallet.contract, event: "BurnSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchBurnSignerRemoved is a free log subscription operation binding the contract event 0x61bf1e503074577bd48a3a903d7b74860236be7560e7324f6614d093a166ff45.
//
// Solidity: event BurnSignerRemoved(address indexed signer)
func (_GatewayWallet *GatewayWalletFilterer) WatchBurnSignerRemoved(opts *bind.WatchOpts, sink chan<- *GatewayWalletBurnSignerRemoved, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "BurnSignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletBurnSignerRemoved)
				if err := _GatewayWallet.contract.UnpackLog(event, "BurnSignerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurnSignerRemoved is a log parse operation binding the contract event 0x61bf1e503074577bd48a3a903d7b74860236be7560e7324f6614d093a166ff45.
//
// Solidity: event BurnSignerRemoved(address indexed signer)
func (_GatewayWallet *GatewayWalletFilterer) ParseBurnSignerRemoved(log types.Log) (*GatewayWalletBurnSignerRemoved, error) {
	event := new(GatewayWalletBurnSignerRemoved)
	if err := _GatewayWallet.contract.UnpackLog(event, "BurnSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletDelegateAddedIterator is returned from FilterDelegateAdded and is used to iterate over the raw logs and unpacked data for DelegateAdded events raised by the GatewayWallet contract.
type GatewayWalletDelegateAddedIterator struct {
	Event *GatewayWalletDelegateAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletDelegateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletDelegateAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletDelegateAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletDelegateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletDelegateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletDelegateAdded represents a DelegateAdded event raised by the GatewayWallet contract.
type GatewayWalletDelegateAdded struct {
	Token     common.Address
	Depositor common.Address
	Delegate  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateAdded is a free log retrieval operation binding the contract event 0xe91cba149cb264c22b997b1a69e785ad1562a07f83318cd0c985ba1be6e66099.
//
// Solidity: event DelegateAdded(address indexed token, address indexed depositor, address delegate)
func (_GatewayWallet *GatewayWalletFilterer) FilterDelegateAdded(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*GatewayWalletDelegateAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "DelegateAdded", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletDelegateAddedIterator{contract: _GatewayWallet.contract, event: "DelegateAdded", logs: logs, sub: sub}, nil
}

// WatchDelegateAdded is a free log subscription operation binding the contract event 0xe91cba149cb264c22b997b1a69e785ad1562a07f83318cd0c985ba1be6e66099.
//
// Solidity: event DelegateAdded(address indexed token, address indexed depositor, address delegate)
func (_GatewayWallet *GatewayWalletFilterer) WatchDelegateAdded(opts *bind.WatchOpts, sink chan<- *GatewayWalletDelegateAdded, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "DelegateAdded", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletDelegateAdded)
				if err := _GatewayWallet.contract.UnpackLog(event, "DelegateAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateAdded is a log parse operation binding the contract event 0xe91cba149cb264c22b997b1a69e785ad1562a07f83318cd0c985ba1be6e66099.
//
// Solidity: event DelegateAdded(address indexed token, address indexed depositor, address delegate)
func (_GatewayWallet *GatewayWalletFilterer) ParseDelegateAdded(log types.Log) (*GatewayWalletDelegateAdded, error) {
	event := new(GatewayWalletDelegateAdded)
	if err := _GatewayWallet.contract.UnpackLog(event, "DelegateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletDelegateRemovedIterator is returned from FilterDelegateRemoved and is used to iterate over the raw logs and unpacked data for DelegateRemoved events raised by the GatewayWallet contract.
type GatewayWalletDelegateRemovedIterator struct {
	Event *GatewayWalletDelegateRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletDelegateRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletDelegateRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletDelegateRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletDelegateRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletDelegateRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletDelegateRemoved represents a DelegateRemoved event raised by the GatewayWallet contract.
type GatewayWalletDelegateRemoved struct {
	Token     common.Address
	Depositor common.Address
	Delegate  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateRemoved is a free log retrieval operation binding the contract event 0xf706b3b70da5d5dada9b3c7d9f14e4f6c280183742a82ea2f951653070b0f195.
//
// Solidity: event DelegateRemoved(address indexed token, address indexed depositor, address delegate)
func (_GatewayWallet *GatewayWalletFilterer) FilterDelegateRemoved(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*GatewayWalletDelegateRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "DelegateRemoved", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletDelegateRemovedIterator{contract: _GatewayWallet.contract, event: "DelegateRemoved", logs: logs, sub: sub}, nil
}

// WatchDelegateRemoved is a free log subscription operation binding the contract event 0xf706b3b70da5d5dada9b3c7d9f14e4f6c280183742a82ea2f951653070b0f195.
//
// Solidity: event DelegateRemoved(address indexed token, address indexed depositor, address delegate)
func (_GatewayWallet *GatewayWalletFilterer) WatchDelegateRemoved(opts *bind.WatchOpts, sink chan<- *GatewayWalletDelegateRemoved, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "DelegateRemoved", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletDelegateRemoved)
				if err := _GatewayWallet.contract.UnpackLog(event, "DelegateRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateRemoved is a log parse operation binding the contract event 0xf706b3b70da5d5dada9b3c7d9f14e4f6c280183742a82ea2f951653070b0f195.
//
// Solidity: event DelegateRemoved(address indexed token, address indexed depositor, address delegate)
func (_GatewayWallet *GatewayWalletFilterer) ParseDelegateRemoved(log types.Log) (*GatewayWalletDelegateRemoved, error) {
	event := new(GatewayWalletDelegateRemoved)
	if err := _GatewayWallet.contract.UnpackLog(event, "DelegateRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletDenylistedIterator is returned from FilterDenylisted and is used to iterate over the raw logs and unpacked data for Denylisted events raised by the GatewayWallet contract.
type GatewayWalletDenylistedIterator struct {
	Event *GatewayWalletDenylisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletDenylistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletDenylisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletDenylisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletDenylistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletDenylistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletDenylisted represents a Denylisted event raised by the GatewayWallet contract.
type GatewayWalletDenylisted struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDenylisted is a free log retrieval operation binding the contract event 0xfa4507bc1f9c730e6e95897024f1fe7d576cf2deb53579d55c14f1ac3439e114.
//
// Solidity: event Denylisted(address indexed addr)
func (_GatewayWallet *GatewayWalletFilterer) FilterDenylisted(opts *bind.FilterOpts, addr []common.Address) (*GatewayWalletDenylistedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "Denylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletDenylistedIterator{contract: _GatewayWallet.contract, event: "Denylisted", logs: logs, sub: sub}, nil
}

// WatchDenylisted is a free log subscription operation binding the contract event 0xfa4507bc1f9c730e6e95897024f1fe7d576cf2deb53579d55c14f1ac3439e114.
//
// Solidity: event Denylisted(address indexed addr)
func (_GatewayWallet *GatewayWalletFilterer) WatchDenylisted(opts *bind.WatchOpts, sink chan<- *GatewayWalletDenylisted, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "Denylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletDenylisted)
				if err := _GatewayWallet.contract.UnpackLog(event, "Denylisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDenylisted is a log parse operation binding the contract event 0xfa4507bc1f9c730e6e95897024f1fe7d576cf2deb53579d55c14f1ac3439e114.
//
// Solidity: event Denylisted(address indexed addr)
func (_GatewayWallet *GatewayWalletFilterer) ParseDenylisted(log types.Log) (*GatewayWalletDenylisted, error) {
	event := new(GatewayWalletDenylisted)
	if err := _GatewayWallet.contract.UnpackLog(event, "Denylisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletDenylisterChangedIterator is returned from FilterDenylisterChanged and is used to iterate over the raw logs and unpacked data for DenylisterChanged events raised by the GatewayWallet contract.
type GatewayWalletDenylisterChangedIterator struct {
	Event *GatewayWalletDenylisterChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletDenylisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletDenylisterChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletDenylisterChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletDenylisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletDenylisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletDenylisterChanged represents a DenylisterChanged event raised by the GatewayWallet contract.
type GatewayWalletDenylisterChanged struct {
	OldDenylister common.Address
	NewDenylister common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDenylisterChanged is a free log retrieval operation binding the contract event 0xe144e84038182cefebda68c192c222085b2c12a85d135d3c938498c0165c01d3.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (_GatewayWallet *GatewayWalletFilterer) FilterDenylisterChanged(opts *bind.FilterOpts, oldDenylister []common.Address, newDenylister []common.Address) (*GatewayWalletDenylisterChangedIterator, error) {

	var oldDenylisterRule []interface{}
	for _, oldDenylisterItem := range oldDenylister {
		oldDenylisterRule = append(oldDenylisterRule, oldDenylisterItem)
	}
	var newDenylisterRule []interface{}
	for _, newDenylisterItem := range newDenylister {
		newDenylisterRule = append(newDenylisterRule, newDenylisterItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "DenylisterChanged", oldDenylisterRule, newDenylisterRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletDenylisterChangedIterator{contract: _GatewayWallet.contract, event: "DenylisterChanged", logs: logs, sub: sub}, nil
}

// WatchDenylisterChanged is a free log subscription operation binding the contract event 0xe144e84038182cefebda68c192c222085b2c12a85d135d3c938498c0165c01d3.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (_GatewayWallet *GatewayWalletFilterer) WatchDenylisterChanged(opts *bind.WatchOpts, sink chan<- *GatewayWalletDenylisterChanged, oldDenylister []common.Address, newDenylister []common.Address) (event.Subscription, error) {

	var oldDenylisterRule []interface{}
	for _, oldDenylisterItem := range oldDenylister {
		oldDenylisterRule = append(oldDenylisterRule, oldDenylisterItem)
	}
	var newDenylisterRule []interface{}
	for _, newDenylisterItem := range newDenylister {
		newDenylisterRule = append(newDenylisterRule, newDenylisterItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "DenylisterChanged", oldDenylisterRule, newDenylisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletDenylisterChanged)
				if err := _GatewayWallet.contract.UnpackLog(event, "DenylisterChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDenylisterChanged is a log parse operation binding the contract event 0xe144e84038182cefebda68c192c222085b2c12a85d135d3c938498c0165c01d3.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (_GatewayWallet *GatewayWalletFilterer) ParseDenylisterChanged(log types.Log) (*GatewayWalletDenylisterChanged, error) {
	event := new(GatewayWalletDenylisterChanged)
	if err := _GatewayWallet.contract.UnpackLog(event, "DenylisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayWallet contract.
type GatewayWalletDepositedIterator struct {
	Event *GatewayWalletDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletDeposited represents a Deposited event raised by the GatewayWallet contract.
type GatewayWalletDeposited struct {
	Token     common.Address
	Depositor common.Address
	Sender    common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x4174a9435a04d04d274c76779cad136a41fde6937c56241c09ab9d3c7064a1a9.
//
// Solidity: event Deposited(address indexed token, address indexed depositor, address indexed sender, uint256 value)
func (_GatewayWallet *GatewayWalletFilterer) FilterDeposited(opts *bind.FilterOpts, token []common.Address, depositor []common.Address, sender []common.Address) (*GatewayWalletDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "Deposited", tokenRule, depositorRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletDepositedIterator{contract: _GatewayWallet.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x4174a9435a04d04d274c76779cad136a41fde6937c56241c09ab9d3c7064a1a9.
//
// Solidity: event Deposited(address indexed token, address indexed depositor, address indexed sender, uint256 value)
func (_GatewayWallet *GatewayWalletFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayWalletDeposited, token []common.Address, depositor []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "Deposited", tokenRule, depositorRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletDeposited)
				if err := _GatewayWallet.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x4174a9435a04d04d274c76779cad136a41fde6937c56241c09ab9d3c7064a1a9.
//
// Solidity: event Deposited(address indexed token, address indexed depositor, address indexed sender, uint256 value)
func (_GatewayWallet *GatewayWalletFilterer) ParseDeposited(log types.Log) (*GatewayWalletDeposited, error) {
	event := new(GatewayWalletDeposited)
	if err := _GatewayWallet.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the GatewayWallet contract.
type GatewayWalletEIP712DomainChangedIterator struct {
	Event *GatewayWalletEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletEIP712DomainChanged represents a EIP712DomainChanged event raised by the GatewayWallet contract.
type GatewayWalletEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GatewayWallet *GatewayWalletFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*GatewayWalletEIP712DomainChangedIterator, error) {

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &GatewayWalletEIP712DomainChangedIterator{contract: _GatewayWallet.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GatewayWallet *GatewayWalletFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *GatewayWalletEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletEIP712DomainChanged)
				if err := _GatewayWallet.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GatewayWallet *GatewayWalletFilterer) ParseEIP712DomainChanged(log types.Log) (*GatewayWalletEIP712DomainChanged, error) {
	event := new(GatewayWalletEIP712DomainChanged)
	if err := _GatewayWallet.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletFeeRecipientChangedIterator is returned from FilterFeeRecipientChanged and is used to iterate over the raw logs and unpacked data for FeeRecipientChanged events raised by the GatewayWallet contract.
type GatewayWalletFeeRecipientChangedIterator struct {
	Event *GatewayWalletFeeRecipientChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletFeeRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletFeeRecipientChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletFeeRecipientChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletFeeRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletFeeRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletFeeRecipientChanged represents a FeeRecipientChanged event raised by the GatewayWallet contract.
type GatewayWalletFeeRecipientChanged struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientChanged is a free log retrieval operation binding the contract event 0x0bc21fe5c3ab742ff1d15b5c4477ffbacf1167e618228078fa625edebe7f331d.
//
// Solidity: event FeeRecipientChanged(address indexed oldFeeRecipient, address indexed newFeeRecipient)
func (_GatewayWallet *GatewayWalletFilterer) FilterFeeRecipientChanged(opts *bind.FilterOpts, oldFeeRecipient []common.Address, newFeeRecipient []common.Address) (*GatewayWalletFeeRecipientChangedIterator, error) {

	var oldFeeRecipientRule []interface{}
	for _, oldFeeRecipientItem := range oldFeeRecipient {
		oldFeeRecipientRule = append(oldFeeRecipientRule, oldFeeRecipientItem)
	}
	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "FeeRecipientChanged", oldFeeRecipientRule, newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletFeeRecipientChangedIterator{contract: _GatewayWallet.contract, event: "FeeRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientChanged is a free log subscription operation binding the contract event 0x0bc21fe5c3ab742ff1d15b5c4477ffbacf1167e618228078fa625edebe7f331d.
//
// Solidity: event FeeRecipientChanged(address indexed oldFeeRecipient, address indexed newFeeRecipient)
func (_GatewayWallet *GatewayWalletFilterer) WatchFeeRecipientChanged(opts *bind.WatchOpts, sink chan<- *GatewayWalletFeeRecipientChanged, oldFeeRecipient []common.Address, newFeeRecipient []common.Address) (event.Subscription, error) {

	var oldFeeRecipientRule []interface{}
	for _, oldFeeRecipientItem := range oldFeeRecipient {
		oldFeeRecipientRule = append(oldFeeRecipientRule, oldFeeRecipientItem)
	}
	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "FeeRecipientChanged", oldFeeRecipientRule, newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletFeeRecipientChanged)
				if err := _GatewayWallet.contract.UnpackLog(event, "FeeRecipientChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeRecipientChanged is a log parse operation binding the contract event 0x0bc21fe5c3ab742ff1d15b5c4477ffbacf1167e618228078fa625edebe7f331d.
//
// Solidity: event FeeRecipientChanged(address indexed oldFeeRecipient, address indexed newFeeRecipient)
func (_GatewayWallet *GatewayWalletFilterer) ParseFeeRecipientChanged(log types.Log) (*GatewayWalletFeeRecipientChanged, error) {
	event := new(GatewayWalletFeeRecipientChanged)
	if err := _GatewayWallet.contract.UnpackLog(event, "FeeRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletGatewayBurnedIterator is returned from FilterGatewayBurned and is used to iterate over the raw logs and unpacked data for GatewayBurned events raised by the GatewayWallet contract.
type GatewayWalletGatewayBurnedIterator struct {
	Event *GatewayWalletGatewayBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletGatewayBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletGatewayBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletGatewayBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletGatewayBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletGatewayBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletGatewayBurned represents a GatewayBurned event raised by the GatewayWallet contract.
type GatewayWalletGatewayBurned struct {
	Token                common.Address
	Depositor            common.Address
	TransferSpecHash     [32]byte
	DestinationDomain    uint32
	DestinationRecipient [32]byte
	Signer               common.Address
	Value                *big.Int
	Fee                  *big.Int
	FromAvailable        *big.Int
	FromWithdrawing      *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGatewayBurned is a free log retrieval operation binding the contract event 0x12ee2719e7e2dec9f2a0041286b66669153dff0d36719f692b8bbaa4dfe0aa87.
//
// Solidity: event GatewayBurned(address indexed token, address indexed depositor, bytes32 indexed transferSpecHash, uint32 destinationDomain, bytes32 destinationRecipient, address signer, uint256 value, uint256 fee, uint256 fromAvailable, uint256 fromWithdrawing)
func (_GatewayWallet *GatewayWalletFilterer) FilterGatewayBurned(opts *bind.FilterOpts, token []common.Address, depositor []common.Address, transferSpecHash [][32]byte) (*GatewayWalletGatewayBurnedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var transferSpecHashRule []interface{}
	for _, transferSpecHashItem := range transferSpecHash {
		transferSpecHashRule = append(transferSpecHashRule, transferSpecHashItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "GatewayBurned", tokenRule, depositorRule, transferSpecHashRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletGatewayBurnedIterator{contract: _GatewayWallet.contract, event: "GatewayBurned", logs: logs, sub: sub}, nil
}

// WatchGatewayBurned is a free log subscription operation binding the contract event 0x12ee2719e7e2dec9f2a0041286b66669153dff0d36719f692b8bbaa4dfe0aa87.
//
// Solidity: event GatewayBurned(address indexed token, address indexed depositor, bytes32 indexed transferSpecHash, uint32 destinationDomain, bytes32 destinationRecipient, address signer, uint256 value, uint256 fee, uint256 fromAvailable, uint256 fromWithdrawing)
func (_GatewayWallet *GatewayWalletFilterer) WatchGatewayBurned(opts *bind.WatchOpts, sink chan<- *GatewayWalletGatewayBurned, token []common.Address, depositor []common.Address, transferSpecHash [][32]byte) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var transferSpecHashRule []interface{}
	for _, transferSpecHashItem := range transferSpecHash {
		transferSpecHashRule = append(transferSpecHashRule, transferSpecHashItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "GatewayBurned", tokenRule, depositorRule, transferSpecHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletGatewayBurned)
				if err := _GatewayWallet.contract.UnpackLog(event, "GatewayBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGatewayBurned is a log parse operation binding the contract event 0x12ee2719e7e2dec9f2a0041286b66669153dff0d36719f692b8bbaa4dfe0aa87.
//
// Solidity: event GatewayBurned(address indexed token, address indexed depositor, bytes32 indexed transferSpecHash, uint32 destinationDomain, bytes32 destinationRecipient, address signer, uint256 value, uint256 fee, uint256 fromAvailable, uint256 fromWithdrawing)
func (_GatewayWallet *GatewayWalletFilterer) ParseGatewayBurned(log types.Log) (*GatewayWalletGatewayBurned, error) {
	event := new(GatewayWalletGatewayBurned)
	if err := _GatewayWallet.contract.UnpackLog(event, "GatewayBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayWallet contract.
type GatewayWalletInitializedIterator struct {
	Event *GatewayWalletInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletInitialized represents a Initialized event raised by the GatewayWallet contract.
type GatewayWalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayWallet *GatewayWalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayWalletInitializedIterator, error) {

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayWalletInitializedIterator{contract: _GatewayWallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayWallet *GatewayWalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayWalletInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletInitialized)
				if err := _GatewayWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayWallet *GatewayWalletFilterer) ParseInitialized(log types.Log) (*GatewayWalletInitialized, error) {
	event := new(GatewayWalletInitialized)
	if err := _GatewayWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletInsufficientBalanceIterator is returned from FilterInsufficientBalance and is used to iterate over the raw logs and unpacked data for InsufficientBalance events raised by the GatewayWallet contract.
type GatewayWalletInsufficientBalanceIterator struct {
	Event *GatewayWalletInsufficientBalance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletInsufficientBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletInsufficientBalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletInsufficientBalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletInsufficientBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletInsufficientBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletInsufficientBalance represents a InsufficientBalance event raised by the GatewayWallet contract.
type GatewayWalletInsufficientBalance struct {
	Token              common.Address
	Depositor          common.Address
	Value              *big.Int
	AvailableBalance   *big.Int
	WithdrawingBalance *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInsufficientBalance is a free log retrieval operation binding the contract event 0x5fab62745de032ad35a05146c51710fd80871be06cb193722c40f011df352f0c.
//
// Solidity: event InsufficientBalance(address indexed token, address indexed depositor, uint256 value, uint256 availableBalance, uint256 withdrawingBalance)
func (_GatewayWallet *GatewayWalletFilterer) FilterInsufficientBalance(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*GatewayWalletInsufficientBalanceIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "InsufficientBalance", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletInsufficientBalanceIterator{contract: _GatewayWallet.contract, event: "InsufficientBalance", logs: logs, sub: sub}, nil
}

// WatchInsufficientBalance is a free log subscription operation binding the contract event 0x5fab62745de032ad35a05146c51710fd80871be06cb193722c40f011df352f0c.
//
// Solidity: event InsufficientBalance(address indexed token, address indexed depositor, uint256 value, uint256 availableBalance, uint256 withdrawingBalance)
func (_GatewayWallet *GatewayWalletFilterer) WatchInsufficientBalance(opts *bind.WatchOpts, sink chan<- *GatewayWalletInsufficientBalance, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "InsufficientBalance", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletInsufficientBalance)
				if err := _GatewayWallet.contract.UnpackLog(event, "InsufficientBalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInsufficientBalance is a log parse operation binding the contract event 0x5fab62745de032ad35a05146c51710fd80871be06cb193722c40f011df352f0c.
//
// Solidity: event InsufficientBalance(address indexed token, address indexed depositor, uint256 value, uint256 availableBalance, uint256 withdrawingBalance)
func (_GatewayWallet *GatewayWalletFilterer) ParseInsufficientBalance(log types.Log) (*GatewayWalletInsufficientBalance, error) {
	event := new(GatewayWalletInsufficientBalance)
	if err := _GatewayWallet.contract.UnpackLog(event, "InsufficientBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the GatewayWallet contract.
type GatewayWalletOwnershipTransferStartedIterator struct {
	Event *GatewayWalletOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the GatewayWallet contract.
type GatewayWalletOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GatewayWallet *GatewayWalletFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayWalletOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletOwnershipTransferStartedIterator{contract: _GatewayWallet.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GatewayWallet *GatewayWalletFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *GatewayWalletOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletOwnershipTransferStarted)
				if err := _GatewayWallet.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GatewayWallet *GatewayWalletFilterer) ParseOwnershipTransferStarted(log types.Log) (*GatewayWalletOwnershipTransferStarted, error) {
	event := new(GatewayWalletOwnershipTransferStarted)
	if err := _GatewayWallet.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayWallet contract.
type GatewayWalletOwnershipTransferredIterator struct {
	Event *GatewayWalletOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayWallet contract.
type GatewayWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayWallet *GatewayWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletOwnershipTransferredIterator{contract: _GatewayWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayWallet *GatewayWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletOwnershipTransferred)
				if err := _GatewayWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayWallet *GatewayWalletFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayWalletOwnershipTransferred, error) {
	event := new(GatewayWalletOwnershipTransferred)
	if err := _GatewayWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayWallet contract.
type GatewayWalletPausedIterator struct {
	Event *GatewayWalletPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletPaused represents a Paused event raised by the GatewayWallet contract.
type GatewayWalletPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayWallet *GatewayWalletFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayWalletPausedIterator, error) {

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayWalletPausedIterator{contract: _GatewayWallet.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayWallet *GatewayWalletFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayWalletPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletPaused)
				if err := _GatewayWallet.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayWallet *GatewayWalletFilterer) ParsePaused(log types.Log) (*GatewayWalletPaused, error) {
	event := new(GatewayWalletPaused)
	if err := _GatewayWallet.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletPauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the GatewayWallet contract.
type GatewayWalletPauserChangedIterator struct {
	Event *GatewayWalletPauserChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletPauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletPauserChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletPauserChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletPauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletPauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletPauserChanged represents a PauserChanged event raised by the GatewayWallet contract.
type GatewayWalletPauserChanged struct {
	OldPauser common.Address
	NewPauser common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed oldPauser, address indexed newPauser)
func (_GatewayWallet *GatewayWalletFilterer) FilterPauserChanged(opts *bind.FilterOpts, oldPauser []common.Address, newPauser []common.Address) (*GatewayWalletPauserChangedIterator, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "PauserChanged", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletPauserChangedIterator{contract: _GatewayWallet.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed oldPauser, address indexed newPauser)
func (_GatewayWallet *GatewayWalletFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *GatewayWalletPauserChanged, oldPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "PauserChanged", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletPauserChanged)
				if err := _GatewayWallet.contract.UnpackLog(event, "PauserChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserChanged is a log parse operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed oldPauser, address indexed newPauser)
func (_GatewayWallet *GatewayWalletFilterer) ParsePauserChanged(log types.Log) (*GatewayWalletPauserChanged, error) {
	event := new(GatewayWalletPauserChanged)
	if err := _GatewayWallet.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletTokenSupportedIterator is returned from FilterTokenSupported and is used to iterate over the raw logs and unpacked data for TokenSupported events raised by the GatewayWallet contract.
type GatewayWalletTokenSupportedIterator struct {
	Event *GatewayWalletTokenSupported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletTokenSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletTokenSupported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletTokenSupported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletTokenSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletTokenSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletTokenSupported represents a TokenSupported event raised by the GatewayWallet contract.
type GatewayWalletTokenSupported struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenSupported is a free log retrieval operation binding the contract event 0xea3145306a87baeba6bb1a8b5c8d3744f840a81cb436b3509f64fc978600cdfb.
//
// Solidity: event TokenSupported(address token)
func (_GatewayWallet *GatewayWalletFilterer) FilterTokenSupported(opts *bind.FilterOpts) (*GatewayWalletTokenSupportedIterator, error) {

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "TokenSupported")
	if err != nil {
		return nil, err
	}
	return &GatewayWalletTokenSupportedIterator{contract: _GatewayWallet.contract, event: "TokenSupported", logs: logs, sub: sub}, nil
}

// WatchTokenSupported is a free log subscription operation binding the contract event 0xea3145306a87baeba6bb1a8b5c8d3744f840a81cb436b3509f64fc978600cdfb.
//
// Solidity: event TokenSupported(address token)
func (_GatewayWallet *GatewayWalletFilterer) WatchTokenSupported(opts *bind.WatchOpts, sink chan<- *GatewayWalletTokenSupported) (event.Subscription, error) {

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "TokenSupported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletTokenSupported)
				if err := _GatewayWallet.contract.UnpackLog(event, "TokenSupported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSupported is a log parse operation binding the contract event 0xea3145306a87baeba6bb1a8b5c8d3744f840a81cb436b3509f64fc978600cdfb.
//
// Solidity: event TokenSupported(address token)
func (_GatewayWallet *GatewayWalletFilterer) ParseTokenSupported(log types.Log) (*GatewayWalletTokenSupported, error) {
	event := new(GatewayWalletTokenSupported)
	if err := _GatewayWallet.contract.UnpackLog(event, "TokenSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletUnDenylistedIterator is returned from FilterUnDenylisted and is used to iterate over the raw logs and unpacked data for UnDenylisted events raised by the GatewayWallet contract.
type GatewayWalletUnDenylistedIterator struct {
	Event *GatewayWalletUnDenylisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletUnDenylistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletUnDenylisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletUnDenylisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletUnDenylistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletUnDenylistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletUnDenylisted represents a UnDenylisted event raised by the GatewayWallet contract.
type GatewayWalletUnDenylisted struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnDenylisted is a free log retrieval operation binding the contract event 0xc904e1b03de0c20d7fcf9dbd056daf1bd3815e93f251199de815fd0f0b96e166.
//
// Solidity: event UnDenylisted(address indexed addr)
func (_GatewayWallet *GatewayWalletFilterer) FilterUnDenylisted(opts *bind.FilterOpts, addr []common.Address) (*GatewayWalletUnDenylistedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "UnDenylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletUnDenylistedIterator{contract: _GatewayWallet.contract, event: "UnDenylisted", logs: logs, sub: sub}, nil
}

// WatchUnDenylisted is a free log subscription operation binding the contract event 0xc904e1b03de0c20d7fcf9dbd056daf1bd3815e93f251199de815fd0f0b96e166.
//
// Solidity: event UnDenylisted(address indexed addr)
func (_GatewayWallet *GatewayWalletFilterer) WatchUnDenylisted(opts *bind.WatchOpts, sink chan<- *GatewayWalletUnDenylisted, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "UnDenylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletUnDenylisted)
				if err := _GatewayWallet.contract.UnpackLog(event, "UnDenylisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnDenylisted is a log parse operation binding the contract event 0xc904e1b03de0c20d7fcf9dbd056daf1bd3815e93f251199de815fd0f0b96e166.
//
// Solidity: event UnDenylisted(address indexed addr)
func (_GatewayWallet *GatewayWalletFilterer) ParseUnDenylisted(log types.Log) (*GatewayWalletUnDenylisted, error) {
	event := new(GatewayWalletUnDenylisted)
	if err := _GatewayWallet.contract.UnpackLog(event, "UnDenylisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayWallet contract.
type GatewayWalletUnpausedIterator struct {
	Event *GatewayWalletUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletUnpaused represents a Unpaused event raised by the GatewayWallet contract.
type GatewayWalletUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayWallet *GatewayWalletFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayWalletUnpausedIterator, error) {

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayWalletUnpausedIterator{contract: _GatewayWallet.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayWallet *GatewayWalletFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayWalletUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletUnpaused)
				if err := _GatewayWallet.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayWallet *GatewayWalletFilterer) ParseUnpaused(log types.Log) (*GatewayWalletUnpaused, error) {
	event := new(GatewayWalletUnpaused)
	if err := _GatewayWallet.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayWallet contract.
type GatewayWalletUpgradedIterator struct {
	Event *GatewayWalletUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletUpgraded represents a Upgraded event raised by the GatewayWallet contract.
type GatewayWalletUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayWallet *GatewayWalletFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayWalletUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletUpgradedIterator{contract: _GatewayWallet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayWallet *GatewayWalletFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayWalletUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletUpgraded)
				if err := _GatewayWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayWallet *GatewayWalletFilterer) ParseUpgraded(log types.Log) (*GatewayWalletUpgraded, error) {
	event := new(GatewayWalletUpgraded)
	if err := _GatewayWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the GatewayWallet contract.
type GatewayWalletWithdrawalCompletedIterator struct {
	Event *GatewayWalletWithdrawalCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletWithdrawalCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletWithdrawalCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletWithdrawalCompleted represents a WithdrawalCompleted event raised by the GatewayWallet contract.
type GatewayWalletWithdrawalCompleted struct {
	Token     common.Address
	Depositor common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xb00382203b46c3b6ad0a2d7af0268e334bd9406256a7c7ba8f7fc8bc47f8cde9.
//
// Solidity: event WithdrawalCompleted(address indexed token, address indexed depositor, uint256 value)
func (_GatewayWallet *GatewayWalletFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*GatewayWalletWithdrawalCompletedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "WithdrawalCompleted", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletWithdrawalCompletedIterator{contract: _GatewayWallet.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xb00382203b46c3b6ad0a2d7af0268e334bd9406256a7c7ba8f7fc8bc47f8cde9.
//
// Solidity: event WithdrawalCompleted(address indexed token, address indexed depositor, uint256 value)
func (_GatewayWallet *GatewayWalletFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *GatewayWalletWithdrawalCompleted, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "WithdrawalCompleted", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletWithdrawalCompleted)
				if err := _GatewayWallet.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xb00382203b46c3b6ad0a2d7af0268e334bd9406256a7c7ba8f7fc8bc47f8cde9.
//
// Solidity: event WithdrawalCompleted(address indexed token, address indexed depositor, uint256 value)
func (_GatewayWallet *GatewayWalletFilterer) ParseWithdrawalCompleted(log types.Log) (*GatewayWalletWithdrawalCompleted, error) {
	event := new(GatewayWalletWithdrawalCompleted)
	if err := _GatewayWallet.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletWithdrawalDelayChangedIterator is returned from FilterWithdrawalDelayChanged and is used to iterate over the raw logs and unpacked data for WithdrawalDelayChanged events raised by the GatewayWallet contract.
type GatewayWalletWithdrawalDelayChangedIterator struct {
	Event *GatewayWalletWithdrawalDelayChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletWithdrawalDelayChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletWithdrawalDelayChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletWithdrawalDelayChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletWithdrawalDelayChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletWithdrawalDelayChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletWithdrawalDelayChanged represents a WithdrawalDelayChanged event raised by the GatewayWallet contract.
type GatewayWalletWithdrawalDelayChanged struct {
	OldDelay *big.Int
	NewDelay *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayChanged is a free log retrieval operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 indexed oldDelay, uint256 indexed newDelay)
func (_GatewayWallet *GatewayWalletFilterer) FilterWithdrawalDelayChanged(opts *bind.FilterOpts, oldDelay []*big.Int, newDelay []*big.Int) (*GatewayWalletWithdrawalDelayChangedIterator, error) {

	var oldDelayRule []interface{}
	for _, oldDelayItem := range oldDelay {
		oldDelayRule = append(oldDelayRule, oldDelayItem)
	}
	var newDelayRule []interface{}
	for _, newDelayItem := range newDelay {
		newDelayRule = append(newDelayRule, newDelayItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "WithdrawalDelayChanged", oldDelayRule, newDelayRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletWithdrawalDelayChangedIterator{contract: _GatewayWallet.contract, event: "WithdrawalDelayChanged", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayChanged is a free log subscription operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 indexed oldDelay, uint256 indexed newDelay)
func (_GatewayWallet *GatewayWalletFilterer) WatchWithdrawalDelayChanged(opts *bind.WatchOpts, sink chan<- *GatewayWalletWithdrawalDelayChanged, oldDelay []*big.Int, newDelay []*big.Int) (event.Subscription, error) {

	var oldDelayRule []interface{}
	for _, oldDelayItem := range oldDelay {
		oldDelayRule = append(oldDelayRule, oldDelayItem)
	}
	var newDelayRule []interface{}
	for _, newDelayItem := range newDelay {
		newDelayRule = append(newDelayRule, newDelayItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "WithdrawalDelayChanged", oldDelayRule, newDelayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletWithdrawalDelayChanged)
				if err := _GatewayWallet.contract.UnpackLog(event, "WithdrawalDelayChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalDelayChanged is a log parse operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 indexed oldDelay, uint256 indexed newDelay)
func (_GatewayWallet *GatewayWalletFilterer) ParseWithdrawalDelayChanged(log types.Log) (*GatewayWalletWithdrawalDelayChanged, error) {
	event := new(GatewayWalletWithdrawalDelayChanged)
	if err := _GatewayWallet.contract.UnpackLog(event, "WithdrawalDelayChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWalletWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the GatewayWallet contract.
type GatewayWalletWithdrawalInitiatedIterator struct {
	Event *GatewayWalletWithdrawalInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayWalletWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWalletWithdrawalInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayWalletWithdrawalInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayWalletWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWalletWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWalletWithdrawalInitiated represents a WithdrawalInitiated event raised by the GatewayWallet contract.
type GatewayWalletWithdrawalInitiated struct {
	Token              common.Address
	Depositor          common.Address
	Value              *big.Int
	RemainingAvailable *big.Int
	TotalWithdrawing   *big.Int
	WithdrawalBlock    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x5f9a559874d8abe05a98d167b78d2012697505ea3e7bcdba906e7b6084014c65.
//
// Solidity: event WithdrawalInitiated(address indexed token, address indexed depositor, uint256 value, uint256 remainingAvailable, uint256 totalWithdrawing, uint256 withdrawalBlock)
func (_GatewayWallet *GatewayWalletFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*GatewayWalletWithdrawalInitiatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.FilterLogs(opts, "WithdrawalInitiated", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &GatewayWalletWithdrawalInitiatedIterator{contract: _GatewayWallet.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x5f9a559874d8abe05a98d167b78d2012697505ea3e7bcdba906e7b6084014c65.
//
// Solidity: event WithdrawalInitiated(address indexed token, address indexed depositor, uint256 value, uint256 remainingAvailable, uint256 totalWithdrawing, uint256 withdrawalBlock)
func (_GatewayWallet *GatewayWalletFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *GatewayWalletWithdrawalInitiated, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _GatewayWallet.contract.WatchLogs(opts, "WithdrawalInitiated", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWalletWithdrawalInitiated)
				if err := _GatewayWallet.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x5f9a559874d8abe05a98d167b78d2012697505ea3e7bcdba906e7b6084014c65.
//
// Solidity: event WithdrawalInitiated(address indexed token, address indexed depositor, uint256 value, uint256 remainingAvailable, uint256 totalWithdrawing, uint256 withdrawalBlock)
func (_GatewayWallet *GatewayWalletFilterer) ParseWithdrawalInitiated(log types.Log) (*GatewayWalletWithdrawalInitiated, error) {
	event := new(GatewayWalletWithdrawalInitiated)
	if err := _GatewayWallet.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
