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

// ExtcodesizeABI is the input ABI used to generate the binding from.
const ExtcodesizeABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExtcodesizeFuncSigs maps the 4-byte function signature to its string representation.
var ExtcodesizeFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// ExtcodesizeBin is the compiled bytecode used for deploying new contracts.
var ExtcodesizeBin = "0x6080604052348015600f57600080fd5b5060888061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336045565b60408051918252519081900360200190f35b6000805a90505a900390509056fea265627a7a7231582040ef0068afc54c54bda1e87362f23e706b65f0d81981a0aef9de49c75d300f4a64736f6c63430005100032"

// DeployExtcodesize deploys a new Ethereum contract, binding an instance of Extcodesize to it.
func DeployExtcodesize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Extcodesize, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtcodesizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExtcodesizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Extcodesize{ExtcodesizeCaller: ExtcodesizeCaller{contract: contract}, ExtcodesizeTransactor: ExtcodesizeTransactor{contract: contract}, ExtcodesizeFilterer: ExtcodesizeFilterer{contract: contract}}, nil
}

// Extcodesize is an auto generated Go binding around an Ethereum contract.
type Extcodesize struct {
	ExtcodesizeCaller     // Read-only binding to the contract
	ExtcodesizeTransactor // Write-only binding to the contract
	ExtcodesizeFilterer   // Log filterer for contract events
}

// ExtcodesizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtcodesizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtcodesizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtcodesizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtcodesizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtcodesizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtcodesizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtcodesizeSession struct {
	Contract     *Extcodesize      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtcodesizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtcodesizeCallerSession struct {
	Contract *ExtcodesizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExtcodesizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtcodesizeTransactorSession struct {
	Contract     *ExtcodesizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExtcodesizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtcodesizeRaw struct {
	Contract *Extcodesize // Generic contract binding to access the raw methods on
}

// ExtcodesizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtcodesizeCallerRaw struct {
	Contract *ExtcodesizeCaller // Generic read-only contract binding to access the raw methods on
}

// ExtcodesizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtcodesizeTransactorRaw struct {
	Contract *ExtcodesizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtcodesize creates a new instance of Extcodesize, bound to a specific deployed contract.
func NewExtcodesize(address common.Address, backend bind.ContractBackend) (*Extcodesize, error) {
	contract, err := bindExtcodesize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Extcodesize{ExtcodesizeCaller: ExtcodesizeCaller{contract: contract}, ExtcodesizeTransactor: ExtcodesizeTransactor{contract: contract}, ExtcodesizeFilterer: ExtcodesizeFilterer{contract: contract}}, nil
}

// NewExtcodesizeCaller creates a new read-only instance of Extcodesize, bound to a specific deployed contract.
func NewExtcodesizeCaller(address common.Address, caller bind.ContractCaller) (*ExtcodesizeCaller, error) {
	contract, err := bindExtcodesize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeCaller{contract: contract}, nil
}

// NewExtcodesizeTransactor creates a new write-only instance of Extcodesize, bound to a specific deployed contract.
func NewExtcodesizeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtcodesizeTransactor, error) {
	contract, err := bindExtcodesize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeTransactor{contract: contract}, nil
}

// NewExtcodesizeFilterer creates a new log filterer instance of Extcodesize, bound to a specific deployed contract.
func NewExtcodesizeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtcodesizeFilterer, error) {
	contract, err := bindExtcodesize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeFilterer{contract: contract}, nil
}

