package pricesubmitter

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

// MetaData contains all meta data concerning the PriceSubmitter contract.
var MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposedGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGoveranance\",\"type\":\"address\"}],\"name\":\"GovernanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIFtsoGenesis[]\",\"name\":\"ftsos\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceHashesSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIFtsoGenesis[]\",\"name\":\"ftsos\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randoms\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PricesRevealed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFtsoManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFtsoRegistry\",\"outputs\":[{\"internalType\":\"contractIFtsoRegistryGenesis\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoterWhitelister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialiseFixedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"proposeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ftsoIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_randoms\",\"type\":\"uint256[]\"}],\"name\":\"revealPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFtsoRegistryGenesis\",\"name\":\"_ftsoRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterWhitelister\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ftsoManager\",\"type\":\"address\"}],\"name\":\"setContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_trustedAddresses\",\"type\":\"address[]\"}],\"name\":\"setTrustedAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ftsoIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"submitPriceHashes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"voterWhitelistBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"voterWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removedVoters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"votersRemovedFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceSubmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceSubmitterMetaData.ABI instead.
//var PriceSubmitterABI = MetaData.ABI

// PriceSubmitter is an auto generated Go binding around an Ethereum contract.
type PriceSubmitter struct {
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
	Contract     *PriceSubmitter   // Generic contract binding to set the session for
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
	Contract *PriceSubmitter // Generic contract binding to access the raw methods on
}

// CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerRaw struct {
	Contract *Caller // Generic read-only contract binding to access the raw methods on
}

// TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactorRaw struct {
	Contract *Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceSubmitter creates a new instance of PriceSubmitter, bound to a specific deployed contract.
func NewPriceSubmitter(address common.Address, backend bind.ContractBackend) (*PriceSubmitter, error) {
	contract, err := bindPriceSubmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitter{Caller: Caller{contract: contract}, Transactor: Transactor{contract: contract}, Filterer: Filterer{contract: contract}}, nil
}

// NewCaller creates a new read-only instance of PriceSubmitter, bound to a specific deployed contract.
func NewCaller(address common.Address, caller bind.ContractCaller) (*Caller, error) {
	contract, err := bindPriceSubmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Caller{contract: contract}, nil
}

// NewTransactor creates a new write-only instance of PriceSubmitter, bound to a specific deployed contract.
func NewTransactor(address common.Address, transactor bind.ContractTransactor) (*Transactor, error) {
	contract, err := bindPriceSubmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Transactor{contract: contract}, nil
}

// NewPriceSubmitterFilterer creates a new log filterer instance of PriceSubmitter, bound to a specific deployed contract.
func NewPriceSubmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*Filterer, error) {
	contract, err := bindPriceSubmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Filterer{contract: contract}, nil
}

