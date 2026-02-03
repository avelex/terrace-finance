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

// TerraceMetaData contains all meta data concerning the Terrace contract.
var TerraceMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyRootHashUpdated\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReceivedFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":false},{\"name\":\"from_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"from_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"to_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"to_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SendFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":false},{\"name\":\"from_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"from_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"to_terrace\",\"type\":\"address\",\"indexed\":false},{\"name\":\"to_domain\",\"type\":\"uint32\",\"indexed\":false},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_root_hash\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"execute_strategy\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_execute\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\"},{\"name\":\"selectors\",\"type\":\"bytes[]\"},{\"name\":\"data\",\"type\":\"bytes[]\"},{\"name\":\"proofs\",\"type\":\"bytes32[][]\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_strategy_allowed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"encode_erc20_approve\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"encode_supply_aave_v3\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"name\":\"referrer\",\"type\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"strategy_root_hash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interface_id\",\"type\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_role_admin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\"},{\"name\":\"admin_role\",\"type\":\"bytes32\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"bytes32\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"usdc_reserve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"usdc_usage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"usdc_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_whitelisted_domain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint32\"},{\"name\":\"terrace\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"receive_message\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"attestation\",\"type\":\"bytes\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_all_to_terrace\",\"inputs\":[{\"name\":\"destination_domain\",\"type\":\"uint32\"},{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_to_terrace\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"destination_domain\",\"type\":\"uint32\"},{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_all_to_hub\",\"inputs\":[{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"send_to_hub\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"max_fee\",\"type\":\"uint256\"},{\"name\":\"min_finality_threshold\",\"type\":\"uint32\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MESSANGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TRANSMITTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"HUB_DOMAIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"whitelisted_domains\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bridge_nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"usdc\",\"type\":\"address\"},{\"name\":\"messanger\",\"type\":\"address\"},{\"name\":\"transmitter\",\"type\":\"address\"},{\"name\":\"hub_domain\",\"type\":\"uint32\"}],\"outputs\":[]}]",
}

// TerraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TerraceMetaData.ABI instead.
var TerraceABI = TerraceMetaData.ABI

// Terrace is an auto generated Go binding around an Ethereum contract.
type Terrace struct {
	TerraceCaller     // Read-only binding to the contract
	TerraceTransactor // Write-only binding to the contract
	TerraceFilterer   // Log filterer for contract events
}

// TerraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TerraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TerraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TerraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TerraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TerraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TerraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TerraceSession struct {
	Contract     *Terrace          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TerraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TerraceCallerSession struct {
	Contract *TerraceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TerraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TerraceTransactorSession struct {
	Contract     *TerraceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TerraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TerraceRaw struct {
	Contract *Terrace // Generic contract binding to access the raw methods on
}

// TerraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TerraceCallerRaw struct {
	Contract *TerraceCaller // Generic read-only contract binding to access the raw methods on
}

// TerraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TerraceTransactorRaw struct {
	Contract *TerraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTerrace creates a new instance of Terrace, bound to a specific deployed contract.
func NewTerrace(address common.Address, backend bind.ContractBackend) (*Terrace, error) {
	contract, err := bindTerrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Terrace{TerraceCaller: TerraceCaller{contract: contract}, TerraceTransactor: TerraceTransactor{contract: contract}, TerraceFilterer: TerraceFilterer{contract: contract}}, nil
}

// NewTerraceCaller creates a new read-only instance of Terrace, bound to a specific deployed contract.
func NewTerraceCaller(address common.Address, caller bind.ContractCaller) (*TerraceCaller, error) {
	contract, err := bindTerrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TerraceCaller{contract: contract}, nil
}

// NewTerraceTransactor creates a new write-only instance of Terrace, bound to a specific deployed contract.
func NewTerraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TerraceTransactor, error) {
	contract, err := bindTerrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TerraceTransactor{contract: contract}, nil
}

// NewTerraceFilterer creates a new log filterer instance of Terrace, bound to a specific deployed contract.
func NewTerraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TerraceFilterer, error) {
	contract, err := bindTerrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TerraceFilterer{contract: contract}, nil
}

// bindTerrace binds a generic wrapper to an already deployed contract.
func bindTerrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TerraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Terrace *TerraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Terrace.Contract.TerraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Terrace *TerraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Terrace.Contract.TerraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Terrace *TerraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Terrace.Contract.TerraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Terrace *TerraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Terrace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Terrace *TerraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Terrace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Terrace *TerraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Terrace.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Terrace *TerraceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Terrace *TerraceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Terrace.Contract.DEFAULTADMINROLE(&_Terrace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Terrace *TerraceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Terrace.Contract.DEFAULTADMINROLE(&_Terrace.CallOpts)
}

