// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// IFtsoRegistryPriceInfo is an auto generated low-level Go binding around an user-defined struct.
type IFtsoRegistryPriceInfo struct {
	FtsoIndex *big.Int
	Price     *big.Int
	Decimals  *big.Int
	Timestamp *big.Int
}

// FtsoRegistryMetaData contains all meta data concerning the FtsoRegistry contract.
var FtsoRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIIFtso\",\"name\":\"_ftsoContract\",\"type\":\"address\"}],\"name\":\"addFtso\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoManager\",\"outputs\":[{\"internalType\":\"contractIIFtsoManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllCurrentPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIFtsoRegistry.PriceInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllFtsos\",\"outputs\":[{\"internalType\":\"contractIIFtso[]\",\"name\":\"_ftsos\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"name\":\"getCurrentPriceWithDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetPriceUsdDecimals\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getCurrentPriceWithDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetPriceUsdDecimals\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"getCurrentPricesByIndices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIFtsoRegistry.PriceInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"}],\"name\":\"getCurrentPricesBySymbols\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIFtsoRegistry.PriceInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"name\":\"getFtso\",\"outputs\":[{\"internalType\":\"contractIIFtso\",\"name\":\"_activeFtso\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getFtsoBySymbol\",\"outputs\":[{\"internalType\":\"contractIIFtso\",\"name\":\"_activeFtso\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"name\":\"getFtsoHistory\",\"outputs\":[{\"internalType\":\"contractIIFtso[5]\",\"name\":\"_ftsoAddressHistory\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getFtsoIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetIndex\",\"type\":\"uint256\"}],\"name\":\"getFtsoSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_assetIndices\",\"type\":\"uint256[]\"}],\"name\":\"getFtsos\",\"outputs\":[{\"internalType\":\"contractIFtsoGenesis[]\",\"name\":\"_ftsos\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedFtsos\",\"outputs\":[{\"internalType\":\"contractIIFtso[]\",\"name\":\"_ftsos\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedIndices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_supportedIndices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedIndicesAndFtsos\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_supportedIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIIFtso[]\",\"name\":\"_ftsos\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedIndicesAndSymbols\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_supportedIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_supportedSymbols\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedIndicesSymbolsAndFtsos\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_supportedIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_supportedSymbols\",\"type\":\"string[]\"},{\"internalType\":\"contractIIFtso[]\",\"name\":\"_ftsos\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedSymbols\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_supportedSymbols\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedSymbolsAndFtsos\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_supportedSymbols\",\"type\":\"string[]\"},{\"internalType\":\"contractIIFtso[]\",\"name\":\"_ftsos\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialiseRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIFtso\",\"name\":\"_ftso\",\"type\":\"address\"}],\"name\":\"removeFtso\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FtsoRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FtsoRegistryMetaData.ABI instead.
var FtsoRegistryABI = FtsoRegistryMetaData.ABI

// FtsoRegistry is an auto generated Go binding around an Ethereum contract.
type FtsoRegistry struct {
	FtsoRegistryCaller     // Read-only binding to the contract
	FtsoRegistryTransactor // Write-only binding to the contract
	FtsoRegistryFilterer   // Log filterer for contract events
}

// FtsoRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FtsoRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FtsoRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FtsoRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtsoRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FtsoRegistrySession struct {
	Contract     *FtsoRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FtsoRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FtsoRegistryCallerSession struct {
	Contract *FtsoRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FtsoRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FtsoRegistryTransactorSession struct {
	Contract     *FtsoRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FtsoRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FtsoRegistryRaw struct {
	Contract *FtsoRegistry // Generic contract binding to access the raw methods on
}

// FtsoRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FtsoRegistryCallerRaw struct {
	Contract *FtsoRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FtsoRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FtsoRegistryTransactorRaw struct {
	Contract *FtsoRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFtsoRegistry creates a new instance of FtsoRegistry, bound to a specific deployed contract.
func NewFtsoRegistry(address common.Address, backend bind.ContractBackend) (*FtsoRegistry, error) {
	contract, err := bindFtsoRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FtsoRegistry{FtsoRegistryCaller: FtsoRegistryCaller{contract: contract}, FtsoRegistryTransactor: FtsoRegistryTransactor{contract: contract}, FtsoRegistryFilterer: FtsoRegistryFilterer{contract: contract}}, nil
}

// NewFtsoRegistryCaller creates a new read-only instance of FtsoRegistry, bound to a specific deployed contract.
func NewFtsoRegistryCaller(address common.Address, caller bind.ContractCaller) (*FtsoRegistryCaller, error) {
	contract, err := bindFtsoRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryCaller{contract: contract}, nil
}

// NewFtsoRegistryTransactor creates a new write-only instance of FtsoRegistry, bound to a specific deployed contract.
func NewFtsoRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FtsoRegistryTransactor, error) {
	contract, err := bindFtsoRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryTransactor{contract: contract}, nil
}

// NewFtsoRegistryFilterer creates a new log filterer instance of FtsoRegistry, bound to a specific deployed contract.
func NewFtsoRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FtsoRegistryFilterer, error) {
	contract, err := bindFtsoRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryFilterer{contract: contract}, nil
}

// bindFtsoRegistry binds a generic wrapper to an already deployed contract.
func bindFtsoRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FtsoRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtsoRegistry *FtsoRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtsoRegistry.Contract.FtsoRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtsoRegistry *FtsoRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.FtsoRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtsoRegistry *FtsoRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.FtsoRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtsoRegistry *FtsoRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtsoRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtsoRegistry *FtsoRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtsoRegistry *FtsoRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.contract.Transact(opts, method, params...)
}

// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
//
// Solidity: function ftsoManager() view returns(address)
func (_FtsoRegistry *FtsoRegistryCaller) FtsoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "ftsoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
//
// Solidity: function ftsoManager() view returns(address)
func (_FtsoRegistry *FtsoRegistrySession) FtsoManager() (common.Address, error) {
	return _FtsoRegistry.Contract.FtsoManager(&_FtsoRegistry.CallOpts)
}

// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
//
// Solidity: function ftsoManager() view returns(address)
func (_FtsoRegistry *FtsoRegistryCallerSession) FtsoManager() (common.Address, error) {
	return _FtsoRegistry.Contract.FtsoManager(&_FtsoRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtsoRegistry *FtsoRegistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtsoRegistry *FtsoRegistrySession) GetAddressUpdater() (common.Address, error) {
	return _FtsoRegistry.Contract.GetAddressUpdater(&_FtsoRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FtsoRegistry.Contract.GetAddressUpdater(&_FtsoRegistry.CallOpts)
}

// GetAllCurrentPrices is a free data retrieval call binding the contract method 0x58f9296f.
//
// Solidity: function getAllCurrentPrices() view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistryCaller) GetAllCurrentPrices(opts *bind.CallOpts) ([]IFtsoRegistryPriceInfo, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getAllCurrentPrices")

	if err != nil {
		return *new([]IFtsoRegistryPriceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFtsoRegistryPriceInfo)).(*[]IFtsoRegistryPriceInfo)

	return out0, err

}

// GetAllCurrentPrices is a free data retrieval call binding the contract method 0x58f9296f.
//
// Solidity: function getAllCurrentPrices() view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistrySession) GetAllCurrentPrices() ([]IFtsoRegistryPriceInfo, error) {
	return _FtsoRegistry.Contract.GetAllCurrentPrices(&_FtsoRegistry.CallOpts)
}

// GetAllCurrentPrices is a free data retrieval call binding the contract method 0x58f9296f.
//
// Solidity: function getAllCurrentPrices() view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistryCallerSession) GetAllCurrentPrices() ([]IFtsoRegistryPriceInfo, error) {
	return _FtsoRegistry.Contract.GetAllCurrentPrices(&_FtsoRegistry.CallOpts)
}

// GetAllFtsos is a free data retrieval call binding the contract method 0x2bcdd6ab.
//
// Solidity: function getAllFtsos() view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCaller) GetAllFtsos(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getAllFtsos")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllFtsos is a free data retrieval call binding the contract method 0x2bcdd6ab.
//
// Solidity: function getAllFtsos() view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistrySession) GetAllFtsos() ([]common.Address, error) {
	return _FtsoRegistry.Contract.GetAllFtsos(&_FtsoRegistry.CallOpts)
}

// GetAllFtsos is a free data retrieval call binding the contract method 0x2bcdd6ab.
//
// Solidity: function getAllFtsos() view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetAllFtsos() ([]common.Address, error) {
	return _FtsoRegistry.Contract.GetAllFtsos(&_FtsoRegistry.CallOpts)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0x42a0f243.
//
// Solidity: function getCurrentPrice(string _symbol) view returns(uint256 _price, uint256 _timestamp)
func (_FtsoRegistry *FtsoRegistryCaller) GetCurrentPrice(opts *bind.CallOpts, _symbol string) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getCurrentPrice", _symbol)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0x42a0f243.
//
// Solidity: function getCurrentPrice(string _symbol) view returns(uint256 _price, uint256 _timestamp)
func (_FtsoRegistry *FtsoRegistrySession) GetCurrentPrice(_symbol string) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPrice(&_FtsoRegistry.CallOpts, _symbol)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0x42a0f243.
//
// Solidity: function getCurrentPrice(string _symbol) view returns(uint256 _price, uint256 _timestamp)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetCurrentPrice(_symbol string) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPrice(&_FtsoRegistry.CallOpts, _symbol)
}