// bindExtcodesize binds a generic wrapper to an already deployed contract.
func bindExtcodesize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtcodesizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extcodesize *ExtcodesizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Extcodesize.Contract.ExtcodesizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extcodesize *ExtcodesizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extcodesize.Contract.ExtcodesizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extcodesize *ExtcodesizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extcodesize.Contract.ExtcodesizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extcodesize *ExtcodesizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Extcodesize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extcodesize *ExtcodesizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extcodesize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extcodesize *ExtcodesizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extcodesize.Contract.contract.Transact(opts, method, params...)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns(uint256)
func (_Extcodesize *ExtcodesizeTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extcodesize.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns(uint256)
func (_Extcodesize *ExtcodesizeSession) Start() (*types.Transaction, error) {
	return _Extcodesize.Contract.Start(&_Extcodesize.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns(uint256)
func (_Extcodesize *ExtcodesizeTransactorSession) Start() (*types.Transaction, error) {
	return _Extcodesize.Contract.Start(&_Extcodesize.TransactOpts)
}

// ExtcodesizeAttackABI is the input ABI used to generate the binding from.
const ExtcodesizeAttackABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"loop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExtcodesizeAttackFuncSigs maps the 4-byte function signature to its string representation.
var ExtcodesizeAttackFuncSigs = map[string]string{
	"a92100cb": "loop()",
}

// ExtcodesizeAttackBin is the compiled bytecode used for deploying new contracts.
var ExtcodesizeAttackBin = "0x6080604052348015600f57600080fd5b5060988061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063a92100cb14602d575b600080fd5b60336045565b60408051918252519081900360200190f35b6000805a90505b61c350811115605c575a9050604c565b5a9150509056fea265627a7a723158208790ef4022e6381ebb8ed6db3bb324ead5dcd68801a77232a6dfc2a8e414a9d264736f6c63430005100032"

// DeployExtcodesizeAttack deploys a new Ethereum contract, binding an instance of ExtcodesizeAttack to it.
func DeployExtcodesizeAttack(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtcodesizeAttack, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtcodesizeAttackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExtcodesizeAttackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtcodesizeAttack{ExtcodesizeAttackCaller: ExtcodesizeAttackCaller{contract: contract}, ExtcodesizeAttackTransactor: ExtcodesizeAttackTransactor{contract: contract}, ExtcodesizeAttackFilterer: ExtcodesizeAttackFilterer{contract: contract}}, nil
}

// ExtcodesizeAttack is an auto generated Go binding around an Ethereum contract.
type ExtcodesizeAttack struct {
	ExtcodesizeAttackCaller     // Read-only binding to the contract
	ExtcodesizeAttackTransactor // Write-only binding to the contract
	ExtcodesizeAttackFilterer   // Log filterer for contract events
}

// ExtcodesizeAttackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtcodesizeAttackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtcodesizeAttackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtcodesizeAttackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtcodesizeAttackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtcodesizeAttackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtcodesizeAttackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtcodesizeAttackSession struct {
	Contract     *ExtcodesizeAttack // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExtcodesizeAttackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtcodesizeAttackCallerSession struct {
	Contract *ExtcodesizeAttackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExtcodesizeAttackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtcodesizeAttackTransactorSession struct {
	Contract     *ExtcodesizeAttackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExtcodesizeAttackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtcodesizeAttackRaw struct {
	Contract *ExtcodesizeAttack // Generic contract binding to access the raw methods on
}

// ExtcodesizeAttackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtcodesizeAttackCallerRaw struct {
	Contract *ExtcodesizeAttackCaller // Generic read-only contract binding to access the raw methods on
}

// ExtcodesizeAttackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtcodesizeAttackTransactorRaw struct {
	Contract *ExtcodesizeAttackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtcodesizeAttack creates a new instance of ExtcodesizeAttack, bound to a specific deployed contract.
func NewExtcodesizeAttack(address common.Address, backend bind.ContractBackend) (*ExtcodesizeAttack, error) {
	contract, err := bindExtcodesizeAttack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeAttack{ExtcodesizeAttackCaller: ExtcodesizeAttackCaller{contract: contract}, ExtcodesizeAttackTransactor: ExtcodesizeAttackTransactor{contract: contract}, ExtcodesizeAttackFilterer: ExtcodesizeAttackFilterer{contract: contract}}, nil
}

// NewExtcodesizeAttackCaller creates a new read-only instance of ExtcodesizeAttack, bound to a specific deployed contract.
func NewExtcodesizeAttackCaller(address common.Address, caller bind.ContractCaller) (*ExtcodesizeAttackCaller, error) {
	contract, err := bindExtcodesizeAttack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeAttackCaller{contract: contract}, nil
}

// NewExtcodesizeAttackTransactor creates a new write-only instance of ExtcodesizeAttack, bound to a specific deployed contract.
func NewExtcodesizeAttackTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtcodesizeAttackTransactor, error) {
	contract, err := bindExtcodesizeAttack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeAttackTransactor{contract: contract}, nil
}

// NewExtcodesizeAttackFilterer creates a new log filterer instance of ExtcodesizeAttack, bound to a specific deployed contract.
func NewExtcodesizeAttackFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtcodesizeAttackFilterer, error) {
	contract, err := bindExtcodesizeAttack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtcodesizeAttackFilterer{contract: contract}, nil
}

// bindExtcodesizeAttack binds a generic wrapper to an already deployed contract.
func bindExtcodesizeAttack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtcodesizeAttackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtcodesizeAttack *ExtcodesizeAttackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExtcodesizeAttack.Contract.ExtcodesizeAttackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtcodesizeAttack *ExtcodesizeAttackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtcodesizeAttack.Contract.ExtcodesizeAttackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtcodesizeAttack *ExtcodesizeAttackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtcodesizeAttack.Contract.ExtcodesizeAttackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtcodesizeAttack *ExtcodesizeAttackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExtcodesizeAttack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtcodesizeAttack *ExtcodesizeAttackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtcodesizeAttack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtcodesizeAttack *ExtcodesizeAttackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtcodesizeAttack.Contract.contract.Transact(opts, method, params...)
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns(uint256)
func (_ExtcodesizeAttack *ExtcodesizeAttackTransactor) Loop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtcodesizeAttack.contract.Transact(opts, "loop")
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns(uint256)
func (_ExtcodesizeAttack *ExtcodesizeAttackSession) Loop() (*types.Transaction, error) {
	return _ExtcodesizeAttack.Contract.Loop(&_ExtcodesizeAttack.TransactOpts)
}

// Loop is a paid mutator transaction binding the contract method 0xa92100cb.
//
// Solidity: function loop() returns(uint256)
func (_ExtcodesizeAttack *ExtcodesizeAttackTransactorSession) Loop() (*types.Transaction, error) {
	return _ExtcodesizeAttack.Contract.Loop(&_ExtcodesizeAttack.TransactOpts)
}
