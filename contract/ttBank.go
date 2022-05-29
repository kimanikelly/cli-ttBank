// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ITTBankBankDetails is an auto generated low-level Go binding around an user-defined struct.
type ITTBankBankDetails struct {
	AccountNumber *big.Int
	AccountName   common.Address
	Balance       *big.Int
}

// TTBankABI is the input ABI used to generate the binding from.
const TTBankABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountName\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingBalance\",\"type\":\"uint256\"}],\"name\":\"AccountOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountName\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountName\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bankBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startingBalance\",\"type\":\"uint256\"}],\"name\":\"openAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accountNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"accountName\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structITTBank.BankDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TTBank is an auto generated Go binding around an Ethereum contract.
type TTBank struct {
	TTBankCaller     // Read-only binding to the contract
	TTBankTransactor // Write-only binding to the contract
	TTBankFilterer   // Log filterer for contract events
}

// TTBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type TTBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TTBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TTBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TTBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TTBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TTBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TTBankSession struct {
	Contract     *TTBank           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TTBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TTBankCallerSession struct {
	Contract *TTBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TTBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TTBankTransactorSession struct {
	Contract     *TTBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TTBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type TTBankRaw struct {
	Contract *TTBank // Generic contract binding to access the raw methods on
}

// TTBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TTBankCallerRaw struct {
	Contract *TTBankCaller // Generic read-only contract binding to access the raw methods on
}

// TTBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TTBankTransactorRaw struct {
	Contract *TTBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTTBank creates a new instance of TTBank, bound to a specific deployed contract.
func NewTTBank(address common.Address, backend bind.ContractBackend) (*TTBank, error) {
	contract, err := bindTTBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TTBank{TTBankCaller: TTBankCaller{contract: contract}, TTBankTransactor: TTBankTransactor{contract: contract}, TTBankFilterer: TTBankFilterer{contract: contract}}, nil
}

// NewTTBankCaller creates a new read-only instance of TTBank, bound to a specific deployed contract.
func NewTTBankCaller(address common.Address, caller bind.ContractCaller) (*TTBankCaller, error) {
	contract, err := bindTTBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TTBankCaller{contract: contract}, nil
}

// NewTTBankTransactor creates a new write-only instance of TTBank, bound to a specific deployed contract.
func NewTTBankTransactor(address common.Address, transactor bind.ContractTransactor) (*TTBankTransactor, error) {
	contract, err := bindTTBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TTBankTransactor{contract: contract}, nil
}

// NewTTBankFilterer creates a new log filterer instance of TTBank, bound to a specific deployed contract.
func NewTTBankFilterer(address common.Address, filterer bind.ContractFilterer) (*TTBankFilterer, error) {
	contract, err := bindTTBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TTBankFilterer{contract: contract}, nil
}

// bindTTBank binds a generic wrapper to an already deployed contract.
func bindTTBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TTBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TTBank *TTBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TTBank.Contract.TTBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TTBank *TTBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TTBank.Contract.TTBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TTBank *TTBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TTBank.Contract.TTBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TTBank *TTBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TTBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TTBank *TTBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TTBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TTBank *TTBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TTBank.Contract.contract.Transact(opts, method, params...)
}

