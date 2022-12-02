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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"current_sequencer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"next_sequencer\",\"type\":\"uint256\"}],\"name\":\"SequencerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AddSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"GetSequencer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161037138038061037183398101604081905261002f9161008b565b60008181555b6010811015610084576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018190558061007c816100a4565b915050610035565b50506100cb565b60006020828403121561009d57600080fd5b5051919050565b6000600182016100c457634e487b7160e01b600052601160045260246000fd5b5060010190565b610297806100da6000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063b70e6be611610050578063b70e6be6146100a5578063d07eb531146100ae578063eed3e6a4146100bd57600080fd5b80632dcaa5961461006c578063af6341f914610092575b600080fd5b61007f61007a3660046101a9565b6100c5565b6040519081526020015b60405180910390f35b61007f6100a03660046101a9565b610188565b61007f60005481565b60405160008152602001610089565b61007f601081565b600080600054836100d691906101f1565b905060006100e5601083610208565b90506000600182815481106100fc576100fc610243565b60009182526020822001549150600160106101178583610272565b6101219190610208565b8154811061013157610131610243565b6000918252602091829020015460408051898152928301859052820181905291507f30d0769c6324f9df783f97d31af6e0a62d6855f83f6e20c63263580a8a5401249060600160405180910390a150949350505050565b6001818154811061019857600080fd5b600091825260209091200154905081565b6000602082840312156101bb57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610203576102036101c2565b500390565b60008261023e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115610285576102856101c2565b50019056fea164736f6c634300080f000a",
}

// SequencerSelectorABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerSelectorMetaData.ABI instead.
var SequencerSelectorABI = SequencerSelectorMetaData.ABI

// SequencerSelectorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerSelectorMetaData.Bin instead.
var SequencerSelectorBin = SequencerSelectorMetaData.Bin

// DeploySequencerSelector deploys a new Ethereum contract, binding an instance of SequencerSelector to it.
func DeploySequencerSelector(auth *bind.TransactOpts, backend bind.ContractBackend, epoch *big.Int) (common.Address, *types.Transaction, *SequencerSelector, error) {
	parsed, err := SequencerSelectorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerSelectorBin), backend, epoch)
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

// GenesisEpoch is a free data retrieval call binding the contract method 0xb70e6be6.
//
// Solidity: function genesisEpoch() view returns(uint256)
func (_SequencerSelector *SequencerSelectorCaller) GenesisEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerSelector.contract.Call(opts, &out, "genesisEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisEpoch is a free data retrieval call binding the contract method 0xb70e6be6.
//
// Solidity: function genesisEpoch() view returns(uint256)
func (_SequencerSelector *SequencerSelectorSession) GenesisEpoch() (*big.Int, error) {
	return _SequencerSelector.Contract.GenesisEpoch(&_SequencerSelector.CallOpts)
}

// GenesisEpoch is a free data retrieval call binding the contract method 0xb70e6be6.
//
// Solidity: function genesisEpoch() view returns(uint256)
func (_SequencerSelector *SequencerSelectorCallerSession) GenesisEpoch() (*big.Int, error) {
	return _SequencerSelector.Contract.GenesisEpoch(&_SequencerSelector.CallOpts)
}

// SequencerIDs is a free data retrieval call binding the contract method 0xaf6341f9.
//
// Solidity: function sequencerIDs(uint256 ) view returns(uint256)
func (_SequencerSelector *SequencerSelectorCaller) SequencerIDs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SequencerSelector.contract.Call(opts, &out, "sequencerIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerIDs is a free data retrieval call binding the contract method 0xaf6341f9.
//
// Solidity: function sequencerIDs(uint256 ) view returns(uint256)
func (_SequencerSelector *SequencerSelectorSession) SequencerIDs(arg0 *big.Int) (*big.Int, error) {
	return _SequencerSelector.Contract.SequencerIDs(&_SequencerSelector.CallOpts, arg0)
}

// SequencerIDs is a free data retrieval call binding the contract method 0xaf6341f9.
//
// Solidity: function sequencerIDs(uint256 ) view returns(uint256)
func (_SequencerSelector *SequencerSelectorCallerSession) SequencerIDs(arg0 *big.Int) (*big.Int, error) {
	return _SequencerSelector.Contract.SequencerIDs(&_SequencerSelector.CallOpts, arg0)
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
// Solidity: function GetSequencer(uint256 epoch) returns(uint256)
func (_SequencerSelector *SequencerSelectorTransactor) GetSequencer(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.contract.Transact(opts, "GetSequencer", epoch)
}

// GetSequencer is a paid mutator transaction binding the contract method 0x2dcaa596.
//
// Solidity: function GetSequencer(uint256 epoch) returns(uint256)
func (_SequencerSelector *SequencerSelectorSession) GetSequencer(epoch *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.Contract.GetSequencer(&_SequencerSelector.TransactOpts, epoch)
}

// GetSequencer is a paid mutator transaction binding the contract method 0x2dcaa596.
//
// Solidity: function GetSequencer(uint256 epoch) returns(uint256)
func (_SequencerSelector *SequencerSelectorTransactorSession) GetSequencer(epoch *big.Int) (*types.Transaction, error) {
	return _SequencerSelector.Contract.GetSequencer(&_SequencerSelector.TransactOpts, epoch)
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
	Epoch            *big.Int
	CurrentSequencer *big.Int
	NextSequencer    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSequencerChanged is a free log retrieval operation binding the contract event 0x30d0769c6324f9df783f97d31af6e0a62d6855f83f6e20c63263580a8a540124.
//
// Solidity: event SequencerChanged(uint256 epoch, uint256 current_sequencer, uint256 next_sequencer)
func (_SequencerSelector *SequencerSelectorFilterer) FilterSequencerChanged(opts *bind.FilterOpts) (*SequencerSelectorSequencerChangedIterator, error) {

	logs, sub, err := _SequencerSelector.contract.FilterLogs(opts, "SequencerChanged")
	if err != nil {
		return nil, err
	}
	return &SequencerSelectorSequencerChangedIterator{contract: _SequencerSelector.contract, event: "SequencerChanged", logs: logs, sub: sub}, nil
}

// WatchSequencerChanged is a free log subscription operation binding the contract event 0x30d0769c6324f9df783f97d31af6e0a62d6855f83f6e20c63263580a8a540124.
//
// Solidity: event SequencerChanged(uint256 epoch, uint256 current_sequencer, uint256 next_sequencer)
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

// ParseSequencerChanged is a log parse operation binding the contract event 0x30d0769c6324f9df783f97d31af6e0a62d6855f83f6e20c63263580a8a540124.
//
// Solidity: event SequencerChanged(uint256 epoch, uint256 current_sequencer, uint256 next_sequencer)
func (_SequencerSelector *SequencerSelectorFilterer) ParseSequencerChanged(log types.Log) (*SequencerSelectorSequencerChanged, error) {
	event := new(SequencerSelectorSequencerChanged)
	if err := _SequencerSelector.contract.UnpackLog(event, "SequencerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
