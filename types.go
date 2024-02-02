package bas_go

import (
	"fmt"
	"strconv"

	types "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/spf13/cast"
)

const (
	BASDOMAINNAME = "BAS ATTESTATION"
	ZEROADDRESS   = "0x0000000000000000000000000000000000000000"
)

type OffchainAttestationType map[string][]types.Type
type OffchainAttestationMessage map[string]interface{}

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

type OffchainAttestationDomain struct {
	Name              string
	Version           string
	ChainId           string
	VerifyingContract string
}

type Signature struct {
	R string
	S string
	V uint8
}

type OffchainAttestationParam struct {
	Domain OffchainAttestationDomain `json:"domain"`

	Message     OffchainAttestationMessage `json:"message"`
	Type        OffchainAttestationType    `json:"types"`
	Signature   Signature                  `json:"signature"`
	PrimaryType string                     `json:"primaryType"`
}
