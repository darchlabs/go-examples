Example(
        uint256 amount,
        address indexed userAddr,
        uint256 amount1,
        uint256 indexed num,
        string indexed lala,
        string lala2
    )// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"lala\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"lala2\",\"type\":\"string\"}],\"name\":\"Example\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"example\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExample\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008055610167806100246000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063477e4a021461003b578063bfdb629c14610050575b600080fd5b61004e6100493660046100f1565b610065565b005b60005460405190815260200160405180910390f35b6000819055604051636e616e6160e01b81526004016040519081900390206005337f52b848661316b3434e299abb1e7191f4dd3f3ff7121adfea56a62366de91af7d846100b381600161010a565b6040516100e69291909182526020820152606060408201819052600490820152636461646160e01b608082015260a00190565b60405180910390a450565b60006020828403121561010357600080fd5b5035919050565b8082018082111561012b57634e487b7160e01b600052601160045260246000fd5b9291505056fea26469706673582212209b76d0f9d052f9dcc0bc69c8978c3848092b616292ca2078248ad65af6647ba164736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetExample is a free data retrieval call binding the contract method 0xbfdb629c.
//
// Solidity: function getExample() view returns(uint256)
func (_Contract *ContractCaller) GetExample(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getExample")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExample is a free data retrieval call binding the contract method 0xbfdb629c.
//
// Solidity: function getExample() view returns(uint256)
func (_Contract *ContractSession) GetExample() (*big.Int, error) {
	return _Contract.Contract.GetExample(&_Contract.CallOpts)
}

// GetExample is a free data retrieval call binding the contract method 0xbfdb629c.
//
// Solidity: function getExample() view returns(uint256)
func (_Contract *ContractCallerSession) GetExample() (*big.Int, error) {
	return _Contract.Contract.GetExample(&_Contract.CallOpts)
}

// Example is a paid mutator transaction binding the contract method 0x477e4a02.
//
// Solidity: function example(uint256 amount) returns()
func (_Contract *ContractTransactor) Example(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "example", amount)
}

// Example is a paid mutator transaction binding the contract method 0x477e4a02.
//
// Solidity: function example(uint256 amount) returns()
func (_Contract *ContractSession) Example(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Example(&_Contract.TransactOpts, amount)
}

// Example is a paid mutator transaction binding the contract method 0x477e4a02.
//
// Solidity: function example(uint256 amount) returns()
func (_Contract *ContractTransactorSession) Example(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Example(&_Contract.TransactOpts, amount)
}

// ContractExampleIterator is returned from FilterExample and is used to iterate over the raw logs and unpacked data for Example events raised by the Contract contract.
type ContractExampleIterator struct {
	Event *ContractExample // Event containing the contract specifics and raw log

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
func (it *ContractExampleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExample)
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
		it.Event = new(ContractExample)
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
func (it *ContractExampleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExampleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExample represents a Example event raised by the Contract contract.
type ContractExample struct {
	Amount   *big.Int
	UserAddr common.Address
	Amount1  *big.Int
	Num      *big.Int
	Lala     common.Hash
	Lala2    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExample is a free log retrieval operation binding the contract event 0x52b848661316b3434e299abb1e7191f4dd3f3ff7121adfea56a62366de91af7d.
//
// Solidity: event Example(uint256 amount, address indexed userAddr, uint256 amount1, uint256 indexed num, string indexed lala, string lala2)
func (_Contract *ContractFilterer) FilterExample(opts *bind.FilterOpts, userAddr []common.Address, num []*big.Int, lala []string) (*ContractExampleIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	var numRule []interface{}
	for _, numItem := range num {
		numRule = append(numRule, numItem)
	}
	var lalaRule []interface{}
	for _, lalaItem := range lala {
		lalaRule = append(lalaRule, lalaItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Example", userAddrRule, numRule, lalaRule)
	if err != nil {
		return nil, err
	}
	return &ContractExampleIterator{contract: _Contract.contract, event: "Example", logs: logs, sub: sub}, nil
}

// WatchExample is a free log subscription operation binding the contract event 0x52b848661316b3434e299abb1e7191f4dd3f3ff7121adfea56a62366de91af7d.
//
// Solidity: event Example(uint256 amount, address indexed userAddr, uint256 amount1, uint256 indexed num, string indexed lala, string lala2)
func (_Contract *ContractFilterer) WatchExample(opts *bind.WatchOpts, sink chan<- *ContractExample, userAddr []common.Address, num []*big.Int, lala []string) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	var numRule []interface{}
	for _, numItem := range num {
		numRule = append(numRule, numItem)
	}
	var lalaRule []interface{}
	for _, lalaItem := range lala {
		lalaRule = append(lalaRule, lalaItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Example", userAddrRule, numRule, lalaRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExample)
				if err := _Contract.contract.UnpackLog(event, "Example", log); err != nil {
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

// ParseExample is a log parse operation binding the contract event 0x52b848661316b3434e299abb1e7191f4dd3f3ff7121adfea56a62366de91af7d.
//
// Solidity: event Example(uint256 amount, address indexed userAddr, uint256 amount1, uint256 indexed num, string indexed lala, string lala2)
func (_Contract *ContractFilterer) ParseExample(log types.Log) (*ContractExample, error) {
	event := new(ContractExample)
	if err := _Contract.contract.UnpackLog(event, "Example", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
