package bas_go

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	types "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/spf13/cast"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/umbracle/ethgo/abi"
	"golang.org/x/crypto/sha3"
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
	Domain      OffchainAttestationDomain  `json:"domain"`
	Type        OffchainAttestationType    `json:"types"`
	Message     OffchainAttestationMessage `json:"message"`
	Signature   Signature                  `json:"signature"`
	PrimaryType string                     `json:"primaryType"`
}

func CheckOffchainAttestation(attestationStr string, signer string, uid string, resolver string, schema string) (bool, error) {
	var offchainAttestation OffchainAttestationParam
	if err := json.Unmarshal([]byte(attestationStr), &offchainAttestation); err != nil {
		return false, fmt.Errorf("error attestation type")
	}

	return _checkOffchainAttestation(offchainAttestation, signer, uid, resolver, schema)

}

func CheckOffchainAttestationRecSigner(attestationStr string, uid string, resolver string, schema string) (bool, string, error) {
	var offchainAttestation OffchainAttestationParam
	if err := json.Unmarshal([]byte(attestationStr), &offchainAttestation); err != nil {
		return false, "", fmt.Errorf("error attestation type")
	}

	return _checkOffchainAttestationRecSigner(offchainAttestation, uid, resolver, schema)

}

func _checkOffchainAttestationRecSigner(attestation OffchainAttestationParam, uid string, resolver string, schema string) (bool, string, error) {
	//step 1: check uid
	var message MessageForUid
	if err := message.Decode(attestation.Message); err != nil {
		return false, "", fmt.Errorf("err for attestation uid: " + err.Error())
	}
	if recUid := getOffChainAttestationUid(message); recUid != uid {
		return false, "", fmt.Errorf("un-matched attestation uid: " + recUid + " for " + uid)
	}

	//step 2: check the signer
	var signer string
	if recSigner, err := getSigner(attestation.Signature, attestation.Domain, attestation.Type, attestation.Message); err != nil {
		return false, "", err
	} else {
		signer = recSigner
	}

	//step 3: resolver
	//TODO

	//step 4: check the data
	if _data, ok := attestation.Message["data"].(string); !ok {
		return false, "", fmt.Errorf("can not find data " + " for " + schema)
	} else if ok := checkSchema(_data, schema); !ok {
		return false, "", fmt.Errorf("un-matched attestation data: " + _data + " for " + schema)
	}

	return true, signer, nil

}

func _checkOffchainAttestation(attestation OffchainAttestationParam, signer string, uid string, resolver string, schema string) (bool, error) {
	//step 1: check uid
	var message MessageForUid
	if err := message.Decode(attestation.Message); err != nil {
		return false, fmt.Errorf("err for attestation uid: " + err.Error())
	}
	if recUid := getOffChainAttestationUid(message); recUid != uid {
		return false, fmt.Errorf("un-matched attestation uid: " + recUid + " for " + uid)
	}

	//step 2: check the signer
	if recSigner, err := getSigner(attestation.Signature, attestation.Domain, attestation.Type, attestation.Message); err != nil {
		return false, err
	} else {
		if recSigner != signer {
			return false, fmt.Errorf("un-matched signer: " + recSigner + " for " + signer)
		}
	}

	//step 3: resolver
	//TODO

	//step 4: check the data
	if _data, ok := attestation.Message["data"].(string); !ok {
		return false, fmt.Errorf("can not find data " + " for " + schema)
	} else if ok := checkSchema(_data, schema); !ok {
		return false, fmt.Errorf("un-matched attestation data: " + _data + " for " + schema)
	}

	return true, nil

}

const (
	ZEROADDRESS = "0x0000000000000000000000000000000000000000"
)

func getOffChainAttestationUid(param MessageForUid) string {

	_schema := hex.EncodeToString([]byte(param.Schema))
	_schema = "0x" + _schema

	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint16", "bytes", "address", "address", "uint64", "uint64", "bool", "bytes32", "bytes", "uint32"},

		// values
		[]interface{}{
			param.Version,
			_schema,
			param.Recipient,
			ZEROADDRESS,
			param.Time,
			param.ExpirationTime,
			param.Revocable,
			param.RefUID,
			param.Data,
			0,
		},
	)

	recUid := hex.EncodeToString(hash)
	return "0x" + recUid

}

