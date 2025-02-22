// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TaikoL2Config is an auto generated low-level Go binding around an user-defined struct.
type TaikoL2Config struct {
	GasTargetPerL1Block       uint32
	BasefeeAdjustmentQuotient uint8
}

// TaikoL2ClientMetaData contains all meta data concerning the TaikoL2Client contract.
var TaikoL2ClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"GOLDEN_TOUCH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOLDEN_TOUCH_PRIVATEKEY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"anchor\",\"inputs\":[{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1SignalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1Height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parentGasUsed\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeCrossChainTransaction\",\"inputs\":[{\"name\":\"txId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"txdata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gasExcess\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBasefee\",\"inputs\":[{\"name\":\"l1Height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parentGasUsed\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"basefee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockHash\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structTaikoL2.Config\",\"components\":[{\"name\":\"gasTargetPerL1Block\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"basefeeAdjustmentQuotient\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSyncedSnippet\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainSync.Snippet\",\"components\":[{\"name\":\"remoteBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_addressManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1ChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_gasExcess\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l2Hashes\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestSyncedL1Height\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTxId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicInputHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signAnchor\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"k\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"s\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"skipFeeCheck\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"snippets\",\"inputs\":[{\"name\":\"l1height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"remoteBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Anchored\",\"inputs\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"gasExcess\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossChainSynced\",\"inputs\":[{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionExecuted\",\"inputs\":[{\"name\":\"txId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EIP1559_INVALID_PARAMS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ETH_TRANSFER_FAILED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_PAUSE_STATUS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_BASEFEE_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_INVALID_CHAIN_ID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_INVALID_GOLDEN_TOUCH_K\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_INVALID_PARAM\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_INVALID_SENDER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_PUBLIC_INPUT_HASH_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2_TOO_LATE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Overflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"REENTRANT_CALL\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_DENIED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_INVALID_MANAGER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_UNEXPECTED_CHAINID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_ZERO_ADDR\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"XCO_INVALID_OWNER_CHAINID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"XCO_INVALID_TX_ID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"XCO_PERMISSION_DENIED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"XCO_TX_REVERTED\",\"inputs\":[]}]",
}

// TaikoL2ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoL2ClientMetaData.ABI instead.
var TaikoL2ClientABI = TaikoL2ClientMetaData.ABI

// TaikoL2Client is an auto generated Go binding around an Ethereum contract.
type TaikoL2Client struct {
	TaikoL2ClientCaller     // Read-only binding to the contract
	TaikoL2ClientTransactor // Write-only binding to the contract
	TaikoL2ClientFilterer   // Log filterer for contract events
}

// TaikoL2ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoL2ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL2ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoL2ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL2ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoL2ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL2ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoL2ClientSession struct {
	Contract     *TaikoL2Client    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaikoL2ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoL2ClientCallerSession struct {
	Contract *TaikoL2ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TaikoL2ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoL2ClientTransactorSession struct {
	Contract     *TaikoL2ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TaikoL2ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoL2ClientRaw struct {
	Contract *TaikoL2Client // Generic contract binding to access the raw methods on
}

// TaikoL2ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoL2ClientCallerRaw struct {
	Contract *TaikoL2ClientCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoL2ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoL2ClientTransactorRaw struct {
	Contract *TaikoL2ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoL2Client creates a new instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2Client(address common.Address, backend bind.ContractBackend) (*TaikoL2Client, error) {
	contract, err := bindTaikoL2Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoL2Client{TaikoL2ClientCaller: TaikoL2ClientCaller{contract: contract}, TaikoL2ClientTransactor: TaikoL2ClientTransactor{contract: contract}, TaikoL2ClientFilterer: TaikoL2ClientFilterer{contract: contract}}, nil
}

// NewTaikoL2ClientCaller creates a new read-only instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2ClientCaller(address common.Address, caller bind.ContractCaller) (*TaikoL2ClientCaller, error) {
	contract, err := bindTaikoL2Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientCaller{contract: contract}, nil
}

// NewTaikoL2ClientTransactor creates a new write-only instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoL2ClientTransactor, error) {
	contract, err := bindTaikoL2Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientTransactor{contract: contract}, nil
}

// NewTaikoL2ClientFilterer creates a new log filterer instance of TaikoL2Client, bound to a specific deployed contract.
func NewTaikoL2ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoL2ClientFilterer, error) {
	contract, err := bindTaikoL2Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientFilterer{contract: contract}, nil
}