// BankBalance is a free data retrieval call binding the contract method 0x28657aa5.
//
// Solidity: function bankBalance() view returns(uint256)
func (_TTBank *TTBankCaller) BankBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TTBank.contract.Call(opts, &out, "bankBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BankBalance is a free data retrieval call binding the contract method 0x28657aa5.
//
// Solidity: function bankBalance() view returns(uint256)
func (_TTBank *TTBankSession) BankBalance() (*big.Int, error) {
	return _TTBank.Contract.BankBalance(&_TTBank.CallOpts)
}

// BankBalance is a free data retrieval call binding the contract method 0x28657aa5.
//
// Solidity: function bankBalance() view returns(uint256)
func (_TTBank *TTBankCallerSession) BankBalance() (*big.Int, error) {
	return _TTBank.Contract.BankBalance(&_TTBank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TTBank *TTBankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TTBank.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TTBank *TTBankSession) Owner() (common.Address, error) {
	return _TTBank.Contract.Owner(&_TTBank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TTBank *TTBankCallerSession) Owner() (common.Address, error) {
	return _TTBank.Contract.Owner(&_TTBank.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TTBank *TTBankCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TTBank.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TTBank *TTBankSession) Token() (common.Address, error) {
	return _TTBank.Contract.Token(&_TTBank.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TTBank *TTBankCallerSession) Token() (common.Address, error) {
	return _TTBank.Contract.Token(&_TTBank.CallOpts)
}

// ViewAccount is a free data retrieval call binding the contract method 0x7f96939d.
//
// Solidity: function viewAccount() view returns((uint256,address,uint256))
func (_TTBank *TTBankCaller) ViewAccount(opts *bind.CallOpts) (ITTBankBankDetails, error) {
	var out []interface{}
	err := _TTBank.contract.Call(opts, &out, "viewAccount")

	if err != nil {
		return *new(ITTBankBankDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(ITTBankBankDetails)).(*ITTBankBankDetails)

	return out0, err

}

// ViewAccount is a free data retrieval call binding the contract method 0x7f96939d.
//
// Solidity: function viewAccount() view returns((uint256,address,uint256))
func (_TTBank *TTBankSession) ViewAccount() (ITTBankBankDetails, error) {
	return _TTBank.Contract.ViewAccount(&_TTBank.CallOpts)
}

// ViewAccount is a free data retrieval call binding the contract method 0x7f96939d.
//
// Solidity: function viewAccount() view returns((uint256,address,uint256))
func (_TTBank *TTBankCallerSession) ViewAccount() (ITTBankBankDetails, error) {
	return _TTBank.Contract.ViewAccount(&_TTBank.CallOpts)
}

// ViewBalance is a free data retrieval call binding the contract method 0x3ff1e05b.
//
// Solidity: function viewBalance() view returns(uint256)
func (_TTBank *TTBankCaller) ViewBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TTBank.contract.Call(opts, &out, "viewBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewBalance is a free data retrieval call binding the contract method 0x3ff1e05b.
//
// Solidity: function viewBalance() view returns(uint256)
func (_TTBank *TTBankSession) ViewBalance() (*big.Int, error) {
	return _TTBank.Contract.ViewBalance(&_TTBank.CallOpts)
}

// ViewBalance is a free data retrieval call binding the contract method 0x3ff1e05b.
//
// Solidity: function viewBalance() view returns(uint256)
func (_TTBank *TTBankCallerSession) ViewBalance() (*big.Int, error) {
	return _TTBank.Contract.ViewBalance(&_TTBank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(bool)
func (_TTBank *TTBankTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TTBank.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(bool)
func (_TTBank *TTBankSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TTBank.Contract.Deposit(&_TTBank.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(bool)
func (_TTBank *TTBankTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TTBank.Contract.Deposit(&_TTBank.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_TTBank *TTBankTransactor) Initialize(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TTBank.contract.Transact(opts, "initialize", tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_TTBank *TTBankSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _TTBank.Contract.Initialize(&_TTBank.TransactOpts, tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_TTBank *TTBankTransactorSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _TTBank.Contract.Initialize(&_TTBank.TransactOpts, tokenAddress)
}

// OpenAccount is a paid mutator transaction binding the contract method 0x1000420d.
//
// Solidity: function openAccount(uint256 startingBalance) returns(bool)
func (_TTBank *TTBankTransactor) OpenAccount(opts *bind.TransactOpts, startingBalance *big.Int) (*types.Transaction, error) {
	return _TTBank.contract.Transact(opts, "openAccount", startingBalance)
}

// OpenAccount is a paid mutator transaction binding the contract method 0x1000420d.
//
// Solidity: function openAccount(uint256 startingBalance) returns(bool)
func (_TTBank *TTBankSession) OpenAccount(startingBalance *big.Int) (*types.Transaction, error) {
	return _TTBank.Contract.OpenAccount(&_TTBank.TransactOpts, startingBalance)
}

// OpenAccount is a paid mutator transaction binding the contract method 0x1000420d.
//
// Solidity: function openAccount(uint256 startingBalance) returns(bool)
func (_TTBank *TTBankTransactorSession) OpenAccount(startingBalance *big.Int) (*types.Transaction, error) {
	return _TTBank.Contract.OpenAccount(&_TTBank.TransactOpts, startingBalance)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TTBank *TTBankTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TTBank.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TTBank *TTBankSession) RenounceOwnership() (*types.Transaction, error) {
	return _TTBank.Contract.RenounceOwnership(&_TTBank.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TTBank *TTBankTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TTBank.Contract.RenounceOwnership(&_TTBank.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TTBank *TTBankTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TTBank.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TTBank *TTBankSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TTBank.Contract.TransferOwnership(&_TTBank.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TTBank *TTBankTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TTBank.Contract.TransferOwnership(&_TTBank.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(bool)
func (_TTBank *TTBankTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TTBank.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(bool)
func (_TTBank *TTBankSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _TTBank.Contract.Withdraw(&_TTBank.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(bool)
func (_TTBank *TTBankTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _TTBank.Contract.Withdraw(&_TTBank.TransactOpts, amount)
}

// TTBankAccountOpenedIterator is returned from FilterAccountOpened and is used to iterate over the raw logs and unpacked data for AccountOpened events raised by the TTBank contract.
type TTBankAccountOpenedIterator struct {
	Event *TTBankAccountOpened // Event containing the contract specifics and raw log

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
func (it *TTBankAccountOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTBankAccountOpened)
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
		it.Event = new(TTBankAccountOpened)
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
func (it *TTBankAccountOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTBankAccountOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTBankAccountOpened represents a AccountOpened event raised by the TTBank contract.
type TTBankAccountOpened struct {
	AccountNumber   *big.Int
	AccountName     common.Address
	StartingBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountOpened is a free log retrieval operation binding the contract event 0x3dcb1c086601de03a591a96777a894297cadcf7af2b0416cceedff1ec49bc7c7.
//
// Solidity: event AccountOpened(uint256 accountNumber, address accountName, uint256 startingBalance)
func (_TTBank *TTBankFilterer) FilterAccountOpened(opts *bind.FilterOpts) (*TTBankAccountOpenedIterator, error) {

	logs, sub, err := _TTBank.contract.FilterLogs(opts, "AccountOpened")
	if err != nil {
		return nil, err
	}
	return &TTBankAccountOpenedIterator{contract: _TTBank.contract, event: "AccountOpened", logs: logs, sub: sub}, nil
}

// WatchAccountOpened is a free log subscription operation binding the contract event 0x3dcb1c086601de03a591a96777a894297cadcf7af2b0416cceedff1ec49bc7c7.
//
// Solidity: event AccountOpened(uint256 accountNumber, address accountName, uint256 startingBalance)
func (_TTBank *TTBankFilterer) WatchAccountOpened(opts *bind.WatchOpts, sink chan<- *TTBankAccountOpened) (event.Subscription, error) {

	logs, sub, err := _TTBank.contract.WatchLogs(opts, "AccountOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTBankAccountOpened)
				if err := _TTBank.contract.UnpackLog(event, "AccountOpened", log); err != nil {
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

// ParseAccountOpened is a log parse operation binding the contract event 0x3dcb1c086601de03a591a96777a894297cadcf7af2b0416cceedff1ec49bc7c7.
//
// Solidity: event AccountOpened(uint256 accountNumber, address accountName, uint256 startingBalance)
func (_TTBank *TTBankFilterer) ParseAccountOpened(log types.Log) (*TTBankAccountOpened, error) {
	event := new(TTBankAccountOpened)
	if err := _TTBank.contract.UnpackLog(event, "AccountOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TTBankDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the TTBank contract.
type TTBankDepositIterator struct {
	Event *TTBankDeposit // Event containing the contract specifics and raw log

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
func (it *TTBankDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTBankDeposit)
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
		it.Event = new(TTBankDeposit)
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
func (it *TTBankDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTBankDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTBankDeposit represents a Deposit event raised by the TTBank contract.
type TTBankDeposit struct {
	AccountNumber *big.Int
	AccountName   common.Address
	DepositAmount *big.Int
	NewBalance    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd36a2f67d06d285786f61a32b052b9ace6b0b7abef5177b54358abdc83a0b69b.
//
// Solidity: event Deposit(uint256 accountNumber, address accountName, uint256 depositAmount, uint256 newBalance)
func (_TTBank *TTBankFilterer) FilterDeposit(opts *bind.FilterOpts) (*TTBankDepositIterator, error) {

	logs, sub, err := _TTBank.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &TTBankDepositIterator{contract: _TTBank.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd36a2f67d06d285786f61a32b052b9ace6b0b7abef5177b54358abdc83a0b69b.
//
// Solidity: event Deposit(uint256 accountNumber, address accountName, uint256 depositAmount, uint256 newBalance)
func (_TTBank *TTBankFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TTBankDeposit) (event.Subscription, error) {

	logs, sub, err := _TTBank.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTBankDeposit)
				if err := _TTBank.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xd36a2f67d06d285786f61a32b052b9ace6b0b7abef5177b54358abdc83a0b69b.
//
// Solidity: event Deposit(uint256 accountNumber, address accountName, uint256 depositAmount, uint256 newBalance)
func (_TTBank *TTBankFilterer) ParseDeposit(log types.Log) (*TTBankDeposit, error) {
	event := new(TTBankDeposit)
	if err := _TTBank.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TTBankOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TTBank contract.
type TTBankOwnershipTransferredIterator struct {
	Event *TTBankOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TTBankOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTBankOwnershipTransferred)
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
		it.Event = new(TTBankOwnershipTransferred)
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
func (it *TTBankOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTBankOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTBankOwnershipTransferred represents a OwnershipTransferred event raised by the TTBank contract.
type TTBankOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TTBank *TTBankFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TTBankOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TTBank.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TTBankOwnershipTransferredIterator{contract: _TTBank.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TTBank *TTBankFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TTBankOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TTBank.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTBankOwnershipTransferred)
				if err := _TTBank.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TTBank *TTBankFilterer) ParseOwnershipTransferred(log types.Log) (*TTBankOwnershipTransferred, error) {
	event := new(TTBankOwnershipTransferred)
	if err := _TTBank.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TTBankWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TTBank contract.
type TTBankWithdrawIterator struct {
	Event *TTBankWithdraw // Event containing the contract specifics and raw log

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
func (it *TTBankWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTBankWithdraw)
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
		it.Event = new(TTBankWithdraw)
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
func (it *TTBankWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTBankWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTBankWithdraw represents a Withdraw event raised by the TTBank contract.
type TTBankWithdraw struct {
	AccountNumber  *big.Int
	AccountName    common.Address
	WithdrawAmount *big.Int
	NewBalance     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xb0ecf14e184effded5473bba77dcfab32b094b77ac1fbb36beec2aef55587970.
//
// Solidity: event Withdraw(uint256 accountNumber, address accountName, uint256 withdrawAmount, uint256 newBalance)
func (_TTBank *TTBankFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TTBankWithdrawIterator, error) {

	logs, sub, err := _TTBank.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TTBankWithdrawIterator{contract: _TTBank.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xb0ecf14e184effded5473bba77dcfab32b094b77ac1fbb36beec2aef55587970.
//
// Solidity: event Withdraw(uint256 accountNumber, address accountName, uint256 withdrawAmount, uint256 newBalance)
func (_TTBank *TTBankFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TTBankWithdraw) (event.Subscription, error) {

	logs, sub, err := _TTBank.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTBankWithdraw)
				if err := _TTBank.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xb0ecf14e184effded5473bba77dcfab32b094b77ac1fbb36beec2aef55587970.
//
// Solidity: event Withdraw(uint256 accountNumber, address accountName, uint256 withdrawAmount, uint256 newBalance)
func (_TTBank *TTBankFilterer) ParseWithdraw(log types.Log) (*TTBankWithdraw, error) {
	event := new(TTBankWithdraw)
	if err := _TTBank.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
