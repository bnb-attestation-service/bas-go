package onchain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	types "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func Sign(adomain OnchainAttestationDomain, atype OnchainAttestationType, amessage OnchainAttestationMessage, privKey *ecdsa.PrivateKey) ([]byte, error) {
	// if adomain.Name != BASDOMAINNAME {
	// 	return nil, fmt.Errorf("not a bas attestation sig")
	// }

	if T, ok := atype["Attest"]; !ok {
		return nil, fmt.Errorf("not a bas attestation sig: no attest type")
	} else {
		for _, typ := range T {
			if _, ok := amessage[typ.Name]; !ok {
				return nil, fmt.Errorf("has mismatch type and value in message: " + typ.Name)
			}
		}
		if len(T) != len(amessage) {
			return nil, fmt.Errorf("message value and type in bas must have same len")
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
		return nil, fmt.Errorf("error chainId: %d", chainId)
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
		return nil, fmt.Errorf("meet err when encode data: " + err.Error())
	}

	if sign, err := crypto.Sign(hash, privKey); err != nil {
		return nil, fmt.Errorf("get invalid sign")
	} else {
		return sign, nil
	}

}

func extractSignature(signature []byte) (sig Signature, err error) {
	// signature 的长度必须至少为 65 字节
	if len(signature) < 65 {
		err = errors.New("invalid signature length")
		return
	}

	// 从 signature 中提取 r、s 和 v
	rBytes := signature[:32]
	sBytes := signature[32:64]
	vByte := signature[64]

	// 将 r、s 转换为十六进制字符串
	sig.R = "0x" + hex.EncodeToString(rBytes)
	sig.S = "0x" + hex.EncodeToString(sBytes)

	// 计算 v
	sig.V = vByte + 27

	return
}