// bindTaikoL2Client binds a generic wrapper to an already deployed contract.
func bindTaikoL2Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoL2ClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL2Client *TaikoL2ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL2Client.Contract.TaikoL2ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL2Client *TaikoL2ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TaikoL2ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL2Client *TaikoL2ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TaikoL2ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL2Client *TaikoL2ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL2Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL2Client *TaikoL2ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL2Client *TaikoL2ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.contract.Transact(opts, method, params...)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) GOLDENTOUCHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "GOLDEN_TOUCH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHADDRESS(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHADDRESS is a free data retrieval call binding the contract method 0x9ee512f2.
//
// Solidity: function GOLDEN_TOUCH_ADDRESS() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GOLDENTOUCHADDRESS() (common.Address, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHADDRESS(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_TaikoL2Client *TaikoL2ClientCaller) GOLDENTOUCHPRIVATEKEY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "GOLDEN_TOUCH_PRIVATEKEY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_TaikoL2Client *TaikoL2ClientSession) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHPRIVATEKEY(&_TaikoL2Client.CallOpts)
}

// GOLDENTOUCHPRIVATEKEY is a free data retrieval call binding the contract method 0x10da3738.
//
// Solidity: function GOLDEN_TOUCH_PRIVATEKEY() view returns(uint256)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GOLDENTOUCHPRIVATEKEY() (*big.Int, error) {
	return _TaikoL2Client.Contract.GOLDENTOUCHPRIVATEKEY(&_TaikoL2Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) AddressManager() (common.Address, error) {
	return _TaikoL2Client.Contract.AddressManager(&_TaikoL2Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) AddressManager() (common.Address, error) {
	return _TaikoL2Client.Contract.AddressManager(&_TaikoL2Client.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCaller) GasExcess(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "gasExcess")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientSession) GasExcess() (uint64, error) {
	return _TaikoL2Client.Contract.GasExcess(&_TaikoL2Client.CallOpts)
}

// GasExcess is a free data retrieval call binding the contract method 0xf535bd56.
//
// Solidity: function gasExcess() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GasExcess() (uint64, error) {
	return _TaikoL2Client.Contract.GasExcess(&_TaikoL2Client.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0xa7e022d1.
//
// Solidity: function getBasefee(uint64 l1Height, uint32 parentGasUsed) view returns(uint256 basefee)
func (_TaikoL2Client *TaikoL2ClientCaller) GetBasefee(opts *bind.CallOpts, l1Height uint64, parentGasUsed uint32) (*big.Int, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getBasefee", l1Height, parentGasUsed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0xa7e022d1.
//
// Solidity: function getBasefee(uint64 l1Height, uint32 parentGasUsed) view returns(uint256 basefee)
func (_TaikoL2Client *TaikoL2ClientSession) GetBasefee(l1Height uint64, parentGasUsed uint32) (*big.Int, error) {
	return _TaikoL2Client.Contract.GetBasefee(&_TaikoL2Client.CallOpts, l1Height, parentGasUsed)
}

// GetBasefee is a free data retrieval call binding the contract method 0xa7e022d1.
//
// Solidity: function getBasefee(uint64 l1Height, uint32 parentGasUsed) view returns(uint256 basefee)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetBasefee(l1Height uint64, parentGasUsed uint32) (*big.Int, error) {
	return _TaikoL2Client.Contract.GetBasefee(&_TaikoL2Client.CallOpts, l1Height, parentGasUsed)
}

// GetBlockHash is a free data retrieval call binding the contract method 0x23ac7136.
//
// Solidity: function getBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCaller) GetBlockHash(opts *bind.CallOpts, blockId uint64) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getBlockHash", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0x23ac7136.
//
// Solidity: function getBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientSession) GetBlockHash(blockId uint64) ([32]byte, error) {
	return _TaikoL2Client.Contract.GetBlockHash(&_TaikoL2Client.CallOpts, blockId)
}

// GetBlockHash is a free data retrieval call binding the contract method 0x23ac7136.
//
// Solidity: function getBlockHash(uint64 blockId) view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetBlockHash(blockId uint64) ([32]byte, error) {
	return _TaikoL2Client.Contract.GetBlockHash(&_TaikoL2Client.CallOpts, blockId)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint32,uint8) config)
