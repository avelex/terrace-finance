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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Amount    *big.Int
	Timestamp *big.Int
}

// StablecoinMetaData contains all meta data concerning the Stablecoin contract.
var StablecoinMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previous_owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_owner\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleMinterChanged\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRootHashUpdated\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DelayUpdated\",\"inputs\":[{\"name\":\"new_delay\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReceivedFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":false},{\"name\":\"from_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"from_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"to_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"to_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SendFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":false},{\"name\":\"from_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"from_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"to_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"to_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RequestWithdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burn_from\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_minter\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer_ownership\",\"inputs\":[{\"name\":\"new_owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounce_ownership\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_minter\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_root_hash\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"execute_strategy\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_execute\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\"},{\"name\":\"selectors\",\"type\":\"bytes[]\"},{\"name\":\"data\",\"type\":\"bytes[]\"},{\"name\":\"proofs\",\"type\":\"bytes32[][]\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_strategy_allowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"encode_erc20_approve\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"encode_supply_aave_v3\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"name\":\"referrer\",\"type\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"strategy_root_hash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interface_id\",\"type\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_role_admin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"admin_role\",\"type\":\"bytes32\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"bytes32\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_withdrawal_delay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"withdrawal_delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pending_withdrawals\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"usdc_reserve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"usdc_usage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"usdc_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_whitelisted_domain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint32\"},{\"name\":\"terrace\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"receive_message\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"attestation\",\"type\":\"bytes\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_all_to_terrace\",\"inputs\":[{\"name\":\"destination_domain\",\"type\":\"uint32\"},{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_to_terrace\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"destination_domain\",\"type\":\"uint32\"},{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_all_to_hub\",\"inputs\":[{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_to_hub\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MESSANGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TRANSMITTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"HUB_DOMAIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"whitelisted_domains\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bridge_nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_staking\",\"inputs\":[{\"name\":\"staking\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"execute_withdraw\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"distribute_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"staking\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"gateway_deposit_and_stake\",\"inputs\":[{\"name\":\"attestation\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"GATEWAY_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"usdc\",\"type\":\"address\"},{\"name\":\"withdrawal_delay\",\"type\":\"uint256\"},{\"name\":\"messanger\",\"type\":\"address\"},{\"name\":\"transmitter\",\"type\":\"address\"},{\"name\":\"hub_domain\",\"type\":\"uint32\"},{\"name\":\"gateway_minter\",\"type\":\"address\"}],\"outputs\":[]}]",
}

// StablecoinABI is the input ABI used to generate the binding from.
// Deprecated: Use StablecoinMetaData.ABI instead.
var StablecoinABI = StablecoinMetaData.ABI

// Stablecoin is an auto generated Go binding around an Ethereum contract.
type Stablecoin struct {
	StablecoinCaller     // Read-only binding to the contract
	StablecoinTransactor // Write-only binding to the contract
	StablecoinFilterer   // Log filterer for contract events
}

// StablecoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type StablecoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StablecoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StablecoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StablecoinSession struct {
	Contract     *Stablecoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StablecoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StablecoinCallerSession struct {
	Contract *StablecoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StablecoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StablecoinTransactorSession struct {
	Contract     *StablecoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StablecoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type StablecoinRaw struct {
	Contract *Stablecoin // Generic contract binding to access the raw methods on
}

// StablecoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StablecoinCallerRaw struct {
	Contract *StablecoinCaller // Generic read-only contract binding to access the raw methods on
}

// StablecoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StablecoinTransactorRaw struct {
	Contract *StablecoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStablecoin creates a new instance of Stablecoin, bound to a specific deployed contract.
func NewStablecoin(address common.Address, backend bind.ContractBackend) (*Stablecoin, error) {
	contract, err := bindStablecoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stablecoin{StablecoinCaller: StablecoinCaller{contract: contract}, StablecoinTransactor: StablecoinTransactor{contract: contract}, StablecoinFilterer: StablecoinFilterer{contract: contract}}, nil
}

// NewStablecoinCaller creates a new read-only instance of Stablecoin, bound to a specific deployed contract.
func NewStablecoinCaller(address common.Address, caller bind.ContractCaller) (*StablecoinCaller, error) {
	contract, err := bindStablecoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StablecoinCaller{contract: contract}, nil
}

// NewStablecoinTransactor creates a new write-only instance of Stablecoin, bound to a specific deployed contract.
func NewStablecoinTransactor(address common.Address, transactor bind.ContractTransactor) (*StablecoinTransactor, error) {
	contract, err := bindStablecoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StablecoinTransactor{contract: contract}, nil
}

// NewStablecoinFilterer creates a new log filterer instance of Stablecoin, bound to a specific deployed contract.
func NewStablecoinFilterer(address common.Address, filterer bind.ContractFilterer) (*StablecoinFilterer, error) {
	contract, err := bindStablecoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StablecoinFilterer{contract: contract}, nil
}

// bindStablecoin binds a generic wrapper to an already deployed contract.
func bindStablecoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StablecoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stablecoin *StablecoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stablecoin.Contract.StablecoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stablecoin *StablecoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stablecoin.Contract.StablecoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stablecoin *StablecoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stablecoin.Contract.StablecoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stablecoin *StablecoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stablecoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stablecoin *StablecoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stablecoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stablecoin *StablecoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stablecoin.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stablecoin *StablecoinCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stablecoin *StablecoinSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stablecoin.Contract.DEFAULTADMINROLE(&_Stablecoin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stablecoin *StablecoinCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stablecoin.Contract.DEFAULTADMINROLE(&_Stablecoin.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Stablecoin *StablecoinCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Stablecoin *StablecoinSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Stablecoin.Contract.DOMAINSEPARATOR(&_Stablecoin.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Stablecoin *StablecoinCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Stablecoin.Contract.DOMAINSEPARATOR(&_Stablecoin.CallOpts)
}

// GATEWAYMINTER is a free data retrieval call binding the contract method 0x6558b981.
//
// Solidity: function GATEWAY_MINTER() view returns(address)
func (_Stablecoin *StablecoinCaller) GATEWAYMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "GATEWAY_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GATEWAYMINTER is a free data retrieval call binding the contract method 0x6558b981.
//
// Solidity: function GATEWAY_MINTER() view returns(address)
func (_Stablecoin *StablecoinSession) GATEWAYMINTER() (common.Address, error) {
	return _Stablecoin.Contract.GATEWAYMINTER(&_Stablecoin.CallOpts)
}

// GATEWAYMINTER is a free data retrieval call binding the contract method 0x6558b981.
//
// Solidity: function GATEWAY_MINTER() view returns(address)
func (_Stablecoin *StablecoinCallerSession) GATEWAYMINTER() (common.Address, error) {
	return _Stablecoin.Contract.GATEWAYMINTER(&_Stablecoin.CallOpts)
}

// HUBDOMAIN is a free data retrieval call binding the contract method 0x17f82014.
//
// Solidity: function HUB_DOMAIN() view returns(uint32)
func (_Stablecoin *StablecoinCaller) HUBDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "HUB_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// HUBDOMAIN is a free data retrieval call binding the contract method 0x17f82014.
//
// Solidity: function HUB_DOMAIN() view returns(uint32)
func (_Stablecoin *StablecoinSession) HUBDOMAIN() (uint32, error) {
	return _Stablecoin.Contract.HUBDOMAIN(&_Stablecoin.CallOpts)
}

// HUBDOMAIN is a free data retrieval call binding the contract method 0x17f82014.
//
// Solidity: function HUB_DOMAIN() view returns(uint32)
func (_Stablecoin *StablecoinCallerSession) HUBDOMAIN() (uint32, error) {
	return _Stablecoin.Contract.HUBDOMAIN(&_Stablecoin.CallOpts)
}

// MESSANGER is a free data retrieval call binding the contract method 0x1404b48f.
//
// Solidity: function MESSANGER() view returns(address)
func (_Stablecoin *StablecoinCaller) MESSANGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "MESSANGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSANGER is a free data retrieval call binding the contract method 0x1404b48f.
//
// Solidity: function MESSANGER() view returns(address)
func (_Stablecoin *StablecoinSession) MESSANGER() (common.Address, error) {
	return _Stablecoin.Contract.MESSANGER(&_Stablecoin.CallOpts)
}

// MESSANGER is a free data retrieval call binding the contract method 0x1404b48f.
//
// Solidity: function MESSANGER() view returns(address)
func (_Stablecoin *StablecoinCallerSession) MESSANGER() (common.Address, error) {
	return _Stablecoin.Contract.MESSANGER(&_Stablecoin.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Stablecoin *StablecoinCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Stablecoin *StablecoinSession) OPERATORROLE() ([32]byte, error) {
	return _Stablecoin.Contract.OPERATORROLE(&_Stablecoin.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Stablecoin *StablecoinCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Stablecoin.Contract.OPERATORROLE(&_Stablecoin.CallOpts)
}

// TRANSMITTER is a free data retrieval call binding the contract method 0xf4a5e62c.
//
// Solidity: function TRANSMITTER() view returns(address)
func (_Stablecoin *StablecoinCaller) TRANSMITTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "TRANSMITTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TRANSMITTER is a free data retrieval call binding the contract method 0xf4a5e62c.
//
// Solidity: function TRANSMITTER() view returns(address)
func (_Stablecoin *StablecoinSession) TRANSMITTER() (common.Address, error) {
	return _Stablecoin.Contract.TRANSMITTER(&_Stablecoin.CallOpts)
}

// TRANSMITTER is a free data retrieval call binding the contract method 0xf4a5e62c.
//
// Solidity: function TRANSMITTER() view returns(address)
func (_Stablecoin *StablecoinCallerSession) TRANSMITTER() (common.Address, error) {
	return _Stablecoin.Contract.TRANSMITTER(&_Stablecoin.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Stablecoin *StablecoinCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Stablecoin *StablecoinSession) USDC() (common.Address, error) {
	return _Stablecoin.Contract.USDC(&_Stablecoin.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Stablecoin *StablecoinCallerSession) USDC() (common.Address, error) {
	return _Stablecoin.Contract.USDC(&_Stablecoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Stablecoin *StablecoinCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Stablecoin *StablecoinSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stablecoin.Contract.Allowance(&_Stablecoin.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stablecoin.Contract.Allowance(&_Stablecoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Stablecoin *StablecoinCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Stablecoin *StablecoinSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Stablecoin.Contract.BalanceOf(&_Stablecoin.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Stablecoin.Contract.BalanceOf(&_Stablecoin.CallOpts, arg0)
}

// BridgeNonce is a free data retrieval call binding the contract method 0x7bde3fa7.
//
// Solidity: function bridge_nonce() view returns(uint256)
func (_Stablecoin *StablecoinCaller) BridgeNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "bridge_nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeNonce is a free data retrieval call binding the contract method 0x7bde3fa7.
//
// Solidity: function bridge_nonce() view returns(uint256)
func (_Stablecoin *StablecoinSession) BridgeNonce() (*big.Int, error) {
	return _Stablecoin.Contract.BridgeNonce(&_Stablecoin.CallOpts)
}

// BridgeNonce is a free data retrieval call binding the contract method 0x7bde3fa7.
//
// Solidity: function bridge_nonce() view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) BridgeNonce() (*big.Int, error) {
	return _Stablecoin.Contract.BridgeNonce(&_Stablecoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stablecoin *StablecoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stablecoin *StablecoinSession) Decimals() (uint8, error) {
	return _Stablecoin.Contract.Decimals(&_Stablecoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Stablecoin *StablecoinCallerSession) Decimals() (uint8, error) {
	return _Stablecoin.Contract.Decimals(&_Stablecoin.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1, string, string, uint256, address, bytes32, uint256[])
func (_Stablecoin *StablecoinCaller) Eip712Domain(opts *bind.CallOpts) ([1]byte, string, string, *big.Int, common.Address, [32]byte, []*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "eip712Domain")

	if err != nil {
		return *new([1]byte), *new(string), *new(string), *new(*big.Int), *new(common.Address), *new([32]byte), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	out6 := *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1, string, string, uint256, address, bytes32, uint256[])
func (_Stablecoin *StablecoinSession) Eip712Domain() ([1]byte, string, string, *big.Int, common.Address, [32]byte, []*big.Int, error) {
	return _Stablecoin.Contract.Eip712Domain(&_Stablecoin.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1, string, string, uint256, address, bytes32, uint256[])
func (_Stablecoin *StablecoinCallerSession) Eip712Domain() ([1]byte, string, string, *big.Int, common.Address, [32]byte, []*big.Int, error) {
	return _Stablecoin.Contract.Eip712Domain(&_Stablecoin.CallOpts)
}

// EncodeErc20Approve is a free data retrieval call binding the contract method 0x70bc332e.
//
// Solidity: function encode_erc20_approve(address target, uint256 amount) pure returns(bytes)
func (_Stablecoin *StablecoinCaller) EncodeErc20Approve(opts *bind.CallOpts, target common.Address, amount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "encode_erc20_approve", target, amount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeErc20Approve is a free data retrieval call binding the contract method 0x70bc332e.
//
// Solidity: function encode_erc20_approve(address target, uint256 amount) pure returns(bytes)
func (_Stablecoin *StablecoinSession) EncodeErc20Approve(target common.Address, amount *big.Int) ([]byte, error) {
	return _Stablecoin.Contract.EncodeErc20Approve(&_Stablecoin.CallOpts, target, amount)
}

// EncodeErc20Approve is a free data retrieval call binding the contract method 0x70bc332e.
//
// Solidity: function encode_erc20_approve(address target, uint256 amount) pure returns(bytes)
func (_Stablecoin *StablecoinCallerSession) EncodeErc20Approve(target common.Address, amount *big.Int) ([]byte, error) {
	return _Stablecoin.Contract.EncodeErc20Approve(&_Stablecoin.CallOpts, target, amount)
}

// EncodeSupplyAaveV3 is a free data retrieval call binding the contract method 0x9d3ac1eb.
//
// Solidity: function encode_supply_aave_v3(address asset, uint256 amount, address onBehalfOf, uint16 referrer) pure returns(bytes)
func (_Stablecoin *StablecoinCaller) EncodeSupplyAaveV3(opts *bind.CallOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referrer uint16) ([]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "encode_supply_aave_v3", asset, amount, onBehalfOf, referrer)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeSupplyAaveV3 is a free data retrieval call binding the contract method 0x9d3ac1eb.
//
// Solidity: function encode_supply_aave_v3(address asset, uint256 amount, address onBehalfOf, uint16 referrer) pure returns(bytes)
func (_Stablecoin *StablecoinSession) EncodeSupplyAaveV3(asset common.Address, amount *big.Int, onBehalfOf common.Address, referrer uint16) ([]byte, error) {
	return _Stablecoin.Contract.EncodeSupplyAaveV3(&_Stablecoin.CallOpts, asset, amount, onBehalfOf, referrer)
}

// EncodeSupplyAaveV3 is a free data retrieval call binding the contract method 0x9d3ac1eb.
//
// Solidity: function encode_supply_aave_v3(address asset, uint256 amount, address onBehalfOf, uint16 referrer) pure returns(bytes)
func (_Stablecoin *StablecoinCallerSession) EncodeSupplyAaveV3(asset common.Address, amount *big.Int, onBehalfOf common.Address, referrer uint16) ([]byte, error) {
	return _Stablecoin.Contract.EncodeSupplyAaveV3(&_Stablecoin.CallOpts, asset, amount, onBehalfOf, referrer)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 arg0) view returns(bytes32)
func (_Stablecoin *StablecoinCaller) GetRoleAdmin(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "getRoleAdmin", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 arg0) view returns(bytes32)
func (_Stablecoin *StablecoinSession) GetRoleAdmin(arg0 [32]byte) ([32]byte, error) {
	return _Stablecoin.Contract.GetRoleAdmin(&_Stablecoin.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 arg0) view returns(bytes32)
func (_Stablecoin *StablecoinCallerSession) GetRoleAdmin(arg0 [32]byte) ([32]byte, error) {
	return _Stablecoin.Contract.GetRoleAdmin(&_Stablecoin.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 arg0, address arg1) view returns(bool)
func (_Stablecoin *StablecoinCaller) HasRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "hasRole", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 arg0, address arg1) view returns(bool)
func (_Stablecoin *StablecoinSession) HasRole(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Stablecoin.Contract.HasRole(&_Stablecoin.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 arg0, address arg1) view returns(bool)
func (_Stablecoin *StablecoinCallerSession) HasRole(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Stablecoin.Contract.HasRole(&_Stablecoin.CallOpts, arg0, arg1)
}

// IsMinter is a free data retrieval call binding the contract method 0x92c94f65.
//
// Solidity: function is_minter(address arg0) view returns(bool)
func (_Stablecoin *StablecoinCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "is_minter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0x92c94f65.
//
// Solidity: function is_minter(address arg0) view returns(bool)
func (_Stablecoin *StablecoinSession) IsMinter(arg0 common.Address) (bool, error) {
	return _Stablecoin.Contract.IsMinter(&_Stablecoin.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0x92c94f65.
//
// Solidity: function is_minter(address arg0) view returns(bool)
func (_Stablecoin *StablecoinCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _Stablecoin.Contract.IsMinter(&_Stablecoin.CallOpts, arg0)
}

// IsStrategyAllowed is a free data retrieval call binding the contract method 0x71fb6c9f.
//
// Solidity: function is_strategy_allowed(address target, bytes selector, bytes32[] proof) view returns(bool)
func (_Stablecoin *StablecoinCaller) IsStrategyAllowed(opts *bind.CallOpts, target common.Address, selector []byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "is_strategy_allowed", target, selector, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrategyAllowed is a free data retrieval call binding the contract method 0x71fb6c9f.
//
// Solidity: function is_strategy_allowed(address target, bytes selector, bytes32[] proof) view returns(bool)
func (_Stablecoin *StablecoinSession) IsStrategyAllowed(target common.Address, selector []byte, proof [][32]byte) (bool, error) {
	return _Stablecoin.Contract.IsStrategyAllowed(&_Stablecoin.CallOpts, target, selector, proof)
}

// IsStrategyAllowed is a free data retrieval call binding the contract method 0x71fb6c9f.
//
// Solidity: function is_strategy_allowed(address target, bytes selector, bytes32[] proof) view returns(bool)
func (_Stablecoin *StablecoinCallerSession) IsStrategyAllowed(target common.Address, selector []byte, proof [][32]byte) (bool, error) {
	return _Stablecoin.Contract.IsStrategyAllowed(&_Stablecoin.CallOpts, target, selector, proof)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stablecoin *StablecoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stablecoin *StablecoinSession) Name() (string, error) {
	return _Stablecoin.Contract.Name(&_Stablecoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stablecoin *StablecoinCallerSession) Name() (string, error) {
	return _Stablecoin.Contract.Name(&_Stablecoin.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Stablecoin *StablecoinCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Stablecoin *StablecoinSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Stablecoin.Contract.Nonces(&_Stablecoin.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Stablecoin.Contract.Nonces(&_Stablecoin.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stablecoin *StablecoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stablecoin *StablecoinSession) Owner() (common.Address, error) {
	return _Stablecoin.Contract.Owner(&_Stablecoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stablecoin *StablecoinCallerSession) Owner() (common.Address, error) {
	return _Stablecoin.Contract.Owner(&_Stablecoin.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0x22e884b8.
//
// Solidity: function pending_withdrawals(address arg0) view returns((uint256,uint256))
func (_Stablecoin *StablecoinCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address) (Struct0, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "pending_withdrawals", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0x22e884b8.
//
// Solidity: function pending_withdrawals(address arg0) view returns((uint256,uint256))
func (_Stablecoin *StablecoinSession) PendingWithdrawals(arg0 common.Address) (Struct0, error) {
	return _Stablecoin.Contract.PendingWithdrawals(&_Stablecoin.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0x22e884b8.
//
// Solidity: function pending_withdrawals(address arg0) view returns((uint256,uint256))
func (_Stablecoin *StablecoinCallerSession) PendingWithdrawals(arg0 common.Address) (Struct0, error) {
	return _Stablecoin.Contract.PendingWithdrawals(&_Stablecoin.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Stablecoin *StablecoinCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Stablecoin *StablecoinSession) Staking() (common.Address, error) {
	return _Stablecoin.Contract.Staking(&_Stablecoin.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Stablecoin *StablecoinCallerSession) Staking() (common.Address, error) {
	return _Stablecoin.Contract.Staking(&_Stablecoin.CallOpts)
}

// StrategyRootHash is a free data retrieval call binding the contract method 0x40e48966.
//
// Solidity: function strategy_root_hash() view returns(bytes32)
func (_Stablecoin *StablecoinCaller) StrategyRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "strategy_root_hash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StrategyRootHash is a free data retrieval call binding the contract method 0x40e48966.
//
// Solidity: function strategy_root_hash() view returns(bytes32)
func (_Stablecoin *StablecoinSession) StrategyRootHash() ([32]byte, error) {
	return _Stablecoin.Contract.StrategyRootHash(&_Stablecoin.CallOpts)
}

// StrategyRootHash is a free data retrieval call binding the contract method 0x40e48966.
//
// Solidity: function strategy_root_hash() view returns(bytes32)
func (_Stablecoin *StablecoinCallerSession) StrategyRootHash() ([32]byte, error) {
	return _Stablecoin.Contract.StrategyRootHash(&_Stablecoin.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interface_id) view returns(bool)
func (_Stablecoin *StablecoinCaller) SupportsInterface(opts *bind.CallOpts, interface_id [4]byte) (bool, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "supportsInterface", interface_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interface_id) view returns(bool)
func (_Stablecoin *StablecoinSession) SupportsInterface(interface_id [4]byte) (bool, error) {
	return _Stablecoin.Contract.SupportsInterface(&_Stablecoin.CallOpts, interface_id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interface_id) view returns(bool)
func (_Stablecoin *StablecoinCallerSession) SupportsInterface(interface_id [4]byte) (bool, error) {
	return _Stablecoin.Contract.SupportsInterface(&_Stablecoin.CallOpts, interface_id)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stablecoin *StablecoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stablecoin *StablecoinSession) Symbol() (string, error) {
	return _Stablecoin.Contract.Symbol(&_Stablecoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stablecoin *StablecoinCallerSession) Symbol() (string, error) {
	return _Stablecoin.Contract.Symbol(&_Stablecoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stablecoin *StablecoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stablecoin *StablecoinSession) TotalSupply() (*big.Int, error) {
	return _Stablecoin.Contract.TotalSupply(&_Stablecoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) TotalSupply() (*big.Int, error) {
	return _Stablecoin.Contract.TotalSupply(&_Stablecoin.CallOpts)
}

// UsdcFees is a free data retrieval call binding the contract method 0x51b5a503.
//
// Solidity: function usdc_fees() view returns(uint256)
func (_Stablecoin *StablecoinCaller) UsdcFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "usdc_fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcFees is a free data retrieval call binding the contract method 0x51b5a503.
//
// Solidity: function usdc_fees() view returns(uint256)
func (_Stablecoin *StablecoinSession) UsdcFees() (*big.Int, error) {
	return _Stablecoin.Contract.UsdcFees(&_Stablecoin.CallOpts)
}

// UsdcFees is a free data retrieval call binding the contract method 0x51b5a503.
//
// Solidity: function usdc_fees() view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) UsdcFees() (*big.Int, error) {
	return _Stablecoin.Contract.UsdcFees(&_Stablecoin.CallOpts)
}

// UsdcReserve is a free data retrieval call binding the contract method 0x29b95248.
//
// Solidity: function usdc_reserve() view returns(uint256)
func (_Stablecoin *StablecoinCaller) UsdcReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "usdc_reserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcReserve is a free data retrieval call binding the contract method 0x29b95248.
//
// Solidity: function usdc_reserve() view returns(uint256)
func (_Stablecoin *StablecoinSession) UsdcReserve() (*big.Int, error) {
	return _Stablecoin.Contract.UsdcReserve(&_Stablecoin.CallOpts)
}

// UsdcReserve is a free data retrieval call binding the contract method 0x29b95248.
//
// Solidity: function usdc_reserve() view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) UsdcReserve() (*big.Int, error) {
	return _Stablecoin.Contract.UsdcReserve(&_Stablecoin.CallOpts)
}

// UsdcUsage is a free data retrieval call binding the contract method 0x1a98e895.
//
// Solidity: function usdc_usage() view returns(uint256)
func (_Stablecoin *StablecoinCaller) UsdcUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "usdc_usage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcUsage is a free data retrieval call binding the contract method 0x1a98e895.
//
// Solidity: function usdc_usage() view returns(uint256)
func (_Stablecoin *StablecoinSession) UsdcUsage() (*big.Int, error) {
	return _Stablecoin.Contract.UsdcUsage(&_Stablecoin.CallOpts)
}

// UsdcUsage is a free data retrieval call binding the contract method 0x1a98e895.
//
// Solidity: function usdc_usage() view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) UsdcUsage() (*big.Int, error) {
	return _Stablecoin.Contract.UsdcUsage(&_Stablecoin.CallOpts)
}

// WhitelistedDomains is a free data retrieval call binding the contract method 0xfaa82272.
//
// Solidity: function whitelisted_domains(uint32 arg0) view returns(address)
func (_Stablecoin *StablecoinCaller) WhitelistedDomains(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "whitelisted_domains", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedDomains is a free data retrieval call binding the contract method 0xfaa82272.
//
// Solidity: function whitelisted_domains(uint32 arg0) view returns(address)
func (_Stablecoin *StablecoinSession) WhitelistedDomains(arg0 uint32) (common.Address, error) {
	return _Stablecoin.Contract.WhitelistedDomains(&_Stablecoin.CallOpts, arg0)
}

// WhitelistedDomains is a free data retrieval call binding the contract method 0xfaa82272.
//
// Solidity: function whitelisted_domains(uint32 arg0) view returns(address)
func (_Stablecoin *StablecoinCallerSession) WhitelistedDomains(arg0 uint32) (common.Address, error) {
	return _Stablecoin.Contract.WhitelistedDomains(&_Stablecoin.CallOpts, arg0)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xeaa26f0f.
//
// Solidity: function withdrawal_delay() view returns(uint256)
func (_Stablecoin *StablecoinCaller) WithdrawalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stablecoin.contract.Call(opts, &out, "withdrawal_delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xeaa26f0f.
//
// Solidity: function withdrawal_delay() view returns(uint256)
func (_Stablecoin *StablecoinSession) WithdrawalDelay() (*big.Int, error) {
	return _Stablecoin.Contract.WithdrawalDelay(&_Stablecoin.CallOpts)
}

// WithdrawalDelay is a free data retrieval call binding the contract method 0xeaa26f0f.
//
// Solidity: function withdrawal_delay() view returns(uint256)
func (_Stablecoin *StablecoinCallerSession) WithdrawalDelay() (*big.Int, error) {
	return _Stablecoin.Contract.WithdrawalDelay(&_Stablecoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Approve(&_Stablecoin.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Approve(&_Stablecoin.TransactOpts, spender, amount)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x3e6aa3dc.
//
// Solidity: function batch_execute(address[] targets, bytes[] selectors, bytes[] data, bytes32[][] proofs) returns()
func (_Stablecoin *StablecoinTransactor) BatchExecute(opts *bind.TransactOpts, targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "batch_execute", targets, selectors, data, proofs)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x3e6aa3dc.
//
// Solidity: function batch_execute(address[] targets, bytes[] selectors, bytes[] data, bytes32[][] proofs) returns()
func (_Stablecoin *StablecoinSession) BatchExecute(targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.BatchExecute(&_Stablecoin.TransactOpts, targets, selectors, data, proofs)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x3e6aa3dc.
//
// Solidity: function batch_execute(address[] targets, bytes[] selectors, bytes[] data, bytes32[][] proofs) returns()
func (_Stablecoin *StablecoinTransactorSession) BatchExecute(targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.BatchExecute(&_Stablecoin.TransactOpts, targets, selectors, data, proofs)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Stablecoin *StablecoinTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Stablecoin *StablecoinSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Burn(&_Stablecoin.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Stablecoin *StablecoinTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Burn(&_Stablecoin.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x0f536f84.
//
// Solidity: function burn_from(address owner, uint256 amount) returns()
func (_Stablecoin *StablecoinTransactor) BurnFrom(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "burn_from", owner, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x0f536f84.
//
// Solidity: function burn_from(address owner, uint256 amount) returns()
func (_Stablecoin *StablecoinSession) BurnFrom(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.BurnFrom(&_Stablecoin.TransactOpts, owner, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x0f536f84.
//
// Solidity: function burn_from(address owner, uint256 amount) returns()
func (_Stablecoin *StablecoinTransactorSession) BurnFrom(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.BurnFrom(&_Stablecoin.TransactOpts, owner, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Stablecoin *StablecoinTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Stablecoin *StablecoinSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Deposit(&_Stablecoin.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Stablecoin *StablecoinTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Deposit(&_Stablecoin.TransactOpts, amount)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x5e844798.
//
// Solidity: function distribute_fees() returns()
func (_Stablecoin *StablecoinTransactor) DistributeFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "distribute_fees")
}

// DistributeFees is a paid mutator transaction binding the contract method 0x5e844798.
//
// Solidity: function distribute_fees() returns()
func (_Stablecoin *StablecoinSession) DistributeFees() (*types.Transaction, error) {
	return _Stablecoin.Contract.DistributeFees(&_Stablecoin.TransactOpts)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x5e844798.
//
// Solidity: function distribute_fees() returns()
func (_Stablecoin *StablecoinTransactorSession) DistributeFees() (*types.Transaction, error) {
	return _Stablecoin.Contract.DistributeFees(&_Stablecoin.TransactOpts)
}

// ExecuteStrategy is a paid mutator transaction binding the contract method 0x4f494399.
//
// Solidity: function execute_strategy(address target, bytes selector, bytes data, bytes32[] proof) returns()
func (_Stablecoin *StablecoinTransactor) ExecuteStrategy(opts *bind.TransactOpts, target common.Address, selector []byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "execute_strategy", target, selector, data, proof)
}

// ExecuteStrategy is a paid mutator transaction binding the contract method 0x4f494399.
//
// Solidity: function execute_strategy(address target, bytes selector, bytes data, bytes32[] proof) returns()
func (_Stablecoin *StablecoinSession) ExecuteStrategy(target common.Address, selector []byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.ExecuteStrategy(&_Stablecoin.TransactOpts, target, selector, data, proof)
}

// ExecuteStrategy is a paid mutator transaction binding the contract method 0x4f494399.
//
// Solidity: function execute_strategy(address target, bytes selector, bytes data, bytes32[] proof) returns()
func (_Stablecoin *StablecoinTransactorSession) ExecuteStrategy(target common.Address, selector []byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.ExecuteStrategy(&_Stablecoin.TransactOpts, target, selector, data, proof)
}

// ExecuteWithdraw is a paid mutator transaction binding the contract method 0xf4d46735.
//
// Solidity: function execute_withdraw() returns()
func (_Stablecoin *StablecoinTransactor) ExecuteWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "execute_withdraw")
}

// ExecuteWithdraw is a paid mutator transaction binding the contract method 0xf4d46735.
//
// Solidity: function execute_withdraw() returns()
func (_Stablecoin *StablecoinSession) ExecuteWithdraw() (*types.Transaction, error) {
	return _Stablecoin.Contract.ExecuteWithdraw(&_Stablecoin.TransactOpts)
}

// ExecuteWithdraw is a paid mutator transaction binding the contract method 0xf4d46735.
//
// Solidity: function execute_withdraw() returns()
func (_Stablecoin *StablecoinTransactorSession) ExecuteWithdraw() (*types.Transaction, error) {
	return _Stablecoin.Contract.ExecuteWithdraw(&_Stablecoin.TransactOpts)
}

// GatewayDepositAndStake is a paid mutator transaction binding the contract method 0x95e4dcec.
//
// Solidity: function gateway_deposit_and_stake(bytes attestation, bytes signature) returns()
func (_Stablecoin *StablecoinTransactor) GatewayDepositAndStake(opts *bind.TransactOpts, attestation []byte, signature []byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "gateway_deposit_and_stake", attestation, signature)
}

// GatewayDepositAndStake is a paid mutator transaction binding the contract method 0x95e4dcec.
//
// Solidity: function gateway_deposit_and_stake(bytes attestation, bytes signature) returns()
func (_Stablecoin *StablecoinSession) GatewayDepositAndStake(attestation []byte, signature []byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.GatewayDepositAndStake(&_Stablecoin.TransactOpts, attestation, signature)
}

// GatewayDepositAndStake is a paid mutator transaction binding the contract method 0x95e4dcec.
//
// Solidity: function gateway_deposit_and_stake(bytes attestation, bytes signature) returns()
func (_Stablecoin *StablecoinTransactorSession) GatewayDepositAndStake(attestation []byte, signature []byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.GatewayDepositAndStake(&_Stablecoin.TransactOpts, attestation, signature)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.GrantRole(&_Stablecoin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.GrantRole(&_Stablecoin.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address owner, uint256 amount) returns()
func (_Stablecoin *StablecoinTransactor) Mint(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "mint", owner, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address owner, uint256 amount) returns()
func (_Stablecoin *StablecoinSession) Mint(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Mint(&_Stablecoin.TransactOpts, owner, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address owner, uint256 amount) returns()
func (_Stablecoin *StablecoinTransactorSession) Mint(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Mint(&_Stablecoin.TransactOpts, owner, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Stablecoin *StablecoinTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "permit", owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Stablecoin *StablecoinSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.Permit(&_Stablecoin.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Stablecoin *StablecoinTransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.Permit(&_Stablecoin.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x600ecd47.
//
// Solidity: function receive_message(bytes message, bytes attestation) returns()
func (_Stablecoin *StablecoinTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "receive_message", message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x600ecd47.
//
// Solidity: function receive_message(bytes message, bytes attestation) returns()
func (_Stablecoin *StablecoinSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.ReceiveMessage(&_Stablecoin.TransactOpts, message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x600ecd47.
//
// Solidity: function receive_message(bytes message, bytes attestation) returns()
func (_Stablecoin *StablecoinTransactorSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.ReceiveMessage(&_Stablecoin.TransactOpts, message, attestation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.RenounceRole(&_Stablecoin.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.RenounceRole(&_Stablecoin.TransactOpts, role, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0xb15e13ee.
//
// Solidity: function renounce_ownership() returns()
func (_Stablecoin *StablecoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "renounce_ownership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0xb15e13ee.
//
// Solidity: function renounce_ownership() returns()
func (_Stablecoin *StablecoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stablecoin.Contract.RenounceOwnership(&_Stablecoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0xb15e13ee.
//
// Solidity: function renounce_ownership() returns()
func (_Stablecoin *StablecoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stablecoin.Contract.RenounceOwnership(&_Stablecoin.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.RevokeRole(&_Stablecoin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stablecoin *StablecoinTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.RevokeRole(&_Stablecoin.TransactOpts, role, account)
}

// SendAllToHub is a paid mutator transaction binding the contract method 0xf289355f.
//
// Solidity: function send_all_to_hub(uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactor) SendAllToHub(opts *bind.TransactOpts, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "send_all_to_hub", max_fee, min_finality_threshold)
}

// SendAllToHub is a paid mutator transaction binding the contract method 0xf289355f.
//
// Solidity: function send_all_to_hub(uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinSession) SendAllToHub(max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendAllToHub(&_Stablecoin.TransactOpts, max_fee, min_finality_threshold)
}

// SendAllToHub is a paid mutator transaction binding the contract method 0xf289355f.
//
// Solidity: function send_all_to_hub(uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactorSession) SendAllToHub(max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendAllToHub(&_Stablecoin.TransactOpts, max_fee, min_finality_threshold)
}

// SendAllToTerrace is a paid mutator transaction binding the contract method 0x56053c3c.
//
// Solidity: function send_all_to_terrace(uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactor) SendAllToTerrace(opts *bind.TransactOpts, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "send_all_to_terrace", destination_domain, max_fee, min_finality_threshold)
}

// SendAllToTerrace is a paid mutator transaction binding the contract method 0x56053c3c.
//
// Solidity: function send_all_to_terrace(uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinSession) SendAllToTerrace(destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendAllToTerrace(&_Stablecoin.TransactOpts, destination_domain, max_fee, min_finality_threshold)
}

// SendAllToTerrace is a paid mutator transaction binding the contract method 0x56053c3c.
//
// Solidity: function send_all_to_terrace(uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactorSession) SendAllToTerrace(destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendAllToTerrace(&_Stablecoin.TransactOpts, destination_domain, max_fee, min_finality_threshold)
}

// SendToHub is a paid mutator transaction binding the contract method 0xd6164993.
//
// Solidity: function send_to_hub(uint256 amount, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactor) SendToHub(opts *bind.TransactOpts, amount *big.Int, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "send_to_hub", amount, max_fee, min_finality_threshold)
}

// SendToHub is a paid mutator transaction binding the contract method 0xd6164993.
//
// Solidity: function send_to_hub(uint256 amount, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinSession) SendToHub(amount *big.Int, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendToHub(&_Stablecoin.TransactOpts, amount, max_fee, min_finality_threshold)
}

// SendToHub is a paid mutator transaction binding the contract method 0xd6164993.
//
// Solidity: function send_to_hub(uint256 amount, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactorSession) SendToHub(amount *big.Int, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendToHub(&_Stablecoin.TransactOpts, amount, max_fee, min_finality_threshold)
}

// SendToTerrace is a paid mutator transaction binding the contract method 0x6d7b50cf.
//
// Solidity: function send_to_terrace(uint256 amount, uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactor) SendToTerrace(opts *bind.TransactOpts, amount *big.Int, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "send_to_terrace", amount, destination_domain, max_fee, min_finality_threshold)
}

// SendToTerrace is a paid mutator transaction binding the contract method 0x6d7b50cf.
//
// Solidity: function send_to_terrace(uint256 amount, uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinSession) SendToTerrace(amount *big.Int, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendToTerrace(&_Stablecoin.TransactOpts, amount, destination_domain, max_fee, min_finality_threshold)
}

// SendToTerrace is a paid mutator transaction binding the contract method 0x6d7b50cf.
//
// Solidity: function send_to_terrace(uint256 amount, uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Stablecoin *StablecoinTransactorSession) SendToTerrace(amount *big.Int, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Stablecoin.Contract.SendToTerrace(&_Stablecoin.TransactOpts, amount, destination_domain, max_fee, min_finality_threshold)
}

// SetMinter is a paid mutator transaction binding the contract method 0x7c3bec3c.
//
// Solidity: function set_minter(address minter, bool status) returns()
func (_Stablecoin *StablecoinTransactor) SetMinter(opts *bind.TransactOpts, minter common.Address, status bool) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "set_minter", minter, status)
}

// SetMinter is a paid mutator transaction binding the contract method 0x7c3bec3c.
//
// Solidity: function set_minter(address minter, bool status) returns()
func (_Stablecoin *StablecoinSession) SetMinter(minter common.Address, status bool) (*types.Transaction, error) {
	return _Stablecoin.Contract.SetMinter(&_Stablecoin.TransactOpts, minter, status)
}

// SetMinter is a paid mutator transaction binding the contract method 0x7c3bec3c.
//
// Solidity: function set_minter(address minter, bool status) returns()
func (_Stablecoin *StablecoinTransactorSession) SetMinter(minter common.Address, status bool) (*types.Transaction, error) {
	return _Stablecoin.Contract.SetMinter(&_Stablecoin.TransactOpts, minter, status)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0xc2f9e63f.
//
// Solidity: function set_role_admin(bytes32 role, bytes32 admin_role) returns()
func (_Stablecoin *StablecoinTransactor) SetRoleAdmin(opts *bind.TransactOpts, role [32]byte, admin_role [32]byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "set_role_admin", role, admin_role)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0xc2f9e63f.
//
// Solidity: function set_role_admin(bytes32 role, bytes32 admin_role) returns()
func (_Stablecoin *StablecoinSession) SetRoleAdmin(role [32]byte, admin_role [32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.SetRoleAdmin(&_Stablecoin.TransactOpts, role, admin_role)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0xc2f9e63f.
//
// Solidity: function set_role_admin(bytes32 role, bytes32 admin_role) returns()
func (_Stablecoin *StablecoinTransactorSession) SetRoleAdmin(role [32]byte, admin_role [32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.SetRoleAdmin(&_Stablecoin.TransactOpts, role, admin_role)
}

// SetStaking is a paid mutator transaction binding the contract method 0x22008d56.
//
// Solidity: function set_staking(address staking) returns()
func (_Stablecoin *StablecoinTransactor) SetStaking(opts *bind.TransactOpts, staking common.Address) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "set_staking", staking)
}

// SetStaking is a paid mutator transaction binding the contract method 0x22008d56.
//
// Solidity: function set_staking(address staking) returns()
func (_Stablecoin *StablecoinSession) SetStaking(staking common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.SetStaking(&_Stablecoin.TransactOpts, staking)
}

// SetStaking is a paid mutator transaction binding the contract method 0x22008d56.
//
// Solidity: function set_staking(address staking) returns()
func (_Stablecoin *StablecoinTransactorSession) SetStaking(staking common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.SetStaking(&_Stablecoin.TransactOpts, staking)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Transfer(&_Stablecoin.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Transfer(&_Stablecoin.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address to, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "transferFrom", owner, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address to, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinSession) TransferFrom(owner common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.TransferFrom(&_Stablecoin.TransactOpts, owner, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address to, uint256 amount) returns(bool)
func (_Stablecoin *StablecoinTransactorSession) TransferFrom(owner common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.TransferFrom(&_Stablecoin.TransactOpts, owner, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf0350c04.
//
// Solidity: function transfer_ownership(address new_owner) returns()
func (_Stablecoin *StablecoinTransactor) TransferOwnership(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "transfer_ownership", new_owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf0350c04.
//
// Solidity: function transfer_ownership(address new_owner) returns()
func (_Stablecoin *StablecoinSession) TransferOwnership(new_owner common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.TransferOwnership(&_Stablecoin.TransactOpts, new_owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf0350c04.
//
// Solidity: function transfer_ownership(address new_owner) returns()
func (_Stablecoin *StablecoinTransactorSession) TransferOwnership(new_owner common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.TransferOwnership(&_Stablecoin.TransactOpts, new_owner)
}

// UpdateRootHash is a paid mutator transaction binding the contract method 0xc069315d.
//
// Solidity: function update_root_hash(bytes32 root) returns()
func (_Stablecoin *StablecoinTransactor) UpdateRootHash(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "update_root_hash", root)
}

// UpdateRootHash is a paid mutator transaction binding the contract method 0xc069315d.
//
// Solidity: function update_root_hash(bytes32 root) returns()
func (_Stablecoin *StablecoinSession) UpdateRootHash(root [32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.UpdateRootHash(&_Stablecoin.TransactOpts, root)
}

// UpdateRootHash is a paid mutator transaction binding the contract method 0xc069315d.
//
// Solidity: function update_root_hash(bytes32 root) returns()
func (_Stablecoin *StablecoinTransactorSession) UpdateRootHash(root [32]byte) (*types.Transaction, error) {
	return _Stablecoin.Contract.UpdateRootHash(&_Stablecoin.TransactOpts, root)
}

// UpdateWhitelistedDomain is a paid mutator transaction binding the contract method 0x1b52da6e.
//
// Solidity: function update_whitelisted_domain(uint32 domain, address terrace) returns()
func (_Stablecoin *StablecoinTransactor) UpdateWhitelistedDomain(opts *bind.TransactOpts, domain uint32, terrace common.Address) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "update_whitelisted_domain", domain, terrace)
}

// UpdateWhitelistedDomain is a paid mutator transaction binding the contract method 0x1b52da6e.
//
// Solidity: function update_whitelisted_domain(uint32 domain, address terrace) returns()
func (_Stablecoin *StablecoinSession) UpdateWhitelistedDomain(domain uint32, terrace common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.UpdateWhitelistedDomain(&_Stablecoin.TransactOpts, domain, terrace)
}

// UpdateWhitelistedDomain is a paid mutator transaction binding the contract method 0x1b52da6e.
//
// Solidity: function update_whitelisted_domain(uint32 domain, address terrace) returns()
func (_Stablecoin *StablecoinTransactorSession) UpdateWhitelistedDomain(domain uint32, terrace common.Address) (*types.Transaction, error) {
	return _Stablecoin.Contract.UpdateWhitelistedDomain(&_Stablecoin.TransactOpts, domain, terrace)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x7e08e9c7.
//
// Solidity: function update_withdrawal_delay(uint256 delay) returns()
func (_Stablecoin *StablecoinTransactor) UpdateWithdrawalDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "update_withdrawal_delay", delay)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x7e08e9c7.
//
// Solidity: function update_withdrawal_delay(uint256 delay) returns()
func (_Stablecoin *StablecoinSession) UpdateWithdrawalDelay(delay *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.UpdateWithdrawalDelay(&_Stablecoin.TransactOpts, delay)
}

// UpdateWithdrawalDelay is a paid mutator transaction binding the contract method 0x7e08e9c7.
//
// Solidity: function update_withdrawal_delay(uint256 delay) returns()
func (_Stablecoin *StablecoinTransactorSession) UpdateWithdrawalDelay(delay *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.UpdateWithdrawalDelay(&_Stablecoin.TransactOpts, delay)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stablecoin *StablecoinTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stablecoin *StablecoinSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Withdraw(&_Stablecoin.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stablecoin *StablecoinTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stablecoin.Contract.Withdraw(&_Stablecoin.TransactOpts, amount)
}

// StablecoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Stablecoin contract.
type StablecoinApprovalIterator struct {
	Event *StablecoinApproval // Event containing the contract specifics and raw log

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
func (it *StablecoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinApproval)
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
		it.Event = new(StablecoinApproval)
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
func (it *StablecoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinApproval represents a Approval event raised by the Stablecoin contract.
type StablecoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Stablecoin *StablecoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StablecoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinApprovalIterator{contract: _Stablecoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Stablecoin *StablecoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StablecoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinApproval)
				if err := _Stablecoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Stablecoin *StablecoinFilterer) ParseApproval(log types.Log) (*StablecoinApproval, error) {
	event := new(StablecoinApproval)
	if err := _Stablecoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinDelayUpdatedIterator is returned from FilterDelayUpdated and is used to iterate over the raw logs and unpacked data for DelayUpdated events raised by the Stablecoin contract.
type StablecoinDelayUpdatedIterator struct {
	Event *StablecoinDelayUpdated // Event containing the contract specifics and raw log

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
func (it *StablecoinDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinDelayUpdated)
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
		it.Event = new(StablecoinDelayUpdated)
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
func (it *StablecoinDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinDelayUpdated represents a DelayUpdated event raised by the Stablecoin contract.
type StablecoinDelayUpdated struct {
	NewDelay *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayUpdated is a free log retrieval operation binding the contract event 0x7ed9288af3fe7320be9af1fcb6714d780e3ffd1a17c1a395594978fd6741ddcb.
//
// Solidity: event DelayUpdated(uint256 new_delay)
func (_Stablecoin *StablecoinFilterer) FilterDelayUpdated(opts *bind.FilterOpts) (*StablecoinDelayUpdatedIterator, error) {

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "DelayUpdated")
	if err != nil {
		return nil, err
	}
	return &StablecoinDelayUpdatedIterator{contract: _Stablecoin.contract, event: "DelayUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayUpdated is a free log subscription operation binding the contract event 0x7ed9288af3fe7320be9af1fcb6714d780e3ffd1a17c1a395594978fd6741ddcb.
//
// Solidity: event DelayUpdated(uint256 new_delay)
func (_Stablecoin *StablecoinFilterer) WatchDelayUpdated(opts *bind.WatchOpts, sink chan<- *StablecoinDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "DelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinDelayUpdated)
				if err := _Stablecoin.contract.UnpackLog(event, "DelayUpdated", log); err != nil {
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

// ParseDelayUpdated is a log parse operation binding the contract event 0x7ed9288af3fe7320be9af1fcb6714d780e3ffd1a17c1a395594978fd6741ddcb.
//
// Solidity: event DelayUpdated(uint256 new_delay)
func (_Stablecoin *StablecoinFilterer) ParseDelayUpdated(log types.Log) (*StablecoinDelayUpdated, error) {
	event := new(StablecoinDelayUpdated)
	if err := _Stablecoin.contract.UnpackLog(event, "DelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Stablecoin contract.
type StablecoinDepositIterator struct {
	Event *StablecoinDeposit // Event containing the contract specifics and raw log

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
func (it *StablecoinDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinDeposit)
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
		it.Event = new(StablecoinDeposit)
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
func (it *StablecoinDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinDeposit represents a Deposit event raised by the Stablecoin contract.
type StablecoinDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*StablecoinDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinDepositIterator{contract: _Stablecoin.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StablecoinDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinDeposit)
				if err := _Stablecoin.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) ParseDeposit(log types.Log) (*StablecoinDeposit, error) {
	event := new(StablecoinDeposit)
	if err := _Stablecoin.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stablecoin contract.
type StablecoinOwnershipTransferredIterator struct {
	Event *StablecoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StablecoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinOwnershipTransferred)
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
		it.Event = new(StablecoinOwnershipTransferred)
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
func (it *StablecoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinOwnershipTransferred represents a OwnershipTransferred event raised by the Stablecoin contract.
type StablecoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previous_owner, address indexed new_owner)
func (_Stablecoin *StablecoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previous_owner []common.Address, new_owner []common.Address) (*StablecoinOwnershipTransferredIterator, error) {

	var previous_ownerRule []interface{}
	for _, previous_ownerItem := range previous_owner {
		previous_ownerRule = append(previous_ownerRule, previous_ownerItem)
	}
	var new_ownerRule []interface{}
	for _, new_ownerItem := range new_owner {
		new_ownerRule = append(new_ownerRule, new_ownerItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "OwnershipTransferred", previous_ownerRule, new_ownerRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinOwnershipTransferredIterator{contract: _Stablecoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previous_owner, address indexed new_owner)
func (_Stablecoin *StablecoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StablecoinOwnershipTransferred, previous_owner []common.Address, new_owner []common.Address) (event.Subscription, error) {

	var previous_ownerRule []interface{}
	for _, previous_ownerItem := range previous_owner {
		previous_ownerRule = append(previous_ownerRule, previous_ownerItem)
	}
	var new_ownerRule []interface{}
	for _, new_ownerItem := range new_owner {
		new_ownerRule = append(new_ownerRule, new_ownerItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "OwnershipTransferred", previous_ownerRule, new_ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinOwnershipTransferred)
				if err := _Stablecoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previous_owner, address indexed new_owner)
func (_Stablecoin *StablecoinFilterer) ParseOwnershipTransferred(log types.Log) (*StablecoinOwnershipTransferred, error) {
	event := new(StablecoinOwnershipTransferred)
	if err := _Stablecoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinReceivedFundsIterator is returned from FilterReceivedFunds and is used to iterate over the raw logs and unpacked data for ReceivedFunds events raised by the Stablecoin contract.
type StablecoinReceivedFundsIterator struct {
	Event *StablecoinReceivedFunds // Event containing the contract specifics and raw log

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
func (it *StablecoinReceivedFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinReceivedFunds)
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
		it.Event = new(StablecoinReceivedFunds)
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
func (it *StablecoinReceivedFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinReceivedFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinReceivedFunds represents a ReceivedFunds event raised by the Stablecoin contract.
type StablecoinReceivedFunds struct {
	Id          [32]byte
	FromTerrace common.Address
	FromDomain  uint32
	ToTerrace   common.Address
	ToDomain    uint32
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedFunds is a free log retrieval operation binding the contract event 0x2b8fa6bde978d1a0c02f660e1cb6f232a39b18669bd77f93ddf5c2324f0445b3.
//
// Solidity: event ReceivedFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Stablecoin *StablecoinFilterer) FilterReceivedFunds(opts *bind.FilterOpts) (*StablecoinReceivedFundsIterator, error) {

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return &StablecoinReceivedFundsIterator{contract: _Stablecoin.contract, event: "ReceivedFunds", logs: logs, sub: sub}, nil
}

// WatchReceivedFunds is a free log subscription operation binding the contract event 0x2b8fa6bde978d1a0c02f660e1cb6f232a39b18669bd77f93ddf5c2324f0445b3.
//
// Solidity: event ReceivedFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Stablecoin *StablecoinFilterer) WatchReceivedFunds(opts *bind.WatchOpts, sink chan<- *StablecoinReceivedFunds) (event.Subscription, error) {

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinReceivedFunds)
				if err := _Stablecoin.contract.UnpackLog(event, "ReceivedFunds", log); err != nil {
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

// ParseReceivedFunds is a log parse operation binding the contract event 0x2b8fa6bde978d1a0c02f660e1cb6f232a39b18669bd77f93ddf5c2324f0445b3.
//
// Solidity: event ReceivedFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Stablecoin *StablecoinFilterer) ParseReceivedFunds(log types.Log) (*StablecoinReceivedFunds, error) {
	event := new(StablecoinReceivedFunds)
	if err := _Stablecoin.contract.UnpackLog(event, "ReceivedFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinRequestWithdrawIterator is returned from FilterRequestWithdraw and is used to iterate over the raw logs and unpacked data for RequestWithdraw events raised by the Stablecoin contract.
type StablecoinRequestWithdrawIterator struct {
	Event *StablecoinRequestWithdraw // Event containing the contract specifics and raw log

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
func (it *StablecoinRequestWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinRequestWithdraw)
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
		it.Event = new(StablecoinRequestWithdraw)
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
func (it *StablecoinRequestWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinRequestWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinRequestWithdraw represents a RequestWithdraw event raised by the Stablecoin contract.
type StablecoinRequestWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdraw is a free log retrieval operation binding the contract event 0x0afd74a2a0a78f6c15e41029f44995ee023fe49276f44a4b2b2cf674829362e6.
//
// Solidity: event RequestWithdraw(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) FilterRequestWithdraw(opts *bind.FilterOpts, user []common.Address) (*StablecoinRequestWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "RequestWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinRequestWithdrawIterator{contract: _Stablecoin.contract, event: "RequestWithdraw", logs: logs, sub: sub}, nil
}

// WatchRequestWithdraw is a free log subscription operation binding the contract event 0x0afd74a2a0a78f6c15e41029f44995ee023fe49276f44a4b2b2cf674829362e6.
//
// Solidity: event RequestWithdraw(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) WatchRequestWithdraw(opts *bind.WatchOpts, sink chan<- *StablecoinRequestWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "RequestWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinRequestWithdraw)
				if err := _Stablecoin.contract.UnpackLog(event, "RequestWithdraw", log); err != nil {
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

// ParseRequestWithdraw is a log parse operation binding the contract event 0x0afd74a2a0a78f6c15e41029f44995ee023fe49276f44a4b2b2cf674829362e6.
//
// Solidity: event RequestWithdraw(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) ParseRequestWithdraw(log types.Log) (*StablecoinRequestWithdraw, error) {
	event := new(StablecoinRequestWithdraw)
	if err := _Stablecoin.contract.UnpackLog(event, "RequestWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Stablecoin contract.
type StablecoinRoleAdminChangedIterator struct {
	Event *StablecoinRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StablecoinRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinRoleAdminChanged)
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
		it.Event = new(StablecoinRoleAdminChanged)
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
func (it *StablecoinRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinRoleAdminChanged represents a RoleAdminChanged event raised by the Stablecoin contract.
type StablecoinRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stablecoin *StablecoinFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StablecoinRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinRoleAdminChangedIterator{contract: _Stablecoin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stablecoin *StablecoinFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StablecoinRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinRoleAdminChanged)
				if err := _Stablecoin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stablecoin *StablecoinFilterer) ParseRoleAdminChanged(log types.Log) (*StablecoinRoleAdminChanged, error) {
	event := new(StablecoinRoleAdminChanged)
	if err := _Stablecoin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Stablecoin contract.
type StablecoinRoleGrantedIterator struct {
	Event *StablecoinRoleGranted // Event containing the contract specifics and raw log

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
func (it *StablecoinRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinRoleGranted)
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
		it.Event = new(StablecoinRoleGranted)
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
func (it *StablecoinRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinRoleGranted represents a RoleGranted event raised by the Stablecoin contract.
type StablecoinRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stablecoin *StablecoinFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StablecoinRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinRoleGrantedIterator{contract: _Stablecoin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stablecoin *StablecoinFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StablecoinRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinRoleGranted)
				if err := _Stablecoin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stablecoin *StablecoinFilterer) ParseRoleGranted(log types.Log) (*StablecoinRoleGranted, error) {
	event := new(StablecoinRoleGranted)
	if err := _Stablecoin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinRoleMinterChangedIterator is returned from FilterRoleMinterChanged and is used to iterate over the raw logs and unpacked data for RoleMinterChanged events raised by the Stablecoin contract.
type StablecoinRoleMinterChangedIterator struct {
	Event *StablecoinRoleMinterChanged // Event containing the contract specifics and raw log

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
func (it *StablecoinRoleMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinRoleMinterChanged)
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
		it.Event = new(StablecoinRoleMinterChanged)
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
func (it *StablecoinRoleMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinRoleMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinRoleMinterChanged represents a RoleMinterChanged event raised by the Stablecoin contract.
type StablecoinRoleMinterChanged struct {
	Minter common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleMinterChanged is a free log retrieval operation binding the contract event 0xbb6e183664bd7425a9e444072cb0f1c7f7c4d5486a36d7d24d0b0735687c2ef4.
//
// Solidity: event RoleMinterChanged(address indexed minter, bool status)
func (_Stablecoin *StablecoinFilterer) FilterRoleMinterChanged(opts *bind.FilterOpts, minter []common.Address) (*StablecoinRoleMinterChangedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "RoleMinterChanged", minterRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinRoleMinterChangedIterator{contract: _Stablecoin.contract, event: "RoleMinterChanged", logs: logs, sub: sub}, nil
}

// WatchRoleMinterChanged is a free log subscription operation binding the contract event 0xbb6e183664bd7425a9e444072cb0f1c7f7c4d5486a36d7d24d0b0735687c2ef4.
//
// Solidity: event RoleMinterChanged(address indexed minter, bool status)
func (_Stablecoin *StablecoinFilterer) WatchRoleMinterChanged(opts *bind.WatchOpts, sink chan<- *StablecoinRoleMinterChanged, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "RoleMinterChanged", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinRoleMinterChanged)
				if err := _Stablecoin.contract.UnpackLog(event, "RoleMinterChanged", log); err != nil {
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

// ParseRoleMinterChanged is a log parse operation binding the contract event 0xbb6e183664bd7425a9e444072cb0f1c7f7c4d5486a36d7d24d0b0735687c2ef4.
//
// Solidity: event RoleMinterChanged(address indexed minter, bool status)
func (_Stablecoin *StablecoinFilterer) ParseRoleMinterChanged(log types.Log) (*StablecoinRoleMinterChanged, error) {
	event := new(StablecoinRoleMinterChanged)
	if err := _Stablecoin.contract.UnpackLog(event, "RoleMinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Stablecoin contract.
type StablecoinRoleRevokedIterator struct {
	Event *StablecoinRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StablecoinRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinRoleRevoked)
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
		it.Event = new(StablecoinRoleRevoked)
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
func (it *StablecoinRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinRoleRevoked represents a RoleRevoked event raised by the Stablecoin contract.
type StablecoinRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stablecoin *StablecoinFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StablecoinRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinRoleRevokedIterator{contract: _Stablecoin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stablecoin *StablecoinFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StablecoinRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinRoleRevoked)
				if err := _Stablecoin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stablecoin *StablecoinFilterer) ParseRoleRevoked(log types.Log) (*StablecoinRoleRevoked, error) {
	event := new(StablecoinRoleRevoked)
	if err := _Stablecoin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinSendFundsIterator is returned from FilterSendFunds and is used to iterate over the raw logs and unpacked data for SendFunds events raised by the Stablecoin contract.
type StablecoinSendFundsIterator struct {
	Event *StablecoinSendFunds // Event containing the contract specifics and raw log

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
func (it *StablecoinSendFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinSendFunds)
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
		it.Event = new(StablecoinSendFunds)
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
func (it *StablecoinSendFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinSendFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinSendFunds represents a SendFunds event raised by the Stablecoin contract.
type StablecoinSendFunds struct {
	Id          [32]byte
	FromTerrace common.Address
	FromDomain  uint32
	ToTerrace   common.Address
	ToDomain    uint32
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSendFunds is a free log retrieval operation binding the contract event 0x68628b1c35f23c5db46c17659a5cdc52eb8e303a889a8d3f29495c558cdd9d22.
//
// Solidity: event SendFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Stablecoin *StablecoinFilterer) FilterSendFunds(opts *bind.FilterOpts) (*StablecoinSendFundsIterator, error) {

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "SendFunds")
	if err != nil {
		return nil, err
	}
	return &StablecoinSendFundsIterator{contract: _Stablecoin.contract, event: "SendFunds", logs: logs, sub: sub}, nil
}

// WatchSendFunds is a free log subscription operation binding the contract event 0x68628b1c35f23c5db46c17659a5cdc52eb8e303a889a8d3f29495c558cdd9d22.
//
// Solidity: event SendFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Stablecoin *StablecoinFilterer) WatchSendFunds(opts *bind.WatchOpts, sink chan<- *StablecoinSendFunds) (event.Subscription, error) {

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "SendFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinSendFunds)
				if err := _Stablecoin.contract.UnpackLog(event, "SendFunds", log); err != nil {
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

// ParseSendFunds is a log parse operation binding the contract event 0x68628b1c35f23c5db46c17659a5cdc52eb8e303a889a8d3f29495c558cdd9d22.
//
// Solidity: event SendFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Stablecoin *StablecoinFilterer) ParseSendFunds(log types.Log) (*StablecoinSendFunds, error) {
	event := new(StablecoinSendFunds)
	if err := _Stablecoin.contract.UnpackLog(event, "SendFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinStrategyRootHashUpdatedIterator is returned from FilterStrategyRootHashUpdated and is used to iterate over the raw logs and unpacked data for StrategyRootHashUpdated events raised by the Stablecoin contract.
type StablecoinStrategyRootHashUpdatedIterator struct {
	Event *StablecoinStrategyRootHashUpdated // Event containing the contract specifics and raw log

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
func (it *StablecoinStrategyRootHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinStrategyRootHashUpdated)
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
		it.Event = new(StablecoinStrategyRootHashUpdated)
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
func (it *StablecoinStrategyRootHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinStrategyRootHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinStrategyRootHashUpdated represents a StrategyRootHashUpdated event raised by the Stablecoin contract.
type StablecoinStrategyRootHashUpdated struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStrategyRootHashUpdated is a free log retrieval operation binding the contract event 0xf2041d1a2718102b7ab7f4756f12e97ddbc269be1fa70d86a6065a0c88921d1c.
//
// Solidity: event StrategyRootHashUpdated(bytes32 indexed root)
func (_Stablecoin *StablecoinFilterer) FilterStrategyRootHashUpdated(opts *bind.FilterOpts, root [][32]byte) (*StablecoinStrategyRootHashUpdatedIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "StrategyRootHashUpdated", rootRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinStrategyRootHashUpdatedIterator{contract: _Stablecoin.contract, event: "StrategyRootHashUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategyRootHashUpdated is a free log subscription operation binding the contract event 0xf2041d1a2718102b7ab7f4756f12e97ddbc269be1fa70d86a6065a0c88921d1c.
//
// Solidity: event StrategyRootHashUpdated(bytes32 indexed root)
func (_Stablecoin *StablecoinFilterer) WatchStrategyRootHashUpdated(opts *bind.WatchOpts, sink chan<- *StablecoinStrategyRootHashUpdated, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "StrategyRootHashUpdated", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinStrategyRootHashUpdated)
				if err := _Stablecoin.contract.UnpackLog(event, "StrategyRootHashUpdated", log); err != nil {
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

// ParseStrategyRootHashUpdated is a log parse operation binding the contract event 0xf2041d1a2718102b7ab7f4756f12e97ddbc269be1fa70d86a6065a0c88921d1c.
//
// Solidity: event StrategyRootHashUpdated(bytes32 indexed root)
func (_Stablecoin *StablecoinFilterer) ParseStrategyRootHashUpdated(log types.Log) (*StablecoinStrategyRootHashUpdated, error) {
	event := new(StablecoinStrategyRootHashUpdated)
	if err := _Stablecoin.contract.UnpackLog(event, "StrategyRootHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Stablecoin contract.
type StablecoinTransferIterator struct {
	Event *StablecoinTransfer // Event containing the contract specifics and raw log

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
func (it *StablecoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinTransfer)
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
		it.Event = new(StablecoinTransfer)
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
func (it *StablecoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinTransfer represents a Transfer event raised by the Stablecoin contract.
type StablecoinTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Stablecoin *StablecoinFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*StablecoinTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinTransferIterator{contract: _Stablecoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Stablecoin *StablecoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StablecoinTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinTransfer)
				if err := _Stablecoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Stablecoin *StablecoinFilterer) ParseTransfer(log types.Log) (*StablecoinTransfer, error) {
	event := new(StablecoinTransfer)
	if err := _Stablecoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Stablecoin contract.
type StablecoinWithdrawIterator struct {
	Event *StablecoinWithdraw // Event containing the contract specifics and raw log

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
func (it *StablecoinWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinWithdraw)
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
		it.Event = new(StablecoinWithdraw)
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
func (it *StablecoinWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinWithdraw represents a Withdraw event raised by the Stablecoin contract.
type StablecoinWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*StablecoinWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stablecoin.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinWithdrawIterator{contract: _Stablecoin.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StablecoinWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stablecoin.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinWithdraw)
				if err := _Stablecoin.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Stablecoin *StablecoinFilterer) ParseWithdraw(log types.Log) (*StablecoinWithdraw, error) {
	event := new(StablecoinWithdraw)
	if err := _Stablecoin.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