// GetCurrentPrice0 is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _assetIndex) view returns(uint256 _price, uint256 _timestamp)
func (_FtsoRegistry *FtsoRegistryCaller) GetCurrentPrice0(opts *bind.CallOpts, _assetIndex *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getCurrentPrice0", _assetIndex)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPrice0 is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _assetIndex) view returns(uint256 _price, uint256 _timestamp)
func (_FtsoRegistry *FtsoRegistrySession) GetCurrentPrice0(_assetIndex *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPrice0(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetCurrentPrice0 is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _assetIndex) view returns(uint256 _price, uint256 _timestamp)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetCurrentPrice0(_assetIndex *big.Int) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPrice0(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetCurrentPriceWithDecimals is a free data retrieval call binding the contract method 0x257cbd3a.
//
// Solidity: function getCurrentPriceWithDecimals(uint256 _assetIndex) view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_FtsoRegistry *FtsoRegistryCaller) GetCurrentPriceWithDecimals(opts *bind.CallOpts, _assetIndex *big.Int) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getCurrentPriceWithDecimals", _assetIndex)

	outstruct := new(struct {
		Price                 *big.Int
		Timestamp             *big.Int
		AssetPriceUsdDecimals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AssetPriceUsdDecimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPriceWithDecimals is a free data retrieval call binding the contract method 0x257cbd3a.
//
// Solidity: function getCurrentPriceWithDecimals(uint256 _assetIndex) view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_FtsoRegistry *FtsoRegistrySession) GetCurrentPriceWithDecimals(_assetIndex *big.Int) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPriceWithDecimals(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetCurrentPriceWithDecimals is a free data retrieval call binding the contract method 0x257cbd3a.
//
// Solidity: function getCurrentPriceWithDecimals(uint256 _assetIndex) view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetCurrentPriceWithDecimals(_assetIndex *big.Int) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPriceWithDecimals(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetCurrentPriceWithDecimals0 is a free data retrieval call binding the contract method 0xa69afdc6.
//
// Solidity: function getCurrentPriceWithDecimals(string _symbol) view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_FtsoRegistry *FtsoRegistryCaller) GetCurrentPriceWithDecimals0(opts *bind.CallOpts, _symbol string) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getCurrentPriceWithDecimals0", _symbol)

	outstruct := new(struct {
		Price                 *big.Int
		Timestamp             *big.Int
		AssetPriceUsdDecimals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AssetPriceUsdDecimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCurrentPriceWithDecimals0 is a free data retrieval call binding the contract method 0xa69afdc6.
//
// Solidity: function getCurrentPriceWithDecimals(string _symbol) view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_FtsoRegistry *FtsoRegistrySession) GetCurrentPriceWithDecimals0(_symbol string) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPriceWithDecimals0(&_FtsoRegistry.CallOpts, _symbol)
}

// GetCurrentPriceWithDecimals0 is a free data retrieval call binding the contract method 0xa69afdc6.
//
// Solidity: function getCurrentPriceWithDecimals(string _symbol) view returns(uint256 _price, uint256 _timestamp, uint256 _assetPriceUsdDecimals)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetCurrentPriceWithDecimals0(_symbol string) (struct {
	Price                 *big.Int
	Timestamp             *big.Int
	AssetPriceUsdDecimals *big.Int
}, error) {
	return _FtsoRegistry.Contract.GetCurrentPriceWithDecimals0(&_FtsoRegistry.CallOpts, _symbol)
}

// GetCurrentPricesByIndices is a free data retrieval call binding the contract method 0x6ba31fa1.
//
// Solidity: function getCurrentPricesByIndices(uint256[] _indices) view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistryCaller) GetCurrentPricesByIndices(opts *bind.CallOpts, _indices []*big.Int) ([]IFtsoRegistryPriceInfo, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getCurrentPricesByIndices", _indices)

	if err != nil {
		return *new([]IFtsoRegistryPriceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFtsoRegistryPriceInfo)).(*[]IFtsoRegistryPriceInfo)

	return out0, err

}

// GetCurrentPricesByIndices is a free data retrieval call binding the contract method 0x6ba31fa1.
//
// Solidity: function getCurrentPricesByIndices(uint256[] _indices) view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistrySession) GetCurrentPricesByIndices(_indices []*big.Int) ([]IFtsoRegistryPriceInfo, error) {
	return _FtsoRegistry.Contract.GetCurrentPricesByIndices(&_FtsoRegistry.CallOpts, _indices)
}

// GetCurrentPricesByIndices is a free data retrieval call binding the contract method 0x6ba31fa1.
//
// Solidity: function getCurrentPricesByIndices(uint256[] _indices) view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistryCallerSession) GetCurrentPricesByIndices(_indices []*big.Int) ([]IFtsoRegistryPriceInfo, error) {
	return _FtsoRegistry.Contract.GetCurrentPricesByIndices(&_FtsoRegistry.CallOpts, _indices)
}

// GetCurrentPricesBySymbols is a free data retrieval call binding the contract method 0x79d5ea4b.
//
// Solidity: function getCurrentPricesBySymbols(string[] _symbols) view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistryCaller) GetCurrentPricesBySymbols(opts *bind.CallOpts, _symbols []string) ([]IFtsoRegistryPriceInfo, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getCurrentPricesBySymbols", _symbols)

	if err != nil {
		return *new([]IFtsoRegistryPriceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFtsoRegistryPriceInfo)).(*[]IFtsoRegistryPriceInfo)

	return out0, err

}

// GetCurrentPricesBySymbols is a free data retrieval call binding the contract method 0x79d5ea4b.
//
// Solidity: function getCurrentPricesBySymbols(string[] _symbols) view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistrySession) GetCurrentPricesBySymbols(_symbols []string) ([]IFtsoRegistryPriceInfo, error) {
	return _FtsoRegistry.Contract.GetCurrentPricesBySymbols(&_FtsoRegistry.CallOpts, _symbols)
}

// GetCurrentPricesBySymbols is a free data retrieval call binding the contract method 0x79d5ea4b.
//
// Solidity: function getCurrentPricesBySymbols(string[] _symbols) view returns((uint256,uint256,uint256,uint256)[])
func (_FtsoRegistry *FtsoRegistryCallerSession) GetCurrentPricesBySymbols(_symbols []string) ([]IFtsoRegistryPriceInfo, error) {
	return _FtsoRegistry.Contract.GetCurrentPricesBySymbols(&_FtsoRegistry.CallOpts, _symbols)
}

