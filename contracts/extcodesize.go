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
const ExtcodesizeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ExtcodesizeFuncSigs maps the 4-byte function signature to its string representation.
var ExtcodesizeFuncSigs = map[string]string{
	"be9a6555": "start()",
}

// ExtcodesizeBin is the compiled bytecode used for deploying new contracts.
var ExtcodesizeBin = "0x6080604052348015600f57600080fd5b5060888061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063be9a655514602d575b600080fd5b60336045565b60408051918252519081900360200190f35b6000805a90505a900390509056fea265627a7a72315820ae8725035a6487afe0bf1ab9e651d8cf659793779d7407c4ebe2db5de915db3364736f6c63430005100032"

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

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_Extcodesize *ExtcodesizeCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Extcodesize.contract.Call(opts, out, "start")
	return *ret0, err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_Extcodesize *ExtcodesizeSession) Start() (*big.Int, error) {
	return _Extcodesize.Contract.Start(&_Extcodesize.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_Extcodesize *ExtcodesizeCallerSession) Start() (*big.Int, error) {
	return _Extcodesize.Contract.Start(&_Extcodesize.CallOpts)
}
