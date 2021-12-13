package wnat

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// MetaData contains all meta data concerning the WNat contract.
var MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"CreatedTotalSupplyCache\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposedGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGoveranance\",\"type\":\"address\"}],\"name\":\"GovernanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_contractType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newContractAddress\",\"type\":\"address\"}],\"name\":\"VotePowerContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"balanceHistoryCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"batchVotePowerOfAt\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bips\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"delegateExplicit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"delegatesOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegateAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_bips\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delegationMode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"delegatesOfAt\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegateAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_bips\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delegationMode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"delegationModeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceVotePower\",\"outputs\":[{\"internalType\":\"contractIGovernanceVotePower\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"needsReplacementVPContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"proposeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readVotePowerContract\",\"outputs\":[{\"internalType\":\"contractIVPContractEvents\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"revokeDelegationAt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cleanerContract\",\"type\":\"address\"}],\"name\":\"setCleanerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"setCleanupBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cleanupBlockNumberManager\",\"type\":\"address\"}],\"name\":\"setCleanupBlockNumberManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIGovernanceVotePower\",\"name\":\"_governanceVotePower\",\"type\":\"address\"}],\"name\":\"setGovernanceVotePower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIVPContract\",\"name\":\"_vpContract\",\"type\":\"address\"}],\"name\":\"setReadVpContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIVPContract\",\"name\":\"_vpContract\",\"type\":\"address\"}],\"name\":\"setWriteVpContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyCacheCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"totalSupplyHistoryCleanup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVotePower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalVotePowerAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalVotePowerAtCached\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegateAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_delegateAddresses\",\"type\":\"address[]\"}],\"name\":\"undelegateAllExplicit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_remainingDelegation\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"undelegatedVotePowerOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"undelegatedVotePowerOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"votePowerFromTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerFromToAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"votePowerOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"votePowerOfAtCached\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"writeVotePowerContract\",\"outputs\":[{\"internalType\":\"contractIVPContractEvents\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WNatABI is the input ABI used to generate the binding from.

var ABI = MetaData.ABI

// WNat is an auto generated Go binding around an Ethereum contract.
type WNat struct {
	Caller     // Read-only binding to the contract
	Transactor // Write-only binding to the contract
	Filterer   // Log filterer for contract events
}