func (_TaikoL2Client *TaikoL2ClientCaller) GetConfig(opts *bind.CallOpts) (TaikoL2Config, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(TaikoL2Config), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoL2Config)).(*TaikoL2Config)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint32,uint8) config)
func (_TaikoL2Client *TaikoL2ClientSession) GetConfig() (TaikoL2Config, error) {
	return _TaikoL2Client.Contract.GetConfig(&_TaikoL2Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint32,uint8) config)
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetConfig() (TaikoL2Config, error) {
	return _TaikoL2Client.Contract.GetConfig(&_TaikoL2Client.CallOpts)
}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((uint64,uint64,bytes32,bytes32))
func (_TaikoL2Client *TaikoL2ClientCaller) GetSyncedSnippet(opts *bind.CallOpts, blockId uint64) (ICrossChainSyncSnippet, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "getSyncedSnippet", blockId)

	if err != nil {
		return *new(ICrossChainSyncSnippet), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainSyncSnippet)).(*ICrossChainSyncSnippet)

	return out0, err

}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((uint64,uint64,bytes32,bytes32))
func (_TaikoL2Client *TaikoL2ClientSession) GetSyncedSnippet(blockId uint64) (ICrossChainSyncSnippet, error) {
	return _TaikoL2Client.Contract.GetSyncedSnippet(&_TaikoL2Client.CallOpts, blockId)
}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((uint64,uint64,bytes32,bytes32))
func (_TaikoL2Client *TaikoL2ClientCallerSession) GetSyncedSnippet(blockId uint64) (ICrossChainSyncSnippet, error) {
	return _TaikoL2Client.Contract.GetSyncedSnippet(&_TaikoL2Client.CallOpts, blockId)
}

// L2Hashes is a free data retrieval call binding the contract method 0x8551f41e.
//
// Solidity: function l2Hashes(uint256 blockId) view returns(bytes32 blockHash)
func (_TaikoL2Client *TaikoL2ClientCaller) L2Hashes(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "l2Hashes", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2Hashes is a free data retrieval call binding the contract method 0x8551f41e.
//
// Solidity: function l2Hashes(uint256 blockId) view returns(bytes32 blockHash)
func (_TaikoL2Client *TaikoL2ClientSession) L2Hashes(blockId *big.Int) ([32]byte, error) {
	return _TaikoL2Client.Contract.L2Hashes(&_TaikoL2Client.CallOpts, blockId)
}

// L2Hashes is a free data retrieval call binding the contract method 0x8551f41e.
//
// Solidity: function l2Hashes(uint256 blockId) view returns(bytes32 blockHash)
func (_TaikoL2Client *TaikoL2ClientCallerSession) L2Hashes(blockId *big.Int) ([32]byte, error) {
	return _TaikoL2Client.Contract.L2Hashes(&_TaikoL2Client.CallOpts, blockId)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCaller) LatestSyncedL1Height(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "latestSyncedL1Height")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientSession) LatestSyncedL1Height() (uint64, error) {
	return _TaikoL2Client.Contract.LatestSyncedL1Height(&_TaikoL2Client.CallOpts)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCallerSession) LatestSyncedL1Height() (uint64, error) {
	return _TaikoL2Client.Contract.LatestSyncedL1Height(&_TaikoL2Client.CallOpts)
}

// NextTxId is a free data retrieval call binding the contract method 0x8aff87b2.
//
// Solidity: function nextTxId() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCaller) NextTxId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "nextTxId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextTxId is a free data retrieval call binding the contract method 0x8aff87b2.
//
// Solidity: function nextTxId() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientSession) NextTxId() (uint64, error) {
	return _TaikoL2Client.Contract.NextTxId(&_TaikoL2Client.CallOpts)
}