// bindPriceSubmitter binds a generic wrapper to an already deployed contract.
func bindPriceSubmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceSubmitter *Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceSubmitter.Contract.Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceSubmitter *Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceSubmitter *Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceSubmitter *CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceSubmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceSubmitter *TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceSubmitter *TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.contract.Transact(opts, method, params...)
}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitter *Caller) GetFtsoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getFtsoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitter *Session) GetFtsoManager() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoManager(&_PriceSubmitter.CallOpts)
}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitter *CallerSession) GetFtsoManager() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoManager(&_PriceSubmitter.CallOpts)
}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitter *Caller) GetFtsoRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getFtsoRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitter *Session) GetFtsoRegistry() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoRegistry(&_PriceSubmitter.CallOpts)
}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitter *CallerSession) GetFtsoRegistry() (common.Address, error) {
	return _PriceSubmitter.Contract.GetFtsoRegistry(&_PriceSubmitter.CallOpts)
}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitter *Caller) GetTrustedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getTrustedAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitter *Session) GetTrustedAddresses() ([]common.Address, error) {
	return _PriceSubmitter.Contract.GetTrustedAddresses(&_PriceSubmitter.CallOpts)
}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitter *CallerSession) GetTrustedAddresses() ([]common.Address, error) {
	return _PriceSubmitter.Contract.GetTrustedAddresses(&_PriceSubmitter.CallOpts)
}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitter *Caller) GetVoterWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "getVoterWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitter *Session) GetVoterWhitelister() (common.Address, error) {
	return _PriceSubmitter.Contract.GetVoterWhitelister(&_PriceSubmitter.CallOpts)
}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitter *CallerSession) GetVoterWhitelister() (common.Address, error) {
	return _PriceSubmitter.Contract.GetVoterWhitelister(&_PriceSubmitter.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitter *Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitter *Session) Governance() (common.Address, error) {
	return _PriceSubmitter.Contract.Governance(&_PriceSubmitter.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitter *CallerSession) Governance() (common.Address, error) {
	return _PriceSubmitter.Contract.Governance(&_PriceSubmitter.CallOpts)
}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitter *Caller) Initialise(opts *bind.CallOpts, _governance common.Address) error {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "initialise", _governance)

	if err != nil {
		return err
	}

	return err

}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitter *Session) Initialise(_governance common.Address) error {
	return _PriceSubmitter.Contract.Initialise(&_PriceSubmitter.CallOpts, _governance)
}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitter *CallerSession) Initialise(_governance common.Address) error {
	return _PriceSubmitter.Contract.Initialise(&_PriceSubmitter.CallOpts, _governance)
}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_PriceSubmitter *Caller) ProposedGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "proposedGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_PriceSubmitter *Session) ProposedGovernance() (common.Address, error) {
	return _PriceSubmitter.Contract.ProposedGovernance(&_PriceSubmitter.CallOpts)
}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_PriceSubmitter *CallerSession) ProposedGovernance() (common.Address, error) {
	return _PriceSubmitter.Contract.ProposedGovernance(&_PriceSubmitter.CallOpts)
}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitter *Caller) VoterWhitelistBitmap(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitter.contract.Call(opts, &out, "voterWhitelistBitmap", _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitter *Session) VoterWhitelistBitmap(_voter common.Address) (*big.Int, error) {
	return _PriceSubmitter.Contract.VoterWhitelistBitmap(&_PriceSubmitter.CallOpts, _voter)
}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitter *CallerSession) VoterWhitelistBitmap(_voter common.Address) (*big.Int, error) {
	return _PriceSubmitter.Contract.VoterWhitelistBitmap(&_PriceSubmitter.CallOpts, _voter)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_PriceSubmitter *Transactor) ClaimGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "claimGovernance")
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_PriceSubmitter *Session) ClaimGovernance() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.ClaimGovernance(&_PriceSubmitter.TransactOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_PriceSubmitter *TransactorSession) ClaimGovernance() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.ClaimGovernance(&_PriceSubmitter.TransactOpts)
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitter *Transactor) InitialiseFixedAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "initialiseFixedAddress")
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitter *Session) InitialiseFixedAddress() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.InitialiseFixedAddress(&_PriceSubmitter.TransactOpts)
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitter *TransactorSession) InitialiseFixedAddress() (*types.Transaction, error) {
	return _PriceSubmitter.Contract.InitialiseFixedAddress(&_PriceSubmitter.TransactOpts)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_PriceSubmitter *Transactor) ProposeGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "proposeGovernance", _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_PriceSubmitter *Session) ProposeGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.ProposeGovernance(&_PriceSubmitter.TransactOpts, _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_PriceSubmitter *TransactorSession) ProposeGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.ProposeGovernance(&_PriceSubmitter.TransactOpts, _governance)
}

// RevealPrices is a paid mutator transaction binding the contract method 0x60848b44.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256[] _randoms) returns()
func (_PriceSubmitter *Transactor) RevealPrices(opts *bind.TransactOpts, _epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _randoms []*big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "revealPrices", _epochId, _ftsoIndices, _prices, _randoms)
}

// RevealPrices is a paid mutator transaction binding the contract method 0x60848b44.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256[] _randoms) returns()
func (_PriceSubmitter *Session) RevealPrices(_epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _randoms []*big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.RevealPrices(&_PriceSubmitter.TransactOpts, _epochId, _ftsoIndices, _prices, _randoms)
}

// RevealPrices is a paid mutator transaction binding the contract method 0x60848b44.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256[] _randoms) returns()
func (_PriceSubmitter *TransactorSession) RevealPrices(_epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _randoms []*big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.RevealPrices(&_PriceSubmitter.TransactOpts, _epochId, _ftsoIndices, _prices, _randoms)
}

