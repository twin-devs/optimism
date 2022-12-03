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
)

// SequencerSelectorMetaData contains all meta data concerning the SequencerSelector contract.
var SequencerSelectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"current_sequencer_pubkey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"next_sequencer_pubkey\",\"type\":\"string\"}],\"name\":\"SequencerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AddSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"GetSequencer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000d2438038062000d248339810160408190526200003491620002e9565b604080516101008101909152606060808201818152829162000c0460a0840139815260200160405180608001604052806060815260200162000ba460609139815260200160405180608001604052806060815260200162000cc460609139815260200160405180608001604052806060815260200162000c64606091399052620000c390600290600462000212565b50620000cf81620000d6565b5062000474565b600054610100900460ff1615808015620000f75750600054600160ff909116105b8062000127575062000114306200020360201b620005091760201c565b15801562000127575060005460ff166001145b6200018f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff191660011790558015620001b3576000805461ff0019166101001790555b60018290558015620001ff576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6001600160a01b03163b151590565b8280548282559060005260206000209081019282156200025d579160200282015b828111156200025d57825182906200024c9082620003a8565b509160200191906001019062000233565b506200026b9291506200026f565b5090565b808211156200026b57600062000286828262000290565b506001016200026f565b5080546200029e9062000319565b6000825580601f10620002af575050565b601f016020900490600052602060002090810190620002cf9190620002d2565b50565b5b808211156200026b5760008155600101620002d3565b600060208284031215620002fc57600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200032e57607f821691505b6020821081036200034f57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620003a357600081815260208120601f850160051c810160208610156200037e5750805b601f850160051c820191505b818110156200039f578281556001016200038a565b5050505b505050565b81516001600160401b03811115620003c457620003c462000303565b620003dc81620003d5845462000319565b8462000355565b602080601f831160018114620004145760008415620003fb5750858301515b600019600386901b1c1916600185901b1785556200039f565b600085815260208120601f198616915b82811015620004455788860151825594840194600190910190840162000424565b5085821015620004645787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61072080620004846000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063d07eb53111610050578063d07eb531146100ca578063eed3e6a4146100d9578063fe4b84df146100e157600080fd5b80632dcaa596146100775780636ba7ccff146100a057806370872aa5146100b3575b600080fd5b61008a610085366004610525565b6100f6565b60405161009791906105a9565b60405180910390f35b61008a6100ae366004610525565b6102c9565b6100bc60015481565b604051908152602001610097565b60405160008152602001610097565b6100bc600481565b6100f46100ef366004610525565b610375565b005b606060006001548361010891906105f2565b90506000610117600483610609565b905060006002828154811061012e5761012e610644565b90600052602060002001805461014390610673565b80601f016020809104026020016040519081016040528092919081815260200182805461016f90610673565b80156101bc5780601f10610191576101008083540402835291602001916101bc565b820191906000526020600020905b81548152906001019060200180831161019f57829003601f168201915b505050505090506000600260048460016101d691906106c6565b6101e09190610609565b815481106101f0576101f0610644565b90600052602060002001805461020590610673565b80601f016020809104026020016040519081016040528092919081815260200182805461023190610673565b801561027e5780601f106102535761010080835404028352916020019161027e565b820191906000526020600020905b81548152906001019060200180831161026157829003601f168201915b505050505090507f416832484bf807592930bba198c0b95699a3b155d0d9f9a4dc572999917262d28683836040516102b8939291906106de565b60405180910390a150949350505050565b600281815481106102d957600080fd5b9060005260206000200160009150905080546102f490610673565b80601f016020809104026020016040519081016040528092919081815260200182805461032090610673565b801561036d5780601f106103425761010080835404028352916020019161036d565b820191906000526020600020905b81548152906001019060200180831161035057829003601f168201915b505050505081565b600054610100900460ff16158080156103955750600054600160ff909116105b806103af5750303b1580156103af575060005460ff166001145b61043f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561049d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001829055801561050557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60006020828403121561053757600080fd5b5035919050565b6000815180845260005b8181101561056457602081850181015186830182015201610548565b81811115610576576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006105bc602083018461053e565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610604576106046105c3565b500390565b60008261063f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c9082168061068757607f821691505b6020821081036106c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600082198211156106d9576106d96105c3565b500190565b8381526060602082015260006106f7606083018561053e565b8281036040840152610709818561053e565b969550505050505056fea164736f6c634300080f000a616230356436653137616461623932613032353239383133333165613232633039636233333039393133356562613638393133636136353437303964653539346361306236333136393431616464303133396634353261663838346465623533386136343666373163663535666162653631373838623739383961313235623063353433653637396332386630376537303433303063393066376239633264393166306131336534663737343163643335613262373165616433346633323932383532656662343033373932666462333839623830633832363630383239343833346435343036623532386336363936346634383064386138366330613536313230386132383135663031373366316130623338393632633330616662626166616633306131393461366235333237303362656231623838303831636335623666326464333537613733396364383966663430303131396432376438643536306133343932663836393637656339363738383462656134373931333939356163",
}

// SequencerSelectorABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerSelectorMetaData.ABI instead.
var SequencerSelectorABI = SequencerSelectorMetaData.ABI

// SequencerSelectorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerSelectorMetaData.Bin instead.
var SequencerSelectorBin = SequencerSelectorMetaData.Bin

// DeploySequencerSelector deploys a new Ethereum contract, binding an instance of SequencerSelector to it.
func DeploySequencerSelector(auth *bind.TransactOpts, backend bind.ContractBackend, _startingBlockNumber *big.Int) (common.Address, *types.Transaction, *SequencerSelector, error) {
	parsed, err := SequencerSelectorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerSelectorBin), backend, _startingBlockNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerSelector{SequencerSelectorCaller: SequencerSelectorCaller{contract: contract}, SequencerSelectorTransactor: SequencerSelectorTransactor{contract: contract}, SequencerSelectorFilterer: SequencerSelectorFilterer{contract: contract}}, nil
}

// SequencerSelector is an auto generated Go binding around an Ethereum contract.
type SequencerSelector struct {
	SequencerSelectorCaller     // Read-only binding to the contract
	SequencerSelectorTransactor // Write-only binding to the contract
	SequencerSelectorFilterer   // Log filterer for contract events
}

// SequencerSelectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerSelectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSelectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerSelectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSelectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerSelectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSelectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSelectorSession struct {
	Contract     *SequencerSelector // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SequencerSelectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerSelectorCallerSession struct {
	Contract *SequencerSelectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SequencerSelectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerSelectorTransactorSession struct {
	Contract     *SequencerSelectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SequencerSelectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerSelectorRaw struct {
	Contract *SequencerSelector // Generic contract binding to access the raw methods on
}

// SequencerSelectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerSelectorCallerRaw struct {
	Contract *SequencerSelectorCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerSelectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerSelectorTransactorRaw struct {
	Contract *SequencerSelectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerSelector creates a new instance of SequencerSelector, bound to a specific deployed contract.
func NewSequencerSelector(address common.Address, backend bind.ContractBackend) (*SequencerSelector, error) {
	contract, err := bindSequencerSelector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerSelector{SequencerSelectorCaller: SequencerSelectorCaller{contract: contract}, SequencerSelectorTransactor: SequencerSelectorTransactor{contract: contract}, SequencerSelectorFilterer: SequencerSelectorFilterer{contract: contract}}, nil
}

// NewSequencerSelectorCaller creates a new read-only instance of SequencerSelector, bound to a specific deployed contract.
func NewSequencerSelectorCaller(address common.Address, caller bind.ContractCaller) (*SequencerSelectorCaller, error) {
	contract, err := bindSequencerSelector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerSelectorCaller{contract: contract}, nil
}

// NewSequencerSelectorTransactor creates a new write-only instance of SequencerSelector, bound to a specific deployed contract.
func NewSequencerSelectorTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerSelectorTransactor, error) {
	contract, err := bindSequencerSelector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerSelectorTransactor{contract: contract}, nil
}

// NewSequencerSelectorFilterer creates a new log filterer instance of SequencerSelector, bound to a specific deployed contract.
func NewSequencerSelectorFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerSelectorFilterer, error) {
	contract, err := bindSequencerSelector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerSelectorFilterer{contract: contract}, nil
}

// bindSequencerSelector binds a generic wrapper to an already deployed contract.
func bindSequencerSelector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerSelectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerSelector *SequencerSelectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerSelector.Contract.SequencerSelectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerSelector *SequencerSelectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSelector.Contract.SequencerSelectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerSelector *SequencerSelectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerSelector.Contract.SequencerSelectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerSelector *SequencerSelectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerSelector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerSelector *SequencerSelectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSelector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerSelector *SequencerSelectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerSelector.Contract.contract.Transact(opts, method, params...)
}