// GetFtso is a free data retrieval call binding the contract method 0xd75f6d81.
//
// Solidity: function getFtso(uint256 _assetIndex) view returns(address _activeFtso)
func (_FtsoRegistry *FtsoRegistryCaller) GetFtso(opts *bind.CallOpts, _assetIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getFtso", _assetIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtso is a free data retrieval call binding the contract method 0xd75f6d81.
//
// Solidity: function getFtso(uint256 _assetIndex) view returns(address _activeFtso)
func (_FtsoRegistry *FtsoRegistrySession) GetFtso(_assetIndex *big.Int) (common.Address, error) {
	return _FtsoRegistry.Contract.GetFtso(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetFtso is a free data retrieval call binding the contract method 0xd75f6d81.
//
// Solidity: function getFtso(uint256 _assetIndex) view returns(address _activeFtso)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetFtso(_assetIndex *big.Int) (common.Address, error) {
	return _FtsoRegistry.Contract.GetFtso(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetFtsoBySymbol is a free data retrieval call binding the contract method 0x97da6af4.
//
// Solidity: function getFtsoBySymbol(string _symbol) view returns(address _activeFtso)
func (_FtsoRegistry *FtsoRegistryCaller) GetFtsoBySymbol(opts *bind.CallOpts, _symbol string) (common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getFtsoBySymbol", _symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoBySymbol is a free data retrieval call binding the contract method 0x97da6af4.
//
// Solidity: function getFtsoBySymbol(string _symbol) view returns(address _activeFtso)
func (_FtsoRegistry *FtsoRegistrySession) GetFtsoBySymbol(_symbol string) (common.Address, error) {
	return _FtsoRegistry.Contract.GetFtsoBySymbol(&_FtsoRegistry.CallOpts, _symbol)
}

// GetFtsoBySymbol is a free data retrieval call binding the contract method 0x97da6af4.
//
// Solidity: function getFtsoBySymbol(string _symbol) view returns(address _activeFtso)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetFtsoBySymbol(_symbol string) (common.Address, error) {
	return _FtsoRegistry.Contract.GetFtsoBySymbol(&_FtsoRegistry.CallOpts, _symbol)
}

// GetFtsoHistory is a free data retrieval call binding the contract method 0xc71a1b20.
//
// Solidity: function getFtsoHistory(uint256 _assetIndex) view returns(address[5] _ftsoAddressHistory)
func (_FtsoRegistry *FtsoRegistryCaller) GetFtsoHistory(opts *bind.CallOpts, _assetIndex *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getFtsoHistory", _assetIndex)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetFtsoHistory is a free data retrieval call binding the contract method 0xc71a1b20.
//
// Solidity: function getFtsoHistory(uint256 _assetIndex) view returns(address[5] _ftsoAddressHistory)
func (_FtsoRegistry *FtsoRegistrySession) GetFtsoHistory(_assetIndex *big.Int) ([5]common.Address, error) {
	return _FtsoRegistry.Contract.GetFtsoHistory(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetFtsoHistory is a free data retrieval call binding the contract method 0xc71a1b20.
//
// Solidity: function getFtsoHistory(uint256 _assetIndex) view returns(address[5] _ftsoAddressHistory)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetFtsoHistory(_assetIndex *big.Int) ([5]common.Address, error) {
	return _FtsoRegistry.Contract.GetFtsoHistory(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetFtsoIndex is a free data retrieval call binding the contract method 0xe848da30.
//
// Solidity: function getFtsoIndex(string _symbol) view returns(uint256 _assetIndex)
func (_FtsoRegistry *FtsoRegistryCaller) GetFtsoIndex(opts *bind.CallOpts, _symbol string) (*big.Int, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getFtsoIndex", _symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFtsoIndex is a free data retrieval call binding the contract method 0xe848da30.
//
// Solidity: function getFtsoIndex(string _symbol) view returns(uint256 _assetIndex)
func (_FtsoRegistry *FtsoRegistrySession) GetFtsoIndex(_symbol string) (*big.Int, error) {
	return _FtsoRegistry.Contract.GetFtsoIndex(&_FtsoRegistry.CallOpts, _symbol)
}

// GetFtsoIndex is a free data retrieval call binding the contract method 0xe848da30.
//
// Solidity: function getFtsoIndex(string _symbol) view returns(uint256 _assetIndex)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetFtsoIndex(_symbol string) (*big.Int, error) {
	return _FtsoRegistry.Contract.GetFtsoIndex(&_FtsoRegistry.CallOpts, _symbol)
}

// GetFtsoSymbol is a free data retrieval call binding the contract method 0x136d3f64.
//
// Solidity: function getFtsoSymbol(uint256 _assetIndex) view returns(string _symbol)
func (_FtsoRegistry *FtsoRegistryCaller) GetFtsoSymbol(opts *bind.CallOpts, _assetIndex *big.Int) (string, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getFtsoSymbol", _assetIndex)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFtsoSymbol is a free data retrieval call binding the contract method 0x136d3f64.
//
// Solidity: function getFtsoSymbol(uint256 _assetIndex) view returns(string _symbol)
func (_FtsoRegistry *FtsoRegistrySession) GetFtsoSymbol(_assetIndex *big.Int) (string, error) {
	return _FtsoRegistry.Contract.GetFtsoSymbol(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetFtsoSymbol is a free data retrieval call binding the contract method 0x136d3f64.
//
// Solidity: function getFtsoSymbol(uint256 _assetIndex) view returns(string _symbol)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetFtsoSymbol(_assetIndex *big.Int) (string, error) {
	return _FtsoRegistry.Contract.GetFtsoSymbol(&_FtsoRegistry.CallOpts, _assetIndex)
}

// GetFtsos is a free data retrieval call binding the contract method 0x9cb47538.
//
// Solidity: function getFtsos(uint256[] _assetIndices) view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCaller) GetFtsos(opts *bind.CallOpts, _assetIndices []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getFtsos", _assetIndices)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFtsos is a free data retrieval call binding the contract method 0x9cb47538.
//
// Solidity: function getFtsos(uint256[] _assetIndices) view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistrySession) GetFtsos(_assetIndices []*big.Int) ([]common.Address, error) {
	return _FtsoRegistry.Contract.GetFtsos(&_FtsoRegistry.CallOpts, _assetIndices)
}

// GetFtsos is a free data retrieval call binding the contract method 0x9cb47538.
//
// Solidity: function getFtsos(uint256[] _assetIndices) view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetFtsos(_assetIndices []*big.Int) ([]common.Address, error) {
	return _FtsoRegistry.Contract.GetFtsos(&_FtsoRegistry.CallOpts, _assetIndices)
}

// GetSupportedFtsos is a free data retrieval call binding the contract method 0xa40060ba.
//
// Solidity: function getSupportedFtsos() view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedFtsos(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedFtsos")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedFtsos is a free data retrieval call binding the contract method 0xa40060ba.
//
// Solidity: function getSupportedFtsos() view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedFtsos() ([]common.Address, error) {
	return _FtsoRegistry.Contract.GetSupportedFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedFtsos is a free data retrieval call binding the contract method 0xa40060ba.
//
// Solidity: function getSupportedFtsos() view returns(address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedFtsos() ([]common.Address, error) {
	return _FtsoRegistry.Contract.GetSupportedFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndices is a free data retrieval call binding the contract method 0x798aac5b.
//
// Solidity: function getSupportedIndices() view returns(uint256[] _supportedIndices)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedIndices(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedIndices")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSupportedIndices is a free data retrieval call binding the contract method 0x798aac5b.
//
// Solidity: function getSupportedIndices() view returns(uint256[] _supportedIndices)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedIndices() ([]*big.Int, error) {
	return _FtsoRegistry.Contract.GetSupportedIndices(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndices is a free data retrieval call binding the contract method 0x798aac5b.
//
// Solidity: function getSupportedIndices() view returns(uint256[] _supportedIndices)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedIndices() ([]*big.Int, error) {
	return _FtsoRegistry.Contract.GetSupportedIndices(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndicesAndFtsos is a free data retrieval call binding the contract method 0x06a2ba29.
//
// Solidity: function getSupportedIndicesAndFtsos() view returns(uint256[] _supportedIndices, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedIndicesAndFtsos(opts *bind.CallOpts) (struct {
	SupportedIndices []*big.Int
	Ftsos            []common.Address
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedIndicesAndFtsos")

	outstruct := new(struct {
		SupportedIndices []*big.Int
		Ftsos            []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupportedIndices = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Ftsos = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSupportedIndicesAndFtsos is a free data retrieval call binding the contract method 0x06a2ba29.
//
// Solidity: function getSupportedIndicesAndFtsos() view returns(uint256[] _supportedIndices, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedIndicesAndFtsos() (struct {
	SupportedIndices []*big.Int
	Ftsos            []common.Address
}, error) {
	return _FtsoRegistry.Contract.GetSupportedIndicesAndFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndicesAndFtsos is a free data retrieval call binding the contract method 0x06a2ba29.
//
// Solidity: function getSupportedIndicesAndFtsos() view returns(uint256[] _supportedIndices, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedIndicesAndFtsos() (struct {
	SupportedIndices []*big.Int
	Ftsos            []common.Address
}, error) {
	return _FtsoRegistry.Contract.GetSupportedIndicesAndFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndicesAndSymbols is a free data retrieval call binding the contract method 0xe68f283b.
//
// Solidity: function getSupportedIndicesAndSymbols() view returns(uint256[] _supportedIndices, string[] _supportedSymbols)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedIndicesAndSymbols(opts *bind.CallOpts) (struct {
	SupportedIndices []*big.Int
	SupportedSymbols []string
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedIndicesAndSymbols")

	outstruct := new(struct {
		SupportedIndices []*big.Int
		SupportedSymbols []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupportedIndices = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SupportedSymbols = *abi.ConvertType(out[1], new([]string)).(*[]string)

	return *outstruct, err

}

// GetSupportedIndicesAndSymbols is a free data retrieval call binding the contract method 0xe68f283b.
//
// Solidity: function getSupportedIndicesAndSymbols() view returns(uint256[] _supportedIndices, string[] _supportedSymbols)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedIndicesAndSymbols() (struct {
	SupportedIndices []*big.Int
	SupportedSymbols []string
}, error) {
	return _FtsoRegistry.Contract.GetSupportedIndicesAndSymbols(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndicesAndSymbols is a free data retrieval call binding the contract method 0xe68f283b.
//
// Solidity: function getSupportedIndicesAndSymbols() view returns(uint256[] _supportedIndices, string[] _supportedSymbols)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedIndicesAndSymbols() (struct {
	SupportedIndices []*big.Int
	SupportedSymbols []string
}, error) {
	return _FtsoRegistry.Contract.GetSupportedIndicesAndSymbols(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndicesSymbolsAndFtsos is a free data retrieval call binding the contract method 0x7687542c.
//
// Solidity: function getSupportedIndicesSymbolsAndFtsos() view returns(uint256[] _supportedIndices, string[] _supportedSymbols, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedIndicesSymbolsAndFtsos(opts *bind.CallOpts) (struct {
	SupportedIndices []*big.Int
	SupportedSymbols []string
	Ftsos            []common.Address
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedIndicesSymbolsAndFtsos")

	outstruct := new(struct {
		SupportedIndices []*big.Int
		SupportedSymbols []string
		Ftsos            []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupportedIndices = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SupportedSymbols = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.Ftsos = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSupportedIndicesSymbolsAndFtsos is a free data retrieval call binding the contract method 0x7687542c.
//
// Solidity: function getSupportedIndicesSymbolsAndFtsos() view returns(uint256[] _supportedIndices, string[] _supportedSymbols, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedIndicesSymbolsAndFtsos() (struct {
	SupportedIndices []*big.Int
	SupportedSymbols []string
	Ftsos            []common.Address
}, error) {
	return _FtsoRegistry.Contract.GetSupportedIndicesSymbolsAndFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedIndicesSymbolsAndFtsos is a free data retrieval call binding the contract method 0x7687542c.
//
// Solidity: function getSupportedIndicesSymbolsAndFtsos() view returns(uint256[] _supportedIndices, string[] _supportedSymbols, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedIndicesSymbolsAndFtsos() (struct {
	SupportedIndices []*big.Int
	SupportedSymbols []string
	Ftsos            []common.Address
}, error) {
	return _FtsoRegistry.Contract.GetSupportedIndicesSymbolsAndFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedSymbols is a free data retrieval call binding the contract method 0xce1c0e4d.
//
// Solidity: function getSupportedSymbols() view returns(string[] _supportedSymbols)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedSymbols(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedSymbols")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetSupportedSymbols is a free data retrieval call binding the contract method 0xce1c0e4d.
//
// Solidity: function getSupportedSymbols() view returns(string[] _supportedSymbols)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedSymbols() ([]string, error) {
	return _FtsoRegistry.Contract.GetSupportedSymbols(&_FtsoRegistry.CallOpts)
}

// GetSupportedSymbols is a free data retrieval call binding the contract method 0xce1c0e4d.
//
// Solidity: function getSupportedSymbols() view returns(string[] _supportedSymbols)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedSymbols() ([]string, error) {
	return _FtsoRegistry.Contract.GetSupportedSymbols(&_FtsoRegistry.CallOpts)
}

// GetSupportedSymbolsAndFtsos is a free data retrieval call binding the contract method 0x0cf48497.
//
// Solidity: function getSupportedSymbolsAndFtsos() view returns(string[] _supportedSymbols, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCaller) GetSupportedSymbolsAndFtsos(opts *bind.CallOpts) (struct {
	SupportedSymbols []string
	Ftsos            []common.Address
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "getSupportedSymbolsAndFtsos")

	outstruct := new(struct {
		SupportedSymbols []string
		Ftsos            []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupportedSymbols = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Ftsos = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSupportedSymbolsAndFtsos is a free data retrieval call binding the contract method 0x0cf48497.
//
// Solidity: function getSupportedSymbolsAndFtsos() view returns(string[] _supportedSymbols, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistrySession) GetSupportedSymbolsAndFtsos() (struct {
	SupportedSymbols []string
	Ftsos            []common.Address
}, error) {
	return _FtsoRegistry.Contract.GetSupportedSymbolsAndFtsos(&_FtsoRegistry.CallOpts)
}

// GetSupportedSymbolsAndFtsos is a free data retrieval call binding the contract method 0x0cf48497.
//
// Solidity: function getSupportedSymbolsAndFtsos() view returns(string[] _supportedSymbols, address[] _ftsos)
func (_FtsoRegistry *FtsoRegistryCallerSession) GetSupportedSymbolsAndFtsos() (struct {
	SupportedSymbols []string
	Ftsos            []common.Address
}, error) {
	return _FtsoRegistry.Contract.GetSupportedSymbolsAndFtsos(&_FtsoRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtsoRegistry *FtsoRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtsoRegistry *FtsoRegistrySession) Governance() (common.Address, error) {
	return _FtsoRegistry.Contract.Governance(&_FtsoRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtsoRegistry *FtsoRegistryCallerSession) Governance() (common.Address, error) {
	return _FtsoRegistry.Contract.Governance(&_FtsoRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtsoRegistry *FtsoRegistryCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtsoRegistry *FtsoRegistrySession) GovernanceSettings() (common.Address, error) {
	return _FtsoRegistry.Contract.GovernanceSettings(&_FtsoRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtsoRegistry *FtsoRegistryCallerSession) GovernanceSettings() (common.Address, error) {
	return _FtsoRegistry.Contract.GovernanceSettings(&_FtsoRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtsoRegistry *FtsoRegistryCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtsoRegistry *FtsoRegistrySession) ProductionMode() (bool, error) {
	return _FtsoRegistry.Contract.ProductionMode(&_FtsoRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtsoRegistry *FtsoRegistryCallerSession) ProductionMode() (bool, error) {
	return _FtsoRegistry.Contract.ProductionMode(&_FtsoRegistry.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoRegistry *FtsoRegistryCaller) TimelockedCalls(opts *bind.CallOpts, arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FtsoRegistry.contract.Call(opts, &out, "timelockedCalls", arg0)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoRegistry *FtsoRegistrySession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtsoRegistry.Contract.TimelockedCalls(&_FtsoRegistry.CallOpts, arg0)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 ) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoRegistry *FtsoRegistryCallerSession) TimelockedCalls(arg0 [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtsoRegistry.Contract.TimelockedCalls(&_FtsoRegistry.CallOpts, arg0)
}

// AddFtso is a paid mutator transaction binding the contract method 0x2663f1b4.
//
// Solidity: function addFtso(address _ftsoContract) returns(uint256 _assetIndex)
func (_FtsoRegistry *FtsoRegistryTransactor) AddFtso(opts *bind.TransactOpts, _ftsoContract common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "addFtso", _ftsoContract)
}

// AddFtso is a paid mutator transaction binding the contract method 0x2663f1b4.
//
// Solidity: function addFtso(address _ftsoContract) returns(uint256 _assetIndex)
func (_FtsoRegistry *FtsoRegistrySession) AddFtso(_ftsoContract common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.AddFtso(&_FtsoRegistry.TransactOpts, _ftsoContract)
}

// AddFtso is a paid mutator transaction binding the contract method 0x2663f1b4.
//
// Solidity: function addFtso(address _ftsoContract) returns(uint256 _assetIndex)
func (_FtsoRegistry *FtsoRegistryTransactorSession) AddFtso(_ftsoContract common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.AddFtso(&_FtsoRegistry.TransactOpts, _ftsoContract)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtsoRegistry *FtsoRegistryTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtsoRegistry *FtsoRegistrySession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.CancelGovernanceCall(&_FtsoRegistry.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.CancelGovernanceCall(&_FtsoRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtsoRegistry *FtsoRegistryTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtsoRegistry *FtsoRegistrySession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.ExecuteGovernanceCall(&_FtsoRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.ExecuteGovernanceCall(&_FtsoRegistry.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_FtsoRegistry *FtsoRegistryTransactor) Initialise(opts *bind.TransactOpts, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "initialise", _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_FtsoRegistry *FtsoRegistrySession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.Initialise(&_FtsoRegistry.TransactOpts, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _initialGovernance) returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) Initialise(_initialGovernance common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.Initialise(&_FtsoRegistry.TransactOpts, _initialGovernance)
}

// InitialiseRegistry is a paid mutator transaction binding the contract method 0xffc880fd.
//
// Solidity: function initialiseRegistry(address _addressUpdater) returns()
func (_FtsoRegistry *FtsoRegistryTransactor) InitialiseRegistry(opts *bind.TransactOpts, _addressUpdater common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "initialiseRegistry", _addressUpdater)
}

// InitialiseRegistry is a paid mutator transaction binding the contract method 0xffc880fd.
//
// Solidity: function initialiseRegistry(address _addressUpdater) returns()
func (_FtsoRegistry *FtsoRegistrySession) InitialiseRegistry(_addressUpdater common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.InitialiseRegistry(&_FtsoRegistry.TransactOpts, _addressUpdater)
}

// InitialiseRegistry is a paid mutator transaction binding the contract method 0xffc880fd.
//
// Solidity: function initialiseRegistry(address _addressUpdater) returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) InitialiseRegistry(_addressUpdater common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.InitialiseRegistry(&_FtsoRegistry.TransactOpts, _addressUpdater)
}

// RemoveFtso is a paid mutator transaction binding the contract method 0xa670ff87.
//
// Solidity: function removeFtso(address _ftso) returns()
func (_FtsoRegistry *FtsoRegistryTransactor) RemoveFtso(opts *bind.TransactOpts, _ftso common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "removeFtso", _ftso)
}

// RemoveFtso is a paid mutator transaction binding the contract method 0xa670ff87.
//
// Solidity: function removeFtso(address _ftso) returns()
func (_FtsoRegistry *FtsoRegistrySession) RemoveFtso(_ftso common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.RemoveFtso(&_FtsoRegistry.TransactOpts, _ftso)
}

// RemoveFtso is a paid mutator transaction binding the contract method 0xa670ff87.
//
// Solidity: function removeFtso(address _ftso) returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) RemoveFtso(_ftso common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.RemoveFtso(&_FtsoRegistry.TransactOpts, _ftso)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtsoRegistry *FtsoRegistryTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtsoRegistry *FtsoRegistrySession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtsoRegistry.Contract.SwitchToProductionMode(&_FtsoRegistry.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtsoRegistry.Contract.SwitchToProductionMode(&_FtsoRegistry.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtsoRegistry *FtsoRegistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtsoRegistry *FtsoRegistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.UpdateContractAddresses(&_FtsoRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtsoRegistry *FtsoRegistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtsoRegistry.Contract.UpdateContractAddresses(&_FtsoRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FtsoRegistryGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FtsoRegistry contract.
type FtsoRegistryGovernanceCallTimelockedIterator struct {
	Event *FtsoRegistryGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FtsoRegistryGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoRegistryGovernanceCallTimelocked)
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
		it.Event = new(FtsoRegistryGovernanceCallTimelocked)
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
func (it *FtsoRegistryGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoRegistryGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoRegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FtsoRegistry contract.
type FtsoRegistryGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoRegistry *FtsoRegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FtsoRegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FtsoRegistry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryGovernanceCallTimelockedIterator{contract: _FtsoRegistry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoRegistry *FtsoRegistryFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FtsoRegistryGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FtsoRegistry.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoRegistryGovernanceCallTimelocked)
				if err := _FtsoRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtsoRegistry *FtsoRegistryFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FtsoRegistryGovernanceCallTimelocked, error) {
	event := new(FtsoRegistryGovernanceCallTimelocked)
	if err := _FtsoRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoRegistryGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FtsoRegistry contract.
type FtsoRegistryGovernanceInitialisedIterator struct {
	Event *FtsoRegistryGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FtsoRegistryGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoRegistryGovernanceInitialised)
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
		it.Event = new(FtsoRegistryGovernanceInitialised)
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
func (it *FtsoRegistryGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoRegistryGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoRegistryGovernanceInitialised represents a GovernanceInitialised event raised by the FtsoRegistry contract.
type FtsoRegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtsoRegistry *FtsoRegistryFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FtsoRegistryGovernanceInitialisedIterator, error) {

	logs, sub, err := _FtsoRegistry.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryGovernanceInitialisedIterator{contract: _FtsoRegistry.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtsoRegistry *FtsoRegistryFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FtsoRegistryGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FtsoRegistry.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoRegistryGovernanceInitialised)
				if err := _FtsoRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtsoRegistry *FtsoRegistryFilterer) ParseGovernanceInitialised(log types.Log) (*FtsoRegistryGovernanceInitialised, error) {
	event := new(FtsoRegistryGovernanceInitialised)
	if err := _FtsoRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoRegistryGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FtsoRegistry contract.
type FtsoRegistryGovernedProductionModeEnteredIterator struct {
	Event *FtsoRegistryGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FtsoRegistryGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoRegistryGovernedProductionModeEntered)
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
		it.Event = new(FtsoRegistryGovernedProductionModeEntered)
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
func (it *FtsoRegistryGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoRegistryGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoRegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FtsoRegistry contract.
type FtsoRegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtsoRegistry *FtsoRegistryFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FtsoRegistryGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FtsoRegistry.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryGovernedProductionModeEnteredIterator{contract: _FtsoRegistry.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtsoRegistry *FtsoRegistryFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FtsoRegistryGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FtsoRegistry.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoRegistryGovernedProductionModeEntered)
				if err := _FtsoRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtsoRegistry *FtsoRegistryFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FtsoRegistryGovernedProductionModeEntered, error) {
	event := new(FtsoRegistryGovernedProductionModeEntered)
	if err := _FtsoRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoRegistryTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FtsoRegistry contract.
type FtsoRegistryTimelockedGovernanceCallCanceledIterator struct {
	Event *FtsoRegistryTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FtsoRegistryTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoRegistryTimelockedGovernanceCallCanceled)
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
		it.Event = new(FtsoRegistryTimelockedGovernanceCallCanceled)
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
func (it *FtsoRegistryTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoRegistryTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoRegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FtsoRegistry contract.
type FtsoRegistryTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtsoRegistry *FtsoRegistryFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FtsoRegistryTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FtsoRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryTimelockedGovernanceCallCanceledIterator{contract: _FtsoRegistry.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtsoRegistry *FtsoRegistryFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FtsoRegistryTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FtsoRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoRegistryTimelockedGovernanceCallCanceled)
				if err := _FtsoRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtsoRegistry *FtsoRegistryFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FtsoRegistryTimelockedGovernanceCallCanceled, error) {
	event := new(FtsoRegistryTimelockedGovernanceCallCanceled)
	if err := _FtsoRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtsoRegistryTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FtsoRegistry contract.
type FtsoRegistryTimelockedGovernanceCallExecutedIterator struct {
	Event *FtsoRegistryTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FtsoRegistryTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtsoRegistryTimelockedGovernanceCallExecuted)
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
		it.Event = new(FtsoRegistryTimelockedGovernanceCallExecuted)
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
func (it *FtsoRegistryTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtsoRegistryTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtsoRegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FtsoRegistry contract.
type FtsoRegistryTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtsoRegistry *FtsoRegistryFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FtsoRegistryTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FtsoRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FtsoRegistryTimelockedGovernanceCallExecutedIterator{contract: _FtsoRegistry.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtsoRegistry *FtsoRegistryFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FtsoRegistryTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FtsoRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtsoRegistryTimelockedGovernanceCallExecuted)
				if err := _FtsoRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtsoRegistry *FtsoRegistryFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FtsoRegistryTimelockedGovernanceCallExecuted, error) {
	event := new(FtsoRegistryTimelockedGovernanceCallExecuted)
	if err := _FtsoRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
