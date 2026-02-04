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

// GatewayMintMetaData contains all meta data concerning the GatewayMint contract.
var GatewayMintMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AccountDenylisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxBlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"AttestationExpiredAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"AttestationValueMustBePositiveAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursorOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"attestationCaller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actualCaller\",\"type\":\"address\"}],\"name\":\"InvalidAttestationDestinationCallerAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"attestationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedContract\",\"type\":\"address\"}],\"name\":\"InvalidAttestationDestinationContractAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"attestationDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expectedDomain\",\"type\":\"uint32\"}],\"name\":\"InvalidAttestationDestinationDomainAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestationSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"}],\"name\":\"InvalidAttestationTokenAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"actualMagic\",\"type\":\"bytes4\"}],\"name\":\"InvalidTransferPayloadMagic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"actualMagic\",\"type\":\"bytes4\"}],\"name\":\"InvalidTransferSpecMagic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actualVersion\",\"type\":\"uint32\"}],\"name\":\"InvalidTransferSpecVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedLengthTokenAndTokenMintAuthorities\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustHaveAtLeastOneAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadDataTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedTotalLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTotalLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadOverallLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"actualSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredOffset\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetElementHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"actualSetLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredOffset\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetElementTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes4\",\"name\":\"actualMagic\",\"type\":\"bytes4\"}],\"name\":\"TransferPayloadSetInvalidElementMagic\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedTotalLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTotalLength\",\"type\":\"uint256\"}],\"name\":\"TransferPayloadSetOverallLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferSpecHash\",\"type\":\"bytes32\"}],\"name\":\"TransferSpecHashUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMinimumLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"TransferSpecHeaderTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedTotalLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTotalLength\",\"type\":\"uint256\"}],\"name\":\"TransferSpecOverallLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"UnauthorizedDenylister\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"}],\"name\":\"UnsupportedTokenAtIndex\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AttestationSignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AttestationSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferSpecHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceDepositor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceSigner\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AttestationUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Denylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDenylister\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"DenylisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMintAuthority\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMintAuthority\",\"type\":\"address\"}],\"name\":\"MintAuthorityChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPauser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"UnDenylisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"addAttestationSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"denylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denylister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attestationPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"gatewayMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauser_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"denylister_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"supportedTokens_\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"domain_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"attestationSigner_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokenMintAuthorities_\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isAttestationSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isDenylisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferSpecHash\",\"type\":\"bytes32\"}],\"name\":\"isTransferSpecHashUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"removeAttestationSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenMintAuthority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"unDenylist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDenylister\",\"type\":\"address\"}],\"name\":\"updateDenylister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newMintAuthority\",\"type\":\"address\"}],\"name\":\"updateMintAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// GatewayMintABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMintMetaData.ABI instead.
var GatewayMintABI = GatewayMintMetaData.ABI

// GatewayMint is an auto generated Go binding around an Ethereum contract.
type GatewayMint struct {
	GatewayMintCaller     // Read-only binding to the contract
	GatewayMintTransactor // Write-only binding to the contract
	GatewayMintFilterer   // Log filterer for contract events
}

// GatewayMintCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayMintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayMintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayMintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayMintSession struct {
	Contract     *GatewayMint      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayMintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayMintCallerSession struct {
	Contract *GatewayMintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GatewayMintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayMintTransactorSession struct {
	Contract     *GatewayMintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayMintRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayMintRaw struct {
	Contract *GatewayMint // Generic contract binding to access the raw methods on
}

// GatewayMintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayMintCallerRaw struct {
	Contract *GatewayMintCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayMintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayMintTransactorRaw struct {
	Contract *GatewayMintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayMint creates a new instance of GatewayMint, bound to a specific deployed contract.
func NewGatewayMint(address common.Address, backend bind.ContractBackend) (*GatewayMint, error) {
	contract, err := bindGatewayMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayMint{GatewayMintCaller: GatewayMintCaller{contract: contract}, GatewayMintTransactor: GatewayMintTransactor{contract: contract}, GatewayMintFilterer: GatewayMintFilterer{contract: contract}}, nil
}

// NewGatewayMintCaller creates a new read-only instance of GatewayMint, bound to a specific deployed contract.
func NewGatewayMintCaller(address common.Address, caller bind.ContractCaller) (*GatewayMintCaller, error) {
	contract, err := bindGatewayMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayMintCaller{contract: contract}, nil
}

// NewGatewayMintTransactor creates a new write-only instance of GatewayMint, bound to a specific deployed contract.
func NewGatewayMintTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayMintTransactor, error) {
	contract, err := bindGatewayMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayMintTransactor{contract: contract}, nil
}

// NewGatewayMintFilterer creates a new log filterer instance of GatewayMint, bound to a specific deployed contract.
func NewGatewayMintFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayMintFilterer, error) {
	contract, err := bindGatewayMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayMintFilterer{contract: contract}, nil
}

// bindGatewayMint binds a generic wrapper to an already deployed contract.
func bindGatewayMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMintMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayMint *GatewayMintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayMint.Contract.GatewayMintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayMint *GatewayMintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMint.Contract.GatewayMintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayMint *GatewayMintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayMint.Contract.GatewayMintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayMint *GatewayMintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayMint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayMint *GatewayMintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayMint *GatewayMintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayMint.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayMint *GatewayMintCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayMint *GatewayMintSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayMint.Contract.UPGRADEINTERFACEVERSION(&_GatewayMint.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayMint *GatewayMintCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayMint.Contract.UPGRADEINTERFACEVERSION(&_GatewayMint.CallOpts)
}

// Denylister is a free data retrieval call binding the contract method 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (_GatewayMint *GatewayMintCaller) Denylister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "denylister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Denylister is a free data retrieval call binding the contract method 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (_GatewayMint *GatewayMintSession) Denylister() (common.Address, error) {
	return _GatewayMint.Contract.Denylister(&_GatewayMint.CallOpts)
}

