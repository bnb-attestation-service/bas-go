package offchain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	types "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"golang.org/x/crypto/sha3"
)

func Sign(adomain OffchainAttestationDomain, atype OffchainAttestationType, amessage OffchainAttestationMessage, privKey *ecdsa.PrivateKey) ([]byte, error) {
	if adomain.Name != BASDOMAINNAME {
		return nil, fmt.Errorf("not a bas attestation sig")
	}

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
func GetSigner(asign Signature, adomain OffchainAttestationDomain, atype OffchainAttestationType, amessage OffchainAttestationMessage) (string, error) {
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
	rBytes, err := hex.DecodeString(r[2:])
	if err != nil {
		return nil, err
	}

	sBytes, err := hex.DecodeString(s[2:])
	if err != nil {
		return nil, err
	}

	vBytes := []byte{v - 27}

	signature := append(rBytes, sBytes...)
	signature = append(signature, vBytes...)

	return signature, nil
}

func extractSignature(signature []byte) (sig Signature, err error) {

	if len(signature) < 65 {
		err = errors.New("invalid signature length")
		return
	}

	rBytes := signature[:32]
	sBytes := signature[32:64]
	vByte := signature[64]

	sig.R = "0x" + hex.EncodeToString(rBytes)
	sig.S = "0x" + hex.EncodeToString(sBytes)

	sig.V = vByte + 27

	return
}

func publicKeyBytesToAddress(pubkeyBytes []byte) ([]byte, error) {

	pubkey, err := crypto.UnmarshalPubkey(pubkeyBytes)
	if err != nil {
		return nil, err
	}

	address := publicKeyToAddress(*pubkey)
	return address, nil
}

func publicKeyToAddress(pubkey ecdsa.PublicKey) []byte {

	pubBytes := crypto.FromECDSAPub(&pubkey)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubBytes[1:])
	address := hash.Sum(nil)[12:]
	return address
}
