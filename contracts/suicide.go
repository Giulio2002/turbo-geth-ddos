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

// SuicideABI is the input ABI used to generate the binding from.
const SuicideABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SuicideFuncSigs maps the 4-byte function signature to its string representation.
var SuicideFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// SuicideBin is the compiled bytecode used for deploying new contracts.
var SuicideBin = "0x6080604052348015600f57600080fd5b50606e80601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336035565b005b6000fffea265627a7a72315820e09b216a1fdc7d1a6a049e09b07068ce6e03a463233b7a0409881767db07ffa964736f6c63430005100032"

// DeploySuicide deploys a new Ethereum contract, binding an instance of Suicide to it.
func DeploySuicide(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Suicide, error) {
	parsed, err := abi.JSON(strings.NewReader(SuicideABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SuicideBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Suicide{SuicideCaller: SuicideCaller{contract: contract}, SuicideTransactor: SuicideTransactor{contract: contract}, SuicideFilterer: SuicideFilterer{contract: contract}}, nil
}

// Suicide is an auto generated Go binding around an Ethereum contract.
type Suicide struct {
	SuicideCaller     // Read-only binding to the contract
	SuicideTransactor // Write-only binding to the contract
	SuicideFilterer   // Log filterer for contract events
}

// SuicideCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuicideCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuicideTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuicideTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuicideFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuicideFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuicideSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuicideSession struct {
	Contract     *Suicide          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuicideCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuicideCallerSession struct {
	Contract *SuicideCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SuicideTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuicideTransactorSession struct {
	Contract     *SuicideTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SuicideRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuicideRaw struct {
	Contract *Suicide // Generic contract binding to access the raw methods on
}

// SuicideCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuicideCallerRaw struct {
	Contract *SuicideCaller // Generic read-only contract binding to access the raw methods on
}

// SuicideTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuicideTransactorRaw struct {
	Contract *SuicideTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuicide creates a new instance of Suicide, bound to a specific deployed contract.
func NewSuicide(address common.Address, backend bind.ContractBackend) (*Suicide, error) {
	contract, err := bindSuicide(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Suicide{SuicideCaller: SuicideCaller{contract: contract}, SuicideTransactor: SuicideTransactor{contract: contract}, SuicideFilterer: SuicideFilterer{contract: contract}}, nil
}

// NewSuicideCaller creates a new read-only instance of Suicide, bound to a specific deployed contract.
func NewSuicideCaller(address common.Address, caller bind.ContractCaller) (*SuicideCaller, error) {
	contract, err := bindSuicide(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuicideCaller{contract: contract}, nil
}

// NewSuicideTransactor creates a new write-only instance of Suicide, bound to a specific deployed contract.
func NewSuicideTransactor(address common.Address, transactor bind.ContractTransactor) (*SuicideTransactor, error) {
	contract, err := bindSuicide(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuicideTransactor{contract: contract}, nil
}

// NewSuicideFilterer creates a new log filterer instance of Suicide, bound to a specific deployed contract.
func NewSuicideFilterer(address common.Address, filterer bind.ContractFilterer) (*SuicideFilterer, error) {
	contract, err := bindSuicide(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuicideFilterer{contract: contract}, nil
}

// bindSuicide binds a generic wrapper to an already deployed contract.
func bindSuicide(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuicideABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Suicide *SuicideRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Suicide.Contract.SuicideCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Suicide *SuicideRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suicide.Contract.SuicideTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Suicide *SuicideRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Suicide.Contract.SuicideTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Suicide *SuicideCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Suicide.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Suicide *SuicideTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suicide.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Suicide *SuicideTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Suicide.Contract.contract.Transact(opts, method, params...)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Suicide *SuicideTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Suicide.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Suicide *SuicideSession) Start() (*types.Transaction, error) {
	return _Suicide.Contract.Start(&_Suicide.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Suicide *SuicideTransactorSession) Start() (*types.Transaction, error) {
	return _Suicide.Contract.Start(&_Suicide.TransactOpts)
}
