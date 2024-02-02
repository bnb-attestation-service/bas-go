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
	"github.com/umbracle/ethgo/abi"
	"golang.org/x/crypto/sha3"
)

func getSigner(asign Signature, adomain OffchainAttestationDomain, atype OffchainAttestationType, amessage OffchainAttestationMessage) (string, error) {
	if adomain.Name != BASDOMAINNAME {
		return "", fmt.Errorf("not a bas attestation sig")
	}

	if T, ok := atype["Attest"]; !ok {
		return "", fmt.Errorf("not a bas attestation sig: no attest type")
	} else {
		for _, typ := range T {
			if _, ok := amessage[typ.Name]; !ok {
				return "", fmt.Errorf("has mismatch type and value in message: " + typ.Name)
			}
		}
		if len(T) != len(amessage) {
			return "", fmt.Errorf("message value and type in bas must have same len")
		}
	}

	domainTypeStr := `{"EIP712Domain":[
		{"name":"name", "type":"string"},
		{"name":"version","type":"string"},
		{"name":"chainId","type":"uint256"},
		{"name":"verifyingContract","type":"address"}]}`

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