// SetContractAddresses is a paid mutator transaction binding the contract method 0x8ab63380.
//
// Solidity: function setContractAddresses(address _ftsoRegistry, address _voterWhitelister, address _ftsoManager) returns()
func (_PriceSubmitter *Transactor) SetContractAddresses(opts *bind.TransactOpts, _ftsoRegistry common.Address, _voterWhitelister common.Address, _ftsoManager common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "setContractAddresses", _ftsoRegistry, _voterWhitelister, _ftsoManager)
}

// SetContractAddresses is a paid mutator transaction binding the contract method 0x8ab63380.
//
// Solidity: function setContractAddresses(address _ftsoRegistry, address _voterWhitelister, address _ftsoManager) returns()
func (_PriceSubmitter *Session) SetContractAddresses(_ftsoRegistry common.Address, _voterWhitelister common.Address, _ftsoManager common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetContractAddresses(&_PriceSubmitter.TransactOpts, _ftsoRegistry, _voterWhitelister, _ftsoManager)
}

// SetContractAddresses is a paid mutator transaction binding the contract method 0x8ab63380.
//
// Solidity: function setContractAddresses(address _ftsoRegistry, address _voterWhitelister, address _ftsoManager) returns()
func (_PriceSubmitter *TransactorSession) SetContractAddresses(_ftsoRegistry common.Address, _voterWhitelister common.Address, _ftsoManager common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetContractAddresses(&_PriceSubmitter.TransactOpts, _ftsoRegistry, _voterWhitelister, _ftsoManager)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitter *Transactor) SetTrustedAddresses(opts *bind.TransactOpts, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "setTrustedAddresses", _trustedAddresses)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitter *Session) SetTrustedAddresses(_trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetTrustedAddresses(&_PriceSubmitter.TransactOpts, _trustedAddresses)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitter *TransactorSession) SetTrustedAddresses(_trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SetTrustedAddresses(&_PriceSubmitter.TransactOpts, _trustedAddresses)
}

// SubmitPriceHashes is a paid mutator transaction binding the contract method 0xc5adc539.
//
// Solidity: function submitPriceHashes(uint256 _epochId, uint256[] _ftsoIndices, bytes32[] _hashes) returns()
func (_PriceSubmitter *Transactor) SubmitPriceHashes(opts *bind.TransactOpts, _epochId *big.Int, _ftsoIndices []*big.Int, _hashes [][32]byte) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "submitPriceHashes", _epochId, _ftsoIndices, _hashes)
}

// SubmitPriceHashes is a paid mutator transaction binding the contract method 0xc5adc539.
//
// Solidity: function submitPriceHashes(uint256 _epochId, uint256[] _ftsoIndices, bytes32[] _hashes) returns()
func (_PriceSubmitter *Session) SubmitPriceHashes(_epochId *big.Int, _ftsoIndices []*big.Int, _hashes [][32]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SubmitPriceHashes(&_PriceSubmitter.TransactOpts, _epochId, _ftsoIndices, _hashes)
}

