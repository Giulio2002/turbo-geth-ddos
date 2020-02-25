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

// BalanceABI is the input ABI used to generate the binding from.
const BalanceABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalanceFuncSigs maps the 4-byte function signature to its string representation.
var BalanceFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// BalanceBin is the compiled bytecode used for deploying new contracts.
var BalanceBin = "0x6080604052348015600f57600080fd5b50606c80601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336035565b005b56fea265627a7a72315820eb768768b43510df596d2985b07085fdcd8e7f0d6537a17438f8f3a91710c45f64736f6c63430005100032"

// DeployBalance deploys a new Ethereum contract, binding an instance of Balance to it.
func DeployBalance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Balance, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BalanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Balance{BalanceCaller: BalanceCaller{contract: contract}, BalanceTransactor: BalanceTransactor{contract: contract}, BalanceFilterer: BalanceFilterer{contract: contract}}, nil
}

// Balance is an auto generated Go binding around an Ethereum contract.
type Balance struct {
	BalanceCaller     // Read-only binding to the contract
	BalanceTransactor // Write-only binding to the contract
	BalanceFilterer   // Log filterer for contract events
}

// BalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceSession struct {
	Contract     *Balance          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceCallerSession struct {
	Contract *BalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceTransactorSession struct {
	Contract     *BalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceRaw struct {
	Contract *Balance // Generic contract binding to access the raw methods on
}

// BalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceCallerRaw struct {
	Contract *BalanceCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceTransactorRaw struct {
	Contract *BalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalance creates a new instance of Balance, bound to a specific deployed contract.
func NewBalance(address common.Address, backend bind.ContractBackend) (*Balance, error) {
	contract, err := bindBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balance{BalanceCaller: BalanceCaller{contract: contract}, BalanceTransactor: BalanceTransactor{contract: contract}, BalanceFilterer: BalanceFilterer{contract: contract}}, nil
}

// NewBalanceCaller creates a new read-only instance of Balance, bound to a specific deployed contract.
func NewBalanceCaller(address common.Address, caller bind.ContractCaller) (*BalanceCaller, error) {
	contract, err := bindBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCaller{contract: contract}, nil
}

// NewBalanceTransactor creates a new write-only instance of Balance, bound to a specific deployed contract.
func NewBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceTransactor, error) {
	contract, err := bindBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceTransactor{contract: contract}, nil
}

// NewBalanceFilterer creates a new log filterer instance of Balance, bound to a specific deployed contract.
func NewBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceFilterer, error) {
	contract, err := bindBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceFilterer{contract: contract}, nil
}

// bindBalance binds a generic wrapper to an already deployed contract.
func bindBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balance *BalanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Balance.Contract.BalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balance *BalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.Contract.BalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balance *BalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balance.Contract.BalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balance *BalanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Balance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balance *BalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balance *BalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balance.Contract.contract.Transact(opts, method, params...)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Balance *BalanceTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Balance *BalanceSession) Start() (*types.Transaction, error) {
	return _Balance.Contract.Start(&_Balance.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Balance *BalanceTransactorSession) Start() (*types.Transaction, error) {
	return _Balance.Contract.Start(&_Balance.TransactOpts)
}

// BalanceCallABI is the input ABI used to generate the binding from.
const BalanceCallABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BalanceCallFuncSigs maps the 4-byte function signature to its string representation.
var BalanceCallFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// BalanceCallBin is the compiled bytecode used for deploying new contracts.
var BalanceCallBin = "0x6080604052348015600f57600080fd5b50606c80601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336035565b005b56fea265627a7a7231582035ad3cff2770fa925e6c9778a3d88fa05c1256eb1b2354286eab0dda26ce620964736f6c63430005100032"

// DeployBalanceCall deploys a new Ethereum contract, binding an instance of BalanceCall to it.
func DeployBalanceCall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceCall, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceCallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BalanceCallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceCall{BalanceCallCaller: BalanceCallCaller{contract: contract}, BalanceCallTransactor: BalanceCallTransactor{contract: contract}, BalanceCallFilterer: BalanceCallFilterer{contract: contract}}, nil
}