// NextTxId is a free data retrieval call binding the contract method 0x8aff87b2.
//
// Solidity: function nextTxId() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCallerSession) NextTxId() (uint64, error) {
	return _TaikoL2Client.Contract.NextTxId(&_TaikoL2Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL2Client *TaikoL2ClientSession) Owner() (common.Address, error) {
	return _TaikoL2Client.Contract.Owner(&_TaikoL2Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Owner() (common.Address, error) {
	return _TaikoL2Client.Contract.Owner(&_TaikoL2Client.CallOpts)
}

// OwnerChainId is a free data retrieval call binding the contract method 0xff4d1815.
//
// Solidity: function ownerChainId() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCaller) OwnerChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "ownerChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OwnerChainId is a free data retrieval call binding the contract method 0xff4d1815.
//
// Solidity: function ownerChainId() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientSession) OwnerChainId() (uint64, error) {
	return _TaikoL2Client.Contract.OwnerChainId(&_TaikoL2Client.CallOpts)
}

// OwnerChainId is a free data retrieval call binding the contract method 0xff4d1815.
//
// Solidity: function ownerChainId() view returns(uint64)
func (_TaikoL2Client *TaikoL2ClientCallerSession) OwnerChainId() (uint64, error) {
	return _TaikoL2Client.Contract.OwnerChainId(&_TaikoL2Client.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoL2Client *TaikoL2ClientCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoL2Client *TaikoL2ClientSession) Paused() (bool, error) {
	return _TaikoL2Client.Contract.Paused(&_TaikoL2Client.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Paused() (bool, error) {
	return _TaikoL2Client.Contract.Paused(&_TaikoL2Client.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientSession) ProxiableUUID() ([32]byte, error) {
	return _TaikoL2Client.Contract.ProxiableUUID(&_TaikoL2Client.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TaikoL2Client.Contract.ProxiableUUID(&_TaikoL2Client.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCaller) PublicInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "publicInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientSession) PublicInputHash() ([32]byte, error) {
	return _TaikoL2Client.Contract.PublicInputHash(&_TaikoL2Client.CallOpts)
}

// PublicInputHash is a free data retrieval call binding the contract method 0xdac5df78.
//
// Solidity: function publicInputHash() view returns(bytes32)
func (_TaikoL2Client *TaikoL2ClientCallerSession) PublicInputHash() ([32]byte, error) {
	return _TaikoL2Client.Contract.PublicInputHash(&_TaikoL2Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCaller) Resolve(opts *bind.CallOpts, chainId uint64, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientSession) Resolve(chainId uint64, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve(&_TaikoL2Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Resolve(chainId uint64, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve(&_TaikoL2Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve0(&_TaikoL2Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL2Client.Contract.Resolve0(&_TaikoL2Client.CallOpts, name, allowZeroAddress)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL2Client *TaikoL2ClientCaller) SignAnchor(opts *bind.CallOpts, digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "signAnchor", digest, k)

	outstruct := new(struct {
		V uint8
		R *big.Int
		S *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL2Client *TaikoL2ClientSession) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL2Client.Contract.SignAnchor(&_TaikoL2Client.CallOpts, digest, k)
}

// SignAnchor is a free data retrieval call binding the contract method 0x591aad8a.
//
// Solidity: function signAnchor(bytes32 digest, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_TaikoL2Client *TaikoL2ClientCallerSession) SignAnchor(digest [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _TaikoL2Client.Contract.SignAnchor(&_TaikoL2Client.CallOpts, digest, k)
}

// SkipFeeCheck is a free data retrieval call binding the contract method 0x2f980473.
//
// Solidity: function skipFeeCheck() pure returns(bool)
func (_TaikoL2Client *TaikoL2ClientCaller) SkipFeeCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "skipFeeCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SkipFeeCheck is a free data retrieval call binding the contract method 0x2f980473.
//
// Solidity: function skipFeeCheck() pure returns(bool)
func (_TaikoL2Client *TaikoL2ClientSession) SkipFeeCheck() (bool, error) {
	return _TaikoL2Client.Contract.SkipFeeCheck(&_TaikoL2Client.CallOpts)
}

// SkipFeeCheck is a free data retrieval call binding the contract method 0x2f980473.
//
// Solidity: function skipFeeCheck() pure returns(bool)
func (_TaikoL2Client *TaikoL2ClientCallerSession) SkipFeeCheck() (bool, error) {
	return _TaikoL2Client.Contract.SkipFeeCheck(&_TaikoL2Client.CallOpts)
}

// Snippets is a free data retrieval call binding the contract method 0xe8e2c5fb.
//
// Solidity: function snippets(uint256 l1height) view returns(uint64 remoteBlockId, uint64 syncedInBlock, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientCaller) Snippets(opts *bind.CallOpts, l1height *big.Int) (struct {
	RemoteBlockId uint64
	SyncedInBlock uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
}, error) {
	var out []interface{}
	err := _TaikoL2Client.contract.Call(opts, &out, "snippets", l1height)

	outstruct := new(struct {
		RemoteBlockId uint64
		SyncedInBlock uint64
		BlockHash     [32]byte
		SignalRoot    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RemoteBlockId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SyncedInBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.SignalRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Snippets is a free data retrieval call binding the contract method 0xe8e2c5fb.
//
// Solidity: function snippets(uint256 l1height) view returns(uint64 remoteBlockId, uint64 syncedInBlock, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientSession) Snippets(l1height *big.Int) (struct {
	RemoteBlockId uint64
	SyncedInBlock uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
}, error) {
	return _TaikoL2Client.Contract.Snippets(&_TaikoL2Client.CallOpts, l1height)
}

// Snippets is a free data retrieval call binding the contract method 0xe8e2c5fb.
//
// Solidity: function snippets(uint256 l1height) view returns(uint64 remoteBlockId, uint64 syncedInBlock, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientCallerSession) Snippets(l1height *big.Int) (struct {
	RemoteBlockId uint64
	SyncedInBlock uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
}, error) {
	return _TaikoL2Client.Contract.Snippets(&_TaikoL2Client.CallOpts, l1height)
}

// Anchor is a paid mutator transaction binding the contract method 0xda69d3db.
//
// Solidity: function anchor(bytes32 l1BlockHash, bytes32 l1SignalRoot, uint64 l1Height, uint32 parentGasUsed) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Anchor(opts *bind.TransactOpts, l1BlockHash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint32) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "anchor", l1BlockHash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0xda69d3db.
//
// Solidity: function anchor(bytes32 l1BlockHash, bytes32 l1SignalRoot, uint64 l1Height, uint32 parentGasUsed) returns()
func (_TaikoL2Client *TaikoL2ClientSession) Anchor(l1BlockHash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint32) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Anchor(&_TaikoL2Client.TransactOpts, l1BlockHash, l1SignalRoot, l1Height, parentGasUsed)
}

// Anchor is a paid mutator transaction binding the contract method 0xda69d3db.
//
// Solidity: function anchor(bytes32 l1BlockHash, bytes32 l1SignalRoot, uint64 l1Height, uint32 parentGasUsed) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Anchor(l1BlockHash [32]byte, l1SignalRoot [32]byte, l1Height uint64, parentGasUsed uint32) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Anchor(&_TaikoL2Client.TransactOpts, l1BlockHash, l1SignalRoot, l1Height, parentGasUsed)
}

// ExecuteCrossChainTransaction is a paid mutator transaction binding the contract method 0xc7a2fc60.
//
// Solidity: function executeCrossChainTransaction(uint64 txId, bytes txdata) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) ExecuteCrossChainTransaction(opts *bind.TransactOpts, txId uint64, txdata []byte) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "executeCrossChainTransaction", txId, txdata)
}

// ExecuteCrossChainTransaction is a paid mutator transaction binding the contract method 0xc7a2fc60.
//
// Solidity: function executeCrossChainTransaction(uint64 txId, bytes txdata) returns()
func (_TaikoL2Client *TaikoL2ClientSession) ExecuteCrossChainTransaction(txId uint64, txdata []byte) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.ExecuteCrossChainTransaction(&_TaikoL2Client.TransactOpts, txId, txdata)
}

// ExecuteCrossChainTransaction is a paid mutator transaction binding the contract method 0xc7a2fc60.
//
// Solidity: function executeCrossChainTransaction(uint64 txId, bytes txdata) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) ExecuteCrossChainTransaction(txId uint64, txdata []byte) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.ExecuteCrossChainTransaction(&_TaikoL2Client.TransactOpts, txId, txdata)
}

// Init is a paid mutator transaction binding the contract method 0xb310e9e9.
//
// Solidity: function init(address _addressManager, uint64 _l1ChainId, uint64 _gasExcess) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _l1ChainId uint64, _gasExcess uint64) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "init", _addressManager, _l1ChainId, _gasExcess)
}

// Init is a paid mutator transaction binding the contract method 0xb310e9e9.
//
// Solidity: function init(address _addressManager, uint64 _l1ChainId, uint64 _gasExcess) returns()
func (_TaikoL2Client *TaikoL2ClientSession) Init(_addressManager common.Address, _l1ChainId uint64, _gasExcess uint64) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Init(&_TaikoL2Client.TransactOpts, _addressManager, _l1ChainId, _gasExcess)
}

