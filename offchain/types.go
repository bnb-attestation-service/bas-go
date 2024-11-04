package offchain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	types "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/spf13/cast"
	"github.com/umbracle/ethgo/abi"
)

const (
	BASDOMAINNAME = "BAS Attestation"
	ZEROADDRESS   = "0x0000000000000000000000000000000000000000"
)

type OffChainAttestation struct {
	MessageForUid
	Attestor string
}
type MessageForUid struct {
	Version        string
	Schema         string
	Recipient      string
	Time           uint64
	ExpirationTime uint64
	Revocable      bool
	RefUID         string
	Data           string
}

func (m *MessageForUid) Decode(message map[string]interface{}) error {

	if v, ok := message["version"].(string); !ok {
		if _v, ok := message["version"].(float64); !ok {
			return fmt.Errorf("decode message missing " + "version")
		} else {
			m.Version = strconv.Itoa(cast.ToInt(_v))
		}
	} else {
		m.Version = v
	}
	if v, ok := message["schema"].(string); !ok {
		return fmt.Errorf("decode message missing " + "schema")
	} else {
		m.Schema = v
	}
	if v, ok := message["recipient"].(string); !ok {
		return fmt.Errorf("decode message missing " + "recipient")
	} else {
		m.Recipient = v
	}
	if v, ok := message["time"].(float64); !ok {
		if _v, ok := message["time"].(string); !ok {
			return fmt.Errorf("decode message missing " + "time")
		} else {
			var err error
			if m.Time, err = strconv.ParseUint(_v, 10, 64); err != nil {
				return fmt.Errorf("message error time " + err.Error())
			}
		}
	} else {
		m.Time = cast.ToUint64(v)
	}
	if v, ok := message["expirationTime"].(float64); !ok {
		if _v, ok := message["expirationTime"].(string); !ok {
			return fmt.Errorf("decode message missing " + "expirationTime")
		} else {
			var err error
			if m.ExpirationTime, err = strconv.ParseUint(_v, 10, 64); err != nil {
				return fmt.Errorf("message error expirationTime: " + err.Error())
			}
		}
	} else {
		m.ExpirationTime = cast.ToUint64(v)
	}

	if v, ok := message["revocable"].(bool); !ok {
		return fmt.Errorf("decode message missing " + "revocable")
	} else {
		m.Revocable = v
	}

	if v, ok := message["refUID"].(string); !ok {
		return fmt.Errorf("decode message missing " + "refUID")
	} else {
		m.RefUID = v
	}

	if v, ok := message["data"].(string); !ok {
		return fmt.Errorf("decode message missing " + "data")
	} else {
		m.Data = v
	}

	return nil

}

type Signature struct {
	R string
	S string
	V uint8
}

type OffchainAttestationDomain struct {
	Name              string
	Version           string
	ChainId           string
	VerifyingContract string
}

type OffchainAttestationType map[string][]types.Type
type OffchainAttestationMessage map[string]interface{}

type OffchainAttestationParam struct {
	Domain OffchainAttestationDomain `json:"domain"`

	Message     OffchainAttestationMessage `json:"message"`
	Type        OffchainAttestationType    `json:"types"`
	Signature   Signature                  `json:"signature"`
	PrimaryType string                     `json:"primaryType"`
	Uid         string                     `json:"uid"`
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

func NewBASOffchainAttestation(schemaUid string, schema string, domain OffchainAttestationDomain, data map[string]interface{}, recipient string, revocable bool, refUid string, nonce uint64, time uint64, expirationTime uint64, version uint16, signer *ecdsa.PrivateKey) (*OffchainAttestationParam, error) {

	attest := OffchainAttestationParam{}
	attest.Domain = domain
	message := OffchainAttestationMessage{}

	var m MessageForUid

	if _data, err := EncodeData(schema, data); err != nil {
		return nil, fmt.Errorf("encode data error: " + err.Error())
	} else {
		message["data"] = "0x" + hex.EncodeToString(_data)
		m.Data = "0x" + hex.EncodeToString(_data)
	}

	message["version"] = big.NewInt(int64(version))
	m.Version = strconv.Itoa(int(version))

	message["recipient"] = recipient
	m.Recipient = recipient

	message["expirationTime"] = big.NewInt(int64(expirationTime))
	m.ExpirationTime = expirationTime

	message["time"] = big.NewInt(int64(time))
	m.Time = time

	message["revocable"] = revocable
	m.Revocable = revocable
	message["nonce"] = big.NewInt(int64(nonce))

	message["schema"] = schemaUid
	m.Schema = schemaUid

	if refUid == "" {
		refUid = "0x0000000000000000000000000000000000000000000000000000000000000000"
	}
	message["refUID"] = refUid
	m.RefUID = refUid

	attest.Message = message

	attest.PrimaryType = "Attest"
	atypes := []types.Type{
		{Name: "version", Type: "uint16"},
		{Name: "schema", Type: "bytes32"},
		{Name: "recipient", Type: "address"},
		{Name: "time", Type: "uint64"},
		{Name: "expirationTime", Type: "uint64"},
		{Name: "revocable", Type: "bool"},
		{Name: "refUID", Type: "bytes32"},
		{Name: "data", Type: "bytes"},
		{Name: "nonce", Type: "uint64"},
	}
	attest.Type = map[string][]types.Type{}
	attest.Type["Attest"] = atypes

	if sig, err := Sign(attest.Domain, attest.Type, attest.Message, signer); err != nil {
		return nil, err
	} else {
		if _sig, err := extractSignature(sig); err != nil {
			return nil, err
		} else {
			attest.Signature = _sig
		}
	}
	uid := GetOffChainAttestationUid(m)
	attest.Uid = uid
	return &attest, nil
}

type SingleBundleObject struct {
	Name string
	Data []byte
}
