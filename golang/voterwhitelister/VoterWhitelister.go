package VoterWhitelister

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





	// VoterWhitelisterMetaData contains all meta data concerning the VoterWhitelister contract.
	var VoterWhitelisterMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"contractIIPriceSubmitter\",\"name\":\"_priceSubmitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_defaultMaxVotersForFtso\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposedGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGoveranance\",\"type\":\"address\"}],\"name\":\"GovernanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"VoterRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"VoterWhitelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"addFtso\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultMaxVotersForFtso\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoRegistry\",\"outputs\":[{\"internalType\":\"contractIFtsoRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"getFtsoWhitelistedPriceProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getFtsoWhitelistedPriceProvidersBySymbol\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxVotersForFtso\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSubmitter\",\"outputs\":[{\"internalType\":\"contractIIPriceSubmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"proposeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"removeFtso\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"removeTrustedAddressFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"requestFullVoterWhitelisting\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_supportedIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_success\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"requestWhitelistingVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFtsoRegistry\",\"name\":\"_ftsoRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ftsoManager\",\"type\":\"address\"}],\"name\":\"setContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultMaxVotersForFtso\",\"type\":\"uint256\"}],\"name\":\"setDefaultMaxVotersForFtso\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newMaxVoters\",\"type\":\"uint256\"}],\"name\":\"setMaxVotersForFtso\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
		
	}
	// VoterWhitelisterABI is the input ABI used to generate the binding from.
	// Deprecated: Use VoterWhitelisterMetaData.ABI instead.
	var VoterWhitelisterABI = VoterWhitelisterMetaData.ABI

	

	

	// VoterWhitelister is an auto generated Go binding around an Ethereum contract.
	type VoterWhitelister struct {
	  VoterWhitelisterCaller     // Read-only binding to the contract
	  VoterWhitelisterTransactor // Write-only binding to the contract
	  VoterWhitelisterFilterer   // Log filterer for contract events
	}

	// VoterWhitelisterCaller is an auto generated read-only Go binding around an Ethereum contract.
	type VoterWhitelisterCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// VoterWhitelisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type VoterWhitelisterTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// VoterWhitelisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type VoterWhitelisterFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// VoterWhitelisterSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type VoterWhitelisterSession struct {
	  Contract     *VoterWhitelister        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// VoterWhitelisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type VoterWhitelisterCallerSession struct {
	  Contract *VoterWhitelisterCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// VoterWhitelisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type VoterWhitelisterTransactorSession struct {
	  Contract     *VoterWhitelisterTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// VoterWhitelisterRaw is an auto generated low-level Go binding around an Ethereum contract.
	type VoterWhitelisterRaw struct {
	  Contract *VoterWhitelister // Generic contract binding to access the raw methods on
	}

	// VoterWhitelisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type VoterWhitelisterCallerRaw struct {
		Contract *VoterWhitelisterCaller // Generic read-only contract binding to access the raw methods on
	}

	// VoterWhitelisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type VoterWhitelisterTransactorRaw struct {
		Contract *VoterWhitelisterTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewVoterWhitelister creates a new instance of VoterWhitelister, bound to a specific deployed contract.
	func NewVoterWhitelister(address common.Address, backend bind.ContractBackend) (*VoterWhitelister, error) {
	  contract, err := bindVoterWhitelister(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &VoterWhitelister{ VoterWhitelisterCaller: VoterWhitelisterCaller{contract: contract}, VoterWhitelisterTransactor: VoterWhitelisterTransactor{contract: contract}, VoterWhitelisterFilterer: VoterWhitelisterFilterer{contract: contract} }, nil
	}

	// NewVoterWhitelisterCaller creates a new read-only instance of VoterWhitelister, bound to a specific deployed contract.
	func NewVoterWhitelisterCaller(address common.Address, caller bind.ContractCaller) (*VoterWhitelisterCaller, error) {
	  contract, err := bindVoterWhitelister(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &VoterWhitelisterCaller{contract: contract}, nil
	}

	// NewVoterWhitelisterTransactor creates a new write-only instance of VoterWhitelister, bound to a specific deployed contract.
	func NewVoterWhitelisterTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterWhitelisterTransactor, error) {
	  contract, err := bindVoterWhitelister(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &VoterWhitelisterTransactor{contract: contract}, nil
	}

	// NewVoterWhitelisterFilterer creates a new log filterer instance of VoterWhitelister, bound to a specific deployed contract.
 	func NewVoterWhitelisterFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterWhitelisterFilterer, error) {
 	  contract, err := bindVoterWhitelister(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &VoterWhitelisterFilterer{contract: contract}, nil
 	}

	// bindVoterWhitelister binds a generic wrapper to an already deployed contract.
	func bindVoterWhitelister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(VoterWhitelisterABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_VoterWhitelister *VoterWhitelisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _VoterWhitelister.Contract.VoterWhitelisterCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_VoterWhitelister *VoterWhitelisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _VoterWhitelister.Contract.VoterWhitelisterTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_VoterWhitelister *VoterWhitelisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _VoterWhitelister.Contract.VoterWhitelisterTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_VoterWhitelister *VoterWhitelisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _VoterWhitelister.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_VoterWhitelister *VoterWhitelisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _VoterWhitelister.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_VoterWhitelister *VoterWhitelisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _VoterWhitelister.Contract.contract.Transact(opts, method, params...)
	}

	
		// DefaultMaxVotersForFtso is a free data retrieval call binding the contract method 0x47ed51b1.
		//
		// Solidity: function defaultMaxVotersForFtso() view returns(uint256)
		func (_VoterWhitelister *VoterWhitelisterCaller) DefaultMaxVotersForFtso(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "defaultMaxVotersForFtso" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// DefaultMaxVotersForFtso is a free data retrieval call binding the contract method 0x47ed51b1.
		//
		// Solidity: function defaultMaxVotersForFtso() view returns(uint256)
		func (_VoterWhitelister *VoterWhitelisterSession) DefaultMaxVotersForFtso() ( *big.Int,  error) {
		  return _VoterWhitelister.Contract.DefaultMaxVotersForFtso(&_VoterWhitelister.CallOpts )
		}

		// DefaultMaxVotersForFtso is a free data retrieval call binding the contract method 0x47ed51b1.
		//
		// Solidity: function defaultMaxVotersForFtso() view returns(uint256)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) DefaultMaxVotersForFtso() ( *big.Int,  error) {
		  return _VoterWhitelister.Contract.DefaultMaxVotersForFtso(&_VoterWhitelister.CallOpts )
		}
	
		// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
		//
		// Solidity: function ftsoManager() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCaller) FtsoManager(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "ftsoManager" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
		//
		// Solidity: function ftsoManager() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterSession) FtsoManager() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.FtsoManager(&_VoterWhitelister.CallOpts )
		}

		// FtsoManager is a free data retrieval call binding the contract method 0x11a7aaaa.
		//
		// Solidity: function ftsoManager() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) FtsoManager() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.FtsoManager(&_VoterWhitelister.CallOpts )
		}
	
		// FtsoRegistry is a free data retrieval call binding the contract method 0x38b5f869.
		//
		// Solidity: function ftsoRegistry() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCaller) FtsoRegistry(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "ftsoRegistry" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// FtsoRegistry is a free data retrieval call binding the contract method 0x38b5f869.
		//
		// Solidity: function ftsoRegistry() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterSession) FtsoRegistry() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.FtsoRegistry(&_VoterWhitelister.CallOpts )
		}

		// FtsoRegistry is a free data retrieval call binding the contract method 0x38b5f869.
		//
		// Solidity: function ftsoRegistry() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) FtsoRegistry() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.FtsoRegistry(&_VoterWhitelister.CallOpts )
		}
	
		// GetFtsoWhitelistedPriceProviders is a free data retrieval call binding the contract method 0x09fcb400.
		//
		// Solidity: function getFtsoWhitelistedPriceProviders(uint256 _ftsoIndex) view returns(address[])
		func (_VoterWhitelister *VoterWhitelisterCaller) GetFtsoWhitelistedPriceProviders(opts *bind.CallOpts , _ftsoIndex *big.Int ) ([]common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "getFtsoWhitelistedPriceProviders" , _ftsoIndex)
			
			if err != nil {
				return *new([]common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
			
			return out0,  err
			
		}

		// GetFtsoWhitelistedPriceProviders is a free data retrieval call binding the contract method 0x09fcb400.
		//
		// Solidity: function getFtsoWhitelistedPriceProviders(uint256 _ftsoIndex) view returns(address[])
		func (_VoterWhitelister *VoterWhitelisterSession) GetFtsoWhitelistedPriceProviders( _ftsoIndex *big.Int ) ( []common.Address,  error) {
		  return _VoterWhitelister.Contract.GetFtsoWhitelistedPriceProviders(&_VoterWhitelister.CallOpts , _ftsoIndex)
		}

		// GetFtsoWhitelistedPriceProviders is a free data retrieval call binding the contract method 0x09fcb400.
		//
		// Solidity: function getFtsoWhitelistedPriceProviders(uint256 _ftsoIndex) view returns(address[])
		func (_VoterWhitelister *VoterWhitelisterCallerSession) GetFtsoWhitelistedPriceProviders( _ftsoIndex *big.Int ) ( []common.Address,  error) {
		  return _VoterWhitelister.Contract.GetFtsoWhitelistedPriceProviders(&_VoterWhitelister.CallOpts , _ftsoIndex)
		}
	
		// GetFtsoWhitelistedPriceProvidersBySymbol is a free data retrieval call binding the contract method 0xaa89dfd4.
		//
		// Solidity: function getFtsoWhitelistedPriceProvidersBySymbol(string _symbol) view returns(address[])
		func (_VoterWhitelister *VoterWhitelisterCaller) GetFtsoWhitelistedPriceProvidersBySymbol(opts *bind.CallOpts , _symbol string ) ([]common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "getFtsoWhitelistedPriceProvidersBySymbol" , _symbol)
			
			if err != nil {
				return *new([]common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
			
			return out0,  err
			
		}

		// GetFtsoWhitelistedPriceProvidersBySymbol is a free data retrieval call binding the contract method 0xaa89dfd4.
		//
		// Solidity: function getFtsoWhitelistedPriceProvidersBySymbol(string _symbol) view returns(address[])
		func (_VoterWhitelister *VoterWhitelisterSession) GetFtsoWhitelistedPriceProvidersBySymbol( _symbol string ) ( []common.Address,  error) {
		  return _VoterWhitelister.Contract.GetFtsoWhitelistedPriceProvidersBySymbol(&_VoterWhitelister.CallOpts , _symbol)
		}

		// GetFtsoWhitelistedPriceProvidersBySymbol is a free data retrieval call binding the contract method 0xaa89dfd4.
		//
		// Solidity: function getFtsoWhitelistedPriceProvidersBySymbol(string _symbol) view returns(address[])
		func (_VoterWhitelister *VoterWhitelisterCallerSession) GetFtsoWhitelistedPriceProvidersBySymbol( _symbol string ) ( []common.Address,  error) {
		  return _VoterWhitelister.Contract.GetFtsoWhitelistedPriceProvidersBySymbol(&_VoterWhitelister.CallOpts , _symbol)
		}
	
		// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
		//
		// Solidity: function governance() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCaller) Governance(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "governance" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
		//
		// Solidity: function governance() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterSession) Governance() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.Governance(&_VoterWhitelister.CallOpts )
		}

		// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
		//
		// Solidity: function governance() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) Governance() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.Governance(&_VoterWhitelister.CallOpts )
		}
	
		// MaxVotersForFtso is a free data retrieval call binding the contract method 0x98dccfc2.
		//
		// Solidity: function maxVotersForFtso(uint256 ) view returns(uint256)
		func (_VoterWhitelister *VoterWhitelisterCaller) MaxVotersForFtso(opts *bind.CallOpts , arg0 *big.Int ) (*big.Int, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "maxVotersForFtso" , arg0)
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// MaxVotersForFtso is a free data retrieval call binding the contract method 0x98dccfc2.
		//
		// Solidity: function maxVotersForFtso(uint256 ) view returns(uint256)
		func (_VoterWhitelister *VoterWhitelisterSession) MaxVotersForFtso( arg0 *big.Int ) ( *big.Int,  error) {
		  return _VoterWhitelister.Contract.MaxVotersForFtso(&_VoterWhitelister.CallOpts , arg0)
		}

		// MaxVotersForFtso is a free data retrieval call binding the contract method 0x98dccfc2.
		//
		// Solidity: function maxVotersForFtso(uint256 ) view returns(uint256)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) MaxVotersForFtso( arg0 *big.Int ) ( *big.Int,  error) {
		  return _VoterWhitelister.Contract.MaxVotersForFtso(&_VoterWhitelister.CallOpts , arg0)
		}
	
		// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
		//
		// Solidity: function priceSubmitter() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCaller) PriceSubmitter(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "priceSubmitter" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
		//
		// Solidity: function priceSubmitter() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterSession) PriceSubmitter() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.PriceSubmitter(&_VoterWhitelister.CallOpts )
		}

		// PriceSubmitter is a free data retrieval call binding the contract method 0xf937d6ad.
		//
		// Solidity: function priceSubmitter() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) PriceSubmitter() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.PriceSubmitter(&_VoterWhitelister.CallOpts )
		}
	
		// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
		//
		// Solidity: function proposedGovernance() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCaller) ProposedGovernance(opts *bind.CallOpts ) (common.Address, error) {
			var out []interface{}
			err := _VoterWhitelister.contract.Call(opts, &out, "proposedGovernance" )
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
		//
		// Solidity: function proposedGovernance() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterSession) ProposedGovernance() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.ProposedGovernance(&_VoterWhitelister.CallOpts )
		}

		// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
		//
		// Solidity: function proposedGovernance() view returns(address)
		func (_VoterWhitelister *VoterWhitelisterCallerSession) ProposedGovernance() ( common.Address,  error) {
		  return _VoterWhitelister.Contract.ProposedGovernance(&_VoterWhitelister.CallOpts )
		}
	

	
		// AddFtso is a paid mutator transaction binding the contract method 0x345705a4.
		//
		// Solidity: function addFtso(uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) AddFtso(opts *bind.TransactOpts , _ftsoIndex *big.Int ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "addFtso" , _ftsoIndex)
		}

		// AddFtso is a paid mutator transaction binding the contract method 0x345705a4.
		//
		// Solidity: function addFtso(uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) AddFtso( _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.AddFtso(&_VoterWhitelister.TransactOpts , _ftsoIndex)
		}

		// AddFtso is a paid mutator transaction binding the contract method 0x345705a4.
		//
		// Solidity: function addFtso(uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) AddFtso( _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.AddFtso(&_VoterWhitelister.TransactOpts , _ftsoIndex)
		}
	
		// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
		//
		// Solidity: function claimGovernance() returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) ClaimGovernance(opts *bind.TransactOpts ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "claimGovernance" )
		}

		// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
		//
		// Solidity: function claimGovernance() returns()
		func (_VoterWhitelister *VoterWhitelisterSession) ClaimGovernance() (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.ClaimGovernance(&_VoterWhitelister.TransactOpts )
		}

		// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
		//
		// Solidity: function claimGovernance() returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) ClaimGovernance() (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.ClaimGovernance(&_VoterWhitelister.TransactOpts )
		}
	
		// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
		//
		// Solidity: function initialise(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) Initialise(opts *bind.TransactOpts , _governance common.Address ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "initialise" , _governance)
		}

		// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
		//
		// Solidity: function initialise(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) Initialise( _governance common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.Initialise(&_VoterWhitelister.TransactOpts , _governance)
		}

		// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
		//
		// Solidity: function initialise(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) Initialise( _governance common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.Initialise(&_VoterWhitelister.TransactOpts , _governance)
		}
	
		// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
		//
		// Solidity: function proposeGovernance(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) ProposeGovernance(opts *bind.TransactOpts , _governance common.Address ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "proposeGovernance" , _governance)
		}

		// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
		//
		// Solidity: function proposeGovernance(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) ProposeGovernance( _governance common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.ProposeGovernance(&_VoterWhitelister.TransactOpts , _governance)
		}

		// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
		//
		// Solidity: function proposeGovernance(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) ProposeGovernance( _governance common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.ProposeGovernance(&_VoterWhitelister.TransactOpts , _governance)
		}
	
		// RemoveFtso is a paid mutator transaction binding the contract method 0xd8736171.
		//
		// Solidity: function removeFtso(uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) RemoveFtso(opts *bind.TransactOpts , _ftsoIndex *big.Int ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "removeFtso" , _ftsoIndex)
		}

		// RemoveFtso is a paid mutator transaction binding the contract method 0xd8736171.
		//
		// Solidity: function removeFtso(uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) RemoveFtso( _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RemoveFtso(&_VoterWhitelister.TransactOpts , _ftsoIndex)
		}

		// RemoveFtso is a paid mutator transaction binding the contract method 0xd8736171.
		//
		// Solidity: function removeFtso(uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) RemoveFtso( _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RemoveFtso(&_VoterWhitelister.TransactOpts , _ftsoIndex)
		}
	
		// RemoveTrustedAddressFromWhitelist is a paid mutator transaction binding the contract method 0x9dc950ab.
		//
		// Solidity: function removeTrustedAddressFromWhitelist(address _trustedAddress, uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) RemoveTrustedAddressFromWhitelist(opts *bind.TransactOpts , _trustedAddress common.Address , _ftsoIndex *big.Int ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "removeTrustedAddressFromWhitelist" , _trustedAddress, _ftsoIndex)
		}

		// RemoveTrustedAddressFromWhitelist is a paid mutator transaction binding the contract method 0x9dc950ab.
		//
		// Solidity: function removeTrustedAddressFromWhitelist(address _trustedAddress, uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) RemoveTrustedAddressFromWhitelist( _trustedAddress common.Address , _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RemoveTrustedAddressFromWhitelist(&_VoterWhitelister.TransactOpts , _trustedAddress, _ftsoIndex)
		}

		// RemoveTrustedAddressFromWhitelist is a paid mutator transaction binding the contract method 0x9dc950ab.
		//
		// Solidity: function removeTrustedAddressFromWhitelist(address _trustedAddress, uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) RemoveTrustedAddressFromWhitelist( _trustedAddress common.Address , _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RemoveTrustedAddressFromWhitelist(&_VoterWhitelister.TransactOpts , _trustedAddress, _ftsoIndex)
		}
	
		// RequestFullVoterWhitelisting is a paid mutator transaction binding the contract method 0xb06cbaf7.
		//
		// Solidity: function requestFullVoterWhitelisting(address _voter) returns(uint256[] _supportedIndices, bool[] _success)
		func (_VoterWhitelister *VoterWhitelisterTransactor) RequestFullVoterWhitelisting(opts *bind.TransactOpts , _voter common.Address ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "requestFullVoterWhitelisting" , _voter)
		}

		// RequestFullVoterWhitelisting is a paid mutator transaction binding the contract method 0xb06cbaf7.
		//
		// Solidity: function requestFullVoterWhitelisting(address _voter) returns(uint256[] _supportedIndices, bool[] _success)
		func (_VoterWhitelister *VoterWhitelisterSession) RequestFullVoterWhitelisting( _voter common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RequestFullVoterWhitelisting(&_VoterWhitelister.TransactOpts , _voter)
		}

		// RequestFullVoterWhitelisting is a paid mutator transaction binding the contract method 0xb06cbaf7.
		//
		// Solidity: function requestFullVoterWhitelisting(address _voter) returns(uint256[] _supportedIndices, bool[] _success)
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) RequestFullVoterWhitelisting( _voter common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RequestFullVoterWhitelisting(&_VoterWhitelister.TransactOpts , _voter)
		}
	
		// RequestWhitelistingVoter is a paid mutator transaction binding the contract method 0x3de2cb1c.
		//
		// Solidity: function requestWhitelistingVoter(address _voter, uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) RequestWhitelistingVoter(opts *bind.TransactOpts , _voter common.Address , _ftsoIndex *big.Int ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "requestWhitelistingVoter" , _voter, _ftsoIndex)
		}

		// RequestWhitelistingVoter is a paid mutator transaction binding the contract method 0x3de2cb1c.
		//
		// Solidity: function requestWhitelistingVoter(address _voter, uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) RequestWhitelistingVoter( _voter common.Address , _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RequestWhitelistingVoter(&_VoterWhitelister.TransactOpts , _voter, _ftsoIndex)
		}

		// RequestWhitelistingVoter is a paid mutator transaction binding the contract method 0x3de2cb1c.
		//
		// Solidity: function requestWhitelistingVoter(address _voter, uint256 _ftsoIndex) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) RequestWhitelistingVoter( _voter common.Address , _ftsoIndex *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.RequestWhitelistingVoter(&_VoterWhitelister.TransactOpts , _voter, _ftsoIndex)
		}
	
		// SetContractAddresses is a paid mutator transaction binding the contract method 0xcd931e40.
		//
		// Solidity: function setContractAddresses(address _ftsoRegistry, address _ftsoManager) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) SetContractAddresses(opts *bind.TransactOpts , _ftsoRegistry common.Address , _ftsoManager common.Address ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "setContractAddresses" , _ftsoRegistry, _ftsoManager)
		}

		// SetContractAddresses is a paid mutator transaction binding the contract method 0xcd931e40.
		//
		// Solidity: function setContractAddresses(address _ftsoRegistry, address _ftsoManager) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) SetContractAddresses( _ftsoRegistry common.Address , _ftsoManager common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.SetContractAddresses(&_VoterWhitelister.TransactOpts , _ftsoRegistry, _ftsoManager)
		}

		// SetContractAddresses is a paid mutator transaction binding the contract method 0xcd931e40.
		//
		// Solidity: function setContractAddresses(address _ftsoRegistry, address _ftsoManager) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) SetContractAddresses( _ftsoRegistry common.Address , _ftsoManager common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.SetContractAddresses(&_VoterWhitelister.TransactOpts , _ftsoRegistry, _ftsoManager)
		}
	
		// SetDefaultMaxVotersForFtso is a paid mutator transaction binding the contract method 0x2ee96140.
		//
		// Solidity: function setDefaultMaxVotersForFtso(uint256 _defaultMaxVotersForFtso) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) SetDefaultMaxVotersForFtso(opts *bind.TransactOpts , _defaultMaxVotersForFtso *big.Int ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "setDefaultMaxVotersForFtso" , _defaultMaxVotersForFtso)
		}

		// SetDefaultMaxVotersForFtso is a paid mutator transaction binding the contract method 0x2ee96140.
		//
		// Solidity: function setDefaultMaxVotersForFtso(uint256 _defaultMaxVotersForFtso) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) SetDefaultMaxVotersForFtso( _defaultMaxVotersForFtso *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.SetDefaultMaxVotersForFtso(&_VoterWhitelister.TransactOpts , _defaultMaxVotersForFtso)
		}

		// SetDefaultMaxVotersForFtso is a paid mutator transaction binding the contract method 0x2ee96140.
		//
		// Solidity: function setDefaultMaxVotersForFtso(uint256 _defaultMaxVotersForFtso) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) SetDefaultMaxVotersForFtso( _defaultMaxVotersForFtso *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.SetDefaultMaxVotersForFtso(&_VoterWhitelister.TransactOpts , _defaultMaxVotersForFtso)
		}
	
		// SetMaxVotersForFtso is a paid mutator transaction binding the contract method 0x7ecfcfa3.
		//
		// Solidity: function setMaxVotersForFtso(uint256 _ftsoIndex, uint256 _newMaxVoters) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) SetMaxVotersForFtso(opts *bind.TransactOpts , _ftsoIndex *big.Int , _newMaxVoters *big.Int ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "setMaxVotersForFtso" , _ftsoIndex, _newMaxVoters)
		}

		// SetMaxVotersForFtso is a paid mutator transaction binding the contract method 0x7ecfcfa3.
		//
		// Solidity: function setMaxVotersForFtso(uint256 _ftsoIndex, uint256 _newMaxVoters) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) SetMaxVotersForFtso( _ftsoIndex *big.Int , _newMaxVoters *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.SetMaxVotersForFtso(&_VoterWhitelister.TransactOpts , _ftsoIndex, _newMaxVoters)
		}

		// SetMaxVotersForFtso is a paid mutator transaction binding the contract method 0x7ecfcfa3.
		//
		// Solidity: function setMaxVotersForFtso(uint256 _ftsoIndex, uint256 _newMaxVoters) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) SetMaxVotersForFtso( _ftsoIndex *big.Int , _newMaxVoters *big.Int ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.SetMaxVotersForFtso(&_VoterWhitelister.TransactOpts , _ftsoIndex, _newMaxVoters)
		}
	
		// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
		//
		// Solidity: function transferGovernance(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactor) TransferGovernance(opts *bind.TransactOpts , _governance common.Address ) (*types.Transaction, error) {
			return _VoterWhitelister.contract.Transact(opts, "transferGovernance" , _governance)
		}

		// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
		//
		// Solidity: function transferGovernance(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterSession) TransferGovernance( _governance common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.TransferGovernance(&_VoterWhitelister.TransactOpts , _governance)
		}

		// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
		//
		// Solidity: function transferGovernance(address _governance) returns()
		func (_VoterWhitelister *VoterWhitelisterTransactorSession) TransferGovernance( _governance common.Address ) (*types.Transaction, error) {
		  return _VoterWhitelister.Contract.TransferGovernance(&_VoterWhitelister.TransactOpts , _governance)
		}
	

	

	

	
		// VoterWhitelisterGovernanceProposedIterator is returned from FilterGovernanceProposed and is used to iterate over the raw logs and unpacked data for GovernanceProposed events raised by the VoterWhitelister contract.
		type VoterWhitelisterGovernanceProposedIterator struct {
			Event *VoterWhitelisterGovernanceProposed // Event containing the contract specifics and raw log

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
		func (it *VoterWhitelisterGovernanceProposedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(VoterWhitelisterGovernanceProposed)
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
				it.Event = new(VoterWhitelisterGovernanceProposed)
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
		func (it *VoterWhitelisterGovernanceProposedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *VoterWhitelisterGovernanceProposedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// VoterWhitelisterGovernanceProposed represents a GovernanceProposed event raised by the VoterWhitelister contract.
		type VoterWhitelisterGovernanceProposed struct { 
			ProposedGovernance common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterGovernanceProposed is a free log retrieval operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
		//
		// Solidity: event GovernanceProposed(address proposedGovernance)
 		func (_VoterWhitelister *VoterWhitelisterFilterer) FilterGovernanceProposed(opts *bind.FilterOpts) (*VoterWhitelisterGovernanceProposedIterator, error) {
			
			

			logs, sub, err := _VoterWhitelister.contract.FilterLogs(opts, "GovernanceProposed")
			if err != nil {
				return nil, err
			}
			return &VoterWhitelisterGovernanceProposedIterator{contract: _VoterWhitelister.contract, event: "GovernanceProposed", logs: logs, sub: sub}, nil
 		}

		// WatchGovernanceProposed is a free log subscription operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
		//
		// Solidity: event GovernanceProposed(address proposedGovernance)
		func (_VoterWhitelister *VoterWhitelisterFilterer) WatchGovernanceProposed(opts *bind.WatchOpts, sink chan<- *VoterWhitelisterGovernanceProposed) (event.Subscription, error) {
			
			

			logs, sub, err := _VoterWhitelister.contract.WatchLogs(opts, "GovernanceProposed")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(VoterWhitelisterGovernanceProposed)
						if err := _VoterWhitelister.contract.UnpackLog(event, "GovernanceProposed", log); err != nil {
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

		// ParseGovernanceProposed is a log parse operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
		//
		// Solidity: event GovernanceProposed(address proposedGovernance)
		func (_VoterWhitelister *VoterWhitelisterFilterer) ParseGovernanceProposed(log types.Log) (*VoterWhitelisterGovernanceProposed, error) {
			event := new(VoterWhitelisterGovernanceProposed)
			if err := _VoterWhitelister.contract.UnpackLog(event, "GovernanceProposed", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// VoterWhitelisterGovernanceUpdatedIterator is returned from FilterGovernanceUpdated and is used to iterate over the raw logs and unpacked data for GovernanceUpdated events raised by the VoterWhitelister contract.
		type VoterWhitelisterGovernanceUpdatedIterator struct {
			Event *VoterWhitelisterGovernanceUpdated // Event containing the contract specifics and raw log

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
		func (it *VoterWhitelisterGovernanceUpdatedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(VoterWhitelisterGovernanceUpdated)
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
				it.Event = new(VoterWhitelisterGovernanceUpdated)
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
		func (it *VoterWhitelisterGovernanceUpdatedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *VoterWhitelisterGovernanceUpdatedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// VoterWhitelisterGovernanceUpdated represents a GovernanceUpdated event raised by the VoterWhitelister contract.
		type VoterWhitelisterGovernanceUpdated struct { 
			OldGovernance common.Address; 
			NewGoveranance common.Address; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterGovernanceUpdated is a free log retrieval operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
		//
		// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
 		func (_VoterWhitelister *VoterWhitelisterFilterer) FilterGovernanceUpdated(opts *bind.FilterOpts) (*VoterWhitelisterGovernanceUpdatedIterator, error) {
			
			
			

			logs, sub, err := _VoterWhitelister.contract.FilterLogs(opts, "GovernanceUpdated")
			if err != nil {
				return nil, err
			}
			return &VoterWhitelisterGovernanceUpdatedIterator{contract: _VoterWhitelister.contract, event: "GovernanceUpdated", logs: logs, sub: sub}, nil
 		}

		// WatchGovernanceUpdated is a free log subscription operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
		//
		// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
		func (_VoterWhitelister *VoterWhitelisterFilterer) WatchGovernanceUpdated(opts *bind.WatchOpts, sink chan<- *VoterWhitelisterGovernanceUpdated) (event.Subscription, error) {
			
			
			

			logs, sub, err := _VoterWhitelister.contract.WatchLogs(opts, "GovernanceUpdated")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(VoterWhitelisterGovernanceUpdated)
						if err := _VoterWhitelister.contract.UnpackLog(event, "GovernanceUpdated", log); err != nil {
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

		// ParseGovernanceUpdated is a log parse operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
		//
		// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
		func (_VoterWhitelister *VoterWhitelisterFilterer) ParseGovernanceUpdated(log types.Log) (*VoterWhitelisterGovernanceUpdated, error) {
			event := new(VoterWhitelisterGovernanceUpdated)
			if err := _VoterWhitelister.contract.UnpackLog(event, "GovernanceUpdated", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// VoterWhitelisterVoterRemovedFromWhitelistIterator is returned from FilterVoterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for VoterRemovedFromWhitelist events raised by the VoterWhitelister contract.
		type VoterWhitelisterVoterRemovedFromWhitelistIterator struct {
			Event *VoterWhitelisterVoterRemovedFromWhitelist // Event containing the contract specifics and raw log

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
		func (it *VoterWhitelisterVoterRemovedFromWhitelistIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(VoterWhitelisterVoterRemovedFromWhitelist)
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
				it.Event = new(VoterWhitelisterVoterRemovedFromWhitelist)
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
		func (it *VoterWhitelisterVoterRemovedFromWhitelistIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *VoterWhitelisterVoterRemovedFromWhitelistIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// VoterWhitelisterVoterRemovedFromWhitelist represents a VoterRemovedFromWhitelist event raised by the VoterWhitelister contract.
		type VoterWhitelisterVoterRemovedFromWhitelist struct { 
			Voter common.Address; 
			FtsoIndex *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterVoterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x33359f2769756ca8d0da4683f25ee440744d6f18bfb166dbfb59315a8c62b016.
		//
		// Solidity: event VoterRemovedFromWhitelist(address voter, uint256 ftsoIndex)
 		func (_VoterWhitelister *VoterWhitelisterFilterer) FilterVoterRemovedFromWhitelist(opts *bind.FilterOpts) (*VoterWhitelisterVoterRemovedFromWhitelistIterator, error) {
			
			
			

			logs, sub, err := _VoterWhitelister.contract.FilterLogs(opts, "VoterRemovedFromWhitelist")
			if err != nil {
				return nil, err
			}
			return &VoterWhitelisterVoterRemovedFromWhitelistIterator{contract: _VoterWhitelister.contract, event: "VoterRemovedFromWhitelist", logs: logs, sub: sub}, nil
 		}

		// WatchVoterRemovedFromWhitelist is a free log subscription operation binding the contract event 0x33359f2769756ca8d0da4683f25ee440744d6f18bfb166dbfb59315a8c62b016.
		//
		// Solidity: event VoterRemovedFromWhitelist(address voter, uint256 ftsoIndex)
		func (_VoterWhitelister *VoterWhitelisterFilterer) WatchVoterRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *VoterWhitelisterVoterRemovedFromWhitelist) (event.Subscription, error) {
			
			
			

			logs, sub, err := _VoterWhitelister.contract.WatchLogs(opts, "VoterRemovedFromWhitelist")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(VoterWhitelisterVoterRemovedFromWhitelist)
						if err := _VoterWhitelister.contract.UnpackLog(event, "VoterRemovedFromWhitelist", log); err != nil {
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

		// ParseVoterRemovedFromWhitelist is a log parse operation binding the contract event 0x33359f2769756ca8d0da4683f25ee440744d6f18bfb166dbfb59315a8c62b016.
		//
		// Solidity: event VoterRemovedFromWhitelist(address voter, uint256 ftsoIndex)
		func (_VoterWhitelister *VoterWhitelisterFilterer) ParseVoterRemovedFromWhitelist(log types.Log) (*VoterWhitelisterVoterRemovedFromWhitelist, error) {
			event := new(VoterWhitelisterVoterRemovedFromWhitelist)
			if err := _VoterWhitelister.contract.UnpackLog(event, "VoterRemovedFromWhitelist", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// VoterWhitelisterVoterWhitelistedIterator is returned from FilterVoterWhitelisted and is used to iterate over the raw logs and unpacked data for VoterWhitelisted events raised by the VoterWhitelister contract.
		type VoterWhitelisterVoterWhitelistedIterator struct {
			Event *VoterWhitelisterVoterWhitelisted // Event containing the contract specifics and raw log

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
		func (it *VoterWhitelisterVoterWhitelistedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(VoterWhitelisterVoterWhitelisted)
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
				it.Event = new(VoterWhitelisterVoterWhitelisted)
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
		func (it *VoterWhitelisterVoterWhitelistedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *VoterWhitelisterVoterWhitelistedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// VoterWhitelisterVoterWhitelisted represents a VoterWhitelisted event raised by the VoterWhitelister contract.
		type VoterWhitelisterVoterWhitelisted struct { 
			Voter common.Address; 
			FtsoIndex *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterVoterWhitelisted is a free log retrieval operation binding the contract event 0x66a8b13abe95391d1851f5bc319f3dde54ce8f2f40a5fe226aa3251d805832e3.
		//
		// Solidity: event VoterWhitelisted(address voter, uint256 ftsoIndex)
 		func (_VoterWhitelister *VoterWhitelisterFilterer) FilterVoterWhitelisted(opts *bind.FilterOpts) (*VoterWhitelisterVoterWhitelistedIterator, error) {
			
			
			

			logs, sub, err := _VoterWhitelister.contract.FilterLogs(opts, "VoterWhitelisted")
			if err != nil {
				return nil, err
			}
			return &VoterWhitelisterVoterWhitelistedIterator{contract: _VoterWhitelister.contract, event: "VoterWhitelisted", logs: logs, sub: sub}, nil
 		}

		// WatchVoterWhitelisted is a free log subscription operation binding the contract event 0x66a8b13abe95391d1851f5bc319f3dde54ce8f2f40a5fe226aa3251d805832e3.
		//
		// Solidity: event VoterWhitelisted(address voter, uint256 ftsoIndex)
		func (_VoterWhitelister *VoterWhitelisterFilterer) WatchVoterWhitelisted(opts *bind.WatchOpts, sink chan<- *VoterWhitelisterVoterWhitelisted) (event.Subscription, error) {
			
			
			

			logs, sub, err := _VoterWhitelister.contract.WatchLogs(opts, "VoterWhitelisted")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(VoterWhitelisterVoterWhitelisted)
						if err := _VoterWhitelister.contract.UnpackLog(event, "VoterWhitelisted", log); err != nil {
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

		// ParseVoterWhitelisted is a log parse operation binding the contract event 0x66a8b13abe95391d1851f5bc319f3dde54ce8f2f40a5fe226aa3251d805832e3.
		//
		// Solidity: event VoterWhitelisted(address voter, uint256 ftsoIndex)
		func (_VoterWhitelister *VoterWhitelisterFilterer) ParseVoterWhitelisted(log types.Log) (*VoterWhitelisterVoterWhitelisted, error) {
			event := new(VoterWhitelisterVoterWhitelisted)
			if err := _VoterWhitelister.contract.UnpackLog(event, "VoterWhitelisted", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	