// Init is a paid mutator transaction binding the contract method 0xb310e9e9.
//
// Solidity: function init(address _addressManager, uint64 _l1ChainId, uint64 _gasExcess) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Init(_addressManager common.Address, _l1ChainId uint64, _gasExcess uint64) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Init(&_TaikoL2Client.TransactOpts, _addressManager, _l1ChainId, _gasExcess)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoL2Client *TaikoL2ClientSession) Pause() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Pause(&_TaikoL2Client.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Pause() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Pause(&_TaikoL2Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL2Client *TaikoL2ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.RenounceOwnership(&_TaikoL2Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.RenounceOwnership(&_TaikoL2Client.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL2Client *TaikoL2ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TransferOwnership(&_TaikoL2Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.TransferOwnership(&_TaikoL2Client.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoL2Client *TaikoL2ClientSession) Unpause() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Unpause(&_TaikoL2Client.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Unpause() (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Unpause(&_TaikoL2Client.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TaikoL2Client *TaikoL2ClientSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.UpgradeTo(&_TaikoL2Client.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.UpgradeTo(&_TaikoL2Client.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TaikoL2Client *TaikoL2ClientSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.UpgradeToAndCall(&_TaikoL2Client.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.UpgradeToAndCall(&_TaikoL2Client.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address token, address to) returns()
func (_TaikoL2Client *TaikoL2ClientTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.contract.Transact(opts, "withdraw", token, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address token, address to) returns()
func (_TaikoL2Client *TaikoL2ClientSession) Withdraw(token common.Address, to common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Withdraw(&_TaikoL2Client.TransactOpts, token, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address token, address to) returns()
func (_TaikoL2Client *TaikoL2ClientTransactorSession) Withdraw(token common.Address, to common.Address) (*types.Transaction, error) {
	return _TaikoL2Client.Contract.Withdraw(&_TaikoL2Client.TransactOpts, token, to)
}

// TaikoL2ClientAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TaikoL2Client contract.
type TaikoL2ClientAdminChangedIterator struct {
	Event *TaikoL2ClientAdminChanged // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientAdminChanged)
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
		it.Event = new(TaikoL2ClientAdminChanged)
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
func (it *TaikoL2ClientAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientAdminChanged represents a AdminChanged event raised by the TaikoL2Client contract.
type TaikoL2ClientAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TaikoL2ClientAdminChangedIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientAdminChangedIterator{contract: _TaikoL2Client.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientAdminChanged)
				if err := _TaikoL2Client.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseAdminChanged(log types.Log) (*TaikoL2ClientAdminChanged, error) {
	event := new(TaikoL2ClientAdminChanged)
	if err := _TaikoL2Client.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientAnchoredIterator is returned from FilterAnchored and is used to iterate over the raw logs and unpacked data for Anchored events raised by the TaikoL2Client contract.
type TaikoL2ClientAnchoredIterator struct {
	Event *TaikoL2ClientAnchored // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientAnchoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientAnchored)
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
		it.Event = new(TaikoL2ClientAnchored)
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
func (it *TaikoL2ClientAnchoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientAnchoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientAnchored represents a Anchored event raised by the TaikoL2Client contract.
type TaikoL2ClientAnchored struct {
	ParentHash [32]byte
	GasExcess  uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAnchored is a free log retrieval operation binding the contract event 0x41c3f410f5c8ac36bb46b1dccef0de0f964087c9e688795fa02ecfa2c20b3fe4.
//
// Solidity: event Anchored(bytes32 parentHash, uint64 gasExcess)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterAnchored(opts *bind.FilterOpts) (*TaikoL2ClientAnchoredIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientAnchoredIterator{contract: _TaikoL2Client.contract, event: "Anchored", logs: logs, sub: sub}, nil
}

// WatchAnchored is a free log subscription operation binding the contract event 0x41c3f410f5c8ac36bb46b1dccef0de0f964087c9e688795fa02ecfa2c20b3fe4.
//
// Solidity: event Anchored(bytes32 parentHash, uint64 gasExcess)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchAnchored(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientAnchored) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Anchored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientAnchored)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Anchored", log); err != nil {
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

// ParseAnchored is a log parse operation binding the contract event 0x41c3f410f5c8ac36bb46b1dccef0de0f964087c9e688795fa02ecfa2c20b3fe4.
//
// Solidity: event Anchored(bytes32 parentHash, uint64 gasExcess)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseAnchored(log types.Log) (*TaikoL2ClientAnchored, error) {
	event := new(TaikoL2ClientAnchored)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Anchored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TaikoL2Client contract.
type TaikoL2ClientBeaconUpgradedIterator struct {
	Event *TaikoL2ClientBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientBeaconUpgraded)
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
		it.Event = new(TaikoL2ClientBeaconUpgraded)
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
func (it *TaikoL2ClientBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientBeaconUpgraded represents a BeaconUpgraded event raised by the TaikoL2Client contract.
type TaikoL2ClientBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TaikoL2ClientBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientBeaconUpgradedIterator{contract: _TaikoL2Client.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientBeaconUpgraded)
				if err := _TaikoL2Client.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseBeaconUpgraded(log types.Log) (*TaikoL2ClientBeaconUpgraded, error) {
	event := new(TaikoL2ClientBeaconUpgraded)
	if err := _TaikoL2Client.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientCrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the TaikoL2Client contract.
type TaikoL2ClientCrossChainSyncedIterator struct {
	Event *TaikoL2ClientCrossChainSynced // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientCrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientCrossChainSynced)
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
		it.Event = new(TaikoL2ClientCrossChainSynced)
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
func (it *TaikoL2ClientCrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientCrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientCrossChainSynced represents a CrossChainSynced event raised by the TaikoL2Client contract.
type TaikoL2ClientCrossChainSynced struct {
	SyncedInBlock uint64
	BlockId       uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, syncedInBlock []uint64, blockId []uint64) (*TaikoL2ClientCrossChainSyncedIterator, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "CrossChainSynced", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientCrossChainSyncedIterator{contract: _TaikoL2Client.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientCrossChainSynced, syncedInBlock []uint64, blockId []uint64) (event.Subscription, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "CrossChainSynced", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientCrossChainSynced)
				if err := _TaikoL2Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
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

// ParseCrossChainSynced is a log parse operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseCrossChainSynced(log types.Log) (*TaikoL2ClientCrossChainSynced, error) {
	event := new(TaikoL2ClientCrossChainSynced)
	if err := _TaikoL2Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaikoL2Client contract.
type TaikoL2ClientInitializedIterator struct {
	Event *TaikoL2ClientInitialized // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientInitialized)
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
		it.Event = new(TaikoL2ClientInitialized)
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
func (it *TaikoL2ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientInitialized represents a Initialized event raised by the TaikoL2Client contract.
type TaikoL2ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaikoL2ClientInitializedIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientInitializedIterator{contract: _TaikoL2Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientInitialized)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseInitialized(log types.Log) (*TaikoL2ClientInitialized, error) {
	event := new(TaikoL2ClientInitialized)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoL2Client contract.
type TaikoL2ClientOwnershipTransferredIterator struct {
	Event *TaikoL2ClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientOwnershipTransferred)
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
		it.Event = new(TaikoL2ClientOwnershipTransferred)
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
func (it *TaikoL2ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoL2Client contract.
type TaikoL2ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoL2ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientOwnershipTransferredIterator{contract: _TaikoL2Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientOwnershipTransferred)
				if err := _TaikoL2Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoL2ClientOwnershipTransferred, error) {
	event := new(TaikoL2ClientOwnershipTransferred)
	if err := _TaikoL2Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TaikoL2Client contract.
type TaikoL2ClientPausedIterator struct {
	Event *TaikoL2ClientPaused // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientPaused)
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
		it.Event = new(TaikoL2ClientPaused)
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
func (it *TaikoL2ClientPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientPaused represents a Paused event raised by the TaikoL2Client contract.
type TaikoL2ClientPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterPaused(opts *bind.FilterOpts) (*TaikoL2ClientPausedIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientPausedIterator{contract: _TaikoL2Client.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientPaused) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientPaused)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParsePaused(log types.Log) (*TaikoL2ClientPaused, error) {
	event := new(TaikoL2ClientPaused)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientTransactionExecutedIterator is returned from FilterTransactionExecuted and is used to iterate over the raw logs and unpacked data for TransactionExecuted events raised by the TaikoL2Client contract.
type TaikoL2ClientTransactionExecutedIterator struct {
	Event *TaikoL2ClientTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientTransactionExecuted)
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
		it.Event = new(TaikoL2ClientTransactionExecuted)
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
func (it *TaikoL2ClientTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientTransactionExecuted represents a TransactionExecuted event raised by the TaikoL2Client contract.
type TaikoL2ClientTransactionExecuted struct {
	TxId     uint64
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransactionExecuted is a free log retrieval operation binding the contract event 0x3c5c4a24a5f3333977c7d675661b0611a16f3c611b9ea63c0be82f4ffa9174c5.
//
// Solidity: event TransactionExecuted(uint64 indexed txId, bytes4 indexed selector)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterTransactionExecuted(opts *bind.FilterOpts, txId []uint64, selector [][4]byte) (*TaikoL2ClientTransactionExecutedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "TransactionExecuted", txIdRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientTransactionExecutedIterator{contract: _TaikoL2Client.contract, event: "TransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchTransactionExecuted is a free log subscription operation binding the contract event 0x3c5c4a24a5f3333977c7d675661b0611a16f3c611b9ea63c0be82f4ffa9174c5.
//
// Solidity: event TransactionExecuted(uint64 indexed txId, bytes4 indexed selector)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchTransactionExecuted(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientTransactionExecuted, txId []uint64, selector [][4]byte) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "TransactionExecuted", txIdRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientTransactionExecuted)
				if err := _TaikoL2Client.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
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

// ParseTransactionExecuted is a log parse operation binding the contract event 0x3c5c4a24a5f3333977c7d675661b0611a16f3c611b9ea63c0be82f4ffa9174c5.
//
// Solidity: event TransactionExecuted(uint64 indexed txId, bytes4 indexed selector)
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseTransactionExecuted(log types.Log) (*TaikoL2ClientTransactionExecuted, error) {
	event := new(TaikoL2ClientTransactionExecuted)
	if err := _TaikoL2Client.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TaikoL2Client contract.
type TaikoL2ClientUnpausedIterator struct {
	Event *TaikoL2ClientUnpaused // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientUnpaused)
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
		it.Event = new(TaikoL2ClientUnpaused)
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
func (it *TaikoL2ClientUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientUnpaused represents a Unpaused event raised by the TaikoL2Client contract.
type TaikoL2ClientUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TaikoL2ClientUnpausedIterator, error) {

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientUnpausedIterator{contract: _TaikoL2Client.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientUnpaused) (event.Subscription, error) {

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientUnpaused)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseUnpaused(log types.Log) (*TaikoL2ClientUnpaused, error) {
	event := new(TaikoL2ClientUnpaused)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL2ClientUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TaikoL2Client contract.
type TaikoL2ClientUpgradedIterator struct {
	Event *TaikoL2ClientUpgraded // Event containing the contract specifics and raw log

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
func (it *TaikoL2ClientUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL2ClientUpgraded)
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
		it.Event = new(TaikoL2ClientUpgraded)
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
func (it *TaikoL2ClientUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL2ClientUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL2ClientUpgraded represents a Upgraded event raised by the TaikoL2Client contract.
type TaikoL2ClientUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TaikoL2Client *TaikoL2ClientFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TaikoL2ClientUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TaikoL2Client.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL2ClientUpgradedIterator{contract: _TaikoL2Client.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TaikoL2Client *TaikoL2ClientFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TaikoL2ClientUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TaikoL2Client.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL2ClientUpgraded)
				if err := _TaikoL2Client.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TaikoL2Client *TaikoL2ClientFilterer) ParseUpgraded(log types.Log) (*TaikoL2ClientUpgraded, error) {
	event := new(TaikoL2ClientUpgraded)
	if err := _TaikoL2Client.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
