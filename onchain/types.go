package onchain

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
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

var BASTESTDOMAIN = OnchainAttestationDomain{
	Name:              "BAS Attestation",
	Version:           "1.3.0",
	ChainId:           "97",
	VerifyingContract: "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD",
}

var OPBNBTESTDOMAIN = OnchainAttestationDomain{
	Name:              "OPBNB ATTESTATION",
	Version:           "1.3.0",
	ChainId:           "5611",
	VerifyingContract: "0x5239d34BDa6b05ee47d1310B7Aaf69BB6e864d36",
}

func EncodeData(schema string, data map[string]interface{}) ([]byte, error) {
	_schema := fmt.Sprintf("tuple(%s)", schema)
	fmt.Println(_schema, data)
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

func NewBASOnchainDelegateAttestation(param OnchainDelegateAttestationParam, domain OnchainAttestationDomain, signer *ecdsa.PrivateKey) (*Signature, error) {

	attest := OnchainAttestationParam{}
	attest.Domain = domain
	message := OnchainAttestationMessage{}

	if _data, err := EncodeData(param.Schema, param.Data); err != nil {
		return nil, fmt.Errorf("encode data error: " + err.Error())
	} else {

		message["data"] = crypto.Keccak256(_data)

	}

	message["value"] = param.Value

	message["recipient"] = param.Recipient

	message["expirationTime"] = big.NewInt(int64(param.ExpirationTime))

	message["attestor"] = param.Attestor

	message["revocable"] = param.Revocable

	message["deadline"] = big.NewInt(int64(param.Deadline))

	message["schema"] = param.SchemaUid

	message["refUID"] = param.RefUid

	attest.Message = message

	attest.PrimaryType = "Attest"
	atypes := []types.Type{
		{Name: "attestor", Type: "address"},
		{Name: "schema", Type: "bytes32"},
		{Name: "recipient", Type: "address"},
		{Name: "expirationTime", Type: "uint64"},
		{Name: "revocable", Type: "bool"},
		{Name: "refUID", Type: "bytes32"},
		{Name: "data", Type: "bytes"},
		{Name: "deadline", Type: "uint64"},
		{Name: "value", Type: "uint256"},
	}
	attest.Type = map[string][]types.Type{}
	attest.Type["Attest"] = atypes

	if sig, err := Sign(attest.Domain, attest.Type, attest.Message, signer); err != nil {
		return nil, err
	} else {
		if _sig, err := extractSignature(sig); err != nil {
			return nil, err
		} else {
			return &_sig, nil
		}
	}
}
