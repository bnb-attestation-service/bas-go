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

func NewBASOnchainAttestation(schemaUid string, schema string, data map[string]interface{}, attestor string, recipient string, revocable bool, refUid string, expirationTime uint64, value string, deadline uint64, signer *ecdsa.PrivateKey) (*Signature, error) {

	attest := OnchainAttestationParam{}
	attest.Domain = BASTESTDOMAIN
	message := OnchainAttestationMessage{}

	if _data, err := EncodeData(schema, data); err != nil {
		return nil, fmt.Errorf("encode data error: " + err.Error())
	} else {

		message["data"] = crypto.Keccak256(_data)

	}

	message["value"] = value

	message["recipient"] = recipient

	message["expirationTime"] = big.NewInt(int64(expirationTime))

	message["attestor"] = attestor

	message["revocable"] = revocable

	message["deadline"] = big.NewInt(int64(deadline))

	message["schema"] = schemaUid

	message["refUID"] = refUid

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
