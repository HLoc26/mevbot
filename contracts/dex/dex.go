// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dex

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

// DexMetaData contains all meta data concerning the Dex contract.
var DexMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token2\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxAmountIn1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAmountIn2\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountToReceive\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createLiquidityPool\",\"inputs\":[{\"name\":\"token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token2\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feePercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSwapPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"liquidityPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidityPools\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeLiquidity\",\"inputs\":[{\"name\":\"token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token2\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountToReceive1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountToReceive2\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapExactInput\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapExactOutput\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxAmountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token1\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token2\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount2\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateLiquidityPool\",\"inputs\":[{\"name\":\"token1\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token2\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"liquidityPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token1\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token2\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount2\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// DexABI is the input ABI used to generate the binding from.
// Deprecated: Use DexMetaData.ABI instead.
var DexABI = DexMetaData.ABI

// Dex is an auto generated Go binding around an Ethereum contract.
type Dex struct {
	DexCaller     // Read-only binding to the contract
	DexTransactor // Write-only binding to the contract
	DexFilterer   // Log filterer for contract events
}

// DexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSession struct {
	Contract     *Dex              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexCallerSession struct {
	Contract *DexCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexTransactorSession struct {
	Contract     *DexTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexRaw struct {
	Contract *Dex // Generic contract binding to access the raw methods on
}

// DexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexCallerRaw struct {
	Contract *DexCaller // Generic read-only contract binding to access the raw methods on
}

// DexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexTransactorRaw struct {
	Contract *DexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDex creates a new instance of Dex, bound to a specific deployed contract.
func NewDex(address common.Address, backend bind.ContractBackend) (*Dex, error) {
	contract, err := bindDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// NewDexCaller creates a new read-only instance of Dex, bound to a specific deployed contract.
func NewDexCaller(address common.Address, caller bind.ContractCaller) (*DexCaller, error) {
	contract, err := bindDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexCaller{contract: contract}, nil
}

// NewDexTransactor creates a new write-only instance of Dex, bound to a specific deployed contract.
func NewDexTransactor(address common.Address, transactor bind.ContractTransactor) (*DexTransactor, error) {
	contract, err := bindDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexTransactor{contract: contract}, nil
}

// NewDexFilterer creates a new log filterer instance of Dex, bound to a specific deployed contract.
func NewDexFilterer(address common.Address, filterer bind.ContractFilterer) (*DexFilterer, error) {
	contract, err := bindDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexFilterer{contract: contract}, nil
}

// bindDex binds a generic wrapper to an already deployed contract.
func bindDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.DexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transact(opts, method, params...)
}

// LiquidityPools is a free data retrieval call binding the contract method 0x565841d1.
//
// Solidity: function liquidityPools(address , address ) view returns(address)
func (_Dex *DexCaller) LiquidityPools(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dex.contract.Call(opts, &out, "liquidityPools", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPools is a free data retrieval call binding the contract method 0x565841d1.
//
// Solidity: function liquidityPools(address , address ) view returns(address)
func (_Dex *DexSession) LiquidityPools(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Dex.Contract.LiquidityPools(&_Dex.CallOpts, arg0, arg1)
}

// LiquidityPools is a free data retrieval call binding the contract method 0x565841d1.
//
// Solidity: function liquidityPools(address , address ) view returns(address)
func (_Dex *DexCallerSession) LiquidityPools(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Dex.Contract.LiquidityPools(&_Dex.CallOpts, arg0, arg1)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x91c98a2a.
//
// Solidity: function addLiquidity(address token1, address token2, uint256 maxAmountIn1, uint256 maxAmountIn2, uint256 minAmountToReceive) returns()
func (_Dex *DexTransactor) AddLiquidity(opts *bind.TransactOpts, token1 common.Address, token2 common.Address, maxAmountIn1 *big.Int, maxAmountIn2 *big.Int, minAmountToReceive *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "addLiquidity", token1, token2, maxAmountIn1, maxAmountIn2, minAmountToReceive)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x91c98a2a.
//
// Solidity: function addLiquidity(address token1, address token2, uint256 maxAmountIn1, uint256 maxAmountIn2, uint256 minAmountToReceive) returns()
func (_Dex *DexSession) AddLiquidity(token1 common.Address, token2 common.Address, maxAmountIn1 *big.Int, maxAmountIn2 *big.Int, minAmountToReceive *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.AddLiquidity(&_Dex.TransactOpts, token1, token2, maxAmountIn1, maxAmountIn2, minAmountToReceive)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x91c98a2a.
//
// Solidity: function addLiquidity(address token1, address token2, uint256 maxAmountIn1, uint256 maxAmountIn2, uint256 minAmountToReceive) returns()
func (_Dex *DexTransactorSession) AddLiquidity(token1 common.Address, token2 common.Address, maxAmountIn1 *big.Int, maxAmountIn2 *big.Int, minAmountToReceive *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.AddLiquidity(&_Dex.TransactOpts, token1, token2, maxAmountIn1, maxAmountIn2, minAmountToReceive)
}

// CreateLiquidityPool is a paid mutator transaction binding the contract method 0xdc6b6d8a.
//
// Solidity: function createLiquidityPool(address token1, address token2, uint256 feePercentage, uint256 maxSwapPercentage, string name, string symbol) returns(address liquidityPool)
func (_Dex *DexTransactor) CreateLiquidityPool(opts *bind.TransactOpts, token1 common.Address, token2 common.Address, feePercentage *big.Int, maxSwapPercentage *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "createLiquidityPool", token1, token2, feePercentage, maxSwapPercentage, name, symbol)
}

// CreateLiquidityPool is a paid mutator transaction binding the contract method 0xdc6b6d8a.
//
// Solidity: function createLiquidityPool(address token1, address token2, uint256 feePercentage, uint256 maxSwapPercentage, string name, string symbol) returns(address liquidityPool)
func (_Dex *DexSession) CreateLiquidityPool(token1 common.Address, token2 common.Address, feePercentage *big.Int, maxSwapPercentage *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Dex.Contract.CreateLiquidityPool(&_Dex.TransactOpts, token1, token2, feePercentage, maxSwapPercentage, name, symbol)
}

// CreateLiquidityPool is a paid mutator transaction binding the contract method 0xdc6b6d8a.
//
// Solidity: function createLiquidityPool(address token1, address token2, uint256 feePercentage, uint256 maxSwapPercentage, string name, string symbol) returns(address liquidityPool)
func (_Dex *DexTransactorSession) CreateLiquidityPool(token1 common.Address, token2 common.Address, feePercentage *big.Int, maxSwapPercentage *big.Int, name string, symbol string) (*types.Transaction, error) {
	return _Dex.Contract.CreateLiquidityPool(&_Dex.TransactOpts, token1, token2, feePercentage, maxSwapPercentage, name, symbol)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address token1, address token2, uint256 burnAmount, uint256 minAmountToReceive1, uint256 minAmountToReceive2) returns()
func (_Dex *DexTransactor) RemoveLiquidity(opts *bind.TransactOpts, token1 common.Address, token2 common.Address, burnAmount *big.Int, minAmountToReceive1 *big.Int, minAmountToReceive2 *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "removeLiquidity", token1, token2, burnAmount, minAmountToReceive1, minAmountToReceive2)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address token1, address token2, uint256 burnAmount, uint256 minAmountToReceive1, uint256 minAmountToReceive2) returns()
func (_Dex *DexSession) RemoveLiquidity(token1 common.Address, token2 common.Address, burnAmount *big.Int, minAmountToReceive1 *big.Int, minAmountToReceive2 *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.RemoveLiquidity(&_Dex.TransactOpts, token1, token2, burnAmount, minAmountToReceive1, minAmountToReceive2)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address token1, address token2, uint256 burnAmount, uint256 minAmountToReceive1, uint256 minAmountToReceive2) returns()
func (_Dex *DexTransactorSession) RemoveLiquidity(token1 common.Address, token2 common.Address, burnAmount *big.Int, minAmountToReceive1 *big.Int, minAmountToReceive2 *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.RemoveLiquidity(&_Dex.TransactOpts, token1, token2, burnAmount, minAmountToReceive1, minAmountToReceive2)
}

// SwapExactInput is a paid mutator transaction binding the contract method 0xbd028142.
//
// Solidity: function swapExactInput(address tokenIn, address tokenOut, uint256 amountIn, uint256 minAmountOut) returns()
func (_Dex *DexTransactor) SwapExactInput(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "swapExactInput", tokenIn, tokenOut, amountIn, minAmountOut)
}

// SwapExactInput is a paid mutator transaction binding the contract method 0xbd028142.
//
// Solidity: function swapExactInput(address tokenIn, address tokenOut, uint256 amountIn, uint256 minAmountOut) returns()
func (_Dex *DexSession) SwapExactInput(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.SwapExactInput(&_Dex.TransactOpts, tokenIn, tokenOut, amountIn, minAmountOut)
}

// SwapExactInput is a paid mutator transaction binding the contract method 0xbd028142.
//
// Solidity: function swapExactInput(address tokenIn, address tokenOut, uint256 amountIn, uint256 minAmountOut) returns()
func (_Dex *DexTransactorSession) SwapExactInput(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.SwapExactInput(&_Dex.TransactOpts, tokenIn, tokenOut, amountIn, minAmountOut)
}

// SwapExactOutput is a paid mutator transaction binding the contract method 0xe09dbe1e.
//
// Solidity: function swapExactOutput(address tokenIn, address tokenOut, uint256 maxAmountIn, uint256 amountOut) returns()
func (_Dex *DexTransactor) SwapExactOutput(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, maxAmountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _Dex.contract.Transact(opts, "swapExactOutput", tokenIn, tokenOut, maxAmountIn, amountOut)
}

// SwapExactOutput is a paid mutator transaction binding the contract method 0xe09dbe1e.
//
// Solidity: function swapExactOutput(address tokenIn, address tokenOut, uint256 maxAmountIn, uint256 amountOut) returns()
func (_Dex *DexSession) SwapExactOutput(tokenIn common.Address, tokenOut common.Address, maxAmountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.SwapExactOutput(&_Dex.TransactOpts, tokenIn, tokenOut, maxAmountIn, amountOut)
}

// SwapExactOutput is a paid mutator transaction binding the contract method 0xe09dbe1e.
//
// Solidity: function swapExactOutput(address tokenIn, address tokenOut, uint256 maxAmountIn, uint256 amountOut) returns()
func (_Dex *DexTransactorSession) SwapExactOutput(tokenIn common.Address, tokenOut common.Address, maxAmountIn *big.Int, amountOut *big.Int) (*types.Transaction, error) {
	return _Dex.Contract.SwapExactOutput(&_Dex.TransactOpts, tokenIn, tokenOut, maxAmountIn, amountOut)
}

// DexAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Dex contract.
type DexAddLiquidityIterator struct {
	Event *DexAddLiquidity // Event containing the contract specifics and raw log

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
func (it *DexAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexAddLiquidity)
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
		it.Event = new(DexAddLiquidity)
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
func (it *DexAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexAddLiquidity represents a AddLiquidity event raised by the Dex contract.
type DexAddLiquidity struct {
	Sender     common.Address
	Token1     common.Address
	Token2     common.Address
	Amount1    *big.Int
	Amount2    *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x9c899bc9e3ccdb62eed42f81290bab167a8d211c543af4e6220e364d65fb6960.
//
// Solidity: event AddLiquidity(address sender, address token1, address token2, uint256 amount1, uint256 amount2, uint256 mintAmount)
func (_Dex *DexFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*DexAddLiquidityIterator, error) {

	logs, sub, err := _Dex.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &DexAddLiquidityIterator{contract: _Dex.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x9c899bc9e3ccdb62eed42f81290bab167a8d211c543af4e6220e364d65fb6960.
//
// Solidity: event AddLiquidity(address sender, address token1, address token2, uint256 amount1, uint256 amount2, uint256 mintAmount)
func (_Dex *DexFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *DexAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _Dex.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexAddLiquidity)
				if err := _Dex.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x9c899bc9e3ccdb62eed42f81290bab167a8d211c543af4e6220e364d65fb6960.
//
// Solidity: event AddLiquidity(address sender, address token1, address token2, uint256 amount1, uint256 amount2, uint256 mintAmount)
func (_Dex *DexFilterer) ParseAddLiquidity(log types.Log) (*DexAddLiquidity, error) {
	event := new(DexAddLiquidity)
	if err := _Dex.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexCreateLiquidityPoolIterator is returned from FilterCreateLiquidityPool and is used to iterate over the raw logs and unpacked data for CreateLiquidityPool events raised by the Dex contract.
type DexCreateLiquidityPoolIterator struct {
	Event *DexCreateLiquidityPool // Event containing the contract specifics and raw log

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
func (it *DexCreateLiquidityPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexCreateLiquidityPool)
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
		it.Event = new(DexCreateLiquidityPool)
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
func (it *DexCreateLiquidityPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexCreateLiquidityPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexCreateLiquidityPool represents a CreateLiquidityPool event raised by the Dex contract.
type DexCreateLiquidityPool struct {
	Token1        common.Address
	Token2        common.Address
	LiquidityPool common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreateLiquidityPool is a free log retrieval operation binding the contract event 0x878fb67339982d5d07dd6ab0e525620d9a3fdefec22e92fce9281683a21b1454.
//
// Solidity: event CreateLiquidityPool(address token1, address token2, address liquidityPool)
func (_Dex *DexFilterer) FilterCreateLiquidityPool(opts *bind.FilterOpts) (*DexCreateLiquidityPoolIterator, error) {

	logs, sub, err := _Dex.contract.FilterLogs(opts, "CreateLiquidityPool")
	if err != nil {
		return nil, err
	}
	return &DexCreateLiquidityPoolIterator{contract: _Dex.contract, event: "CreateLiquidityPool", logs: logs, sub: sub}, nil
}

// WatchCreateLiquidityPool is a free log subscription operation binding the contract event 0x878fb67339982d5d07dd6ab0e525620d9a3fdefec22e92fce9281683a21b1454.
//
// Solidity: event CreateLiquidityPool(address token1, address token2, address liquidityPool)
func (_Dex *DexFilterer) WatchCreateLiquidityPool(opts *bind.WatchOpts, sink chan<- *DexCreateLiquidityPool) (event.Subscription, error) {

	logs, sub, err := _Dex.contract.WatchLogs(opts, "CreateLiquidityPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexCreateLiquidityPool)
				if err := _Dex.contract.UnpackLog(event, "CreateLiquidityPool", log); err != nil {
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

// ParseCreateLiquidityPool is a log parse operation binding the contract event 0x878fb67339982d5d07dd6ab0e525620d9a3fdefec22e92fce9281683a21b1454.
//
// Solidity: event CreateLiquidityPool(address token1, address token2, address liquidityPool)
func (_Dex *DexFilterer) ParseCreateLiquidityPool(log types.Log) (*DexCreateLiquidityPool, error) {
	event := new(DexCreateLiquidityPool)
	if err := _Dex.contract.UnpackLog(event, "CreateLiquidityPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Dex contract.
type DexRemoveLiquidityIterator struct {
	Event *DexRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *DexRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexRemoveLiquidity)
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
		it.Event = new(DexRemoveLiquidity)
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
func (it *DexRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexRemoveLiquidity represents a RemoveLiquidity event raised by the Dex contract.
type DexRemoveLiquidity struct {
	Sender     common.Address
	Token1     common.Address
	Token2     common.Address
	Amount1    *big.Int
	Amount2    *big.Int
	BurnAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x094577f52c15f75d98e8b7e35fd170e61ac73ad9fb3135eae6674b39828d4d6e.
//
// Solidity: event RemoveLiquidity(address sender, address token1, address token2, uint256 amount1, uint256 amount2, uint256 burnAmount)
func (_Dex *DexFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*DexRemoveLiquidityIterator, error) {

	logs, sub, err := _Dex.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &DexRemoveLiquidityIterator{contract: _Dex.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x094577f52c15f75d98e8b7e35fd170e61ac73ad9fb3135eae6674b39828d4d6e.
//
// Solidity: event RemoveLiquidity(address sender, address token1, address token2, uint256 amount1, uint256 amount2, uint256 burnAmount)
func (_Dex *DexFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *DexRemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _Dex.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexRemoveLiquidity)
				if err := _Dex.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x094577f52c15f75d98e8b7e35fd170e61ac73ad9fb3135eae6674b39828d4d6e.
//
// Solidity: event RemoveLiquidity(address sender, address token1, address token2, uint256 amount1, uint256 amount2, uint256 burnAmount)
func (_Dex *DexFilterer) ParseRemoveLiquidity(log types.Log) (*DexRemoveLiquidity, error) {
	event := new(DexRemoveLiquidity)
	if err := _Dex.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Dex contract.
type DexSwapIterator struct {
	Event *DexSwap // Event containing the contract specifics and raw log

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
func (it *DexSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexSwap)
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
		it.Event = new(DexSwap)
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
func (it *DexSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexSwap represents a Swap event raised by the Dex contract.
type DexSwap struct {
	Sender    common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address sender, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Dex *DexFilterer) FilterSwap(opts *bind.FilterOpts) (*DexSwapIterator, error) {

	logs, sub, err := _Dex.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &DexSwapIterator{contract: _Dex.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address sender, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Dex *DexFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *DexSwap) (event.Subscription, error) {

	logs, sub, err := _Dex.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexSwap)
				if err := _Dex.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address sender, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Dex *DexFilterer) ParseSwap(log types.Log) (*DexSwap, error) {
	event := new(DexSwap)
	if err := _Dex.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
