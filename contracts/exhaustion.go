// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ledgerwatch/turbo-geth"
	"github.com/ledgerwatch/turbo-geth/accounts/abi"
	"github.com/ledgerwatch/turbo-geth/accounts/abi/bind"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/core/types"
	"github.com/ledgerwatch/turbo-geth/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExhaustionAttackABI is the input ABI used to generate the binding from.
const ExhaustionAttackABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"loop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExhaustionAttackFuncSigs maps the 4-byte function signature to its string representation.
var ExhaustionAttackFuncSigs = map[string]string{
	"a92100cb": "loop()",
}

// ExhaustionAttackBin is the compiled bytecode used for deploying new contracts.
var ExhaustionAttackBin = "0x6080604052348015600f57600080fd5b5060898061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063a92100cb14602d575b600080fd5b60336035565b005b60005a90505b61c350811115605157601f80601f375a9050603b565b5056fea265627a7a7231582061d3abd1882697c26052a3e863ca13fc7b9ddadab3e01c61856c8027cf8edddf64736f6c63430005100032"

// DeployExhaustionAttack deploys a new Ethereum contract, binding an instance of ExhaustionAttack to it.
func DeployExhaustionAttack(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExhaustionAttack, error) {
	parsed, err := abi.JSON(strings.NewReader(ExhaustionAttackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExhaustionAttackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExhaustionAttack{ExhaustionAttackCaller: ExhaustionAttackCaller{contract: contract}, ExhaustionAttackTransactor: ExhaustionAttackTransactor{contract: contract}, ExhaustionAttackFilterer: ExhaustionAttackFilterer{contract: contract}}, nil
}

// ExhaustionAttack is an auto generated Go binding around an Ethereum contract.
type ExhaustionAttack struct {
	ExhaustionAttackCaller     // Read-only binding to the contract
	ExhaustionAttackTransactor // Write-only binding to the contract
	ExhaustionAttackFilterer   // Log filterer for contract events
}

// ExhaustionAttackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExhaustionAttackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExhaustionAttackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExhaustionAttackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExhaustionAttackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExhaustionAttackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExhaustionAttackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExhaustionAttackSession struct {
	Contract     *ExhaustionAttack // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExhaustionAttackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExhaustionAttackCallerSession struct {
	Contract *ExhaustionAttackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExhaustionAttackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExhaustionAttackTransactorSession struct {
	Contract     *ExhaustionAttackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExhaustionAttackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExhaustionAttackRaw struct {
	Contract *ExhaustionAttack // Generic contract binding to access the raw methods on
}

// ExhaustionAttackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExhaustionAttackCallerRaw struct {
	Contract *ExhaustionAttackCaller // Generic read-only contract binding to access the raw methods on
}

// ExhaustionAttackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExhaustionAttackTransactorRaw struct {
	Contract *ExhaustionAttackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExhaustionAttack creates a new instance of ExhaustionAttack, bound to a specific deployed contract.
func NewExhaustionAttack(address common.Address, backend bind.ContractBackend) (*ExhaustionAttack, error) {
	contract, err := bindExhaustionAttack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExhaustionAttack{ExhaustionAttackCaller: ExhaustionAttackCaller{contract: contract}, ExhaustionAttackTransactor: ExhaustionAttackTransactor{contract: contract}, ExhaustionAttackFilterer: ExhaustionAttackFilterer{contract: contract}}, nil
}

// NewExhaustionAttackCaller creates a new read-only instance of ExhaustionAttack, bound to a specific deployed contract.
func NewExhaustionAttackCaller(address common.Address, caller bind.ContractCaller) (*ExhaustionAttackCaller, error) {
	contract, err := bindExhaustionAttack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExhaustionAttackCaller{contract: contract}, nil
}

// NewExhaustionAttackTransactor creates a new write-only instance of ExhaustionAttack, bound to a specific deployed contract.
func NewExhaustionAttackTransactor(address common.Address, transactor bind.ContractTransactor) (*ExhaustionAttackTransactor, error) {
	contract, err := bindExhaustionAttack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExhaustionAttackTransactor{contract: contract}, nil
}

// NewExhaustionAttackFilterer creates a new log filterer instance of ExhaustionAttack, bound to a specific deployed contract.
func NewExhaustionAttackFilterer(address common.Address, filterer bind.ContractFilterer) (*ExhaustionAttackFilterer, error) {
	contract, err := bindExhaustionAttack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExhaustionAttackFilterer{contract: contract}, nil
}

// bindExhaustionAttack binds a generic wrapper to an already deployed contract.
func bindExhaustionAttack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExhaustionAttackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExhaustionAttack *ExhaustionAttackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExhaustionAttack.Contract.ExhaustionAttackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExhaustionAttack *ExhaustionAttackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExhaustionAttack.Contract.ExhaustionAttackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExhaustionAttack *ExhaustionAttackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExhaustionAttack.Contract.ExhaustionAttackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExhaustionAttack *ExhaustionAttackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExhaustionAttack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExhaustionAttack *ExhaustionAttackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExhaustionAttack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExhaustionAttack *ExhaustionAttackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExhaustionAttack.Contract.contract.Transact(opts, method, params...)
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns()
func (_ExhaustionAttack *ExhaustionAttackTransactor) Loop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExhaustionAttack.contract.Transact(opts, "loop")
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns()
func (_ExhaustionAttack *ExhaustionAttackSession) Loop() (*types.Transaction, error) {
	return _ExhaustionAttack.Contract.Loop(&_ExhaustionAttack.TransactOpts)
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns()
func (_ExhaustionAttack *ExhaustionAttackTransactorSession) Loop() (*types.Transaction, error) {
	return _ExhaustionAttack.Contract.Loop(&_ExhaustionAttack.TransactOpts)
}