// HUBDOMAIN is a free data retrieval call binding the contract method 0x17f82014.
//
// Solidity: function HUB_DOMAIN() view returns(uint32)
func (_Terrace *TerraceCaller) HUBDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "HUB_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// HUBDOMAIN is a free data retrieval call binding the contract method 0x17f82014.
//
// Solidity: function HUB_DOMAIN() view returns(uint32)
func (_Terrace *TerraceSession) HUBDOMAIN() (uint32, error) {
	return _Terrace.Contract.HUBDOMAIN(&_Terrace.CallOpts)
}

// HUBDOMAIN is a free data retrieval call binding the contract method 0x17f82014.
//
// Solidity: function HUB_DOMAIN() view returns(uint32)
func (_Terrace *TerraceCallerSession) HUBDOMAIN() (uint32, error) {
	return _Terrace.Contract.HUBDOMAIN(&_Terrace.CallOpts)
}

// MESSANGER is a free data retrieval call binding the contract method 0x1404b48f.
//
// Solidity: function MESSANGER() view returns(address)
func (_Terrace *TerraceCaller) MESSANGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "MESSANGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSANGER is a free data retrieval call binding the contract method 0x1404b48f.
//
// Solidity: function MESSANGER() view returns(address)
func (_Terrace *TerraceSession) MESSANGER() (common.Address, error) {
	return _Terrace.Contract.MESSANGER(&_Terrace.CallOpts)
}

