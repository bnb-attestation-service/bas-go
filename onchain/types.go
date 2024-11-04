package onchain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/umbracle/ethgo/abi"
)

const (
	BASDOMAINNAME = "BAS Attestation"
	ZEROADDRESS   = "0x0000000000000000000000000000000000000000"

	DELEGATEATTESTPREFIX = "0xea02ffba7dcb45f6fc649714d23f315eef12e3b27f9a7735d8d8bf41eb2b1af1"
)

type Signature struct {
	R string
	S string
	V uint8
}

type OnchainAttestationDomain struct {
	Name              string
	Version           string
	ChainId           string
	VerifyingContract string
}

type OnchainAttestationType map[string][]types.Type
type OnchainAttestationMessage map[string]interface{}

type OnchainAttestationParam struct {
	Domain OnchainAttestationDomain `json:"domain"`

	Message     OnchainAttestationMessage `json:"message"`
	Type        OnchainAttestationType    `json:"types"`
	Signature   Signature                 `json:"signature"`
	PrimaryType string                    `json:"primaryType"`
	Uid         string                    `json:"uid"`
}

func EncodeData(schema string, data map[string]interface{}) ([]byte, error) {
	_schema := fmt.Sprintf("tuple(%s)", schema)

	typ := abi.MustNewType(_schema)
	if res, err := typ.Encode(data); err != nil {
		return nil, err
	} else {
		return res, nil
	}

}

type OnchainDelegateAttestationParam struct {
	Hash           string
	Attestor       string
	Schema         string
	SchemaUid      string
	Recipient      string
	ExpirationTime uint64
	Revocable      bool
	RefUid         string
	Data           map[string]interface{}
	Value          string
	Deadline       uint64
}

type AttestationRequestData struct {
	Recipient      common.Address
	ExpirationTime uint64
	Revocable      bool
	RefUID         [32]byte
	Data           []byte
	Value          *big.Int
}
type DelegatedProxyAttestation struct {
	Schema    [32]byte
	Data      AttestationRequestData
	Signature Signature
	Attester  common.Address
	Deadline  uint64
}

func NewBASOnchainDelegateAttestation(param OnchainDelegateAttestationParam, domain OnchainAttestationDomain, signer *ecdsa.PrivateKey) (*DelegatedProxyAttestation, error) {

	var data AttestationRequestData
	var result DelegatedProxyAttestation
	attest := OnchainAttestationParam{}
	attest.Domain = domain
	message := OnchainAttestationMessage{}

	var _data []byte
	var err error
	if _data, err = EncodeData(param.Schema, param.Data); err != nil {
		return nil, fmt.Errorf("encode data error: " + err.Error())
	} else {

		message["data"] = _data

	}

	message["value"] = param.Value

	message["recipient"] = param.Recipient

	message["expirationTime"] = big.NewInt(int64(param.ExpirationTime))

	message["attester"] = param.Attestor

	message["revocable"] = param.Revocable

	message["deadline"] = big.NewInt(int64(param.Deadline))

	message["schema"] = param.SchemaUid

	message["refUID"] = param.RefUid

	attest.Message = message

	attest.PrimaryType = "Attest"
	atypes := []types.Type{
		{Name: "attester", Type: "address"},
		{Name: "schema", Type: "bytes32"},
		{Name: "recipient", Type: "address"},
		{Name: "expirationTime", Type: "uint64"},
		{Name: "revocable", Type: "bool"},
		{Name: "refUID", Type: "bytes32"},
		{Name: "data", Type: "bytes"},
		{Name: "value", Type: "uint256"},
		{Name: "deadline", Type: "uint64"},
	}
	attest.Type = map[string][]types.Type{}
	attest.Type["Attest"] = atypes

	data.Data = _data
	data.ExpirationTime = param.ExpirationTime
	data.Recipient = common.HexToAddress(param.Recipient)
	if param.RefUid[:2] == "0x" {
		param.RefUid = param.RefUid[2:]
	}
	refUid, err := hex.DecodeString(param.RefUid)
	if len(refUid) != 32 || err != nil {
		return nil, fmt.Errorf("invalid refuid")
	}
	data.RefUID = [32]byte(refUid)
	data.Revocable = param.Revocable
	value := new(big.Int)
	if value, ok := value.SetString(param.Value, 10); !ok {
		return nil, fmt.Errorf("invalid value")
	} else {
		data.Value = value
	}

	result.Attester = common.HexToAddress(param.Attestor)
	result.Data = data
	result.Deadline = param.Deadline
	if param.SchemaUid[:2] == "0x" {
		param.SchemaUid = param.SchemaUid[2:]
	}
	if _uid, err := hex.DecodeString(param.SchemaUid); err != nil || len(_uid) != 32 {
		return nil, err
	} else {
		result.Schema = [32]byte(_uid)
	}

	if sig, err := Sign(attest.Domain, attest.Type, attest.Message, signer); err != nil {
		return nil, err
	} else {
		if _sig, err := extractSignature(sig); err != nil {
			return nil, err
		} else {
			result.Signature = _sig
		}
	}

	return &result, nil
}