// SEQUENCERCOUNT is a free data retrieval call binding the contract method 0xeed3e6a4.
//
// Solidity: function SEQUENCER_COUNT() view returns(uint256)
func (_SequencerSelector *SequencerSelectorCaller) SEQUENCERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerSelector.contract.Call(opts, &out, "SEQUENCER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SEQUENCERCOUNT is a free data retrieval call binding the contract method 0xeed3e6a4.
//
// Solidity: function SEQUENCER_COUNT() view returns(uint256)
func (_SequencerSelector *SequencerSelectorSession) SEQUENCERCOUNT() (*big.Int, error) {
	return _SequencerSelector.Contract.SEQUENCERCOUNT(&_SequencerSelector.CallOpts)
}

// SEQUENCERCOUNT is a free data retrieval call binding the contract method 0xeed3e6a4.
//
// Solidity: function SEQUENCER_COUNT() view returns(uint256)
func (_SequencerSelector *SequencerSelectorCallerSession) SEQUENCERCOUNT() (*big.Int, error) {
	return _SequencerSelector.Contract.SEQUENCERCOUNT(&_SequencerSelector.CallOpts)
}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 ) view returns(string)
func (_SequencerSelector *SequencerSelectorCaller) Sequencers(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _SequencerSelector.contract.Call(opts, &out, "sequencers", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 ) view returns(string)
func (_SequencerSelector *SequencerSelectorSession) Sequencers(arg0 *big.Int) (string, error) {
	return _SequencerSelector.Contract.Sequencers(&_SequencerSelector.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 ) view returns(string)
func (_SequencerSelector *SequencerSelectorCallerSession) Sequencers(arg0 *big.Int) (string, error) {
	return _SequencerSelector.Contract.Sequencers(&_SequencerSelector.CallOpts, arg0)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_SequencerSelector *SequencerSelectorCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerSelector.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_SequencerSelector *SequencerSelectorSession) StartingBlockNumber() (*big.Int, error) {
	return _SequencerSelector.Contract.StartingBlockNumber(&_SequencerSelector.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_SequencerSelector *SequencerSelectorCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _SequencerSelector.Contract.StartingBlockNumber(&_SequencerSelector.CallOpts)
}

// AddSequencer is a paid mutator transaction binding the contract method 0xd07eb531.
//
// Solidity: function AddSequencer() returns(bool)
func (_SequencerSelector *SequencerSelectorTransactor) AddSequencer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSelector.contract.Transact(opts, "AddSequencer")
}

// AddSequencer is a paid mutator transaction binding the contract method 0xd07eb531.
//
// Solidity: function AddSequencer() returns(bool)
func (_SequencerSelector *SequencerSelectorSession) AddSequencer() (*types.Transaction, error) {
	return _SequencerSelector.Contract.AddSequencer(&_SequencerSelector.TransactOpts)
}

// AddSequencer is a paid mutator transaction binding the contract method 0xd07eb531.
//
// Solidity: function AddSequencer() returns(bool)
func (_SequencerSelector *SequencerSelectorTransactorSession) AddSequencer() (*types.Transaction, error) {
	return _SequencerSelector.Contract.AddSequencer(&_SequencerSelector.TransactOpts)
}

// GetSequencer is a paid mutator transaction binding the contract method 0x2dcaa596.
//
// Solidity: function GetSequencer(uint256 blockNumber) returns(string)
func (_SequencerSelector *SequencerSelectorTransactor) GetSequencer(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.contract.Transact(opts, "GetSequencer", blockNumber)
}

// GetSequencer is a paid mutator transaction binding the contract method 0x2dcaa596.
//
// Solidity: function GetSequencer(uint256 blockNumber) returns(string)
func (_SequencerSelector *SequencerSelectorSession) GetSequencer(blockNumber *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.Contract.GetSequencer(&_SequencerSelector.TransactOpts, blockNumber)
}

// GetSequencer is a paid mutator transaction binding the contract method 0x2dcaa596.
//
// Solidity: function GetSequencer(uint256 blockNumber) returns(string)
func (_SequencerSelector *SequencerSelectorTransactorSession) GetSequencer(blockNumber *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.Contract.GetSequencer(&_SequencerSelector.TransactOpts, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _startingBlockNumber) returns()
func (_SequencerSelector *SequencerSelectorTransactor) Initialize(opts *bind.TransactOpts, _startingBlockNumber *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.contract.Transact(opts, "initialize", _startingBlockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _startingBlockNumber) returns()
func (_SequencerSelector *SequencerSelectorSession) Initialize(_startingBlockNumber *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.Contract.Initialize(&_SequencerSelector.TransactOpts, _startingBlockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _startingBlockNumber) returns()
func (_SequencerSelector *SequencerSelectorTransactorSession) Initialize(_startingBlockNumber *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.Contract.Initialize(&_SequencerSelector.TransactOpts, _startingBlockNumber)
}

// SequencerSelectorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SequencerSelector contract.
type SequencerSelectorInitializedIterator struct {
	Event *SequencerSelectorInitialized // Event containing the contract specifics and raw log

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
func (it *SequencerSelectorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSelectorInitialized)
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
		it.Event = new(SequencerSelectorInitialized)
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
func (it *SequencerSelectorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSelectorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSelectorInitialized represents a Initialized event raised by the SequencerSelector contract.
type SequencerSelectorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerSelector *SequencerSelectorFilterer) FilterInitialized(opts *bind.FilterOpts) (*SequencerSelectorInitializedIterator, error) {

	logs, sub, err := _SequencerSelector.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SequencerSelectorInitializedIterator{contract: _SequencerSelector.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerSelector *SequencerSelectorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SequencerSelectorInitialized) (event.Subscription, error) {

	logs, sub, err := _SequencerSelector.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSelectorInitialized)
				if err := _SequencerSelector.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SequencerSelector *SequencerSelectorFilterer) ParseInitialized(log types.Log) (*SequencerSelectorInitialized, error) {
	event := new(SequencerSelectorInitialized)
	if err := _SequencerSelector.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerSelectorSequencerChangedIterator is returned from FilterSequencerChanged and is used to iterate over the raw logs and unpacked data for SequencerChanged events raised by the SequencerSelector contract.
type SequencerSelectorSequencerChangedIterator struct {
	Event *SequencerSelectorSequencerChanged // Event containing the contract specifics and raw log

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
func (it *SequencerSelectorSequencerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerSelectorSequencerChanged)
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
		it.Event = new(SequencerSelectorSequencerChanged)
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
func (it *SequencerSelectorSequencerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerSelectorSequencerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerSelectorSequencerChanged represents a SequencerChanged event raised by the SequencerSelector contract.
type SequencerSelectorSequencerChanged struct {
	BlockNumber            *big.Int
	CurrentSequencerPubkey string
	NextSequencerPubkey    string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSequencerChanged is a free log retrieval operation binding the contract event 0x416832484bf807592930bba198c0b95699a3b155d0d9f9a4dc572999917262d2.
//
// Solidity: event SequencerChanged(uint256 blockNumber, string current_sequencer_pubkey, string next_sequencer_pubkey)
func (_SequencerSelector *SequencerSelectorFilterer) FilterSequencerChanged(opts *bind.FilterOpts) (*SequencerSelectorSequencerChangedIterator, error) {

	logs, sub, err := _SequencerSelector.contract.FilterLogs(opts, "SequencerChanged")
	if err != nil {
		return nil, err
	}
	return &SequencerSelectorSequencerChangedIterator{contract: _SequencerSelector.contract, event: "SequencerChanged", logs: logs, sub: sub}, nil
}

// WatchSequencerChanged is a free log subscription operation binding the contract event 0x416832484bf807592930bba198c0b95699a3b155d0d9f9a4dc572999917262d2.
//
// Solidity: event SequencerChanged(uint256 blockNumber, string current_sequencer_pubkey, string next_sequencer_pubkey)
func (_SequencerSelector *SequencerSelectorFilterer) WatchSequencerChanged(opts *bind.WatchOpts, sink chan<- *SequencerSelectorSequencerChanged) (event.Subscription, error) {

	logs, sub, err := _SequencerSelector.contract.WatchLogs(opts, "SequencerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerSelectorSequencerChanged)
				if err := _SequencerSelector.contract.UnpackLog(event, "SequencerChanged", log); err != nil {
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

// ParseSequencerChanged is a log parse operation binding the contract event 0x416832484bf807592930bba198c0b95699a3b155d0d9f9a4dc572999917262d2.
//
// Solidity: event SequencerChanged(uint256 blockNumber, string current_sequencer_pubkey, string next_sequencer_pubkey)
func (_SequencerSelector *SequencerSelectorFilterer) ParseSequencerChanged(log types.Log) (*SequencerSelectorSequencerChanged, error) {
	event := new(SequencerSelectorSequencerChanged)
	if err := _SequencerSelector.contract.UnpackLog(event, "SequencerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