// Denylister is a free data retrieval call binding the contract method 0xbcc76c60.
//
// Solidity: function denylister() view returns(address)
func (_GatewayMint *GatewayMintCallerSession) Denylister() (common.Address, error) {
	return _GatewayMint.Contract.Denylister(&_GatewayMint.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(uint32)
func (_GatewayMint *GatewayMintCaller) Domain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "domain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(uint32)
func (_GatewayMint *GatewayMintSession) Domain() (uint32, error) {
	return _GatewayMint.Contract.Domain(&_GatewayMint.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(uint32)
func (_GatewayMint *GatewayMintCallerSession) Domain() (uint32, error) {
	return _GatewayMint.Contract.Domain(&_GatewayMint.CallOpts)
}

// IsAttestationSigner is a free data retrieval call binding the contract method 0xc418fac3.
//
// Solidity: function isAttestationSigner(address signer) view returns(bool)
func (_GatewayMint *GatewayMintCaller) IsAttestationSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "isAttestationSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAttestationSigner is a free data retrieval call binding the contract method 0xc418fac3.
//
// Solidity: function isAttestationSigner(address signer) view returns(bool)
func (_GatewayMint *GatewayMintSession) IsAttestationSigner(signer common.Address) (bool, error) {
	return _GatewayMint.Contract.IsAttestationSigner(&_GatewayMint.CallOpts, signer)
}

// IsAttestationSigner is a free data retrieval call binding the contract method 0xc418fac3.
//
// Solidity: function isAttestationSigner(address signer) view returns(bool)
func (_GatewayMint *GatewayMintCallerSession) IsAttestationSigner(signer common.Address) (bool, error) {
	return _GatewayMint.Contract.IsAttestationSigner(&_GatewayMint.CallOpts, signer)
}

// IsDenylisted is a free data retrieval call binding the contract method 0xe877a526.
//
// Solidity: function isDenylisted(address addr) view returns(bool)
func (_GatewayMint *GatewayMintCaller) IsDenylisted(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "isDenylisted", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDenylisted is a free data retrieval call binding the contract method 0xe877a526.
//
// Solidity: function isDenylisted(address addr) view returns(bool)
func (_GatewayMint *GatewayMintSession) IsDenylisted(addr common.Address) (bool, error) {
	return _GatewayMint.Contract.IsDenylisted(&_GatewayMint.CallOpts, addr)
}

// IsDenylisted is a free data retrieval call binding the contract method 0xe877a526.
//
// Solidity: function isDenylisted(address addr) view returns(bool)
func (_GatewayMint *GatewayMintCallerSession) IsDenylisted(addr common.Address) (bool, error) {
	return _GatewayMint.Contract.IsDenylisted(&_GatewayMint.CallOpts, addr)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address token) view returns(bool)
func (_GatewayMint *GatewayMintCaller) IsTokenSupported(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "isTokenSupported", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address token) view returns(bool)
func (_GatewayMint *GatewayMintSession) IsTokenSupported(token common.Address) (bool, error) {
	return _GatewayMint.Contract.IsTokenSupported(&_GatewayMint.CallOpts, token)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address token) view returns(bool)
func (_GatewayMint *GatewayMintCallerSession) IsTokenSupported(token common.Address) (bool, error) {
	return _GatewayMint.Contract.IsTokenSupported(&_GatewayMint.CallOpts, token)
}

// IsTransferSpecHashUsed is a free data retrieval call binding the contract method 0x17580158.
//
// Solidity: function isTransferSpecHashUsed(bytes32 transferSpecHash) view returns(bool)
func (_GatewayMint *GatewayMintCaller) IsTransferSpecHashUsed(opts *bind.CallOpts, transferSpecHash [32]byte) (bool, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "isTransferSpecHashUsed", transferSpecHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferSpecHashUsed is a free data retrieval call binding the contract method 0x17580158.
//
// Solidity: function isTransferSpecHashUsed(bytes32 transferSpecHash) view returns(bool)
func (_GatewayMint *GatewayMintSession) IsTransferSpecHashUsed(transferSpecHash [32]byte) (bool, error) {
	return _GatewayMint.Contract.IsTransferSpecHashUsed(&_GatewayMint.CallOpts, transferSpecHash)
}

// IsTransferSpecHashUsed is a free data retrieval call binding the contract method 0x17580158.
//
// Solidity: function isTransferSpecHashUsed(bytes32 transferSpecHash) view returns(bool)
func (_GatewayMint *GatewayMintCallerSession) IsTransferSpecHashUsed(transferSpecHash [32]byte) (bool, error) {
	return _GatewayMint.Contract.IsTransferSpecHashUsed(&_GatewayMint.CallOpts, transferSpecHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayMint *GatewayMintCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayMint *GatewayMintSession) Owner() (common.Address, error) {
	return _GatewayMint.Contract.Owner(&_GatewayMint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayMint *GatewayMintCallerSession) Owner() (common.Address, error) {
	return _GatewayMint.Contract.Owner(&_GatewayMint.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayMint *GatewayMintCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayMint *GatewayMintSession) Paused() (bool, error) {
	return _GatewayMint.Contract.Paused(&_GatewayMint.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayMint *GatewayMintCallerSession) Paused() (bool, error) {
	return _GatewayMint.Contract.Paused(&_GatewayMint.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_GatewayMint *GatewayMintCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_GatewayMint *GatewayMintSession) Pauser() (common.Address, error) {
	return _GatewayMint.Contract.Pauser(&_GatewayMint.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_GatewayMint *GatewayMintCallerSession) Pauser() (common.Address, error) {
	return _GatewayMint.Contract.Pauser(&_GatewayMint.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GatewayMint *GatewayMintCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GatewayMint *GatewayMintSession) PendingOwner() (common.Address, error) {
	return _GatewayMint.Contract.PendingOwner(&_GatewayMint.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_GatewayMint *GatewayMintCallerSession) PendingOwner() (common.Address, error) {
	return _GatewayMint.Contract.PendingOwner(&_GatewayMint.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayMint *GatewayMintCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayMint *GatewayMintSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayMint.Contract.ProxiableUUID(&_GatewayMint.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayMint *GatewayMintCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayMint.Contract.ProxiableUUID(&_GatewayMint.CallOpts)
}

// TokenMintAuthority is a free data retrieval call binding the contract method 0x3e5de4df.
//
// Solidity: function tokenMintAuthority(address token) view returns(address)
func (_GatewayMint *GatewayMintCaller) TokenMintAuthority(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _GatewayMint.contract.Call(opts, &out, "tokenMintAuthority", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMintAuthority is a free data retrieval call binding the contract method 0x3e5de4df.
//
// Solidity: function tokenMintAuthority(address token) view returns(address)
func (_GatewayMint *GatewayMintSession) TokenMintAuthority(token common.Address) (common.Address, error) {
	return _GatewayMint.Contract.TokenMintAuthority(&_GatewayMint.CallOpts, token)
}

// TokenMintAuthority is a free data retrieval call binding the contract method 0x3e5de4df.
//
// Solidity: function tokenMintAuthority(address token) view returns(address)
func (_GatewayMint *GatewayMintCallerSession) TokenMintAuthority(token common.Address) (common.Address, error) {
	return _GatewayMint.Contract.TokenMintAuthority(&_GatewayMint.CallOpts, token)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GatewayMint *GatewayMintTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GatewayMint *GatewayMintSession) AcceptOwnership() (*types.Transaction, error) {
	return _GatewayMint.Contract.AcceptOwnership(&_GatewayMint.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_GatewayMint *GatewayMintTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _GatewayMint.Contract.AcceptOwnership(&_GatewayMint.TransactOpts)
}

// AddAttestationSigner is a paid mutator transaction binding the contract method 0x08307b78.
//
// Solidity: function addAttestationSigner(address signer) returns()
func (_GatewayMint *GatewayMintTransactor) AddAttestationSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "addAttestationSigner", signer)
}

// AddAttestationSigner is a paid mutator transaction binding the contract method 0x08307b78.
//
// Solidity: function addAttestationSigner(address signer) returns()
func (_GatewayMint *GatewayMintSession) AddAttestationSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.AddAttestationSigner(&_GatewayMint.TransactOpts, signer)
}

// AddAttestationSigner is a paid mutator transaction binding the contract method 0x08307b78.
//
// Solidity: function addAttestationSigner(address signer) returns()
func (_GatewayMint *GatewayMintTransactorSession) AddAttestationSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.AddAttestationSigner(&_GatewayMint.TransactOpts, signer)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address token) returns()
func (_GatewayMint *GatewayMintTransactor) AddSupportedToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "addSupportedToken", token)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address token) returns()
func (_GatewayMint *GatewayMintSession) AddSupportedToken(token common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.AddSupportedToken(&_GatewayMint.TransactOpts, token)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address token) returns()
func (_GatewayMint *GatewayMintTransactorSession) AddSupportedToken(token common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.AddSupportedToken(&_GatewayMint.TransactOpts, token)
}

// Denylist is a paid mutator transaction binding the contract method 0x3371bfff.
//
// Solidity: function denylist(address addr) returns()
func (_GatewayMint *GatewayMintTransactor) Denylist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "denylist", addr)
}

// Denylist is a paid mutator transaction binding the contract method 0x3371bfff.
//
// Solidity: function denylist(address addr) returns()
func (_GatewayMint *GatewayMintSession) Denylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.Denylist(&_GatewayMint.TransactOpts, addr)
}

// Denylist is a paid mutator transaction binding the contract method 0x3371bfff.
//
// Solidity: function denylist(address addr) returns()
func (_GatewayMint *GatewayMintTransactorSession) Denylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.Denylist(&_GatewayMint.TransactOpts, addr)
}

// GatewayMint is a paid mutator transaction binding the contract method 0x9fb01cc5.
//
// Solidity: function gatewayMint(bytes attestationPayload, bytes signature) returns()
func (_GatewayMint *GatewayMintTransactor) GatewayMint(opts *bind.TransactOpts, attestationPayload []byte, signature []byte) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "gatewayMint", attestationPayload, signature)
}

// GatewayMint is a paid mutator transaction binding the contract method 0x9fb01cc5.
//
// Solidity: function gatewayMint(bytes attestationPayload, bytes signature) returns()
func (_GatewayMint *GatewayMintSession) GatewayMint(attestationPayload []byte, signature []byte) (*types.Transaction, error) {
	return _GatewayMint.Contract.GatewayMint(&_GatewayMint.TransactOpts, attestationPayload, signature)
}

// GatewayMint is a paid mutator transaction binding the contract method 0x9fb01cc5.
//
// Solidity: function gatewayMint(bytes attestationPayload, bytes signature) returns()
func (_GatewayMint *GatewayMintTransactorSession) GatewayMint(attestationPayload []byte, signature []byte) (*types.Transaction, error) {
	return _GatewayMint.Contract.GatewayMint(&_GatewayMint.TransactOpts, attestationPayload, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a827f6b.
//
// Solidity: function initialize(address pauser_, address denylister_, address[] supportedTokens_, uint32 domain_, address attestationSigner_, address[] tokenMintAuthorities_) returns()
func (_GatewayMint *GatewayMintTransactor) Initialize(opts *bind.TransactOpts, pauser_ common.Address, denylister_ common.Address, supportedTokens_ []common.Address, domain_ uint32, attestationSigner_ common.Address, tokenMintAuthorities_ []common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "initialize", pauser_, denylister_, supportedTokens_, domain_, attestationSigner_, tokenMintAuthorities_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a827f6b.
//
// Solidity: function initialize(address pauser_, address denylister_, address[] supportedTokens_, uint32 domain_, address attestationSigner_, address[] tokenMintAuthorities_) returns()
func (_GatewayMint *GatewayMintSession) Initialize(pauser_ common.Address, denylister_ common.Address, supportedTokens_ []common.Address, domain_ uint32, attestationSigner_ common.Address, tokenMintAuthorities_ []common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.Initialize(&_GatewayMint.TransactOpts, pauser_, denylister_, supportedTokens_, domain_, attestationSigner_, tokenMintAuthorities_)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a827f6b.
//
// Solidity: function initialize(address pauser_, address denylister_, address[] supportedTokens_, uint32 domain_, address attestationSigner_, address[] tokenMintAuthorities_) returns()
func (_GatewayMint *GatewayMintTransactorSession) Initialize(pauser_ common.Address, denylister_ common.Address, supportedTokens_ []common.Address, domain_ uint32, attestationSigner_ common.Address, tokenMintAuthorities_ []common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.Initialize(&_GatewayMint.TransactOpts, pauser_, denylister_, supportedTokens_, domain_, attestationSigner_, tokenMintAuthorities_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayMint *GatewayMintTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayMint *GatewayMintSession) Pause() (*types.Transaction, error) {
	return _GatewayMint.Contract.Pause(&_GatewayMint.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayMint *GatewayMintTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayMint.Contract.Pause(&_GatewayMint.TransactOpts)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x3a57609e.
//
// Solidity: function removeAttestationSigner(address signer) returns()
func (_GatewayMint *GatewayMintTransactor) RemoveAttestationSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "removeAttestationSigner", signer)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x3a57609e.
//
// Solidity: function removeAttestationSigner(address signer) returns()
func (_GatewayMint *GatewayMintSession) RemoveAttestationSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.RemoveAttestationSigner(&_GatewayMint.TransactOpts, signer)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x3a57609e.
//
// Solidity: function removeAttestationSigner(address signer) returns()
func (_GatewayMint *GatewayMintTransactorSession) RemoveAttestationSigner(signer common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.RemoveAttestationSigner(&_GatewayMint.TransactOpts, signer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayMint *GatewayMintTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayMint *GatewayMintSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayMint.Contract.RenounceOwnership(&_GatewayMint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayMint *GatewayMintTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayMint.Contract.RenounceOwnership(&_GatewayMint.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayMint *GatewayMintTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayMint *GatewayMintSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.TransferOwnership(&_GatewayMint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayMint *GatewayMintTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.TransferOwnership(&_GatewayMint.TransactOpts, newOwner)
}

// UnDenylist is a paid mutator transaction binding the contract method 0x9cab0c1c.
//
// Solidity: function unDenylist(address addr) returns()
func (_GatewayMint *GatewayMintTransactor) UnDenylist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "unDenylist", addr)
}

// UnDenylist is a paid mutator transaction binding the contract method 0x9cab0c1c.
//
// Solidity: function unDenylist(address addr) returns()
func (_GatewayMint *GatewayMintSession) UnDenylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UnDenylist(&_GatewayMint.TransactOpts, addr)
}

// UnDenylist is a paid mutator transaction binding the contract method 0x9cab0c1c.
//
// Solidity: function unDenylist(address addr) returns()
func (_GatewayMint *GatewayMintTransactorSession) UnDenylist(addr common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UnDenylist(&_GatewayMint.TransactOpts, addr)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayMint *GatewayMintTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayMint *GatewayMintSession) Unpause() (*types.Transaction, error) {
	return _GatewayMint.Contract.Unpause(&_GatewayMint.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayMint *GatewayMintTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayMint.Contract.Unpause(&_GatewayMint.TransactOpts)
}

// UpdateDenylister is a paid mutator transaction binding the contract method 0xa946de04.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (_GatewayMint *GatewayMintTransactor) UpdateDenylister(opts *bind.TransactOpts, newDenylister common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "updateDenylister", newDenylister)
}

// UpdateDenylister is a paid mutator transaction binding the contract method 0xa946de04.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (_GatewayMint *GatewayMintSession) UpdateDenylister(newDenylister common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpdateDenylister(&_GatewayMint.TransactOpts, newDenylister)
}

// UpdateDenylister is a paid mutator transaction binding the contract method 0xa946de04.
//
// Solidity: function updateDenylister(address newDenylister) returns()
func (_GatewayMint *GatewayMintTransactorSession) UpdateDenylister(newDenylister common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpdateDenylister(&_GatewayMint.TransactOpts, newDenylister)
}

// UpdateMintAuthority is a paid mutator transaction binding the contract method 0x1ad20c51.
//
// Solidity: function updateMintAuthority(address token, address newMintAuthority) returns()
func (_GatewayMint *GatewayMintTransactor) UpdateMintAuthority(opts *bind.TransactOpts, token common.Address, newMintAuthority common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "updateMintAuthority", token, newMintAuthority)
}

// UpdateMintAuthority is a paid mutator transaction binding the contract method 0x1ad20c51.
//
// Solidity: function updateMintAuthority(address token, address newMintAuthority) returns()
func (_GatewayMint *GatewayMintSession) UpdateMintAuthority(token common.Address, newMintAuthority common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpdateMintAuthority(&_GatewayMint.TransactOpts, token, newMintAuthority)
}

// UpdateMintAuthority is a paid mutator transaction binding the contract method 0x1ad20c51.
//
// Solidity: function updateMintAuthority(address token, address newMintAuthority) returns()
func (_GatewayMint *GatewayMintTransactorSession) UpdateMintAuthority(token common.Address, newMintAuthority common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpdateMintAuthority(&_GatewayMint.TransactOpts, token, newMintAuthority)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address newPauser) returns()
func (_GatewayMint *GatewayMintTransactor) UpdatePauser(opts *bind.TransactOpts, newPauser common.Address) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "updatePauser", newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address newPauser) returns()
func (_GatewayMint *GatewayMintSession) UpdatePauser(newPauser common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpdatePauser(&_GatewayMint.TransactOpts, newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address newPauser) returns()
func (_GatewayMint *GatewayMintTransactorSession) UpdatePauser(newPauser common.Address) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpdatePauser(&_GatewayMint.TransactOpts, newPauser)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayMint *GatewayMintTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayMint.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayMint *GatewayMintSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpgradeToAndCall(&_GatewayMint.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayMint *GatewayMintTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayMint.Contract.UpgradeToAndCall(&_GatewayMint.TransactOpts, newImplementation, data)
}

// GatewayMintAttestationSignerAddedIterator is returned from FilterAttestationSignerAdded and is used to iterate over the raw logs and unpacked data for AttestationSignerAdded events raised by the GatewayMint contract.
type GatewayMintAttestationSignerAddedIterator struct {
	Event *GatewayMintAttestationSignerAdded // Event containing the contract specifics and raw log

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
func (it *GatewayMintAttestationSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintAttestationSignerAdded)
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
		it.Event = new(GatewayMintAttestationSignerAdded)
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
func (it *GatewayMintAttestationSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintAttestationSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintAttestationSignerAdded represents a AttestationSignerAdded event raised by the GatewayMint contract.
type GatewayMintAttestationSignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAttestationSignerAdded is a free log retrieval operation binding the contract event 0x09b9e0aeb0d8dd390c8ab392ed51b6cbc0e6f3a2ddae0149cec4dd4acf13d7a9.
//
// Solidity: event AttestationSignerAdded(address indexed signer)
func (_GatewayMint *GatewayMintFilterer) FilterAttestationSignerAdded(opts *bind.FilterOpts, signer []common.Address) (*GatewayMintAttestationSignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "AttestationSignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintAttestationSignerAddedIterator{contract: _GatewayMint.contract, event: "AttestationSignerAdded", logs: logs, sub: sub}, nil
}

// WatchAttestationSignerAdded is a free log subscription operation binding the contract event 0x09b9e0aeb0d8dd390c8ab392ed51b6cbc0e6f3a2ddae0149cec4dd4acf13d7a9.
//
// Solidity: event AttestationSignerAdded(address indexed signer)
func (_GatewayMint *GatewayMintFilterer) WatchAttestationSignerAdded(opts *bind.WatchOpts, sink chan<- *GatewayMintAttestationSignerAdded, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "AttestationSignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintAttestationSignerAdded)
				if err := _GatewayMint.contract.UnpackLog(event, "AttestationSignerAdded", log); err != nil {
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

// ParseAttestationSignerAdded is a log parse operation binding the contract event 0x09b9e0aeb0d8dd390c8ab392ed51b6cbc0e6f3a2ddae0149cec4dd4acf13d7a9.
//
// Solidity: event AttestationSignerAdded(address indexed signer)
func (_GatewayMint *GatewayMintFilterer) ParseAttestationSignerAdded(log types.Log) (*GatewayMintAttestationSignerAdded, error) {
	event := new(GatewayMintAttestationSignerAdded)
	if err := _GatewayMint.contract.UnpackLog(event, "AttestationSignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintAttestationSignerRemovedIterator is returned from FilterAttestationSignerRemoved and is used to iterate over the raw logs and unpacked data for AttestationSignerRemoved events raised by the GatewayMint contract.
type GatewayMintAttestationSignerRemovedIterator struct {
	Event *GatewayMintAttestationSignerRemoved // Event containing the contract specifics and raw log

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
func (it *GatewayMintAttestationSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintAttestationSignerRemoved)
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
		it.Event = new(GatewayMintAttestationSignerRemoved)
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
func (it *GatewayMintAttestationSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintAttestationSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintAttestationSignerRemoved represents a AttestationSignerRemoved event raised by the GatewayMint contract.
type GatewayMintAttestationSignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAttestationSignerRemoved is a free log retrieval operation binding the contract event 0xfdf1ac7bfb61a45532dcc5b24ecc2b3373a97f45b89df6489ac493f873dd6548.
//
// Solidity: event AttestationSignerRemoved(address indexed signer)
func (_GatewayMint *GatewayMintFilterer) FilterAttestationSignerRemoved(opts *bind.FilterOpts, signer []common.Address) (*GatewayMintAttestationSignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "AttestationSignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintAttestationSignerRemovedIterator{contract: _GatewayMint.contract, event: "AttestationSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchAttestationSignerRemoved is a free log subscription operation binding the contract event 0xfdf1ac7bfb61a45532dcc5b24ecc2b3373a97f45b89df6489ac493f873dd6548.
//
// Solidity: event AttestationSignerRemoved(address indexed signer)
func (_GatewayMint *GatewayMintFilterer) WatchAttestationSignerRemoved(opts *bind.WatchOpts, sink chan<- *GatewayMintAttestationSignerRemoved, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "AttestationSignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintAttestationSignerRemoved)
				if err := _GatewayMint.contract.UnpackLog(event, "AttestationSignerRemoved", log); err != nil {
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

// ParseAttestationSignerRemoved is a log parse operation binding the contract event 0xfdf1ac7bfb61a45532dcc5b24ecc2b3373a97f45b89df6489ac493f873dd6548.
//
// Solidity: event AttestationSignerRemoved(address indexed signer)
func (_GatewayMint *GatewayMintFilterer) ParseAttestationSignerRemoved(log types.Log) (*GatewayMintAttestationSignerRemoved, error) {
	event := new(GatewayMintAttestationSignerRemoved)
	if err := _GatewayMint.contract.UnpackLog(event, "AttestationSignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintAttestationUsedIterator is returned from FilterAttestationUsed and is used to iterate over the raw logs and unpacked data for AttestationUsed events raised by the GatewayMint contract.
type GatewayMintAttestationUsedIterator struct {
	Event *GatewayMintAttestationUsed // Event containing the contract specifics and raw log

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
func (it *GatewayMintAttestationUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintAttestationUsed)
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
		it.Event = new(GatewayMintAttestationUsed)
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
func (it *GatewayMintAttestationUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintAttestationUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintAttestationUsed represents a AttestationUsed event raised by the GatewayMint contract.
type GatewayMintAttestationUsed struct {
	Token            common.Address
	Recipient        common.Address
	TransferSpecHash [32]byte
	SourceDomain     uint32
	SourceDepositor  [32]byte
	SourceSigner     [32]byte
	Value            *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAttestationUsed is a free log retrieval operation binding the contract event 0xbb312ce0cc311b2cb0746e09ccd2f91fdb9e2ac755d2f11c65300eb0d0fffd63.
//
// Solidity: event AttestationUsed(address indexed token, address indexed recipient, bytes32 indexed transferSpecHash, uint32 sourceDomain, bytes32 sourceDepositor, bytes32 sourceSigner, uint256 value)
func (_GatewayMint *GatewayMintFilterer) FilterAttestationUsed(opts *bind.FilterOpts, token []common.Address, recipient []common.Address, transferSpecHash [][32]byte) (*GatewayMintAttestationUsedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var transferSpecHashRule []interface{}
	for _, transferSpecHashItem := range transferSpecHash {
		transferSpecHashRule = append(transferSpecHashRule, transferSpecHashItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "AttestationUsed", tokenRule, recipientRule, transferSpecHashRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintAttestationUsedIterator{contract: _GatewayMint.contract, event: "AttestationUsed", logs: logs, sub: sub}, nil
}

// WatchAttestationUsed is a free log subscription operation binding the contract event 0xbb312ce0cc311b2cb0746e09ccd2f91fdb9e2ac755d2f11c65300eb0d0fffd63.
//
// Solidity: event AttestationUsed(address indexed token, address indexed recipient, bytes32 indexed transferSpecHash, uint32 sourceDomain, bytes32 sourceDepositor, bytes32 sourceSigner, uint256 value)
func (_GatewayMint *GatewayMintFilterer) WatchAttestationUsed(opts *bind.WatchOpts, sink chan<- *GatewayMintAttestationUsed, token []common.Address, recipient []common.Address, transferSpecHash [][32]byte) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var transferSpecHashRule []interface{}
	for _, transferSpecHashItem := range transferSpecHash {
		transferSpecHashRule = append(transferSpecHashRule, transferSpecHashItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "AttestationUsed", tokenRule, recipientRule, transferSpecHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintAttestationUsed)
				if err := _GatewayMint.contract.UnpackLog(event, "AttestationUsed", log); err != nil {
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

// ParseAttestationUsed is a log parse operation binding the contract event 0xbb312ce0cc311b2cb0746e09ccd2f91fdb9e2ac755d2f11c65300eb0d0fffd63.
//
// Solidity: event AttestationUsed(address indexed token, address indexed recipient, bytes32 indexed transferSpecHash, uint32 sourceDomain, bytes32 sourceDepositor, bytes32 sourceSigner, uint256 value)
func (_GatewayMint *GatewayMintFilterer) ParseAttestationUsed(log types.Log) (*GatewayMintAttestationUsed, error) {
	event := new(GatewayMintAttestationUsed)
	if err := _GatewayMint.contract.UnpackLog(event, "AttestationUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintDenylistedIterator is returned from FilterDenylisted and is used to iterate over the raw logs and unpacked data for Denylisted events raised by the GatewayMint contract.
type GatewayMintDenylistedIterator struct {
	Event *GatewayMintDenylisted // Event containing the contract specifics and raw log

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
func (it *GatewayMintDenylistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintDenylisted)
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
		it.Event = new(GatewayMintDenylisted)
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
func (it *GatewayMintDenylistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintDenylistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintDenylisted represents a Denylisted event raised by the GatewayMint contract.
type GatewayMintDenylisted struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDenylisted is a free log retrieval operation binding the contract event 0xfa4507bc1f9c730e6e95897024f1fe7d576cf2deb53579d55c14f1ac3439e114.
//
// Solidity: event Denylisted(address indexed addr)
func (_GatewayMint *GatewayMintFilterer) FilterDenylisted(opts *bind.FilterOpts, addr []common.Address) (*GatewayMintDenylistedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "Denylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintDenylistedIterator{contract: _GatewayMint.contract, event: "Denylisted", logs: logs, sub: sub}, nil
}

// WatchDenylisted is a free log subscription operation binding the contract event 0xfa4507bc1f9c730e6e95897024f1fe7d576cf2deb53579d55c14f1ac3439e114.
//
// Solidity: event Denylisted(address indexed addr)
func (_GatewayMint *GatewayMintFilterer) WatchDenylisted(opts *bind.WatchOpts, sink chan<- *GatewayMintDenylisted, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "Denylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintDenylisted)
				if err := _GatewayMint.contract.UnpackLog(event, "Denylisted", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseDenylisted(log types.Log) (*GatewayMintDenylisted, error) {
	event := new(GatewayMintDenylisted)
	if err := _GatewayMint.contract.UnpackLog(event, "Denylisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintDenylisterChangedIterator is returned from FilterDenylisterChanged and is used to iterate over the raw logs and unpacked data for DenylisterChanged events raised by the GatewayMint contract.
type GatewayMintDenylisterChangedIterator struct {
	Event *GatewayMintDenylisterChanged // Event containing the contract specifics and raw log

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
func (it *GatewayMintDenylisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintDenylisterChanged)
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
		it.Event = new(GatewayMintDenylisterChanged)
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
func (it *GatewayMintDenylisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintDenylisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintDenylisterChanged represents a DenylisterChanged event raised by the GatewayMint contract.
type GatewayMintDenylisterChanged struct {
	OldDenylister common.Address
	NewDenylister common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDenylisterChanged is a free log retrieval operation binding the contract event 0xe144e84038182cefebda68c192c222085b2c12a85d135d3c938498c0165c01d3.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (_GatewayMint *GatewayMintFilterer) FilterDenylisterChanged(opts *bind.FilterOpts, oldDenylister []common.Address, newDenylister []common.Address) (*GatewayMintDenylisterChangedIterator, error) {

	var oldDenylisterRule []interface{}
	for _, oldDenylisterItem := range oldDenylister {
		oldDenylisterRule = append(oldDenylisterRule, oldDenylisterItem)
	}
	var newDenylisterRule []interface{}
	for _, newDenylisterItem := range newDenylister {
		newDenylisterRule = append(newDenylisterRule, newDenylisterItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "DenylisterChanged", oldDenylisterRule, newDenylisterRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintDenylisterChangedIterator{contract: _GatewayMint.contract, event: "DenylisterChanged", logs: logs, sub: sub}, nil
}

// WatchDenylisterChanged is a free log subscription operation binding the contract event 0xe144e84038182cefebda68c192c222085b2c12a85d135d3c938498c0165c01d3.
//
// Solidity: event DenylisterChanged(address indexed oldDenylister, address indexed newDenylister)
func (_GatewayMint *GatewayMintFilterer) WatchDenylisterChanged(opts *bind.WatchOpts, sink chan<- *GatewayMintDenylisterChanged, oldDenylister []common.Address, newDenylister []common.Address) (event.Subscription, error) {

	var oldDenylisterRule []interface{}
	for _, oldDenylisterItem := range oldDenylister {
		oldDenylisterRule = append(oldDenylisterRule, oldDenylisterItem)
	}
	var newDenylisterRule []interface{}
	for _, newDenylisterItem := range newDenylister {
		newDenylisterRule = append(newDenylisterRule, newDenylisterItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "DenylisterChanged", oldDenylisterRule, newDenylisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintDenylisterChanged)
				if err := _GatewayMint.contract.UnpackLog(event, "DenylisterChanged", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseDenylisterChanged(log types.Log) (*GatewayMintDenylisterChanged, error) {
	event := new(GatewayMintDenylisterChanged)
	if err := _GatewayMint.contract.UnpackLog(event, "DenylisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayMint contract.
type GatewayMintInitializedIterator struct {
	Event *GatewayMintInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayMintInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintInitialized)
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
		it.Event = new(GatewayMintInitialized)
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
func (it *GatewayMintInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintInitialized represents a Initialized event raised by the GatewayMint contract.
type GatewayMintInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayMint *GatewayMintFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayMintInitializedIterator, error) {

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayMintInitializedIterator{contract: _GatewayMint.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayMint *GatewayMintFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayMintInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintInitialized)
				if err := _GatewayMint.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseInitialized(log types.Log) (*GatewayMintInitialized, error) {
	event := new(GatewayMintInitialized)
	if err := _GatewayMint.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintMintAuthorityChangedIterator is returned from FilterMintAuthorityChanged and is used to iterate over the raw logs and unpacked data for MintAuthorityChanged events raised by the GatewayMint contract.
type GatewayMintMintAuthorityChangedIterator struct {
	Event *GatewayMintMintAuthorityChanged // Event containing the contract specifics and raw log

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
func (it *GatewayMintMintAuthorityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintMintAuthorityChanged)
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
		it.Event = new(GatewayMintMintAuthorityChanged)
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
func (it *GatewayMintMintAuthorityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintMintAuthorityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintMintAuthorityChanged represents a MintAuthorityChanged event raised by the GatewayMint contract.
type GatewayMintMintAuthorityChanged struct {
	Token            common.Address
	OldMintAuthority common.Address
	NewMintAuthority common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMintAuthorityChanged is a free log retrieval operation binding the contract event 0x5c3cb1b340c7f4e8dfc2058a7e8992362e063c9ded5e14aa21e23e8856ee1831.
//
// Solidity: event MintAuthorityChanged(address indexed token, address indexed oldMintAuthority, address indexed newMintAuthority)
func (_GatewayMint *GatewayMintFilterer) FilterMintAuthorityChanged(opts *bind.FilterOpts, token []common.Address, oldMintAuthority []common.Address, newMintAuthority []common.Address) (*GatewayMintMintAuthorityChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldMintAuthorityRule []interface{}
	for _, oldMintAuthorityItem := range oldMintAuthority {
		oldMintAuthorityRule = append(oldMintAuthorityRule, oldMintAuthorityItem)
	}
	var newMintAuthorityRule []interface{}
	for _, newMintAuthorityItem := range newMintAuthority {
		newMintAuthorityRule = append(newMintAuthorityRule, newMintAuthorityItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "MintAuthorityChanged", tokenRule, oldMintAuthorityRule, newMintAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintMintAuthorityChangedIterator{contract: _GatewayMint.contract, event: "MintAuthorityChanged", logs: logs, sub: sub}, nil
}

// WatchMintAuthorityChanged is a free log subscription operation binding the contract event 0x5c3cb1b340c7f4e8dfc2058a7e8992362e063c9ded5e14aa21e23e8856ee1831.
//
// Solidity: event MintAuthorityChanged(address indexed token, address indexed oldMintAuthority, address indexed newMintAuthority)
func (_GatewayMint *GatewayMintFilterer) WatchMintAuthorityChanged(opts *bind.WatchOpts, sink chan<- *GatewayMintMintAuthorityChanged, token []common.Address, oldMintAuthority []common.Address, newMintAuthority []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldMintAuthorityRule []interface{}
	for _, oldMintAuthorityItem := range oldMintAuthority {
		oldMintAuthorityRule = append(oldMintAuthorityRule, oldMintAuthorityItem)
	}
	var newMintAuthorityRule []interface{}
	for _, newMintAuthorityItem := range newMintAuthority {
		newMintAuthorityRule = append(newMintAuthorityRule, newMintAuthorityItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "MintAuthorityChanged", tokenRule, oldMintAuthorityRule, newMintAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintMintAuthorityChanged)
				if err := _GatewayMint.contract.UnpackLog(event, "MintAuthorityChanged", log); err != nil {
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

// ParseMintAuthorityChanged is a log parse operation binding the contract event 0x5c3cb1b340c7f4e8dfc2058a7e8992362e063c9ded5e14aa21e23e8856ee1831.
//
// Solidity: event MintAuthorityChanged(address indexed token, address indexed oldMintAuthority, address indexed newMintAuthority)
func (_GatewayMint *GatewayMintFilterer) ParseMintAuthorityChanged(log types.Log) (*GatewayMintMintAuthorityChanged, error) {
	event := new(GatewayMintMintAuthorityChanged)
	if err := _GatewayMint.contract.UnpackLog(event, "MintAuthorityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the GatewayMint contract.
type GatewayMintOwnershipTransferStartedIterator struct {
	Event *GatewayMintOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *GatewayMintOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintOwnershipTransferStarted)
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
		it.Event = new(GatewayMintOwnershipTransferStarted)
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
func (it *GatewayMintOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the GatewayMint contract.
type GatewayMintOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GatewayMint *GatewayMintFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayMintOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintOwnershipTransferStartedIterator{contract: _GatewayMint.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_GatewayMint *GatewayMintFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *GatewayMintOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintOwnershipTransferStarted)
				if err := _GatewayMint.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseOwnershipTransferStarted(log types.Log) (*GatewayMintOwnershipTransferStarted, error) {
	event := new(GatewayMintOwnershipTransferStarted)
	if err := _GatewayMint.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayMint contract.
type GatewayMintOwnershipTransferredIterator struct {
	Event *GatewayMintOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayMintOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintOwnershipTransferred)
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
		it.Event = new(GatewayMintOwnershipTransferred)
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
func (it *GatewayMintOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayMint contract.
type GatewayMintOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayMint *GatewayMintFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayMintOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintOwnershipTransferredIterator{contract: _GatewayMint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayMint *GatewayMintFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayMintOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintOwnershipTransferred)
				if err := _GatewayMint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayMintOwnershipTransferred, error) {
	event := new(GatewayMintOwnershipTransferred)
	if err := _GatewayMint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayMint contract.
type GatewayMintPausedIterator struct {
	Event *GatewayMintPaused // Event containing the contract specifics and raw log

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
func (it *GatewayMintPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintPaused)
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
		it.Event = new(GatewayMintPaused)
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
func (it *GatewayMintPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintPaused represents a Paused event raised by the GatewayMint contract.
type GatewayMintPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayMint *GatewayMintFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayMintPausedIterator, error) {

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayMintPausedIterator{contract: _GatewayMint.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayMint *GatewayMintFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayMintPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintPaused)
				if err := _GatewayMint.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParsePaused(log types.Log) (*GatewayMintPaused, error) {
	event := new(GatewayMintPaused)
	if err := _GatewayMint.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintPauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the GatewayMint contract.
type GatewayMintPauserChangedIterator struct {
	Event *GatewayMintPauserChanged // Event containing the contract specifics and raw log

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
func (it *GatewayMintPauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintPauserChanged)
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
		it.Event = new(GatewayMintPauserChanged)
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
func (it *GatewayMintPauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintPauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintPauserChanged represents a PauserChanged event raised by the GatewayMint contract.
type GatewayMintPauserChanged struct {
	OldPauser common.Address
	NewPauser common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed oldPauser, address indexed newPauser)
func (_GatewayMint *GatewayMintFilterer) FilterPauserChanged(opts *bind.FilterOpts, oldPauser []common.Address, newPauser []common.Address) (*GatewayMintPauserChangedIterator, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "PauserChanged", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintPauserChangedIterator{contract: _GatewayMint.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0x95bb211a5a393c4d30c3edc9a745825fba4e6ad3e3bb949e6bf8ccdfe431a811.
//
// Solidity: event PauserChanged(address indexed oldPauser, address indexed newPauser)
func (_GatewayMint *GatewayMintFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *GatewayMintPauserChanged, oldPauser []common.Address, newPauser []common.Address) (event.Subscription, error) {

	var oldPauserRule []interface{}
	for _, oldPauserItem := range oldPauser {
		oldPauserRule = append(oldPauserRule, oldPauserItem)
	}
	var newPauserRule []interface{}
	for _, newPauserItem := range newPauser {
		newPauserRule = append(newPauserRule, newPauserItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "PauserChanged", oldPauserRule, newPauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintPauserChanged)
				if err := _GatewayMint.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParsePauserChanged(log types.Log) (*GatewayMintPauserChanged, error) {
	event := new(GatewayMintPauserChanged)
	if err := _GatewayMint.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintTokenSupportedIterator is returned from FilterTokenSupported and is used to iterate over the raw logs and unpacked data for TokenSupported events raised by the GatewayMint contract.
type GatewayMintTokenSupportedIterator struct {
	Event *GatewayMintTokenSupported // Event containing the contract specifics and raw log

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
func (it *GatewayMintTokenSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintTokenSupported)
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
		it.Event = new(GatewayMintTokenSupported)
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
func (it *GatewayMintTokenSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintTokenSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintTokenSupported represents a TokenSupported event raised by the GatewayMint contract.
type GatewayMintTokenSupported struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenSupported is a free log retrieval operation binding the contract event 0xea3145306a87baeba6bb1a8b5c8d3744f840a81cb436b3509f64fc978600cdfb.
//
// Solidity: event TokenSupported(address token)
func (_GatewayMint *GatewayMintFilterer) FilterTokenSupported(opts *bind.FilterOpts) (*GatewayMintTokenSupportedIterator, error) {

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "TokenSupported")
	if err != nil {
		return nil, err
	}
	return &GatewayMintTokenSupportedIterator{contract: _GatewayMint.contract, event: "TokenSupported", logs: logs, sub: sub}, nil
}

// WatchTokenSupported is a free log subscription operation binding the contract event 0xea3145306a87baeba6bb1a8b5c8d3744f840a81cb436b3509f64fc978600cdfb.
//
// Solidity: event TokenSupported(address token)
func (_GatewayMint *GatewayMintFilterer) WatchTokenSupported(opts *bind.WatchOpts, sink chan<- *GatewayMintTokenSupported) (event.Subscription, error) {

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "TokenSupported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintTokenSupported)
				if err := _GatewayMint.contract.UnpackLog(event, "TokenSupported", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseTokenSupported(log types.Log) (*GatewayMintTokenSupported, error) {
	event := new(GatewayMintTokenSupported)
	if err := _GatewayMint.contract.UnpackLog(event, "TokenSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintUnDenylistedIterator is returned from FilterUnDenylisted and is used to iterate over the raw logs and unpacked data for UnDenylisted events raised by the GatewayMint contract.
type GatewayMintUnDenylistedIterator struct {
	Event *GatewayMintUnDenylisted // Event containing the contract specifics and raw log

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
func (it *GatewayMintUnDenylistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintUnDenylisted)
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
		it.Event = new(GatewayMintUnDenylisted)
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
func (it *GatewayMintUnDenylistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintUnDenylistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintUnDenylisted represents a UnDenylisted event raised by the GatewayMint contract.
type GatewayMintUnDenylisted struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnDenylisted is a free log retrieval operation binding the contract event 0xc904e1b03de0c20d7fcf9dbd056daf1bd3815e93f251199de815fd0f0b96e166.
//
// Solidity: event UnDenylisted(address indexed addr)
func (_GatewayMint *GatewayMintFilterer) FilterUnDenylisted(opts *bind.FilterOpts, addr []common.Address) (*GatewayMintUnDenylistedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "UnDenylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintUnDenylistedIterator{contract: _GatewayMint.contract, event: "UnDenylisted", logs: logs, sub: sub}, nil
}

// WatchUnDenylisted is a free log subscription operation binding the contract event 0xc904e1b03de0c20d7fcf9dbd056daf1bd3815e93f251199de815fd0f0b96e166.
//
// Solidity: event UnDenylisted(address indexed addr)
func (_GatewayMint *GatewayMintFilterer) WatchUnDenylisted(opts *bind.WatchOpts, sink chan<- *GatewayMintUnDenylisted, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "UnDenylisted", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintUnDenylisted)
				if err := _GatewayMint.contract.UnpackLog(event, "UnDenylisted", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseUnDenylisted(log types.Log) (*GatewayMintUnDenylisted, error) {
	event := new(GatewayMintUnDenylisted)
	if err := _GatewayMint.contract.UnpackLog(event, "UnDenylisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayMint contract.
type GatewayMintUnpausedIterator struct {
	Event *GatewayMintUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayMintUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintUnpaused)
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
		it.Event = new(GatewayMintUnpaused)
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
func (it *GatewayMintUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintUnpaused represents a Unpaused event raised by the GatewayMint contract.
type GatewayMintUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayMint *GatewayMintFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayMintUnpausedIterator, error) {

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayMintUnpausedIterator{contract: _GatewayMint.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayMint *GatewayMintFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayMintUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintUnpaused)
				if err := _GatewayMint.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseUnpaused(log types.Log) (*GatewayMintUnpaused, error) {
	event := new(GatewayMintUnpaused)
	if err := _GatewayMint.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMintUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayMint contract.
type GatewayMintUpgradedIterator struct {
	Event *GatewayMintUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayMintUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMintUpgraded)
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
		it.Event = new(GatewayMintUpgraded)
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
func (it *GatewayMintUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMintUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMintUpgraded represents a Upgraded event raised by the GatewayMint contract.
type GatewayMintUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayMint *GatewayMintFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayMintUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayMint.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMintUpgradedIterator{contract: _GatewayMint.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayMint *GatewayMintFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayMintUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayMint.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMintUpgraded)
				if err := _GatewayMint.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayMint *GatewayMintFilterer) ParseUpgraded(log types.Log) (*GatewayMintUpgraded, error) {
	event := new(GatewayMintUpgraded)
	if err := _GatewayMint.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
