// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eip712

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

// AttestationRequestData is an auto generated low-level Go binding around an user-defined struct.
type AttestationRequestData struct {
	Recipient      common.Address
	ExpirationTime uint64
	Revocable      bool
	RefUID         [32]byte
	Data           []byte
	Value          *big.Int
}

// DelegatedProxyAttestationRequest is an auto generated low-level Go binding around an user-defined struct.
type DelegatedProxyAttestationRequest struct {
	Schema    [32]byte
	Data      AttestationRequestData
	Signature Signature
	Attester  common.Address
	Deadline  uint64
}

// DelegatedProxyRevocationRequest is an auto generated low-level Go binding around an user-defined struct.
type DelegatedProxyRevocationRequest struct {
	Schema    [32]byte
	Data      RevocationRequestData
	Signature Signature
	Revoker   common.Address
	Deadline  uint64
}

// MultiDelegatedProxyAttestationRequest is an auto generated low-level Go binding around an user-defined struct.
type MultiDelegatedProxyAttestationRequest struct {
	Schema     [32]byte
	Data       []AttestationRequestData
	Signatures []Signature
	Attester   common.Address
	Deadline   uint64
}

// MultiDelegatedProxyRevocationRequest is an auto generated low-level Go binding around an user-defined struct.
type MultiDelegatedProxyRevocationRequest struct {
	Schema     [32]byte
	Data       []RevocationRequestData
	Signatures []Signature
	Revoker    common.Address
	Deadline   uint64
}