// Caller is an auto generated read-only Go binding around an Ethereum contract.
type Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Session struct {
	Contract     *WNat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallerSession struct {
	Contract *Caller       // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactorSession struct {
	Contract     *Transactor       // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Raw is an auto generated low-level Go binding around an Ethereum contract.
type Raw struct {
	Contract *WNat // Generic contract binding to access the raw methods on
}

// CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerRaw struct {
	Contract *Caller // Generic read-only contract binding to access the raw methods on
}

// TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactorRaw struct {
	Contract *Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWNat creates a new instance of WNat, bound to a specific deployed contract.
func NewWNat(address common.Address, backend bind.ContractBackend) (*WNat, error) {
	contract, err := bindWNat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WNat{Caller: Caller{contract: contract}, Transactor: Transactor{contract: contract}, Filterer: Filterer{contract: contract}}, nil
}

// NewWNatCaller creates a new read-only instance of WNat, bound to a specific deployed contract.
func NewWNatCaller(address common.Address, caller bind.ContractCaller) (*Caller, error) {
	contract, err := bindWNat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Caller{contract: contract}, nil
}

// NewWNatTransactor creates a new write-only instance of WNat, bound to a specific deployed contract.
func NewWNatTransactor(address common.Address, transactor bind.ContractTransactor) (*Transactor, error) {
	contract, err := bindWNat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Transactor{contract: contract}, nil
}

// NewWNatFilterer creates a new log filterer instance of WNat, bound to a specific deployed contract.
func NewWNatFilterer(address common.Address, filterer bind.ContractFilterer) (*Filterer, error) {
	contract, err := bindWNat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Filterer{contract: contract}, nil
}

// bindWNat binds a generic wrapper to an already deployed contract.
func bindWNat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WNat *Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WNat.Contract.Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WNat *Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNat.Contract.Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WNat *Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WNat.Contract.Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WNat *CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WNat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WNat *TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WNat *TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WNat.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WNat *Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WNat *Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WNat.Contract.Allowance(&_WNat.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WNat *CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WNat.Contract.Allowance(&_WNat.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WNat *Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WNat *Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _WNat.Contract.BalanceOf(&_WNat.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WNat *CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WNat.Contract.BalanceOf(&_WNat.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *Caller) BalanceOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "balanceOfAt", _owner, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *Session) BalanceOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.BalanceOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *CallerSession) BalanceOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.BalanceOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// BatchVotePowerOfAt is a free data retrieval call binding the contract method 0x49e3c7e5.
//
// Solidity: function batchVotePowerOfAt(address[] _owners, uint256 _blockNumber) view returns(uint256[])
func (_WNat *Caller) BatchVotePowerOfAt(opts *bind.CallOpts, _owners []common.Address, _blockNumber *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "batchVotePowerOfAt", _owners, _blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BatchVotePowerOfAt is a free data retrieval call binding the contract method 0x49e3c7e5.
//
// Solidity: function batchVotePowerOfAt(address[] _owners, uint256 _blockNumber) view returns(uint256[])
func (_WNat *Session) BatchVotePowerOfAt(_owners []common.Address, _blockNumber *big.Int) ([]*big.Int, error) {
	return _WNat.Contract.BatchVotePowerOfAt(&_WNat.CallOpts, _owners, _blockNumber)
}

// BatchVotePowerOfAt is a free data retrieval call binding the contract method 0x49e3c7e5.
//
// Solidity: function batchVotePowerOfAt(address[] _owners, uint256 _blockNumber) view returns(uint256[])
func (_WNat *CallerSession) BatchVotePowerOfAt(_owners []common.Address, _blockNumber *big.Int) ([]*big.Int, error) {
	return _WNat.Contract.BatchVotePowerOfAt(&_WNat.CallOpts, _owners, _blockNumber)
}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_WNat *Caller) CleanupBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "cleanupBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_WNat *Session) CleanupBlockNumber() (*big.Int, error) {
	return _WNat.Contract.CleanupBlockNumber(&_WNat.CallOpts)
}

// CleanupBlockNumber is a free data retrieval call binding the contract method 0xdeea13e7.
//
// Solidity: function cleanupBlockNumber() view returns(uint256)
func (_WNat *CallerSession) CleanupBlockNumber() (*big.Int, error) {
	return _WNat.Contract.CleanupBlockNumber(&_WNat.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WNat *Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WNat *Session) Decimals() (uint8, error) {
	return _WNat.Contract.Decimals(&_WNat.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WNat *CallerSession) Decimals() (uint8, error) {
	return _WNat.Contract.Decimals(&_WNat.CallOpts)
}

// DelegatesOf is a free data retrieval call binding the contract method 0x7de5b8ed.
//
// Solidity: function delegatesOf(address _owner) view returns(address[] _delegateAddresses, uint256[] _bips, uint256 _count, uint256 _delegationMode)
func (_WNat *Caller) DelegatesOf(opts *bind.CallOpts, _owner common.Address) (struct {
	DelegateAddresses []common.Address
	Bips              []*big.Int
	Count             *big.Int
	DelegationMode    *big.Int
}, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "delegatesOf", _owner)

	outstruct := new(struct {
		DelegateAddresses []common.Address
		Bips              []*big.Int
		Count             *big.Int
		DelegationMode    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegateAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Bips = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Count = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DelegationMode = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegatesOf is a free data retrieval call binding the contract method 0x7de5b8ed.
//
// Solidity: function delegatesOf(address _owner) view returns(address[] _delegateAddresses, uint256[] _bips, uint256 _count, uint256 _delegationMode)
func (_WNat *Session) DelegatesOf(_owner common.Address) (struct {
	DelegateAddresses []common.Address
	Bips              []*big.Int
	Count             *big.Int
	DelegationMode    *big.Int
}, error) {
	return _WNat.Contract.DelegatesOf(&_WNat.CallOpts, _owner)
}

// DelegatesOf is a free data retrieval call binding the contract method 0x7de5b8ed.
//
// Solidity: function delegatesOf(address _owner) view returns(address[] _delegateAddresses, uint256[] _bips, uint256 _count, uint256 _delegationMode)
func (_WNat *CallerSession) DelegatesOf(_owner common.Address) (struct {
	DelegateAddresses []common.Address
	Bips              []*big.Int
	Count             *big.Int
	DelegationMode    *big.Int
}, error) {
	return _WNat.Contract.DelegatesOf(&_WNat.CallOpts, _owner)
}

// DelegatesOfAt is a free data retrieval call binding the contract method 0xed475a79.
//
// Solidity: function delegatesOfAt(address _owner, uint256 _blockNumber) view returns(address[] _delegateAddresses, uint256[] _bips, uint256 _count, uint256 _delegationMode)
func (_WNat *Caller) DelegatesOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (struct {
	DelegateAddresses []common.Address
	Bips              []*big.Int
	Count             *big.Int
	DelegationMode    *big.Int
}, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "delegatesOfAt", _owner, _blockNumber)

	outstruct := new(struct {
		DelegateAddresses []common.Address
		Bips              []*big.Int
		Count             *big.Int
		DelegationMode    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegateAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Bips = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Count = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DelegationMode = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegatesOfAt is a free data retrieval call binding the contract method 0xed475a79.
//
// Solidity: function delegatesOfAt(address _owner, uint256 _blockNumber) view returns(address[] _delegateAddresses, uint256[] _bips, uint256 _count, uint256 _delegationMode)
func (_WNat *Session) DelegatesOfAt(_owner common.Address, _blockNumber *big.Int) (struct {
	DelegateAddresses []common.Address
	Bips              []*big.Int
	Count             *big.Int
	DelegationMode    *big.Int
}, error) {
	return _WNat.Contract.DelegatesOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// DelegatesOfAt is a free data retrieval call binding the contract method 0xed475a79.
//
// Solidity: function delegatesOfAt(address _owner, uint256 _blockNumber) view returns(address[] _delegateAddresses, uint256[] _bips, uint256 _count, uint256 _delegationMode)
func (_WNat *CallerSession) DelegatesOfAt(_owner common.Address, _blockNumber *big.Int) (struct {
	DelegateAddresses []common.Address
	Bips              []*big.Int
	Count             *big.Int
	DelegationMode    *big.Int
}, error) {
	return _WNat.Contract.DelegatesOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// DelegationModeOf is a free data retrieval call binding the contract method 0xf6837767.
//
// Solidity: function delegationModeOf(address _who) view returns(uint256)
func (_WNat *Caller) DelegationModeOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "delegationModeOf", _who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationModeOf is a free data retrieval call binding the contract method 0xf6837767.
//
// Solidity: function delegationModeOf(address _who) view returns(uint256)
func (_WNat *Session) DelegationModeOf(_who common.Address) (*big.Int, error) {
	return _WNat.Contract.DelegationModeOf(&_WNat.CallOpts, _who)
}

// DelegationModeOf is a free data retrieval call binding the contract method 0xf6837767.
//
// Solidity: function delegationModeOf(address _who) view returns(uint256)
func (_WNat *CallerSession) DelegationModeOf(_who common.Address) (*big.Int, error) {
	return _WNat.Contract.DelegationModeOf(&_WNat.CallOpts, _who)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_WNat *Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_WNat *Session) Governance() (common.Address, error) {
	return _WNat.Contract.Governance(&_WNat.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_WNat *CallerSession) Governance() (common.Address, error) {
	return _WNat.Contract.Governance(&_WNat.CallOpts)
}

// GovernanceVotePower is a free data retrieval call binding the contract method 0x8c2b8ae1.
//
// Solidity: function governanceVotePower() view returns(address)
func (_WNat *Caller) GovernanceVotePower(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "governanceVotePower")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceVotePower is a free data retrieval call binding the contract method 0x8c2b8ae1.
//
// Solidity: function governanceVotePower() view returns(address)
func (_WNat *Session) GovernanceVotePower() (common.Address, error) {
	return _WNat.Contract.GovernanceVotePower(&_WNat.CallOpts)
}

// GovernanceVotePower is a free data retrieval call binding the contract method 0x8c2b8ae1.
//
// Solidity: function governanceVotePower() view returns(address)
func (_WNat *CallerSession) GovernanceVotePower() (common.Address, error) {
	return _WNat.Contract.GovernanceVotePower(&_WNat.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WNat *Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WNat *Session) Name() (string, error) {
	return _WNat.Contract.Name(&_WNat.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WNat *CallerSession) Name() (string, error) {
	return _WNat.Contract.Name(&_WNat.CallOpts)
}

// NeedsReplacementVPContract is a free data retrieval call binding the contract method 0xd582cef4.
//
// Solidity: function needsReplacementVPContract() view returns(bool)
func (_WNat *Caller) NeedsReplacementVPContract(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "needsReplacementVPContract")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NeedsReplacementVPContract is a free data retrieval call binding the contract method 0xd582cef4.
//
// Solidity: function needsReplacementVPContract() view returns(bool)
func (_WNat *Session) NeedsReplacementVPContract() (bool, error) {
	return _WNat.Contract.NeedsReplacementVPContract(&_WNat.CallOpts)
}

// NeedsReplacementVPContract is a free data retrieval call binding the contract method 0xd582cef4.
//
// Solidity: function needsReplacementVPContract() view returns(bool)
func (_WNat *CallerSession) NeedsReplacementVPContract() (bool, error) {
	return _WNat.Contract.NeedsReplacementVPContract(&_WNat.CallOpts)
}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_WNat *Caller) ProposedGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "proposedGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_WNat *Session) ProposedGovernance() (common.Address, error) {
	return _WNat.Contract.ProposedGovernance(&_WNat.CallOpts)
}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_WNat *CallerSession) ProposedGovernance() (common.Address, error) {
	return _WNat.Contract.ProposedGovernance(&_WNat.CallOpts)
}

// ReadVotePowerContract is a free data retrieval call binding the contract method 0x9b3baa0e.
//
// Solidity: function readVotePowerContract() view returns(address)
func (_WNat *Caller) ReadVotePowerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "readVotePowerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadVotePowerContract is a free data retrieval call binding the contract method 0x9b3baa0e.
//
// Solidity: function readVotePowerContract() view returns(address)
func (_WNat *Session) ReadVotePowerContract() (common.Address, error) {
	return _WNat.Contract.ReadVotePowerContract(&_WNat.CallOpts)
}

// ReadVotePowerContract is a free data retrieval call binding the contract method 0x9b3baa0e.
//
// Solidity: function readVotePowerContract() view returns(address)
func (_WNat *CallerSession) ReadVotePowerContract() (common.Address, error) {
	return _WNat.Contract.ReadVotePowerContract(&_WNat.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WNat *Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WNat *Session) Symbol() (string, error) {
	return _WNat.Contract.Symbol(&_WNat.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WNat *CallerSession) Symbol() (string, error) {
	return _WNat.Contract.Symbol(&_WNat.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WNat *Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WNat *Session) TotalSupply() (*big.Int, error) {
	return _WNat.Contract.TotalSupply(&_WNat.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WNat *CallerSession) TotalSupply() (*big.Int, error) {
	return _WNat.Contract.TotalSupply(&_WNat.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_WNat *Caller) TotalSupplyAt(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "totalSupplyAt", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_WNat *Session) TotalSupplyAt(_blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.TotalSupplyAt(&_WNat.CallOpts, _blockNumber)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _blockNumber) view returns(uint256)
func (_WNat *CallerSession) TotalSupplyAt(_blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.TotalSupplyAt(&_WNat.CallOpts, _blockNumber)
}

// TotalVotePower is a free data retrieval call binding the contract method 0xf5f3d4f7.
//
// Solidity: function totalVotePower() view returns(uint256)
func (_WNat *Caller) TotalVotePower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "totalVotePower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotePower is a free data retrieval call binding the contract method 0xf5f3d4f7.
//
// Solidity: function totalVotePower() view returns(uint256)
func (_WNat *Session) TotalVotePower() (*big.Int, error) {
	return _WNat.Contract.TotalVotePower(&_WNat.CallOpts)
}

// TotalVotePower is a free data retrieval call binding the contract method 0xf5f3d4f7.
//
// Solidity: function totalVotePower() view returns(uint256)
func (_WNat *CallerSession) TotalVotePower() (*big.Int, error) {
	return _WNat.Contract.TotalVotePower(&_WNat.CallOpts)
}

// TotalVotePowerAt is a free data retrieval call binding the contract method 0x3e5aa26a.
//
// Solidity: function totalVotePowerAt(uint256 _blockNumber) view returns(uint256)
func (_WNat *Caller) TotalVotePowerAt(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "totalVotePowerAt", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVotePowerAt is a free data retrieval call binding the contract method 0x3e5aa26a.
//
// Solidity: function totalVotePowerAt(uint256 _blockNumber) view returns(uint256)
func (_WNat *Session) TotalVotePowerAt(_blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.TotalVotePowerAt(&_WNat.CallOpts, _blockNumber)
}

// TotalVotePowerAt is a free data retrieval call binding the contract method 0x3e5aa26a.
//
// Solidity: function totalVotePowerAt(uint256 _blockNumber) view returns(uint256)
func (_WNat *CallerSession) TotalVotePowerAt(_blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.TotalVotePowerAt(&_WNat.CallOpts, _blockNumber)
}

// UndelegatedVotePowerOf is a free data retrieval call binding the contract method 0xd6aa0b77.
//
// Solidity: function undelegatedVotePowerOf(address _owner) view returns(uint256)
func (_WNat *Caller) UndelegatedVotePowerOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "undelegatedVotePowerOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndelegatedVotePowerOf is a free data retrieval call binding the contract method 0xd6aa0b77.
//
// Solidity: function undelegatedVotePowerOf(address _owner) view returns(uint256)
func (_WNat *Session) UndelegatedVotePowerOf(_owner common.Address) (*big.Int, error) {
	return _WNat.Contract.UndelegatedVotePowerOf(&_WNat.CallOpts, _owner)
}

// UndelegatedVotePowerOf is a free data retrieval call binding the contract method 0xd6aa0b77.
//
// Solidity: function undelegatedVotePowerOf(address _owner) view returns(uint256)
func (_WNat *CallerSession) UndelegatedVotePowerOf(_owner common.Address) (*big.Int, error) {
	return _WNat.Contract.UndelegatedVotePowerOf(&_WNat.CallOpts, _owner)
}

// UndelegatedVotePowerOfAt is a free data retrieval call binding the contract method 0x83035a82.
//
// Solidity: function undelegatedVotePowerOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *Caller) UndelegatedVotePowerOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "undelegatedVotePowerOfAt", _owner, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndelegatedVotePowerOfAt is a free data retrieval call binding the contract method 0x83035a82.
//
// Solidity: function undelegatedVotePowerOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *Session) UndelegatedVotePowerOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.UndelegatedVotePowerOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// UndelegatedVotePowerOfAt is a free data retrieval call binding the contract method 0x83035a82.
//
// Solidity: function undelegatedVotePowerOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *CallerSession) UndelegatedVotePowerOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.UndelegatedVotePowerOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// VotePowerFromTo is a free data retrieval call binding the contract method 0xbe0ca747.
//
// Solidity: function votePowerFromTo(address _from, address _to) view returns(uint256)
func (_WNat *Caller) VotePowerFromTo(opts *bind.CallOpts, _from common.Address, _to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "votePowerFromTo", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerFromTo is a free data retrieval call binding the contract method 0xbe0ca747.
//
// Solidity: function votePowerFromTo(address _from, address _to) view returns(uint256)
func (_WNat *Session) VotePowerFromTo(_from common.Address, _to common.Address) (*big.Int, error) {
	return _WNat.Contract.VotePowerFromTo(&_WNat.CallOpts, _from, _to)
}

// VotePowerFromTo is a free data retrieval call binding the contract method 0xbe0ca747.
//
// Solidity: function votePowerFromTo(address _from, address _to) view returns(uint256)
func (_WNat *CallerSession) VotePowerFromTo(_from common.Address, _to common.Address) (*big.Int, error) {
	return _WNat.Contract.VotePowerFromTo(&_WNat.CallOpts, _from, _to)
}

// VotePowerFromToAt is a free data retrieval call binding the contract method 0xe64767aa.
//
// Solidity: function votePowerFromToAt(address _from, address _to, uint256 _blockNumber) view returns(uint256)
func (_WNat *Caller) VotePowerFromToAt(opts *bind.CallOpts, _from common.Address, _to common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "votePowerFromToAt", _from, _to, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerFromToAt is a free data retrieval call binding the contract method 0xe64767aa.
//
// Solidity: function votePowerFromToAt(address _from, address _to, uint256 _blockNumber) view returns(uint256)
func (_WNat *Session) VotePowerFromToAt(_from common.Address, _to common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.VotePowerFromToAt(&_WNat.CallOpts, _from, _to, _blockNumber)
}

// VotePowerFromToAt is a free data retrieval call binding the contract method 0xe64767aa.
//
// Solidity: function votePowerFromToAt(address _from, address _to, uint256 _blockNumber) view returns(uint256)
func (_WNat *CallerSession) VotePowerFromToAt(_from common.Address, _to common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.VotePowerFromToAt(&_WNat.CallOpts, _from, _to, _blockNumber)
}

// VotePowerOf is a free data retrieval call binding the contract method 0x142d1018.
//
// Solidity: function votePowerOf(address _owner) view returns(uint256)
func (_WNat *Caller) VotePowerOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "votePowerOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerOf is a free data retrieval call binding the contract method 0x142d1018.
//
// Solidity: function votePowerOf(address _owner) view returns(uint256)
func (_WNat *Session) VotePowerOf(_owner common.Address) (*big.Int, error) {
	return _WNat.Contract.VotePowerOf(&_WNat.CallOpts, _owner)
}

// VotePowerOf is a free data retrieval call binding the contract method 0x142d1018.
//
// Solidity: function votePowerOf(address _owner) view returns(uint256)
func (_WNat *CallerSession) VotePowerOf(_owner common.Address) (*big.Int, error) {
	return _WNat.Contract.VotePowerOf(&_WNat.CallOpts, _owner)
}

// VotePowerOfAt is a free data retrieval call binding the contract method 0x92bfe6d8.
//
// Solidity: function votePowerOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *Caller) VotePowerOfAt(opts *bind.CallOpts, _owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "votePowerOfAt", _owner, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotePowerOfAt is a free data retrieval call binding the contract method 0x92bfe6d8.
//
// Solidity: function votePowerOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *Session) VotePowerOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.VotePowerOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// VotePowerOfAt is a free data retrieval call binding the contract method 0x92bfe6d8.
//
// Solidity: function votePowerOfAt(address _owner, uint256 _blockNumber) view returns(uint256)
func (_WNat *CallerSession) VotePowerOfAt(_owner common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _WNat.Contract.VotePowerOfAt(&_WNat.CallOpts, _owner, _blockNumber)
}

// WriteVotePowerContract is a free data retrieval call binding the contract method 0x1fec092a.
//
// Solidity: function writeVotePowerContract() view returns(address)
func (_WNat *Caller) WriteVotePowerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WNat.contract.Call(opts, &out, "writeVotePowerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WriteVotePowerContract is a free data retrieval call binding the contract method 0x1fec092a.
//
// Solidity: function writeVotePowerContract() view returns(address)
func (_WNat *Session) WriteVotePowerContract() (common.Address, error) {
	return _WNat.Contract.WriteVotePowerContract(&_WNat.CallOpts)
}

// WriteVotePowerContract is a free data retrieval call binding the contract method 0x1fec092a.
//
// Solidity: function writeVotePowerContract() view returns(address)
func (_WNat *CallerSession) WriteVotePowerContract() (common.Address, error) {
	return _WNat.Contract.WriteVotePowerContract(&_WNat.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WNat *Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WNat *Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Approve(&_WNat.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WNat *TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Approve(&_WNat.TransactOpts, spender, amount)
}

// BalanceHistoryCleanup is a paid mutator transaction binding the contract method 0xf0e292c9.
//
// Solidity: function balanceHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_WNat *Transactor) BalanceHistoryCleanup(opts *bind.TransactOpts, _owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "balanceHistoryCleanup", _owner, _count)
}

// BalanceHistoryCleanup is a paid mutator transaction binding the contract method 0xf0e292c9.
//
// Solidity: function balanceHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_WNat *Session) BalanceHistoryCleanup(_owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.BalanceHistoryCleanup(&_WNat.TransactOpts, _owner, _count)
}

// BalanceHistoryCleanup is a paid mutator transaction binding the contract method 0xf0e292c9.
//
// Solidity: function balanceHistoryCleanup(address _owner, uint256 _count) returns(uint256)
func (_WNat *TransactorSession) BalanceHistoryCleanup(_owner common.Address, _count *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.BalanceHistoryCleanup(&_WNat.TransactOpts, _owner, _count)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_WNat *Transactor) ClaimGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "claimGovernance")
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_WNat *Session) ClaimGovernance() (*types.Transaction, error) {
	return _WNat.Contract.ClaimGovernance(&_WNat.TransactOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_WNat *TransactorSession) ClaimGovernance() (*types.Transaction, error) {
	return _WNat.Contract.ClaimGovernance(&_WNat.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WNat *Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WNat *Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.DecreaseAllowance(&_WNat.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WNat *TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.DecreaseAllowance(&_WNat.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _to, uint256 _bips) returns()
func (_WNat *Transactor) Delegate(opts *bind.TransactOpts, _to common.Address, _bips *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "delegate", _to, _bips)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _to, uint256 _bips) returns()
func (_WNat *Session) Delegate(_to common.Address, _bips *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Delegate(&_WNat.TransactOpts, _to, _bips)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _to, uint256 _bips) returns()
func (_WNat *TransactorSession) Delegate(_to common.Address, _bips *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Delegate(&_WNat.TransactOpts, _to, _bips)
}

// DelegateExplicit is a paid mutator transaction binding the contract method 0xd06dc3ad.
//
// Solidity: function delegateExplicit(address _to, uint256 _amount) returns()
func (_WNat *Transactor) DelegateExplicit(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "delegateExplicit", _to, _amount)
}

// DelegateExplicit is a paid mutator transaction binding the contract method 0xd06dc3ad.
//
// Solidity: function delegateExplicit(address _to, uint256 _amount) returns()
func (_WNat *Session) DelegateExplicit(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.DelegateExplicit(&_WNat.TransactOpts, _to, _amount)
}

// DelegateExplicit is a paid mutator transaction binding the contract method 0xd06dc3ad.
//
// Solidity: function delegateExplicit(address _to, uint256 _amount) returns()
func (_WNat *TransactorSession) DelegateExplicit(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.DelegateExplicit(&_WNat.TransactOpts, _to, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WNat *Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WNat *Session) Deposit() (*types.Transaction, error) {
	return _WNat.Contract.Deposit(&_WNat.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WNat *TransactorSession) Deposit() (*types.Transaction, error) {
	return _WNat.Contract.Deposit(&_WNat.TransactOpts)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address recipient) payable returns()
func (_WNat *Transactor) DepositTo(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "depositTo", recipient)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address recipient) payable returns()
func (_WNat *Session) DepositTo(recipient common.Address) (*types.Transaction, error) {
	return _WNat.Contract.DepositTo(&_WNat.TransactOpts, recipient)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address recipient) payable returns()
func (_WNat *TransactorSession) DepositTo(recipient common.Address) (*types.Transaction, error) {
	return _WNat.Contract.DepositTo(&_WNat.TransactOpts, recipient)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WNat *Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WNat *Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.IncreaseAllowance(&_WNat.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WNat *TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.IncreaseAllowance(&_WNat.TransactOpts, spender, addedValue)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) returns()
func (_WNat *Transactor) Initialise(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "initialise", _governance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) returns()
func (_WNat *Session) Initialise(_governance common.Address) (*types.Transaction, error) {
	return _WNat.Contract.Initialise(&_WNat.TransactOpts, _governance)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) returns()
func (_WNat *TransactorSession) Initialise(_governance common.Address) (*types.Transaction, error) {
	return _WNat.Contract.Initialise(&_WNat.TransactOpts, _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_WNat *Transactor) ProposeGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "proposeGovernance", _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_WNat *Session) ProposeGovernance(_governance common.Address) (*types.Transaction, error) {
	return _WNat.Contract.ProposeGovernance(&_WNat.TransactOpts, _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_WNat *TransactorSession) ProposeGovernance(_governance common.Address) (*types.Transaction, error) {
	return _WNat.Contract.ProposeGovernance(&_WNat.TransactOpts, _governance)
}

// RevokeDelegationAt is a paid mutator transaction binding the contract method 0xbbd6fbf8.
//
// Solidity: function revokeDelegationAt(address _who, uint256 _blockNumber) returns()
func (_WNat *Transactor) RevokeDelegationAt(opts *bind.TransactOpts, _who common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "revokeDelegationAt", _who, _blockNumber)
}

// RevokeDelegationAt is a paid mutator transaction binding the contract method 0xbbd6fbf8.
//
// Solidity: function revokeDelegationAt(address _who, uint256 _blockNumber) returns()
func (_WNat *Session) RevokeDelegationAt(_who common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.RevokeDelegationAt(&_WNat.TransactOpts, _who, _blockNumber)
}

// RevokeDelegationAt is a paid mutator transaction binding the contract method 0xbbd6fbf8.
//
// Solidity: function revokeDelegationAt(address _who, uint256 _blockNumber) returns()
func (_WNat *TransactorSession) RevokeDelegationAt(_who common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.RevokeDelegationAt(&_WNat.TransactOpts, _who, _blockNumber)
}

// SetCleanerContract is a paid mutator transaction binding the contract method 0xf6a494af.
//
// Solidity: function setCleanerContract(address _cleanerContract) returns()
func (_WNat *Transactor) SetCleanerContract(opts *bind.TransactOpts, _cleanerContract common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "setCleanerContract", _cleanerContract)
}

// SetCleanerContract is a paid mutator transaction binding the contract method 0xf6a494af.
//
// Solidity: function setCleanerContract(address _cleanerContract) returns()
func (_WNat *Session) SetCleanerContract(_cleanerContract common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetCleanerContract(&_WNat.TransactOpts, _cleanerContract)
}

// SetCleanerContract is a paid mutator transaction binding the contract method 0xf6a494af.
//
// Solidity: function setCleanerContract(address _cleanerContract) returns()
func (_WNat *TransactorSession) SetCleanerContract(_cleanerContract common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetCleanerContract(&_WNat.TransactOpts, _cleanerContract)
}

// SetCleanupBlockNumber is a paid mutator transaction binding the contract method 0x13de97f5.
//
// Solidity: function setCleanupBlockNumber(uint256 _blockNumber) returns()
func (_WNat *Transactor) SetCleanupBlockNumber(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "setCleanupBlockNumber", _blockNumber)
}

// SetCleanupBlockNumber is a paid mutator transaction binding the contract method 0x13de97f5.
//
// Solidity: function setCleanupBlockNumber(uint256 _blockNumber) returns()
func (_WNat *Session) SetCleanupBlockNumber(_blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.SetCleanupBlockNumber(&_WNat.TransactOpts, _blockNumber)
}

// SetCleanupBlockNumber is a paid mutator transaction binding the contract method 0x13de97f5.
//
// Solidity: function setCleanupBlockNumber(uint256 _blockNumber) returns()
func (_WNat *TransactorSession) SetCleanupBlockNumber(_blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.SetCleanupBlockNumber(&_WNat.TransactOpts, _blockNumber)
}

// SetCleanupBlockNumberManager is a paid mutator transaction binding the contract method 0x7f4fcaa9.
//
// Solidity: function setCleanupBlockNumberManager(address _cleanupBlockNumberManager) returns()
func (_WNat *Transactor) SetCleanupBlockNumberManager(opts *bind.TransactOpts, _cleanupBlockNumberManager common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "setCleanupBlockNumberManager", _cleanupBlockNumberManager)
}

// SetCleanupBlockNumberManager is a paid mutator transaction binding the contract method 0x7f4fcaa9.
//
// Solidity: function setCleanupBlockNumberManager(address _cleanupBlockNumberManager) returns()
func (_WNat *Session) SetCleanupBlockNumberManager(_cleanupBlockNumberManager common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetCleanupBlockNumberManager(&_WNat.TransactOpts, _cleanupBlockNumberManager)
}

// SetCleanupBlockNumberManager is a paid mutator transaction binding the contract method 0x7f4fcaa9.
//
// Solidity: function setCleanupBlockNumberManager(address _cleanupBlockNumberManager) returns()
func (_WNat *TransactorSession) SetCleanupBlockNumberManager(_cleanupBlockNumberManager common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetCleanupBlockNumberManager(&_WNat.TransactOpts, _cleanupBlockNumberManager)
}

// SetGovernanceVotePower is a paid mutator transaction binding the contract method 0x9ca2231a.
//
// Solidity: function setGovernanceVotePower(address _governanceVotePower) returns()
func (_WNat *Transactor) SetGovernanceVotePower(opts *bind.TransactOpts, _governanceVotePower common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "setGovernanceVotePower", _governanceVotePower)
}

// SetGovernanceVotePower is a paid mutator transaction binding the contract method 0x9ca2231a.
//
// Solidity: function setGovernanceVotePower(address _governanceVotePower) returns()
func (_WNat *Session) SetGovernanceVotePower(_governanceVotePower common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetGovernanceVotePower(&_WNat.TransactOpts, _governanceVotePower)
}

// SetGovernanceVotePower is a paid mutator transaction binding the contract method 0x9ca2231a.
//
// Solidity: function setGovernanceVotePower(address _governanceVotePower) returns()
func (_WNat *TransactorSession) SetGovernanceVotePower(_governanceVotePower common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetGovernanceVotePower(&_WNat.TransactOpts, _governanceVotePower)
}

// SetReadVpContract is a paid mutator transaction binding the contract method 0x31d12a16.
//
// Solidity: function setReadVpContract(address _vpContract) returns()
func (_WNat *Transactor) SetReadVpContract(opts *bind.TransactOpts, _vpContract common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "setReadVpContract", _vpContract)
}

// SetReadVpContract is a paid mutator transaction binding the contract method 0x31d12a16.
//
// Solidity: function setReadVpContract(address _vpContract) returns()
func (_WNat *Session) SetReadVpContract(_vpContract common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetReadVpContract(&_WNat.TransactOpts, _vpContract)
}

// SetReadVpContract is a paid mutator transaction binding the contract method 0x31d12a16.
//
// Solidity: function setReadVpContract(address _vpContract) returns()
func (_WNat *TransactorSession) SetReadVpContract(_vpContract common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetReadVpContract(&_WNat.TransactOpts, _vpContract)
}

// SetWriteVpContract is a paid mutator transaction binding the contract method 0x755d10a4.
//
// Solidity: function setWriteVpContract(address _vpContract) returns()
func (_WNat *Transactor) SetWriteVpContract(opts *bind.TransactOpts, _vpContract common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "setWriteVpContract", _vpContract)
}

// SetWriteVpContract is a paid mutator transaction binding the contract method 0x755d10a4.
//
// Solidity: function setWriteVpContract(address _vpContract) returns()
func (_WNat *Session) SetWriteVpContract(_vpContract common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetWriteVpContract(&_WNat.TransactOpts, _vpContract)
}

// SetWriteVpContract is a paid mutator transaction binding the contract method 0x755d10a4.
//
// Solidity: function setWriteVpContract(address _vpContract) returns()
func (_WNat *TransactorSession) SetWriteVpContract(_vpContract common.Address) (*types.Transaction, error) {
	return _WNat.Contract.SetWriteVpContract(&_WNat.TransactOpts, _vpContract)
}

// TotalSupplyCacheCleanup is a paid mutator transaction binding the contract method 0x43ea370b.
//
// Solidity: function totalSupplyCacheCleanup(uint256 _blockNumber) returns(uint256)
func (_WNat *Transactor) TotalSupplyCacheCleanup(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "totalSupplyCacheCleanup", _blockNumber)
}

// TotalSupplyCacheCleanup is a paid mutator transaction binding the contract method 0x43ea370b.
//
// Solidity: function totalSupplyCacheCleanup(uint256 _blockNumber) returns(uint256)
func (_WNat *Session) TotalSupplyCacheCleanup(_blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TotalSupplyCacheCleanup(&_WNat.TransactOpts, _blockNumber)
}

// TotalSupplyCacheCleanup is a paid mutator transaction binding the contract method 0x43ea370b.
//
// Solidity: function totalSupplyCacheCleanup(uint256 _blockNumber) returns(uint256)
func (_WNat *TransactorSession) TotalSupplyCacheCleanup(_blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TotalSupplyCacheCleanup(&_WNat.TransactOpts, _blockNumber)
}

// TotalSupplyHistoryCleanup is a paid mutator transaction binding the contract method 0xf62f8f3a.
//
// Solidity: function totalSupplyHistoryCleanup(uint256 _count) returns(uint256)
func (_WNat *Transactor) TotalSupplyHistoryCleanup(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "totalSupplyHistoryCleanup", _count)
}

// TotalSupplyHistoryCleanup is a paid mutator transaction binding the contract method 0xf62f8f3a.
//
// Solidity: function totalSupplyHistoryCleanup(uint256 _count) returns(uint256)
func (_WNat *Session) TotalSupplyHistoryCleanup(_count *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TotalSupplyHistoryCleanup(&_WNat.TransactOpts, _count)
}

// TotalSupplyHistoryCleanup is a paid mutator transaction binding the contract method 0xf62f8f3a.
//
// Solidity: function totalSupplyHistoryCleanup(uint256 _count) returns(uint256)
func (_WNat *TransactorSession) TotalSupplyHistoryCleanup(_count *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TotalSupplyHistoryCleanup(&_WNat.TransactOpts, _count)
}

// TotalVotePowerAtCached is a paid mutator transaction binding the contract method 0xcaeb942b.
//
// Solidity: function totalVotePowerAtCached(uint256 _blockNumber) returns(uint256)
func (_WNat *Transactor) TotalVotePowerAtCached(opts *bind.TransactOpts, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "totalVotePowerAtCached", _blockNumber)
}

// TotalVotePowerAtCached is a paid mutator transaction binding the contract method 0xcaeb942b.
//
// Solidity: function totalVotePowerAtCached(uint256 _blockNumber) returns(uint256)
func (_WNat *Session) TotalVotePowerAtCached(_blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TotalVotePowerAtCached(&_WNat.TransactOpts, _blockNumber)
}

// TotalVotePowerAtCached is a paid mutator transaction binding the contract method 0xcaeb942b.
//
// Solidity: function totalVotePowerAtCached(uint256 _blockNumber) returns(uint256)
func (_WNat *TransactorSession) TotalVotePowerAtCached(_blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TotalVotePowerAtCached(&_WNat.TransactOpts, _blockNumber)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WNat *Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WNat *Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Transfer(&_WNat.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WNat *TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Transfer(&_WNat.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WNat *Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WNat *Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TransferFrom(&_WNat.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WNat *TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.TransferFrom(&_WNat.TransactOpts, sender, recipient, amount)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_WNat *Transactor) TransferGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "transferGovernance", _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_WNat *Session) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _WNat.Contract.TransferGovernance(&_WNat.TransactOpts, _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_WNat *TransactorSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _WNat.Contract.TransferGovernance(&_WNat.TransactOpts, _governance)
}

// UndelegateAll is a paid mutator transaction binding the contract method 0xb302f393.
//
// Solidity: function undelegateAll() returns()
func (_WNat *Transactor) UndelegateAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "undelegateAll")
}

// UndelegateAll is a paid mutator transaction binding the contract method 0xb302f393.
//
// Solidity: function undelegateAll() returns()
func (_WNat *Session) UndelegateAll() (*types.Transaction, error) {
	return _WNat.Contract.UndelegateAll(&_WNat.TransactOpts)
}

// UndelegateAll is a paid mutator transaction binding the contract method 0xb302f393.
//
// Solidity: function undelegateAll() returns()
func (_WNat *TransactorSession) UndelegateAll() (*types.Transaction, error) {
	return _WNat.Contract.UndelegateAll(&_WNat.TransactOpts)
}

// UndelegateAllExplicit is a paid mutator transaction binding the contract method 0x5d6d11eb.
//
// Solidity: function undelegateAllExplicit(address[] _delegateAddresses) returns(uint256 _remainingDelegation)
func (_WNat *Transactor) UndelegateAllExplicit(opts *bind.TransactOpts, _delegateAddresses []common.Address) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "undelegateAllExplicit", _delegateAddresses)
}

// UndelegateAllExplicit is a paid mutator transaction binding the contract method 0x5d6d11eb.
//
// Solidity: function undelegateAllExplicit(address[] _delegateAddresses) returns(uint256 _remainingDelegation)
func (_WNat *Session) UndelegateAllExplicit(_delegateAddresses []common.Address) (*types.Transaction, error) {
	return _WNat.Contract.UndelegateAllExplicit(&_WNat.TransactOpts, _delegateAddresses)
}

// UndelegateAllExplicit is a paid mutator transaction binding the contract method 0x5d6d11eb.
//
// Solidity: function undelegateAllExplicit(address[] _delegateAddresses) returns(uint256 _remainingDelegation)
func (_WNat *TransactorSession) UndelegateAllExplicit(_delegateAddresses []common.Address) (*types.Transaction, error) {
	return _WNat.Contract.UndelegateAllExplicit(&_WNat.TransactOpts, _delegateAddresses)
}

// VotePowerOfAtCached is a paid mutator transaction binding the contract method 0xe587497e.
//
// Solidity: function votePowerOfAtCached(address _owner, uint256 _blockNumber) returns(uint256)
func (_WNat *Transactor) VotePowerOfAtCached(opts *bind.TransactOpts, _owner common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "votePowerOfAtCached", _owner, _blockNumber)
}

// VotePowerOfAtCached is a paid mutator transaction binding the contract method 0xe587497e.
//
// Solidity: function votePowerOfAtCached(address _owner, uint256 _blockNumber) returns(uint256)
func (_WNat *Session) VotePowerOfAtCached(_owner common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.VotePowerOfAtCached(&_WNat.TransactOpts, _owner, _blockNumber)
}

// VotePowerOfAtCached is a paid mutator transaction binding the contract method 0xe587497e.
//
// Solidity: function votePowerOfAtCached(address _owner, uint256 _blockNumber) returns(uint256)
func (_WNat *TransactorSession) VotePowerOfAtCached(_owner common.Address, _blockNumber *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.VotePowerOfAtCached(&_WNat.TransactOpts, _owner, _blockNumber)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WNat *Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WNat *Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Withdraw(&_WNat.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WNat *TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.Withdraw(&_WNat.TransactOpts, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address owner, uint256 amount) returns()
func (_WNat *Transactor) WithdrawFrom(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.contract.Transact(opts, "withdrawFrom", owner, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address owner, uint256 amount) returns()
func (_WNat *Session) WithdrawFrom(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.WithdrawFrom(&_WNat.TransactOpts, owner, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address owner, uint256 amount) returns()
func (_WNat *TransactorSession) WithdrawFrom(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WNat.Contract.WithdrawFrom(&_WNat.TransactOpts, owner, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WNat *Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WNat.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WNat *Session) Receive() (*types.Transaction, error) {
	return _WNat.Contract.Receive(&_WNat.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WNat *TransactorSession) Receive() (*types.Transaction, error) {
	return _WNat.Contract.Receive(&_WNat.TransactOpts)
}

// ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WNat contract.
type ApprovalIterator struct {
	Event *Approval // Event containing the contract specifics and raw log

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
func (it *ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Approval)
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
		it.Event = new(Approval)
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
func (it *ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Approval represents a Approval event raised by the WNat contract.
type Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WNat *Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WNat.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ApprovalIterator{contract: _WNat.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WNat *Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WNat.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(Approval)
				if err := _WNat.contract.UnpackLog(wnEvent, "Approval", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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
func (_WNat *Filterer) ParseApproval(log types.Log) (*Approval, error) {
	wnEvent := new(Approval)
	if err := _WNat.contract.UnpackLog(wnEvent, "Approval", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// CreatedTotalSupplyCacheIterator is returned from FilterCreatedTotalSupplyCache and is used to iterate over the raw logs and unpacked data for CreatedTotalSupplyCache events raised by the WNat contract.
type CreatedTotalSupplyCacheIterator struct {
	Event *CreatedTotalSupplyCache // Event containing the contract specifics and raw log

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
func (it *CreatedTotalSupplyCacheIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatedTotalSupplyCache)
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
		it.Event = new(CreatedTotalSupplyCache)
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
func (it *CreatedTotalSupplyCacheIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatedTotalSupplyCacheIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatedTotalSupplyCache represents a CreatedTotalSupplyCache event raised by the WNat contract.
type CreatedTotalSupplyCache struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreatedTotalSupplyCache is a free log retrieval operation binding the contract event 0xfec477a10b4fcdfdf1114eb32b3caf6118b2d76d20e1fcb70007274bb4b616be.
//
// Solidity: event CreatedTotalSupplyCache(uint256 _blockNumber)
func (_WNat *Filterer) FilterCreatedTotalSupplyCache(opts *bind.FilterOpts) (*CreatedTotalSupplyCacheIterator, error) {

	logs, sub, err := _WNat.contract.FilterLogs(opts, "CreatedTotalSupplyCache")
	if err != nil {
		return nil, err
	}
	return &CreatedTotalSupplyCacheIterator{contract: _WNat.contract, event: "CreatedTotalSupplyCache", logs: logs, sub: sub}, nil
}

// WatchCreatedTotalSupplyCache is a free log subscription operation binding the contract event 0xfec477a10b4fcdfdf1114eb32b3caf6118b2d76d20e1fcb70007274bb4b616be.
//
// Solidity: event CreatedTotalSupplyCache(uint256 _blockNumber)
func (_WNat *Filterer) WatchCreatedTotalSupplyCache(opts *bind.WatchOpts, sink chan<- *CreatedTotalSupplyCache) (event.Subscription, error) {

	logs, sub, err := _WNat.contract.WatchLogs(opts, "CreatedTotalSupplyCache")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(CreatedTotalSupplyCache)
				if err := _WNat.contract.UnpackLog(wnEvent, "CreatedTotalSupplyCache", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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

// ParseCreatedTotalSupplyCache is a log parse operation binding the contract event 0xfec477a10b4fcdfdf1114eb32b3caf6118b2d76d20e1fcb70007274bb4b616be.
//
// Solidity: event CreatedTotalSupplyCache(uint256 _blockNumber)
func (_WNat *Filterer) ParseCreatedTotalSupplyCache(log types.Log) (*CreatedTotalSupplyCache, error) {
	wnEvent := new(CreatedTotalSupplyCache)
	if err := _WNat.contract.UnpackLog(wnEvent, "CreatedTotalSupplyCache", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WNat contract.
type DepositIterator struct {
	Event *Deposit // Event containing the contract specifics and raw log

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
func (it *DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Deposit)
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
		it.Event = new(Deposit)
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
func (it *DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Deposit represents a Deposit event raised by the WNat contract.
type Deposit struct {
	Dst    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 amount)
func (_WNat *Filterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*DepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WNat.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &DepositIterator{contract: _WNat.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 amount)
func (_WNat *Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Deposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WNat.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(Deposit)
				if err := _WNat.contract.UnpackLog(wnEvent, "Deposit", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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
// Solidity: event Deposit(address indexed dst, uint256 amount)
func (_WNat *Filterer) ParseDeposit(log types.Log) (*Deposit, error) {
	wnEvent := new(Deposit)
	if err := _WNat.contract.UnpackLog(wnEvent, "Deposit", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// GovernanceProposedIterator is returned from FilterGovernanceProposed and is used to iterate over the raw logs and unpacked data for GovernanceProposed events raised by the WNat contract.
type GovernanceProposedIterator struct {
	Event *GovernanceProposed // Event containing the contract specifics and raw log

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
func (it *GovernanceProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposed)
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
		it.Event = new(GovernanceProposed)
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
func (it *GovernanceProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposed represents a GovernanceProposed event raised by the WNat contract.
type GovernanceProposed struct {
	ProposedGovernance common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceProposed is a free log retrieval operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_WNat *Filterer) FilterGovernanceProposed(opts *bind.FilterOpts) (*GovernanceProposedIterator, error) {

	logs, sub, err := _WNat.contract.FilterLogs(opts, "GovernanceProposed")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposedIterator{contract: _WNat.contract, event: "GovernanceProposed", logs: logs, sub: sub}, nil
}

// WatchGovernanceProposed is a free log subscription operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_WNat *Filterer) WatchGovernanceProposed(opts *bind.WatchOpts, sink chan<- *GovernanceProposed) (event.Subscription, error) {

	logs, sub, err := _WNat.contract.WatchLogs(opts, "GovernanceProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(GovernanceProposed)
				if err := _WNat.contract.UnpackLog(wnEvent, "GovernanceProposed", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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

// ParseGovernanceProposed is a log parse operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_WNat *Filterer) ParseGovernanceProposed(log types.Log) (*GovernanceProposed, error) {
	wnEvent := new(GovernanceProposed)
	if err := _WNat.contract.UnpackLog(wnEvent, "GovernanceProposed", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// GovernanceUpdatedIterator is returned from FilterGovernanceUpdated and is used to iterate over the raw logs and unpacked data for GovernanceUpdated events raised by the WNat contract.
type GovernanceUpdatedIterator struct {
	Event *GovernanceUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceUpdated)
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
		it.Event = new(GovernanceUpdated)
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
func (it *GovernanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceUpdated represents a GovernanceUpdated event raised by the WNat contract.
type GovernanceUpdated struct {
	OldGovernance  common.Address
	NewGoveranance common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceUpdated is a free log retrieval operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_WNat *Filterer) FilterGovernanceUpdated(opts *bind.FilterOpts) (*GovernanceUpdatedIterator, error) {

	logs, sub, err := _WNat.contract.FilterLogs(opts, "GovernanceUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceUpdatedIterator{contract: _WNat.contract, event: "GovernanceUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernanceUpdated is a free log subscription operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_WNat *Filterer) WatchGovernanceUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceUpdated) (event.Subscription, error) {

	logs, sub, err := _WNat.contract.WatchLogs(opts, "GovernanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(GovernanceUpdated)
				if err := _WNat.contract.UnpackLog(wnEvent, "GovernanceUpdated", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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

// ParseGovernanceUpdated is a log parse operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_WNat *Filterer) ParseGovernanceUpdated(log types.Log) (*GovernanceUpdated, error) {
	wnEvent := new(GovernanceUpdated)
	if err := _WNat.contract.UnpackLog(wnEvent, "GovernanceUpdated", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WNat contract.
type TransferIterator struct {
	Event *Transfer // Event containing the contract specifics and raw log

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
func (it *TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Transfer)
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
		it.Event = new(Transfer)
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
func (it *TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Transfer represents a Transfer event raised by the WNat contract.
type Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WNat *Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WNat.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TransferIterator{contract: _WNat.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WNat *Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WNat.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(Transfer)
				if err := _WNat.contract.UnpackLog(wnEvent, "Transfer", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WNat *Filterer) ParseTransfer(log types.Log) (*Transfer, error) {
	wnEvent := new(Transfer)
	if err := _WNat.contract.UnpackLog(wnEvent, "Transfer", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// VotePowerContractChangedIterator is returned from FilterVotePowerContractChanged and is used to iterate over the raw logs and unpacked data for VotePowerContractChanged events raised by the WNat contract.
type VotePowerContractChangedIterator struct {
	Event *VotePowerContractChanged // Event containing the contract specifics and raw log

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
func (it *VotePowerContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotePowerContractChanged)
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
		it.Event = new(VotePowerContractChanged)
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
func (it *VotePowerContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotePowerContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotePowerContractChanged represents a VotePowerContractChanged event raised by the WNat contract.
type VotePowerContractChanged struct {
	ContractType       *big.Int
	OldContractAddress common.Address
	NewContractAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVotePowerContractChanged is a free log retrieval operation binding the contract event 0xbec98a6c0f609cda6b9f23b95824ba6ac733cb57edd17d344f2f279684400739.
//
// Solidity: event VotePowerContractChanged(uint256 _contractType, address _oldContractAddress, address _newContractAddress)
func (_WNat *Filterer) FilterVotePowerContractChanged(opts *bind.FilterOpts) (*VotePowerContractChangedIterator, error) {

	logs, sub, err := _WNat.contract.FilterLogs(opts, "VotePowerContractChanged")
	if err != nil {
		return nil, err
	}
	return &VotePowerContractChangedIterator{contract: _WNat.contract, event: "VotePowerContractChanged", logs: logs, sub: sub}, nil
}

// WatchVotePowerContractChanged is a free log subscription operation binding the contract event 0xbec98a6c0f609cda6b9f23b95824ba6ac733cb57edd17d344f2f279684400739.
//
// Solidity: event VotePowerContractChanged(uint256 _contractType, address _oldContractAddress, address _newContractAddress)
func (_WNat *Filterer) WatchVotePowerContractChanged(opts *bind.WatchOpts, sink chan<- *VotePowerContractChanged) (event.Subscription, error) {

	logs, sub, err := _WNat.contract.WatchLogs(opts, "VotePowerContractChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(VotePowerContractChanged)
				if err := _WNat.contract.UnpackLog(wnEvent, "VotePowerContractChanged", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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

// ParseVotePowerContractChanged is a log parse operation binding the contract event 0xbec98a6c0f609cda6b9f23b95824ba6ac733cb57edd17d344f2f279684400739.
//
// Solidity: event VotePowerContractChanged(uint256 _contractType, address _oldContractAddress, address _newContractAddress)
func (_WNat *Filterer) ParseVotePowerContractChanged(log types.Log) (*VotePowerContractChanged, error) {
	wnEvent := new(VotePowerContractChanged)
	if err := _WNat.contract.UnpackLog(wnEvent, "VotePowerContractChanged", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}

// WithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WNat contract.
type WithdrawalIterator struct {
	Event *Withdrawal // Event containing the contract specifics and raw log

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
func (it *WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Withdrawal)
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
		it.Event = new(Withdrawal)
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
func (it *WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Withdrawal represents a Withdrawal event raised by the WNat contract.
type Withdrawal struct {
	Src    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 amount)
func (_WNat *Filterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WNat.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalIterator{contract: _WNat.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 amount)
func (_WNat *Filterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *Withdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WNat.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the wnEvent and forward to the user
				wnEvent := new(Withdrawal)
				if err := _WNat.contract.UnpackLog(wnEvent, "Withdrawal", log); err != nil {
					return err
				}
				wnEvent.Raw = log

				select {
				case sink <- wnEvent:
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 amount)
func (_WNat *Filterer) ParseWithdrawal(log types.Log) (*Withdrawal, error) {
	wnEvent := new(Withdrawal)
	if err := _WNat.contract.UnpackLog(wnEvent, "Withdrawal", log); err != nil {
		return nil, err
	}
	wnEvent.Raw = log
	return wnEvent, nil
}