// SubmitPriceHashes is a paid mutator transaction binding the contract method 0xc5adc539.
//
// Solidity: function submitPriceHashes(uint256 _epochId, uint256[] _ftsoIndices, bytes32[] _hashes) returns()
func (_PriceSubmitter *TransactorSession) SubmitPriceHashes(_epochId *big.Int, _ftsoIndices []*big.Int, _hashes [][32]byte) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.SubmitPriceHashes(&_PriceSubmitter.TransactOpts, _epochId, _ftsoIndices, _hashes)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_PriceSubmitter *Transactor) TransferGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "transferGovernance", _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_PriceSubmitter *Session) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.TransferGovernance(&_PriceSubmitter.TransactOpts, _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_PriceSubmitter *TransactorSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.TransferGovernance(&_PriceSubmitter.TransactOpts, _governance)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *Transactor) VoterWhitelisted(opts *bind.TransactOpts, _voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "voterWhitelisted", _voter, _ftsoIndex)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *Session) VoterWhitelisted(_voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VoterWhitelisted(&_PriceSubmitter.TransactOpts, _voter, _ftsoIndex)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *TransactorSession) VoterWhitelisted(_voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VoterWhitelisted(&_PriceSubmitter.TransactOpts, _voter, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *Transactor) VotersRemovedFromWhitelist(opts *bind.TransactOpts, _removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.contract.Transact(opts, "votersRemovedFromWhitelist", _removedVoters, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *Session) VotersRemovedFromWhitelist(_removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VotersRemovedFromWhitelist(&_PriceSubmitter.TransactOpts, _removedVoters, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitter *TransactorSession) VotersRemovedFromWhitelist(_removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitter.Contract.VotersRemovedFromWhitelist(&_PriceSubmitter.TransactOpts, _removedVoters, _ftsoIndex)
}

// GovernanceProposedIterator is returned from FilterGovernanceProposed and is used to iterate over the raw logs and unpacked data for GovernanceProposed events raised by the PriceSubmitter contract.
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

// GovernanceProposed represents a GovernanceProposed event raised by the PriceSubmitter contract.
type GovernanceProposed struct {
	ProposedGovernance common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceProposed is a free log retrieval operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_PriceSubmitter *Filterer) FilterGovernanceProposed(opts *bind.FilterOpts) (*GovernanceProposedIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "GovernanceProposed")
	if err != nil {
		return nil, err
	}
	return &GovernanceProposedIterator{contract: _PriceSubmitter.contract, event: "GovernanceProposed", logs: logs, sub: sub}, nil
}

// WatchGovernanceProposed is a free log subscription operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_PriceSubmitter *Filterer) WatchGovernanceProposed(opts *bind.WatchOpts, sink chan<- *GovernanceProposed) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "GovernanceProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the psEvent and forward to the user
				psEvent := new(GovernanceProposed)
				if err := _PriceSubmitter.contract.UnpackLog(psEvent, "GovernanceProposed", log); err != nil {
					return err
				}
				psEvent.Raw = log

				select {
				case sink <- psEvent:
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
func (_PriceSubmitter *Filterer) ParseGovernanceProposed(log types.Log) (*GovernanceProposed, error) {
	psEvent := new(GovernanceProposed)
	if err := _PriceSubmitter.contract.UnpackLog(psEvent, "GovernanceProposed", log); err != nil {
		return nil, err
	}
	psEvent.Raw = log
	return psEvent, nil
}

// GovernanceUpdatedIterator is returned from FilterGovernanceUpdated and is used to iterate over the raw logs and unpacked data for GovernanceUpdated events raised by the PriceSubmitter contract.
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

// GovernanceUpdated represents a GovernanceUpdated event raised by the PriceSubmitter contract.
type GovernanceUpdated struct {
	OldGovernance  common.Address
	NewGoveranance common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceUpdated is a free log retrieval operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_PriceSubmitter *Filterer) FilterGovernanceUpdated(opts *bind.FilterOpts) (*GovernanceUpdatedIterator, error) {

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "GovernanceUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceUpdatedIterator{contract: _PriceSubmitter.contract, event: "GovernanceUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernanceUpdated is a free log subscription operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_PriceSubmitter *Filterer) WatchGovernanceUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "GovernanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the psEvent and forward to the user
				psEvent := new(GovernanceUpdated)
				if err := _PriceSubmitter.contract.UnpackLog(psEvent, "GovernanceUpdated", log); err != nil {
					return err
				}
				psEvent.Raw = log

				select {
				case sink <- psEvent:
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
func (_PriceSubmitter *Filterer) ParseGovernanceUpdated(log types.Log) (*GovernanceUpdated, error) {
	psEvent := new(GovernanceUpdated)
	if err := _PriceSubmitter.contract.UnpackLog(psEvent, "GovernanceUpdated", log); err != nil {
		return nil, err
	}
	psEvent.Raw = log
	return psEvent, nil
}

// PriceHashesSubmittedIterator is returned from FilterPriceHashesSubmitted and is used to iterate over the raw logs and unpacked data for PriceHashesSubmitted events raised by the PriceSubmitter contract.
type PriceHashesSubmittedIterator struct {
	Event *PriceHashesSubmitted // Event containing the contract specifics and raw log

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
func (it *PriceHashesSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceHashesSubmitted)
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
		it.Event = new(PriceHashesSubmitted)
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
func (it *PriceHashesSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceHashesSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceHashesSubmitted represents a PriceHashesSubmitted event raised by the PriceSubmitter contract.
type PriceHashesSubmitted struct {
	Submitter common.Address
	EpochId   *big.Int
	Ftsos     []common.Address
	Hashes    [][32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceHashesSubmitted is a free log retrieval operation binding the contract event 0x90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434.
//
// Solidity: event PriceHashesSubmitted(address indexed submitter, uint256 indexed epochId, address[] ftsos, bytes32[] hashes, uint256 timestamp)
func (_PriceSubmitter *Filterer) FilterPriceHashesSubmitted(opts *bind.FilterOpts, submitter []common.Address, epochId []*big.Int) (*PriceHashesSubmittedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "PriceHashesSubmitted", submitterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceHashesSubmittedIterator{contract: _PriceSubmitter.contract, event: "PriceHashesSubmitted", logs: logs, sub: sub}, nil
}

// WatchPriceHashesSubmitted is a free log subscription operation binding the contract event 0x90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434.
//
// Solidity: event PriceHashesSubmitted(address indexed submitter, uint256 indexed epochId, address[] ftsos, bytes32[] hashes, uint256 timestamp)
func (_PriceSubmitter *Filterer) WatchPriceHashesSubmitted(opts *bind.WatchOpts, sink chan<- *PriceHashesSubmitted, submitter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "PriceHashesSubmitted", submitterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the psEvent and forward to the user
				psEvent := new(PriceHashesSubmitted)
				if err := _PriceSubmitter.contract.UnpackLog(psEvent, "PriceHashesSubmitted", log); err != nil {
					return err
				}
				psEvent.Raw = log

				select {
				case sink <- psEvent:
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

// ParsePriceHashesSubmitted is a log parse operation binding the contract event 0x90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434.
//
// Solidity: event PriceHashesSubmitted(address indexed submitter, uint256 indexed epochId, address[] ftsos, bytes32[] hashes, uint256 timestamp)
func (_PriceSubmitter *Filterer) ParsePriceHashesSubmitted(log types.Log) (*PriceHashesSubmitted, error) {
	psEvent := new(PriceHashesSubmitted)
	if err := _PriceSubmitter.contract.UnpackLog(psEvent, "PriceHashesSubmitted", log); err != nil {
		return nil, err
	}
	psEvent.Raw = log
	return psEvent, nil
}

// PricesRevealedIterator is returned from FilterPricesRevealed and is used to iterate over the raw logs and unpacked data for PricesRevealed events raised by the PriceSubmitter contract.
type PricesRevealedIterator struct {
	Event *PricesRevealed // Event containing the contract specifics and raw log

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
func (it *PricesRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PricesRevealed)
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
		it.Event = new(PricesRevealed)
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
func (it *PricesRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PricesRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PricesRevealed represents a PricesRevealed event raised by the PriceSubmitter contract.
type PricesRevealed struct {
	Voter     common.Address
	EpochId   *big.Int
	Ftsos     []common.Address
	Prices    []*big.Int
	Randoms   []*big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPricesRevealed is a free log retrieval operation binding the contract event 0xa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256[] randoms, uint256 timestamp)
func (_PriceSubmitter *Filterer) FilterPricesRevealed(opts *bind.FilterOpts, voter []common.Address, epochId []*big.Int) (*PricesRevealedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.FilterLogs(opts, "PricesRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &PricesRevealedIterator{contract: _PriceSubmitter.contract, event: "PricesRevealed", logs: logs, sub: sub}, nil
}

// WatchPricesRevealed is a free log subscription operation binding the contract event 0xa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256[] randoms, uint256 timestamp)
func (_PriceSubmitter *Filterer) WatchPricesRevealed(opts *bind.WatchOpts, sink chan<- *PricesRevealed, voter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitter.contract.WatchLogs(opts, "PricesRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the psEvent and forward to the user
				psEvent := new(PricesRevealed)
				if err := _PriceSubmitter.contract.UnpackLog(psEvent, "PricesRevealed", log); err != nil {
					return err
				}
				psEvent.Raw = log

				select {
				case sink <- psEvent:
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

// ParsePricesRevealed is a log parse operation binding the contract event 0xa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256[] randoms, uint256 timestamp)
func (_PriceSubmitter *Filterer) ParsePricesRevealed(log types.Log) (*PricesRevealed, error) {
	psEvent := new(PricesRevealed)
	if err := _PriceSubmitter.contract.UnpackLog(psEvent, "PricesRevealed", log); err != nil {
		return nil, err
	}
	psEvent.Raw = log
	return psEvent, nil
}
