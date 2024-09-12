// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package schemaRegistry

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

// Attestation is an auto generated low-level Go binding around an user-defined struct.
type Attestation struct {
	SchemaId [32]byte
	Name     string
}

// SchemaNameMetaData contains all meta data concerning the SchemaName contract.
var SchemaNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schemaId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structAttestation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SchemaNameABI is the input ABI used to generate the binding from.
// Deprecated: Use SchemaNameMetaData.ABI instead.
var SchemaNameABI = SchemaNameMetaData.ABI

// SchemaName is an auto generated Go binding around an Ethereum contract.
type SchemaName struct {
	SchemaNameCaller     // Read-only binding to the contract
	SchemaNameTransactor // Write-only binding to the contract
	SchemaNameFilterer   // Log filterer for contract events
}

// SchemaNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type SchemaNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SchemaNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SchemaNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchemaNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SchemaNameSession struct {
	Contract     *SchemaName       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SchemaNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SchemaNameCallerSession struct {
	Contract *SchemaNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SchemaNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SchemaNameTransactorSession struct {
	Contract     *SchemaNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SchemaNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type SchemaNameRaw struct {
	Contract *SchemaName // Generic contract binding to access the raw methods on
}

// SchemaNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SchemaNameCallerRaw struct {
	Contract *SchemaNameCaller // Generic read-only contract binding to access the raw methods on
}

// SchemaNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SchemaNameTransactorRaw struct {
	Contract *SchemaNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSchemaName creates a new instance of SchemaName, bound to a specific deployed contract.
func NewSchemaName(address common.Address, backend bind.ContractBackend) (*SchemaName, error) {
	contract, err := bindSchemaName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SchemaName{SchemaNameCaller: SchemaNameCaller{contract: contract}, SchemaNameTransactor: SchemaNameTransactor{contract: contract}, SchemaNameFilterer: SchemaNameFilterer{contract: contract}}, nil
}

// NewSchemaNameCaller creates a new read-only instance of SchemaName, bound to a specific deployed contract.
func NewSchemaNameCaller(address common.Address, caller bind.ContractCaller) (*SchemaNameCaller, error) {
	contract, err := bindSchemaName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaNameCaller{contract: contract}, nil
}

// NewSchemaNameTransactor creates a new write-only instance of SchemaName, bound to a specific deployed contract.
func NewSchemaNameTransactor(address common.Address, transactor bind.ContractTransactor) (*SchemaNameTransactor, error) {
	contract, err := bindSchemaName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SchemaNameTransactor{contract: contract}, nil
}

// NewSchemaNameFilterer creates a new log filterer instance of SchemaName, bound to a specific deployed contract.
func NewSchemaNameFilterer(address common.Address, filterer bind.ContractFilterer) (*SchemaNameFilterer, error) {
	contract, err := bindSchemaName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SchemaNameFilterer{contract: contract}, nil
}

// bindSchemaName binds a generic wrapper to an already deployed contract.
func bindSchemaName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SchemaNameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaName *SchemaNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SchemaName.Contract.SchemaNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaName *SchemaNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaName.Contract.SchemaNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaName *SchemaNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaName.Contract.SchemaNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SchemaName *SchemaNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SchemaName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SchemaName *SchemaNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SchemaName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SchemaName *SchemaNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SchemaName.Contract.contract.Transact(opts, method, params...)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa3112a64.
//
// Solidity: function getAttestation(bytes32 uid) view returns((bytes32,string))
func (_SchemaName *SchemaNameCaller) GetAttestation(opts *bind.CallOpts, uid [32]byte) (Attestation, error) {
	var out []interface{}
	err := _SchemaName.contract.Call(opts, &out, "getAttestation", uid)

	if err != nil {
		return *new(Attestation), err
	}

	out0 := *abi.ConvertType(out[0], new(Attestation)).(*Attestation)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa3112a64.
//
// Solidity: function getAttestation(bytes32 uid) view returns((bytes32,string))
func (_SchemaName *SchemaNameSession) GetAttestation(uid [32]byte) (Attestation, error) {
	return _SchemaName.Contract.GetAttestation(&_SchemaName.CallOpts, uid)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa3112a64.
//
// Solidity: function getAttestation(bytes32 uid) view returns((bytes32,string))
func (_SchemaName *SchemaNameCallerSession) GetAttestation(uid [32]byte) (Attestation, error) {
	return _SchemaName.Contract.GetAttestation(&_SchemaName.CallOpts, uid)
}