func combineSignature(r, s string, v uint8) ([]byte, error) {
	// 将 r、s 转换为字节数组
	rBytes, err := hex.DecodeString(r[2:])
	if err != nil {
		return nil, err
	}

	sBytes, err := hex.DecodeString(s[2:])
	if err != nil {
		return nil, err
	}

	// 将 v 转换为字节数组
	vBytes := []byte{v - 27}

	// 拼接 r、s、v
	signature := append(rBytes, sBytes...)
	signature = append(signature, vBytes...)

	return signature, nil
}

func publicKeyBytesToAddress(pubkeyBytes []byte) ([]byte, error) {
	// 将字节数组转换为 ecdsa.PublicKey
	pubkey, err := crypto.UnmarshalPubkey(pubkeyBytes)
	if err != nil {
		return nil, err
	}

	// 计算以太坊地址
	address := publicKeyToAddress(*pubkey)
	return address, nil
}

func publicKeyToAddress(pubkey ecdsa.PublicKey) []byte {
	// 以太坊地址是公钥的最后 20 个字节的 Keccak256 哈希值
	pubBytes := crypto.FromECDSAPub(&pubkey)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubBytes[1:])      // 去掉公钥前缀字节 0x04
	address := hash.Sum(nil)[12:] // 取后 20 个字节
	return address
}

func getSigner(asign Signature, adomain OffchainAttestationDomain, atype OffchainAttestationType, amessage OffchainAttestationMessage) (string, error) {

	// typesStr := `{"Attest":[{"name":"version","type":"uint16"},{"name":"schema","type":"bytes32"},{"name":"recipient","type":"address"},{"name":"time","type":"uint64"},{"name":"expirationTime","type":"uint64"},{"name":"revocable","type":"bool"},{"name":"refUID","type":"bytes32"},{"name":"data","type":"bytes"},{"name":"nonce","type":"uint256"}]}`
	// typesStr := `{"Attest":[
	// 	{"name":"version","type":"uint16"},
	// 	{"name":"schema","type":"bytes32"},
	// 	{"name":"recipient","type":"address"},
	// 	{"name":"time","type":"uint64"},
	// 	{"name":"expirationTime","type":"uint64"},
	// 	{"name":"revocable","type":"bool"},
	// 	{"name":"refUID","type":"bytes32"},
	// 	{"name":"data","type":"bytes"}]}`
	domainTypeStr := `{"EIP712Domain":[
		{"name":"name", "type":"string"},
		{"name":"version","type":"string"},
		{"name":"chainId","type":"uint256"},
		{"name":"verifyingContract","type":"address"}]}`

	// myType := map[string][]types.Type{}
	// if err := json.Unmarshal([]byte(typesStr), &myType); err != nil {
	// 	panic(err)
	// }
	myType := atype

	extraType := map[string][]types.Type{}
	if err := json.Unmarshal([]byte(domainTypeStr), &extraType); err != nil {
		panic(err)
	}

	myType["EIP712Domain"] = extraType["EIP712Domain"]

	data := types.TypedData{}
	data.Types = types.Types(myType)

	chainId, err := strconv.Atoi(adomain.ChainId)
	if err != nil {
		return "", fmt.Errorf("error chainId: %d", chainId)
	}
	domain := types.TypedDataDomain{
		Name:              adomain.Name,
		Version:           adomain.Version,
		ChainId:           math.NewHexOrDecimal256(int64(chainId)),
		VerifyingContract: adomain.VerifyingContract,
	}
	data.Domain = domain

	data.Message = amessage

	data.PrimaryType = "Attest"

	hash, _, err := types.TypedDataAndHash(data)
	if err != nil {
		return "", fmt.Errorf("meet err when encode data: " + err.Error())
	}

	sig, err := combineSignature(asign.R, asign.S, asign.V)
	if err != nil {
		return "", fmt.Errorf("meet error when combine sig: " + err.Error())
	}

	sigPublicKey, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return "", fmt.Errorf("meet error when recover pubkey: " + err.Error())
	}

	address, err := publicKeyBytesToAddress(sigPublicKey)
	if err != nil {
		return "", fmt.Errorf("meet error when recover address from: " + string(sigPublicKey))
	}
	addr := common.BytesToAddress(address).String()
	return addr, nil
}

func checkSchema(dataStr string, schema string) bool {
	_schema := fmt.Sprintf("tuple(%s)", schema)
	fmt.Println(_schema)
	typ := abi.MustNewType(_schema)
	if dataStr[:2] == "0x" {
		dataStr = dataStr[2:]
	}
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return false
	}

	v, err := typ.Decode(data)
	fmt.Println(v, err)
	if err != nil {
		return false
	}
	return true

}