// MESSANGER is a free data retrieval call binding the contract method 0x1404b48f.
//
// Solidity: function MESSANGER() view returns(address)
func (_Terrace *TerraceCallerSession) MESSANGER() (common.Address, error) {
	return _Terrace.Contract.MESSANGER(&_Terrace.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Terrace *TerraceCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Terrace *TerraceSession) OPERATORROLE() ([32]byte, error) {
	return _Terrace.Contract.OPERATORROLE(&_Terrace.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Terrace *TerraceCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Terrace.Contract.OPERATORROLE(&_Terrace.CallOpts)
}

// TRANSMITTER is a free data retrieval call binding the contract method 0xf4a5e62c.
//
// Solidity: function TRANSMITTER() view returns(address)
func (_Terrace *TerraceCaller) TRANSMITTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "TRANSMITTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TRANSMITTER is a free data retrieval call binding the contract method 0xf4a5e62c.
//
// Solidity: function TRANSMITTER() view returns(address)
func (_Terrace *TerraceSession) TRANSMITTER() (common.Address, error) {
	return _Terrace.Contract.TRANSMITTER(&_Terrace.CallOpts)
}

// TRANSMITTER is a free data retrieval call binding the contract method 0xf4a5e62c.
//
// Solidity: function TRANSMITTER() view returns(address)
func (_Terrace *TerraceCallerSession) TRANSMITTER() (common.Address, error) {
	return _Terrace.Contract.TRANSMITTER(&_Terrace.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Terrace *TerraceCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Terrace *TerraceSession) USDC() (common.Address, error) {
	return _Terrace.Contract.USDC(&_Terrace.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_Terrace *TerraceCallerSession) USDC() (common.Address, error) {
	return _Terrace.Contract.USDC(&_Terrace.CallOpts)
}

// BridgeNonce is a free data retrieval call binding the contract method 0x7bde3fa7.
//
// Solidity: function bridge_nonce() view returns(uint256)
func (_Terrace *TerraceCaller) BridgeNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "bridge_nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeNonce is a free data retrieval call binding the contract method 0x7bde3fa7.
//
// Solidity: function bridge_nonce() view returns(uint256)
func (_Terrace *TerraceSession) BridgeNonce() (*big.Int, error) {
	return _Terrace.Contract.BridgeNonce(&_Terrace.CallOpts)
}

// BridgeNonce is a free data retrieval call binding the contract method 0x7bde3fa7.
//
// Solidity: function bridge_nonce() view returns(uint256)
func (_Terrace *TerraceCallerSession) BridgeNonce() (*big.Int, error) {
	return _Terrace.Contract.BridgeNonce(&_Terrace.CallOpts)
}

// EncodeErc20Approve is a free data retrieval call binding the contract method 0x70bc332e.
//
// Solidity: function encode_erc20_approve(address target, uint256 amount) pure returns(bytes)
func (_Terrace *TerraceCaller) EncodeErc20Approve(opts *bind.CallOpts, target common.Address, amount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "encode_erc20_approve", target, amount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeErc20Approve is a free data retrieval call binding the contract method 0x70bc332e.
//
// Solidity: function encode_erc20_approve(address target, uint256 amount) pure returns(bytes)
func (_Terrace *TerraceSession) EncodeErc20Approve(target common.Address, amount *big.Int) ([]byte, error) {
	return _Terrace.Contract.EncodeErc20Approve(&_Terrace.CallOpts, target, amount)
}

// EncodeErc20Approve is a free data retrieval call binding the contract method 0x70bc332e.
//
// Solidity: function encode_erc20_approve(address target, uint256 amount) pure returns(bytes)
func (_Terrace *TerraceCallerSession) EncodeErc20Approve(target common.Address, amount *big.Int) ([]byte, error) {
	return _Terrace.Contract.EncodeErc20Approve(&_Terrace.CallOpts, target, amount)
}

// EncodeSupplyAaveV3 is a free data retrieval call binding the contract method 0x9d3ac1eb.
//
// Solidity: function encode_supply_aave_v3(address asset, uint256 amount, address onBehalfOf, uint16 referrer) pure returns(bytes)
func (_Terrace *TerraceCaller) EncodeSupplyAaveV3(opts *bind.CallOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referrer uint16) ([]byte, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "encode_supply_aave_v3", asset, amount, onBehalfOf, referrer)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeSupplyAaveV3 is a free data retrieval call binding the contract method 0x9d3ac1eb.
//
// Solidity: function encode_supply_aave_v3(address asset, uint256 amount, address onBehalfOf, uint16 referrer) pure returns(bytes)
func (_Terrace *TerraceSession) EncodeSupplyAaveV3(asset common.Address, amount *big.Int, onBehalfOf common.Address, referrer uint16) ([]byte, error) {
	return _Terrace.Contract.EncodeSupplyAaveV3(&_Terrace.CallOpts, asset, amount, onBehalfOf, referrer)
}

// EncodeSupplyAaveV3 is a free data retrieval call binding the contract method 0x9d3ac1eb.
//
// Solidity: function encode_supply_aave_v3(address asset, uint256 amount, address onBehalfOf, uint16 referrer) pure returns(bytes)
func (_Terrace *TerraceCallerSession) EncodeSupplyAaveV3(asset common.Address, amount *big.Int, onBehalfOf common.Address, referrer uint16) ([]byte, error) {
	return _Terrace.Contract.EncodeSupplyAaveV3(&_Terrace.CallOpts, asset, amount, onBehalfOf, referrer)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 arg0) view returns(bytes32)
func (_Terrace *TerraceCaller) GetRoleAdmin(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "getRoleAdmin", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 arg0) view returns(bytes32)
func (_Terrace *TerraceSession) GetRoleAdmin(arg0 [32]byte) ([32]byte, error) {
	return _Terrace.Contract.GetRoleAdmin(&_Terrace.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 arg0) view returns(bytes32)
func (_Terrace *TerraceCallerSession) GetRoleAdmin(arg0 [32]byte) ([32]byte, error) {
	return _Terrace.Contract.GetRoleAdmin(&_Terrace.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 arg0, address arg1) view returns(bool)
func (_Terrace *TerraceCaller) HasRole(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "hasRole", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 arg0, address arg1) view returns(bool)
func (_Terrace *TerraceSession) HasRole(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Terrace.Contract.HasRole(&_Terrace.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 arg0, address arg1) view returns(bool)
func (_Terrace *TerraceCallerSession) HasRole(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Terrace.Contract.HasRole(&_Terrace.CallOpts, arg0, arg1)
}

// IsStrategyAllowed is a free data retrieval call binding the contract method 0x71fb6c9f.
//
// Solidity: function is_strategy_allowed(address target, bytes selector, bytes32[] proof) view returns(bool)
func (_Terrace *TerraceCaller) IsStrategyAllowed(opts *bind.CallOpts, target common.Address, selector []byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "is_strategy_allowed", target, selector, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStrategyAllowed is a free data retrieval call binding the contract method 0x71fb6c9f.
//
// Solidity: function is_strategy_allowed(address target, bytes selector, bytes32[] proof) view returns(bool)
func (_Terrace *TerraceSession) IsStrategyAllowed(target common.Address, selector []byte, proof [][32]byte) (bool, error) {
	return _Terrace.Contract.IsStrategyAllowed(&_Terrace.CallOpts, target, selector, proof)
}

// IsStrategyAllowed is a free data retrieval call binding the contract method 0x71fb6c9f.
//
// Solidity: function is_strategy_allowed(address target, bytes selector, bytes32[] proof) view returns(bool)
func (_Terrace *TerraceCallerSession) IsStrategyAllowed(target common.Address, selector []byte, proof [][32]byte) (bool, error) {
	return _Terrace.Contract.IsStrategyAllowed(&_Terrace.CallOpts, target, selector, proof)
}

// StrategyRootHash is a free data retrieval call binding the contract method 0x40e48966.
//
// Solidity: function strategy_root_hash() view returns(bytes32)
func (_Terrace *TerraceCaller) StrategyRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "strategy_root_hash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StrategyRootHash is a free data retrieval call binding the contract method 0x40e48966.
//
// Solidity: function strategy_root_hash() view returns(bytes32)
func (_Terrace *TerraceSession) StrategyRootHash() ([32]byte, error) {
	return _Terrace.Contract.StrategyRootHash(&_Terrace.CallOpts)
}

// StrategyRootHash is a free data retrieval call binding the contract method 0x40e48966.
//
// Solidity: function strategy_root_hash() view returns(bytes32)
func (_Terrace *TerraceCallerSession) StrategyRootHash() ([32]byte, error) {
	return _Terrace.Contract.StrategyRootHash(&_Terrace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interface_id) view returns(bool)
func (_Terrace *TerraceCaller) SupportsInterface(opts *bind.CallOpts, interface_id [4]byte) (bool, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "supportsInterface", interface_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interface_id) view returns(bool)
func (_Terrace *TerraceSession) SupportsInterface(interface_id [4]byte) (bool, error) {
	return _Terrace.Contract.SupportsInterface(&_Terrace.CallOpts, interface_id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interface_id) view returns(bool)
func (_Terrace *TerraceCallerSession) SupportsInterface(interface_id [4]byte) (bool, error) {
	return _Terrace.Contract.SupportsInterface(&_Terrace.CallOpts, interface_id)
}

// UsdcFees is a free data retrieval call binding the contract method 0x51b5a503.
//
// Solidity: function usdc_fees() view returns(uint256)
func (_Terrace *TerraceCaller) UsdcFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "usdc_fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcFees is a free data retrieval call binding the contract method 0x51b5a503.
//
// Solidity: function usdc_fees() view returns(uint256)
func (_Terrace *TerraceSession) UsdcFees() (*big.Int, error) {
	return _Terrace.Contract.UsdcFees(&_Terrace.CallOpts)
}

// UsdcFees is a free data retrieval call binding the contract method 0x51b5a503.
//
// Solidity: function usdc_fees() view returns(uint256)
func (_Terrace *TerraceCallerSession) UsdcFees() (*big.Int, error) {
	return _Terrace.Contract.UsdcFees(&_Terrace.CallOpts)
}

// UsdcReserve is a free data retrieval call binding the contract method 0x29b95248.
//
// Solidity: function usdc_reserve() view returns(uint256)
func (_Terrace *TerraceCaller) UsdcReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "usdc_reserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcReserve is a free data retrieval call binding the contract method 0x29b95248.
//
// Solidity: function usdc_reserve() view returns(uint256)
func (_Terrace *TerraceSession) UsdcReserve() (*big.Int, error) {
	return _Terrace.Contract.UsdcReserve(&_Terrace.CallOpts)
}

// UsdcReserve is a free data retrieval call binding the contract method 0x29b95248.
//
// Solidity: function usdc_reserve() view returns(uint256)
func (_Terrace *TerraceCallerSession) UsdcReserve() (*big.Int, error) {
	return _Terrace.Contract.UsdcReserve(&_Terrace.CallOpts)
}

// UsdcUsage is a free data retrieval call binding the contract method 0x1a98e895.
//
// Solidity: function usdc_usage() view returns(uint256)
func (_Terrace *TerraceCaller) UsdcUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "usdc_usage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcUsage is a free data retrieval call binding the contract method 0x1a98e895.
//
// Solidity: function usdc_usage() view returns(uint256)
func (_Terrace *TerraceSession) UsdcUsage() (*big.Int, error) {
	return _Terrace.Contract.UsdcUsage(&_Terrace.CallOpts)
}

// UsdcUsage is a free data retrieval call binding the contract method 0x1a98e895.
//
// Solidity: function usdc_usage() view returns(uint256)
func (_Terrace *TerraceCallerSession) UsdcUsage() (*big.Int, error) {
	return _Terrace.Contract.UsdcUsage(&_Terrace.CallOpts)
}

// WhitelistedDomains is a free data retrieval call binding the contract method 0xfaa82272.
//
// Solidity: function whitelisted_domains(uint32 arg0) view returns(address)
func (_Terrace *TerraceCaller) WhitelistedDomains(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _Terrace.contract.Call(opts, &out, "whitelisted_domains", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedDomains is a free data retrieval call binding the contract method 0xfaa82272.
//
// Solidity: function whitelisted_domains(uint32 arg0) view returns(address)
func (_Terrace *TerraceSession) WhitelistedDomains(arg0 uint32) (common.Address, error) {
	return _Terrace.Contract.WhitelistedDomains(&_Terrace.CallOpts, arg0)
}

// WhitelistedDomains is a free data retrieval call binding the contract method 0xfaa82272.
//
// Solidity: function whitelisted_domains(uint32 arg0) view returns(address)
func (_Terrace *TerraceCallerSession) WhitelistedDomains(arg0 uint32) (common.Address, error) {
	return _Terrace.Contract.WhitelistedDomains(&_Terrace.CallOpts, arg0)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x3e6aa3dc.
//
// Solidity: function batch_execute(address[] targets, bytes[] selectors, bytes[] data, bytes32[][] proofs) returns()
func (_Terrace *TerraceTransactor) BatchExecute(opts *bind.TransactOpts, targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "batch_execute", targets, selectors, data, proofs)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x3e6aa3dc.
//
// Solidity: function batch_execute(address[] targets, bytes[] selectors, bytes[] data, bytes32[][] proofs) returns()
func (_Terrace *TerraceSession) BatchExecute(targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.BatchExecute(&_Terrace.TransactOpts, targets, selectors, data, proofs)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x3e6aa3dc.
//
// Solidity: function batch_execute(address[] targets, bytes[] selectors, bytes[] data, bytes32[][] proofs) returns()
func (_Terrace *TerraceTransactorSession) BatchExecute(targets []common.Address, selectors [][]byte, data [][]byte, proofs [][][32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.BatchExecute(&_Terrace.TransactOpts, targets, selectors, data, proofs)
}

// ExecuteStrategy is a paid mutator transaction binding the contract method 0x4f494399.
//
// Solidity: function execute_strategy(address target, bytes selector, bytes data, bytes32[] proof) returns()
func (_Terrace *TerraceTransactor) ExecuteStrategy(opts *bind.TransactOpts, target common.Address, selector []byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "execute_strategy", target, selector, data, proof)
}

// ExecuteStrategy is a paid mutator transaction binding the contract method 0x4f494399.
//
// Solidity: function execute_strategy(address target, bytes selector, bytes data, bytes32[] proof) returns()
func (_Terrace *TerraceSession) ExecuteStrategy(target common.Address, selector []byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.ExecuteStrategy(&_Terrace.TransactOpts, target, selector, data, proof)
}

// ExecuteStrategy is a paid mutator transaction binding the contract method 0x4f494399.
//
// Solidity: function execute_strategy(address target, bytes selector, bytes data, bytes32[] proof) returns()
func (_Terrace *TerraceTransactorSession) ExecuteStrategy(target common.Address, selector []byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.ExecuteStrategy(&_Terrace.TransactOpts, target, selector, data, proof)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Terrace *TerraceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Terrace *TerraceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.GrantRole(&_Terrace.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Terrace *TerraceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.GrantRole(&_Terrace.TransactOpts, role, account)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x600ecd47.
//
// Solidity: function receive_message(bytes message, bytes attestation) returns()
func (_Terrace *TerraceTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "receive_message", message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x600ecd47.
//
// Solidity: function receive_message(bytes message, bytes attestation) returns()
func (_Terrace *TerraceSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Terrace.Contract.ReceiveMessage(&_Terrace.TransactOpts, message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x600ecd47.
//
// Solidity: function receive_message(bytes message, bytes attestation) returns()
func (_Terrace *TerraceTransactorSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _Terrace.Contract.ReceiveMessage(&_Terrace.TransactOpts, message, attestation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Terrace *TerraceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Terrace *TerraceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.RenounceRole(&_Terrace.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Terrace *TerraceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.RenounceRole(&_Terrace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Terrace *TerraceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Terrace *TerraceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.RevokeRole(&_Terrace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Terrace *TerraceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.RevokeRole(&_Terrace.TransactOpts, role, account)
}

// SendAllToHub is a paid mutator transaction binding the contract method 0xf289355f.
//
// Solidity: function send_all_to_hub(uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactor) SendAllToHub(opts *bind.TransactOpts, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "send_all_to_hub", max_fee, min_finality_threshold)
}

// SendAllToHub is a paid mutator transaction binding the contract method 0xf289355f.
//
// Solidity: function send_all_to_hub(uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceSession) SendAllToHub(max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendAllToHub(&_Terrace.TransactOpts, max_fee, min_finality_threshold)
}

// SendAllToHub is a paid mutator transaction binding the contract method 0xf289355f.
//
// Solidity: function send_all_to_hub(uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactorSession) SendAllToHub(max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendAllToHub(&_Terrace.TransactOpts, max_fee, min_finality_threshold)
}

// SendAllToTerrace is a paid mutator transaction binding the contract method 0x56053c3c.
//
// Solidity: function send_all_to_terrace(uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactor) SendAllToTerrace(opts *bind.TransactOpts, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "send_all_to_terrace", destination_domain, max_fee, min_finality_threshold)
}

// SendAllToTerrace is a paid mutator transaction binding the contract method 0x56053c3c.
//
// Solidity: function send_all_to_terrace(uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceSession) SendAllToTerrace(destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendAllToTerrace(&_Terrace.TransactOpts, destination_domain, max_fee, min_finality_threshold)
}

// SendAllToTerrace is a paid mutator transaction binding the contract method 0x56053c3c.
//
// Solidity: function send_all_to_terrace(uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactorSession) SendAllToTerrace(destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendAllToTerrace(&_Terrace.TransactOpts, destination_domain, max_fee, min_finality_threshold)
}

// SendToHub is a paid mutator transaction binding the contract method 0xd6164993.
//
// Solidity: function send_to_hub(uint256 amount, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactor) SendToHub(opts *bind.TransactOpts, amount *big.Int, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "send_to_hub", amount, max_fee, min_finality_threshold)
}

// SendToHub is a paid mutator transaction binding the contract method 0xd6164993.
//
// Solidity: function send_to_hub(uint256 amount, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceSession) SendToHub(amount *big.Int, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendToHub(&_Terrace.TransactOpts, amount, max_fee, min_finality_threshold)
}

// SendToHub is a paid mutator transaction binding the contract method 0xd6164993.
//
// Solidity: function send_to_hub(uint256 amount, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactorSession) SendToHub(amount *big.Int, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendToHub(&_Terrace.TransactOpts, amount, max_fee, min_finality_threshold)
}

// SendToTerrace is a paid mutator transaction binding the contract method 0x6d7b50cf.
//
// Solidity: function send_to_terrace(uint256 amount, uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactor) SendToTerrace(opts *bind.TransactOpts, amount *big.Int, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "send_to_terrace", amount, destination_domain, max_fee, min_finality_threshold)
}

// SendToTerrace is a paid mutator transaction binding the contract method 0x6d7b50cf.
//
// Solidity: function send_to_terrace(uint256 amount, uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceSession) SendToTerrace(amount *big.Int, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendToTerrace(&_Terrace.TransactOpts, amount, destination_domain, max_fee, min_finality_threshold)
}

// SendToTerrace is a paid mutator transaction binding the contract method 0x6d7b50cf.
//
// Solidity: function send_to_terrace(uint256 amount, uint32 destination_domain, uint256 max_fee, uint32 min_finality_threshold) returns()
func (_Terrace *TerraceTransactorSession) SendToTerrace(amount *big.Int, destination_domain uint32, max_fee *big.Int, min_finality_threshold uint32) (*types.Transaction, error) {
	return _Terrace.Contract.SendToTerrace(&_Terrace.TransactOpts, amount, destination_domain, max_fee, min_finality_threshold)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0xc2f9e63f.
//
// Solidity: function set_role_admin(bytes32 role, bytes32 admin_role) returns()
func (_Terrace *TerraceTransactor) SetRoleAdmin(opts *bind.TransactOpts, role [32]byte, admin_role [32]byte) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "set_role_admin", role, admin_role)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0xc2f9e63f.
//
// Solidity: function set_role_admin(bytes32 role, bytes32 admin_role) returns()
func (_Terrace *TerraceSession) SetRoleAdmin(role [32]byte, admin_role [32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.SetRoleAdmin(&_Terrace.TransactOpts, role, admin_role)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0xc2f9e63f.
//
// Solidity: function set_role_admin(bytes32 role, bytes32 admin_role) returns()
func (_Terrace *TerraceTransactorSession) SetRoleAdmin(role [32]byte, admin_role [32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.SetRoleAdmin(&_Terrace.TransactOpts, role, admin_role)
}

// UpdateRootHash is a paid mutator transaction binding the contract method 0xc069315d.
//
// Solidity: function update_root_hash(bytes32 root) returns()
func (_Terrace *TerraceTransactor) UpdateRootHash(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "update_root_hash", root)
}

// UpdateRootHash is a paid mutator transaction binding the contract method 0xc069315d.
//
// Solidity: function update_root_hash(bytes32 root) returns()
func (_Terrace *TerraceSession) UpdateRootHash(root [32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.UpdateRootHash(&_Terrace.TransactOpts, root)
}

// UpdateRootHash is a paid mutator transaction binding the contract method 0xc069315d.
//
// Solidity: function update_root_hash(bytes32 root) returns()
func (_Terrace *TerraceTransactorSession) UpdateRootHash(root [32]byte) (*types.Transaction, error) {
	return _Terrace.Contract.UpdateRootHash(&_Terrace.TransactOpts, root)
}

// UpdateWhitelistedDomain is a paid mutator transaction binding the contract method 0x1b52da6e.
//
// Solidity: function update_whitelisted_domain(uint32 domain, address terrace) returns()
func (_Terrace *TerraceTransactor) UpdateWhitelistedDomain(opts *bind.TransactOpts, domain uint32, terrace common.Address) (*types.Transaction, error) {
	return _Terrace.contract.Transact(opts, "update_whitelisted_domain", domain, terrace)
}

// UpdateWhitelistedDomain is a paid mutator transaction binding the contract method 0x1b52da6e.
//
// Solidity: function update_whitelisted_domain(uint32 domain, address terrace) returns()
func (_Terrace *TerraceSession) UpdateWhitelistedDomain(domain uint32, terrace common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.UpdateWhitelistedDomain(&_Terrace.TransactOpts, domain, terrace)
}

// UpdateWhitelistedDomain is a paid mutator transaction binding the contract method 0x1b52da6e.
//
// Solidity: function update_whitelisted_domain(uint32 domain, address terrace) returns()
func (_Terrace *TerraceTransactorSession) UpdateWhitelistedDomain(domain uint32, terrace common.Address) (*types.Transaction, error) {
	return _Terrace.Contract.UpdateWhitelistedDomain(&_Terrace.TransactOpts, domain, terrace)
}

// TerraceReceivedFundsIterator is returned from FilterReceivedFunds and is used to iterate over the raw logs and unpacked data for ReceivedFunds events raised by the Terrace contract.
type TerraceReceivedFundsIterator struct {
	Event *TerraceReceivedFunds // Event containing the contract specifics and raw log

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
func (it *TerraceReceivedFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TerraceReceivedFunds)
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
		it.Event = new(TerraceReceivedFunds)
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
func (it *TerraceReceivedFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TerraceReceivedFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TerraceReceivedFunds represents a ReceivedFunds event raised by the Terrace contract.
type TerraceReceivedFunds struct {
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
func (_Terrace *TerraceFilterer) FilterReceivedFunds(opts *bind.FilterOpts) (*TerraceReceivedFundsIterator, error) {

	logs, sub, err := _Terrace.contract.FilterLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return &TerraceReceivedFundsIterator{contract: _Terrace.contract, event: "ReceivedFunds", logs: logs, sub: sub}, nil
}

// WatchReceivedFunds is a free log subscription operation binding the contract event 0x2b8fa6bde978d1a0c02f660e1cb6f232a39b18669bd77f93ddf5c2324f0445b3.
//
// Solidity: event ReceivedFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Terrace *TerraceFilterer) WatchReceivedFunds(opts *bind.WatchOpts, sink chan<- *TerraceReceivedFunds) (event.Subscription, error) {

	logs, sub, err := _Terrace.contract.WatchLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TerraceReceivedFunds)
				if err := _Terrace.contract.UnpackLog(event, "ReceivedFunds", log); err != nil {
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
func (_Terrace *TerraceFilterer) ParseReceivedFunds(log types.Log) (*TerraceReceivedFunds, error) {
	event := new(TerraceReceivedFunds)
	if err := _Terrace.contract.UnpackLog(event, "ReceivedFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TerraceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Terrace contract.
type TerraceRoleAdminChangedIterator struct {
	Event *TerraceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TerraceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TerraceRoleAdminChanged)
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
		it.Event = new(TerraceRoleAdminChanged)
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
func (it *TerraceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TerraceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TerraceRoleAdminChanged represents a RoleAdminChanged event raised by the Terrace contract.
type TerraceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Terrace *TerraceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TerraceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Terrace.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TerraceRoleAdminChangedIterator{contract: _Terrace.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Terrace *TerraceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TerraceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Terrace.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TerraceRoleAdminChanged)
				if err := _Terrace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Terrace *TerraceFilterer) ParseRoleAdminChanged(log types.Log) (*TerraceRoleAdminChanged, error) {
	event := new(TerraceRoleAdminChanged)
	if err := _Terrace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TerraceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Terrace contract.
type TerraceRoleGrantedIterator struct {
	Event *TerraceRoleGranted // Event containing the contract specifics and raw log

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
func (it *TerraceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TerraceRoleGranted)
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
		it.Event = new(TerraceRoleGranted)
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
func (it *TerraceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TerraceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TerraceRoleGranted represents a RoleGranted event raised by the Terrace contract.
type TerraceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Terrace *TerraceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TerraceRoleGrantedIterator, error) {

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

	logs, sub, err := _Terrace.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TerraceRoleGrantedIterator{contract: _Terrace.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Terrace *TerraceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TerraceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Terrace.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TerraceRoleGranted)
				if err := _Terrace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Terrace *TerraceFilterer) ParseRoleGranted(log types.Log) (*TerraceRoleGranted, error) {
	event := new(TerraceRoleGranted)
	if err := _Terrace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TerraceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Terrace contract.
type TerraceRoleRevokedIterator struct {
	Event *TerraceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TerraceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TerraceRoleRevoked)
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
		it.Event = new(TerraceRoleRevoked)
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
func (it *TerraceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TerraceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TerraceRoleRevoked represents a RoleRevoked event raised by the Terrace contract.
type TerraceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Terrace *TerraceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TerraceRoleRevokedIterator, error) {

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

	logs, sub, err := _Terrace.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TerraceRoleRevokedIterator{contract: _Terrace.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Terrace *TerraceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TerraceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Terrace.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TerraceRoleRevoked)
				if err := _Terrace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Terrace *TerraceFilterer) ParseRoleRevoked(log types.Log) (*TerraceRoleRevoked, error) {
	event := new(TerraceRoleRevoked)
	if err := _Terrace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TerraceSendFundsIterator is returned from FilterSendFunds and is used to iterate over the raw logs and unpacked data for SendFunds events raised by the Terrace contract.
type TerraceSendFundsIterator struct {
	Event *TerraceSendFunds // Event containing the contract specifics and raw log

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
func (it *TerraceSendFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TerraceSendFunds)
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
		it.Event = new(TerraceSendFunds)
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
func (it *TerraceSendFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TerraceSendFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TerraceSendFunds represents a SendFunds event raised by the Terrace contract.
type TerraceSendFunds struct {
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
func (_Terrace *TerraceFilterer) FilterSendFunds(opts *bind.FilterOpts) (*TerraceSendFundsIterator, error) {

	logs, sub, err := _Terrace.contract.FilterLogs(opts, "SendFunds")
	if err != nil {
		return nil, err
	}
	return &TerraceSendFundsIterator{contract: _Terrace.contract, event: "SendFunds", logs: logs, sub: sub}, nil
}

// WatchSendFunds is a free log subscription operation binding the contract event 0x68628b1c35f23c5db46c17659a5cdc52eb8e303a889a8d3f29495c558cdd9d22.
//
// Solidity: event SendFunds(bytes32 id, address from_terrace, uint32 from_domain, address to_terrace, uint32 to_domain, uint256 amount)
func (_Terrace *TerraceFilterer) WatchSendFunds(opts *bind.WatchOpts, sink chan<- *TerraceSendFunds) (event.Subscription, error) {

	logs, sub, err := _Terrace.contract.WatchLogs(opts, "SendFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TerraceSendFunds)
				if err := _Terrace.contract.UnpackLog(event, "SendFunds", log); err != nil {
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
func (_Terrace *TerraceFilterer) ParseSendFunds(log types.Log) (*TerraceSendFunds, error) {
	event := new(TerraceSendFunds)
	if err := _Terrace.contract.UnpackLog(event, "SendFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TerraceStrategyRootHashUpdatedIterator is returned from FilterStrategyRootHashUpdated and is used to iterate over the raw logs and unpacked data for StrategyRootHashUpdated events raised by the Terrace contract.
type TerraceStrategyRootHashUpdatedIterator struct {
	Event *TerraceStrategyRootHashUpdated // Event containing the contract specifics and raw log

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
func (it *TerraceStrategyRootHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TerraceStrategyRootHashUpdated)
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
		it.Event = new(TerraceStrategyRootHashUpdated)
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
func (it *TerraceStrategyRootHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TerraceStrategyRootHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TerraceStrategyRootHashUpdated represents a StrategyRootHashUpdated event raised by the Terrace contract.
type TerraceStrategyRootHashUpdated struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStrategyRootHashUpdated is a free log retrieval operation binding the contract event 0xf2041d1a2718102b7ab7f4756f12e97ddbc269be1fa70d86a6065a0c88921d1c.
//
// Solidity: event StrategyRootHashUpdated(bytes32 indexed root)
func (_Terrace *TerraceFilterer) FilterStrategyRootHashUpdated(opts *bind.FilterOpts, root [][32]byte) (*TerraceStrategyRootHashUpdatedIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Terrace.contract.FilterLogs(opts, "StrategyRootHashUpdated", rootRule)
	if err != nil {
		return nil, err
	}
	return &TerraceStrategyRootHashUpdatedIterator{contract: _Terrace.contract, event: "StrategyRootHashUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategyRootHashUpdated is a free log subscription operation binding the contract event 0xf2041d1a2718102b7ab7f4756f12e97ddbc269be1fa70d86a6065a0c88921d1c.
//
// Solidity: event StrategyRootHashUpdated(bytes32 indexed root)
func (_Terrace *TerraceFilterer) WatchStrategyRootHashUpdated(opts *bind.WatchOpts, sink chan<- *TerraceStrategyRootHashUpdated, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Terrace.contract.WatchLogs(opts, "StrategyRootHashUpdated", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TerraceStrategyRootHashUpdated)
				if err := _Terrace.contract.UnpackLog(event, "StrategyRootHashUpdated", log); err != nil {
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
func (_Terrace *TerraceFilterer) ParseStrategyRootHashUpdated(log types.Log) (*TerraceStrategyRootHashUpdated, error) {
	event := new(TerraceStrategyRootHashUpdated)
	if err := _Terrace.contract.UnpackLog(event, "StrategyRootHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