// BalanceCall is an auto generated Go binding around an Ethereum contract.
type BalanceCall struct {
	BalanceCallCaller     // Read-only binding to the contract
	BalanceCallTransactor // Write-only binding to the contract
	BalanceCallFilterer   // Log filterer for contract events
}

// BalanceCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceCallSession struct {
	Contract     *BalanceCall      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceCallCallerSession struct {
	Contract *BalanceCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BalanceCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceCallTransactorSession struct {
	Contract     *BalanceCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BalanceCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceCallRaw struct {
	Contract *BalanceCall // Generic contract binding to access the raw methods on
}

// BalanceCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceCallCallerRaw struct {
	Contract *BalanceCallCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceCallTransactorRaw struct {
	Contract *BalanceCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceCall creates a new instance of BalanceCall, bound to a specific deployed contract.
func NewBalanceCall(address common.Address, backend bind.ContractBackend) (*BalanceCall, error) {
	contract, err := bindBalanceCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceCall{BalanceCallCaller: BalanceCallCaller{contract: contract}, BalanceCallTransactor: BalanceCallTransactor{contract: contract}, BalanceCallFilterer: BalanceCallFilterer{contract: contract}}, nil
}

// NewBalanceCallCaller creates a new read-only instance of BalanceCall, bound to a specific deployed contract.
func NewBalanceCallCaller(address common.Address, caller bind.ContractCaller) (*BalanceCallCaller, error) {
	contract, err := bindBalanceCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCallCaller{contract: contract}, nil
}

// NewBalanceCallTransactor creates a new write-only instance of BalanceCall, bound to a specific deployed contract.
func NewBalanceCallTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceCallTransactor, error) {
	contract, err := bindBalanceCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCallTransactor{contract: contract}, nil
}

// NewBalanceCallFilterer creates a new log filterer instance of BalanceCall, bound to a specific deployed contract.
func NewBalanceCallFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceCallFilterer, error) {
	contract, err := bindBalanceCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceCallFilterer{contract: contract}, nil
}

// bindBalanceCall binds a generic wrapper to an already deployed contract.
func bindBalanceCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceCall *BalanceCallRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalanceCall.Contract.BalanceCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceCall *BalanceCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceCall.Contract.BalanceCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceCall *BalanceCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceCall.Contract.BalanceCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceCall *BalanceCallCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalanceCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceCall *BalanceCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceCall *BalanceCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceCall.Contract.contract.Transact(opts, method, params...)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns()
func (_BalanceCall *BalanceCallCaller) Start(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _BalanceCall.contract.Call(opts, out, "start")
	return err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns()
func (_BalanceCall *BalanceCallSession) Start() error {
	return _BalanceCall.Contract.Start(&_BalanceCall.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns()
func (_BalanceCall *BalanceCallCallerSession) Start() error {
	return _BalanceCall.Contract.Start(&_BalanceCall.CallOpts)
}

// ExhaustionAttackABI is the input ABI used to generate the binding from.
const ExhaustionAttackABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"loop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExhaustionAttackFuncSigs maps the 4-byte function signature to its string representation.
var ExhaustionAttackFuncSigs = map[string]string{
	"a92100cb": "loop()",
}

// ExhaustionAttackBin is the compiled bytecode used for deploying new contracts.
var ExhaustionAttackBin = "0x6080604052348015600f57600080fd5b5060898061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063a92100cb14602d575b600080fd5b60336035565b005b60005a90505b61c350811115605157601f80601f375a9050603b565b5056fea265627a7a72315820664256302ab449b8614eb2ded427f9572849363dee6684b00125b909294e9a8a64736f6c63430005100032"

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

// SloadABI is the input ABI used to generate the binding from.
const SloadABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SloadFuncSigs maps the 4-byte function signature to its string representation.
var SloadFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// SloadBin is the compiled bytecode used for deploying new contracts.
var SloadBin = "0x6080604052348015600f57600080fd5b50606c80601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336035565b005b56fea265627a7a723158203e7adf290e01847404d54f125fcdb2620874473a366ee504ec3d368a350afff064736f6c63430005100032"

// DeploySload deploys a new Ethereum contract, binding an instance of Sload to it.
func DeploySload(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sload, error) {
	parsed, err := abi.JSON(strings.NewReader(SloadABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SloadBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sload{SloadCaller: SloadCaller{contract: contract}, SloadTransactor: SloadTransactor{contract: contract}, SloadFilterer: SloadFilterer{contract: contract}}, nil
}

// Sload is an auto generated Go binding around an Ethereum contract.
type Sload struct {
	SloadCaller     // Read-only binding to the contract
	SloadTransactor // Write-only binding to the contract
	SloadFilterer   // Log filterer for contract events
}

// SloadCaller is an auto generated read-only Go binding around an Ethereum contract.
type SloadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SloadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SloadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SloadSession struct {
	Contract     *Sload            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SloadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SloadCallerSession struct {
	Contract *SloadCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SloadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SloadTransactorSession struct {
	Contract     *SloadTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SloadRaw is an auto generated low-level Go binding around an Ethereum contract.
type SloadRaw struct {
	Contract *Sload // Generic contract binding to access the raw methods on
}

// SloadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SloadCallerRaw struct {
	Contract *SloadCaller // Generic read-only contract binding to access the raw methods on
}

// SloadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SloadTransactorRaw struct {
	Contract *SloadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSload creates a new instance of Sload, bound to a specific deployed contract.
func NewSload(address common.Address, backend bind.ContractBackend) (*Sload, error) {
	contract, err := bindSload(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sload{SloadCaller: SloadCaller{contract: contract}, SloadTransactor: SloadTransactor{contract: contract}, SloadFilterer: SloadFilterer{contract: contract}}, nil
}

// NewSloadCaller creates a new read-only instance of Sload, bound to a specific deployed contract.
func NewSloadCaller(address common.Address, caller bind.ContractCaller) (*SloadCaller, error) {
	contract, err := bindSload(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SloadCaller{contract: contract}, nil
}

// NewSloadTransactor creates a new write-only instance of Sload, bound to a specific deployed contract.
func NewSloadTransactor(address common.Address, transactor bind.ContractTransactor) (*SloadTransactor, error) {
	contract, err := bindSload(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SloadTransactor{contract: contract}, nil
}

// NewSloadFilterer creates a new log filterer instance of Sload, bound to a specific deployed contract.
func NewSloadFilterer(address common.Address, filterer bind.ContractFilterer) (*SloadFilterer, error) {
	contract, err := bindSload(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SloadFilterer{contract: contract}, nil
}

// bindSload binds a generic wrapper to an already deployed contract.
func bindSload(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SloadABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sload *SloadRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sload.Contract.SloadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sload *SloadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.Contract.SloadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sload *SloadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sload.Contract.SloadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sload *SloadCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sload.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sload *SloadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sload *SloadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sload.Contract.contract.Transact(opts, method, params...)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Sload *SloadTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Sload *SloadSession) Start() (*types.Transaction, error) {
	return _Sload.Contract.Start(&_Sload.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Sload *SloadTransactorSession) Start() (*types.Transaction, error) {
	return _Sload.Contract.Start(&_Sload.TransactOpts)
}

// SloadCallABI is the input ABI used to generate the binding from.
const SloadCallABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SloadCallFuncSigs maps the 4-byte function signature to its string representation.
var SloadCallFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// SloadCallBin is the compiled bytecode used for deploying new contracts.
var SloadCallBin = "0x6080604052348015600f57600080fd5b50606c80601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336035565b005b56fea265627a7a723158201a550f1fd87a4f8b5bbda4ae8607b1d6328a8058b2687e8a39ef2f998c64a4a964736f6c63430005100032"

// DeploySloadCall deploys a new Ethereum contract, binding an instance of SloadCall to it.
func DeploySloadCall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SloadCall, error) {
	parsed, err := abi.JSON(strings.NewReader(SloadCallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SloadCallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SloadCall{SloadCallCaller: SloadCallCaller{contract: contract}, SloadCallTransactor: SloadCallTransactor{contract: contract}, SloadCallFilterer: SloadCallFilterer{contract: contract}}, nil
}

// SloadCall is an auto generated Go binding around an Ethereum contract.
type SloadCall struct {
	SloadCallCaller     // Read-only binding to the contract
	SloadCallTransactor // Write-only binding to the contract
	SloadCallFilterer   // Log filterer for contract events
}

// SloadCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type SloadCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SloadCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SloadCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SloadCallSession struct {
	Contract     *SloadCall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SloadCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SloadCallCallerSession struct {
	Contract *SloadCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SloadCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SloadCallTransactorSession struct {
	Contract     *SloadCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SloadCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type SloadCallRaw struct {
	Contract *SloadCall // Generic contract binding to access the raw methods on
}

// SloadCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SloadCallCallerRaw struct {
	Contract *SloadCallCaller // Generic read-only contract binding to access the raw methods on
}

// SloadCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SloadCallTransactorRaw struct {
	Contract *SloadCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSloadCall creates a new instance of SloadCall, bound to a specific deployed contract.
func NewSloadCall(address common.Address, backend bind.ContractBackend) (*SloadCall, error) {
	contract, err := bindSloadCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SloadCall{SloadCallCaller: SloadCallCaller{contract: contract}, SloadCallTransactor: SloadCallTransactor{contract: contract}, SloadCallFilterer: SloadCallFilterer{contract: contract}}, nil
}

// NewSloadCallCaller creates a new read-only instance of SloadCall, bound to a specific deployed contract.
func NewSloadCallCaller(address common.Address, caller bind.ContractCaller) (*SloadCallCaller, error) {
	contract, err := bindSloadCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SloadCallCaller{contract: contract}, nil
}

// NewSloadCallTransactor creates a new write-only instance of SloadCall, bound to a specific deployed contract.
func NewSloadCallTransactor(address common.Address, transactor bind.ContractTransactor) (*SloadCallTransactor, error) {
	contract, err := bindSloadCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SloadCallTransactor{contract: contract}, nil
}

// NewSloadCallFilterer creates a new log filterer instance of SloadCall, bound to a specific deployed contract.
func NewSloadCallFilterer(address common.Address, filterer bind.ContractFilterer) (*SloadCallFilterer, error) {
	contract, err := bindSloadCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SloadCallFilterer{contract: contract}, nil
}

// bindSloadCall binds a generic wrapper to an already deployed contract.
func bindSloadCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SloadCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SloadCall *SloadCallRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SloadCall.Contract.SloadCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SloadCall *SloadCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SloadCall.Contract.SloadCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SloadCall *SloadCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SloadCall.Contract.SloadCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SloadCall *SloadCallCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SloadCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SloadCall *SloadCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SloadCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SloadCall *SloadCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SloadCall.Contract.contract.Transact(opts, method, params...)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns()
func (_SloadCall *SloadCallCaller) Start(opts *bind.CallOpts) error {
	var ()
	out := &[]interface{}{}
	err := _SloadCall.contract.Call(opts, out, "start")
	return err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns()
func (_SloadCall *SloadCallSession) Start() error {
	return _SloadCall.Contract.Start(&_SloadCall.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns()
func (_SloadCall *SloadCallCallerSession) Start() error {
	return _SloadCall.Contract.Start(&_SloadCall.CallOpts)
}