// RevocationRequestData is an auto generated low-level Go binding around an user-defined struct.
type RevocationRequestData struct {
	Uid   [32]byte
	Value *big.Int
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// EIP712MetaData contains all meta data concerning the EIP712 contract.
var EIP712MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEAS\",\"name\":\"eas\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEAS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsedSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATTEST_PROXY_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyAttestationRequest\",\"name\":\"delegatedRequest\",\"type\":\"tuple\"}],\"name\":\"attestByDelegation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAttestTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyAttestationRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"getAttestationData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyAttestationRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"getAttestationDataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"name\":\"getAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEAS\",\"outputs\":[{\"internalType\":\"contractIEAS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevokeTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyAttestationRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"getTypedData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"typedData\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyAttestationRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"getTypedData1\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData[]\",\"name\":\"data\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structMultiDelegatedProxyAttestationRequest[]\",\"name\":\"multiDelegatedRequests\",\"type\":\"tuple[]\"}],\"name\":\"multiAttestByDelegation\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structRevocationRequestData[]\",\"name\":\"data\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structMultiDelegatedProxyRevocationRequest[]\",\"name\":\"multiDelegatedRequests\",\"type\":\"tuple[]\"}],\"name\":\"multiRevokeByDelegation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structRevocationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyRevocationRequest\",\"name\":\"delegatedRequest\",\"type\":\"tuple\"}],\"name\":\"revokeByDelegation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expirationTime\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"refUID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structAttestationRequestData\",\"name\":\"data\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structDelegatedProxyAttestationRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pass\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EIP712ABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP712MetaData.ABI instead.
var EIP712ABI = EIP712MetaData.ABI

// EIP712 is an auto generated Go binding around an Ethereum contract.
type EIP712 struct {
	EIP712Caller     // Read-only binding to the contract
	EIP712Transactor // Write-only binding to the contract
	EIP712Filterer   // Log filterer for contract events
}

// EIP712Caller is an auto generated read-only Go binding around an Ethereum contract.
type EIP712Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP712Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP712Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP712Session struct {
	Contract     *EIP712           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP712CallerSession struct {
	Contract *EIP712Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EIP712TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP712TransactorSession struct {
	Contract     *EIP712Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712Raw is an auto generated low-level Go binding around an Ethereum contract.
type EIP712Raw struct {
	Contract *EIP712 // Generic contract binding to access the raw methods on
}

// EIP712CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP712CallerRaw struct {
	Contract *EIP712Caller // Generic read-only contract binding to access the raw methods on
}

// EIP712TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP712TransactorRaw struct {
	Contract *EIP712Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP712 creates a new instance of EIP712, bound to a specific deployed contract.
func NewEIP712(address common.Address, backend bind.ContractBackend) (*EIP712, error) {
	contract, err := bindEIP712(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP712{EIP712Caller: EIP712Caller{contract: contract}, EIP712Transactor: EIP712Transactor{contract: contract}, EIP712Filterer: EIP712Filterer{contract: contract}}, nil
}

// NewEIP712Caller creates a new read-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Caller(address common.Address, caller bind.ContractCaller) (*EIP712Caller, error) {
	contract, err := bindEIP712(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Caller{contract: contract}, nil
}

// NewEIP712Transactor creates a new write-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Transactor(address common.Address, transactor bind.ContractTransactor) (*EIP712Transactor, error) {
	contract, err := bindEIP712(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Transactor{contract: contract}, nil
}

// NewEIP712Filterer creates a new log filterer instance of EIP712, bound to a specific deployed contract.
func NewEIP712Filterer(address common.Address, filterer bind.ContractFilterer) (*EIP712Filterer, error) {
	contract, err := bindEIP712(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP712Filterer{contract: contract}, nil
}

// bindEIP712 binds a generic wrapper to an already deployed contract.
func bindEIP712(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EIP712MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.EIP712Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transact(opts, method, params...)
}

// ATTESTPROXYTYPEHASH is a free data retrieval call binding the contract method 0x48e1d352.
//
// Solidity: function ATTEST_PROXY_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Caller) ATTESTPROXYTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "ATTEST_PROXY_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ATTESTPROXYTYPEHASH is a free data retrieval call binding the contract method 0x48e1d352.
//
// Solidity: function ATTEST_PROXY_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712Session) ATTESTPROXYTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ATTESTPROXYTYPEHASH(&_EIP712.CallOpts)
}

// ATTESTPROXYTYPEHASH is a free data retrieval call binding the contract method 0x48e1d352.
//
// Solidity: function ATTEST_PROXY_TYPEHASH() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) ATTESTPROXYTYPEHASH() ([32]byte, error) {
	return _EIP712.Contract.ATTESTPROXYTYPEHASH(&_EIP712.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EIP712 *EIP712Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EIP712 *EIP712Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EIP712.Contract.Eip712Domain(&_EIP712.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EIP712 *EIP712CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EIP712.Contract.Eip712Domain(&_EIP712.CallOpts)
}

// GetAttestTypeHash is a free data retrieval call binding the contract method 0x12b11a17.
//
// Solidity: function getAttestTypeHash() pure returns(bytes32)
func (_EIP712 *EIP712Caller) GetAttestTypeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getAttestTypeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAttestTypeHash is a free data retrieval call binding the contract method 0x12b11a17.
//
// Solidity: function getAttestTypeHash() pure returns(bytes32)
func (_EIP712 *EIP712Session) GetAttestTypeHash() ([32]byte, error) {
	return _EIP712.Contract.GetAttestTypeHash(&_EIP712.CallOpts)
}

// GetAttestTypeHash is a free data retrieval call binding the contract method 0x12b11a17.
//
// Solidity: function getAttestTypeHash() pure returns(bytes32)
func (_EIP712 *EIP712CallerSession) GetAttestTypeHash() ([32]byte, error) {
	return _EIP712.Contract.GetAttestTypeHash(&_EIP712.CallOpts)
}

// GetAttestationData is a free data retrieval call binding the contract method 0xc01fa67e.
//
// Solidity: function getAttestationData((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes result)
func (_EIP712 *EIP712Caller) GetAttestationData(opts *bind.CallOpts, request DelegatedProxyAttestationRequest) ([]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getAttestationData", request)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestationData is a free data retrieval call binding the contract method 0xc01fa67e.
//
// Solidity: function getAttestationData((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes result)
func (_EIP712 *EIP712Session) GetAttestationData(request DelegatedProxyAttestationRequest) ([]byte, error) {
	return _EIP712.Contract.GetAttestationData(&_EIP712.CallOpts, request)
}

// GetAttestationData is a free data retrieval call binding the contract method 0xc01fa67e.
//
// Solidity: function getAttestationData((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes result)
func (_EIP712 *EIP712CallerSession) GetAttestationData(request DelegatedProxyAttestationRequest) ([]byte, error) {
	return _EIP712.Contract.GetAttestationData(&_EIP712.CallOpts, request)
}

// GetAttestationDataHash is a free data retrieval call binding the contract method 0x58a9c7c1.
//
// Solidity: function getAttestationDataHash((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes32 result)
func (_EIP712 *EIP712Caller) GetAttestationDataHash(opts *bind.CallOpts, request DelegatedProxyAttestationRequest) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getAttestationDataHash", request)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAttestationDataHash is a free data retrieval call binding the contract method 0x58a9c7c1.
//
// Solidity: function getAttestationDataHash((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes32 result)
func (_EIP712 *EIP712Session) GetAttestationDataHash(request DelegatedProxyAttestationRequest) ([32]byte, error) {
	return _EIP712.Contract.GetAttestationDataHash(&_EIP712.CallOpts, request)
}

// GetAttestationDataHash is a free data retrieval call binding the contract method 0x58a9c7c1.
//
// Solidity: function getAttestationDataHash((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes32 result)
func (_EIP712 *EIP712CallerSession) GetAttestationDataHash(request DelegatedProxyAttestationRequest) ([32]byte, error) {
	return _EIP712.Contract.GetAttestationDataHash(&_EIP712.CallOpts, request)
}

// GetAttester is a free data retrieval call binding the contract method 0x10d736d5.
//
// Solidity: function getAttester(bytes32 uid) view returns(address)
func (_EIP712 *EIP712Caller) GetAttester(opts *bind.CallOpts, uid [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getAttester", uid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAttester is a free data retrieval call binding the contract method 0x10d736d5.
//
// Solidity: function getAttester(bytes32 uid) view returns(address)
func (_EIP712 *EIP712Session) GetAttester(uid [32]byte) (common.Address, error) {
	return _EIP712.Contract.GetAttester(&_EIP712.CallOpts, uid)
}

// GetAttester is a free data retrieval call binding the contract method 0x10d736d5.
//
// Solidity: function getAttester(bytes32 uid) view returns(address)
func (_EIP712 *EIP712CallerSession) GetAttester(uid [32]byte) (common.Address, error) {
	return _EIP712.Contract.GetAttester(&_EIP712.CallOpts, uid)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_EIP712 *EIP712Caller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_EIP712 *EIP712Session) GetDomainSeparator() ([32]byte, error) {
	return _EIP712.Contract.GetDomainSeparator(&_EIP712.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_EIP712 *EIP712CallerSession) GetDomainSeparator() ([32]byte, error) {
	return _EIP712.Contract.GetDomainSeparator(&_EIP712.CallOpts)
}

// GetEAS is a free data retrieval call binding the contract method 0x65c40b9c.
//
// Solidity: function getEAS() view returns(address)
func (_EIP712 *EIP712Caller) GetEAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getEAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEAS is a free data retrieval call binding the contract method 0x65c40b9c.
//
// Solidity: function getEAS() view returns(address)
func (_EIP712 *EIP712Session) GetEAS() (common.Address, error) {
	return _EIP712.Contract.GetEAS(&_EIP712.CallOpts)
}

// GetEAS is a free data retrieval call binding the contract method 0x65c40b9c.
//
// Solidity: function getEAS() view returns(address)
func (_EIP712 *EIP712CallerSession) GetEAS() (common.Address, error) {
	return _EIP712.Contract.GetEAS(&_EIP712.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_EIP712 *EIP712Caller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_EIP712 *EIP712Session) GetName() (string, error) {
	return _EIP712.Contract.GetName(&_EIP712.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_EIP712 *EIP712CallerSession) GetName() (string, error) {
	return _EIP712.Contract.GetName(&_EIP712.CallOpts)
}

// GetRevokeTypeHash is a free data retrieval call binding the contract method 0xb83010d3.
//
// Solidity: function getRevokeTypeHash() pure returns(bytes32)
func (_EIP712 *EIP712Caller) GetRevokeTypeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getRevokeTypeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRevokeTypeHash is a free data retrieval call binding the contract method 0xb83010d3.
//
// Solidity: function getRevokeTypeHash() pure returns(bytes32)
func (_EIP712 *EIP712Session) GetRevokeTypeHash() ([32]byte, error) {
	return _EIP712.Contract.GetRevokeTypeHash(&_EIP712.CallOpts)
}

// GetRevokeTypeHash is a free data retrieval call binding the contract method 0xb83010d3.
//
// Solidity: function getRevokeTypeHash() pure returns(bytes32)
func (_EIP712 *EIP712CallerSession) GetRevokeTypeHash() ([32]byte, error) {
	return _EIP712.Contract.GetRevokeTypeHash(&_EIP712.CallOpts)
}

// GetTypedData is a free data retrieval call binding the contract method 0xba3326f2.
//
// Solidity: function getTypedData((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) view returns(bytes32 typedData)
func (_EIP712 *EIP712Caller) GetTypedData(opts *bind.CallOpts, request DelegatedProxyAttestationRequest) ([32]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getTypedData", request)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTypedData is a free data retrieval call binding the contract method 0xba3326f2.
//
// Solidity: function getTypedData((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) view returns(bytes32 typedData)
func (_EIP712 *EIP712Session) GetTypedData(request DelegatedProxyAttestationRequest) ([32]byte, error) {
	return _EIP712.Contract.GetTypedData(&_EIP712.CallOpts, request)
}

// GetTypedData is a free data retrieval call binding the contract method 0xba3326f2.
//
// Solidity: function getTypedData((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) view returns(bytes32 typedData)
func (_EIP712 *EIP712CallerSession) GetTypedData(request DelegatedProxyAttestationRequest) ([32]byte, error) {
	return _EIP712.Contract.GetTypedData(&_EIP712.CallOpts, request)
}

// GetTypedData1 is a free data retrieval call binding the contract method 0xe1706de9.
//
// Solidity: function getTypedData1((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes result)
func (_EIP712 *EIP712Caller) GetTypedData1(opts *bind.CallOpts, request DelegatedProxyAttestationRequest) ([]byte, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "getTypedData1", request)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTypedData1 is a free data retrieval call binding the contract method 0xe1706de9.
//
// Solidity: function getTypedData1((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes result)
func (_EIP712 *EIP712Session) GetTypedData1(request DelegatedProxyAttestationRequest) ([]byte, error) {
	return _EIP712.Contract.GetTypedData1(&_EIP712.CallOpts, request)
}

// GetTypedData1 is a free data retrieval call binding the contract method 0xe1706de9.
//
// Solidity: function getTypedData1((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) pure returns(bytes result)
func (_EIP712 *EIP712CallerSession) GetTypedData1(request DelegatedProxyAttestationRequest) ([]byte, error) {
	return _EIP712.Contract.GetTypedData1(&_EIP712.CallOpts, request)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EIP712 *EIP712Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EIP712 *EIP712Session) Owner() (common.Address, error) {
	return _EIP712.Contract.Owner(&_EIP712.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EIP712 *EIP712CallerSession) Owner() (common.Address, error) {
	return _EIP712.Contract.Owner(&_EIP712.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1fccbcf7.
//
// Solidity: function verifyAttestation((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) view returns(bool pass)
func (_EIP712 *EIP712Caller) VerifyAttestation(opts *bind.CallOpts, request DelegatedProxyAttestationRequest) (bool, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "verifyAttestation", request)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1fccbcf7.
//
// Solidity: function verifyAttestation((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) view returns(bool pass)
func (_EIP712 *EIP712Session) VerifyAttestation(request DelegatedProxyAttestationRequest) (bool, error) {
	return _EIP712.Contract.VerifyAttestation(&_EIP712.CallOpts, request)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1fccbcf7.
//
// Solidity: function verifyAttestation((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) request) view returns(bool pass)
func (_EIP712 *EIP712CallerSession) VerifyAttestation(request DelegatedProxyAttestationRequest) (bool, error) {
	return _EIP712.Contract.VerifyAttestation(&_EIP712.CallOpts, request)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EIP712 *EIP712Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EIP712.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EIP712 *EIP712Session) Version() (string, error) {
	return _EIP712.Contract.Version(&_EIP712.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EIP712 *EIP712CallerSession) Version() (string, error) {
	return _EIP712.Contract.Version(&_EIP712.CallOpts)
}

// AttestByDelegation is a paid mutator transaction binding the contract method 0x3c042715.
//
// Solidity: function attestByDelegation((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) delegatedRequest) payable returns(bytes32)
func (_EIP712 *EIP712Transactor) AttestByDelegation(opts *bind.TransactOpts, delegatedRequest DelegatedProxyAttestationRequest) (*types.Transaction, error) {
	return _EIP712.contract.Transact(opts, "attestByDelegation", delegatedRequest)
}

// AttestByDelegation is a paid mutator transaction binding the contract method 0x3c042715.
//
// Solidity: function attestByDelegation((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) delegatedRequest) payable returns(bytes32)
func (_EIP712 *EIP712Session) AttestByDelegation(delegatedRequest DelegatedProxyAttestationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.AttestByDelegation(&_EIP712.TransactOpts, delegatedRequest)
}

// AttestByDelegation is a paid mutator transaction binding the contract method 0x3c042715.
//
// Solidity: function attestByDelegation((bytes32,(address,uint64,bool,bytes32,bytes,uint256),(uint8,bytes32,bytes32),address,uint64) delegatedRequest) payable returns(bytes32)
func (_EIP712 *EIP712TransactorSession) AttestByDelegation(delegatedRequest DelegatedProxyAttestationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.AttestByDelegation(&_EIP712.TransactOpts, delegatedRequest)
}

// MultiAttestByDelegation is a paid mutator transaction binding the contract method 0x95411525.
//
// Solidity: function multiAttestByDelegation((bytes32,(address,uint64,bool,bytes32,bytes,uint256)[],(uint8,bytes32,bytes32)[],address,uint64)[] multiDelegatedRequests) payable returns(bytes32[])
func (_EIP712 *EIP712Transactor) MultiAttestByDelegation(opts *bind.TransactOpts, multiDelegatedRequests []MultiDelegatedProxyAttestationRequest) (*types.Transaction, error) {
	return _EIP712.contract.Transact(opts, "multiAttestByDelegation", multiDelegatedRequests)
}

// MultiAttestByDelegation is a paid mutator transaction binding the contract method 0x95411525.
//
// Solidity: function multiAttestByDelegation((bytes32,(address,uint64,bool,bytes32,bytes,uint256)[],(uint8,bytes32,bytes32)[],address,uint64)[] multiDelegatedRequests) payable returns(bytes32[])
func (_EIP712 *EIP712Session) MultiAttestByDelegation(multiDelegatedRequests []MultiDelegatedProxyAttestationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.MultiAttestByDelegation(&_EIP712.TransactOpts, multiDelegatedRequests)
}

// MultiAttestByDelegation is a paid mutator transaction binding the contract method 0x95411525.
//
// Solidity: function multiAttestByDelegation((bytes32,(address,uint64,bool,bytes32,bytes,uint256)[],(uint8,bytes32,bytes32)[],address,uint64)[] multiDelegatedRequests) payable returns(bytes32[])
func (_EIP712 *EIP712TransactorSession) MultiAttestByDelegation(multiDelegatedRequests []MultiDelegatedProxyAttestationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.MultiAttestByDelegation(&_EIP712.TransactOpts, multiDelegatedRequests)
}

// MultiRevokeByDelegation is a paid mutator transaction binding the contract method 0x0eabf660.
//
// Solidity: function multiRevokeByDelegation((bytes32,(bytes32,uint256)[],(uint8,bytes32,bytes32)[],address,uint64)[] multiDelegatedRequests) payable returns()
func (_EIP712 *EIP712Transactor) MultiRevokeByDelegation(opts *bind.TransactOpts, multiDelegatedRequests []MultiDelegatedProxyRevocationRequest) (*types.Transaction, error) {
	return _EIP712.contract.Transact(opts, "multiRevokeByDelegation", multiDelegatedRequests)
}

// MultiRevokeByDelegation is a paid mutator transaction binding the contract method 0x0eabf660.
//
// Solidity: function multiRevokeByDelegation((bytes32,(bytes32,uint256)[],(uint8,bytes32,bytes32)[],address,uint64)[] multiDelegatedRequests) payable returns()
func (_EIP712 *EIP712Session) MultiRevokeByDelegation(multiDelegatedRequests []MultiDelegatedProxyRevocationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.MultiRevokeByDelegation(&_EIP712.TransactOpts, multiDelegatedRequests)
}

// MultiRevokeByDelegation is a paid mutator transaction binding the contract method 0x0eabf660.
//
// Solidity: function multiRevokeByDelegation((bytes32,(bytes32,uint256)[],(uint8,bytes32,bytes32)[],address,uint64)[] multiDelegatedRequests) payable returns()
func (_EIP712 *EIP712TransactorSession) MultiRevokeByDelegation(multiDelegatedRequests []MultiDelegatedProxyRevocationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.MultiRevokeByDelegation(&_EIP712.TransactOpts, multiDelegatedRequests)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EIP712 *EIP712Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EIP712 *EIP712Session) RenounceOwnership() (*types.Transaction, error) {
	return _EIP712.Contract.RenounceOwnership(&_EIP712.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EIP712 *EIP712TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EIP712.Contract.RenounceOwnership(&_EIP712.TransactOpts)
}

// RevokeByDelegation is a paid mutator transaction binding the contract method 0xa6d4dbc7.
//
// Solidity: function revokeByDelegation((bytes32,(bytes32,uint256),(uint8,bytes32,bytes32),address,uint64) delegatedRequest) payable returns()
func (_EIP712 *EIP712Transactor) RevokeByDelegation(opts *bind.TransactOpts, delegatedRequest DelegatedProxyRevocationRequest) (*types.Transaction, error) {
	return _EIP712.contract.Transact(opts, "revokeByDelegation", delegatedRequest)
}

// RevokeByDelegation is a paid mutator transaction binding the contract method 0xa6d4dbc7.
//
// Solidity: function revokeByDelegation((bytes32,(bytes32,uint256),(uint8,bytes32,bytes32),address,uint64) delegatedRequest) payable returns()
func (_EIP712 *EIP712Session) RevokeByDelegation(delegatedRequest DelegatedProxyRevocationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.RevokeByDelegation(&_EIP712.TransactOpts, delegatedRequest)
}

// RevokeByDelegation is a paid mutator transaction binding the contract method 0xa6d4dbc7.
//
// Solidity: function revokeByDelegation((bytes32,(bytes32,uint256),(uint8,bytes32,bytes32),address,uint64) delegatedRequest) payable returns()
func (_EIP712 *EIP712TransactorSession) RevokeByDelegation(delegatedRequest DelegatedProxyRevocationRequest) (*types.Transaction, error) {
	return _EIP712.Contract.RevokeByDelegation(&_EIP712.TransactOpts, delegatedRequest)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EIP712 *EIP712Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EIP712.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EIP712 *EIP712Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EIP712.Contract.TransferOwnership(&_EIP712.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EIP712 *EIP712TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EIP712.Contract.TransferOwnership(&_EIP712.TransactOpts, newOwner)
}

// EIP712EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the EIP712 contract.
type EIP712EIP712DomainChangedIterator struct {
	Event *EIP712EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EIP712EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP712EIP712DomainChanged)
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
		it.Event = new(EIP712EIP712DomainChanged)
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
func (it *EIP712EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP712EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP712EIP712DomainChanged represents a EIP712DomainChanged event raised by the EIP712 contract.
type EIP712EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EIP712 *EIP712Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EIP712EIP712DomainChangedIterator, error) {

	logs, sub, err := _EIP712.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EIP712EIP712DomainChangedIterator{contract: _EIP712.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EIP712 *EIP712Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EIP712EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _EIP712.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP712EIP712DomainChanged)
				if err := _EIP712.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EIP712 *EIP712Filterer) ParseEIP712DomainChanged(log types.Log) (*EIP712EIP712DomainChanged, error) {
	event := new(EIP712EIP712DomainChanged)
	if err := _EIP712.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EIP712OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EIP712 contract.
type EIP712OwnershipTransferredIterator struct {
	Event *EIP712OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EIP712OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP712OwnershipTransferred)
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
		it.Event = new(EIP712OwnershipTransferred)
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
func (it *EIP712OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP712OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP712OwnershipTransferred represents a OwnershipTransferred event raised by the EIP712 contract.
type EIP712OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EIP712 *EIP712Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EIP712OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EIP712.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EIP712OwnershipTransferredIterator{contract: _EIP712.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EIP712 *EIP712Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EIP712OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EIP712.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP712OwnershipTransferred)
				if err := _EIP712.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EIP712 *EIP712Filterer) ParseOwnershipTransferred(log types.Log) (*EIP712OwnershipTransferred, error) {
	event := new(EIP712OwnershipTransferred)
	if err := _EIP712.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
