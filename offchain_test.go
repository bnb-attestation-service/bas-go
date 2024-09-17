// package main

// import (
// 	"fmt"
// 	"testing"
// )

// // NOTE!!!!!!!!!
// // POINT ONE: MODIFY go/src/github.com/go-ethereum/signer/core/apitypes/types.go to remove line 347/348 like:
// // *************************
// // Verify extra data
// //
// //	if exp, got := len(typedData.Types[primaryType]), len(data); exp < got {
// //		return nil, fmt.Errorf("there is extra data provided in the message (%d < %d)", exp, got)
// //	}
// //
// // *************************

// // POINT TWO: MODIFY go/pkg/mod/github.com/miguelmota/go-solidity-sha3@v0.1.1/utils.go to add at line 43 like:
// // *************************
// // switch typ {
// // case "address":
// // 	if _isArray {
// // 		return padZeros(Address(value), 32)
// // 	}

// // 	return Address(value)
// // case "bytes":
// // 	bytes, err := hex.DecodeString(value.(string)[2:])
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	return bytes

// // case "string":
// // 	return String(value)
// // case "bool":
// // 	if _isArray {
// // 		return padZeros(Bool(value), 32)
// // 	}

// //		return Bool(value)
// //	}
// //
// // *************************
// // Example 1: PADO
// func Test_checkOffchainAttestation(t *testing.T) {
// 	attestationStr := `{
// 		"version": 2,
// 		"uid": "0xc5ae92a79c7ca46bde999f628d27aea1a910e4af825541318115b84c8c6e2a55",
// 		"domain": {
// 		  "name": "BAS Attestation",
// 		  "version": "1.3.0",
// 		  "chainId": "97",
// 		  "verifyingContract": "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"
// 		},
// 		"primaryType": "Attest",
// 		"message": {
// 		  "version": 1,
// 		  "recipient": "0xf5eaB61f3A738A4B37d5d9B95f638a0A7db5Cb5f",
// 		  "expirationTime": 0,
// 		  "time": 1707419147,
// 		  "revocable": true,
// 		  "nonce": 0,
// 		  "schema": "0x3969bb076acfb992af54d51274c5c868641ca5344e1aacd0b1f5e4f80ac0822f",
// 		  "refUID": "0x0000000000000000000000000000000000000000000000000000000000000000",
// 		  "data": "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000b48656c6c6f2c2045415321000000000000000000000000000000000000000000",
// 		  "salt": "0x50bd672cccbf7e1450067cd4af7a441ace046692a22f2453559bbd261125cf29"
// 		},
// 		"types": {
// 		  "Attest": [
// 			{ "name": "version", "type": "uint16" },
// 			{ "name": "schema", "type": "bytes32" },
// 			{ "name": "recipient", "type": "address" },
// 			{ "name": "time", "type": "uint64" },
// 			{ "name": "expirationTime", "type": "uint64" },
// 			{ "name": "revocable", "type": "bool" },
// 			{ "name": "refUID", "type": "bytes32" },
// 			{ "name": "data", "type": "bytes" },
// 			{ "name": "salt", "type": "bytes32" },
// 			{ "name": "nonce", "type": "uint64" }
// 		  ]
// 		},
// 		"signature": {
// 		  "v": 27,
// 		  "r": "0x9c2bb4d7883a6df9466faada5b054692af44cb91e5a0862c13e1ebe380412d86",
// 		  "s": "0x30773bd71491aadbf0b2ba04796116fbc998ee4264848ef85997abc26dad1b12"
// 		}
// 	  }`
// 	if pass, signer, err := CheckOffchainAttestationRecSigner(
// 		attestationStr,
// 		"0xc5ae92a79c7ca46bde999f628d27aea1a910e4af825541318115b84c8c6e2a55",
// 		"",
// 		"string ProofType,string Source,string Content,string Condition,bytes32 SourceUserIdHash,bool Result,uint64 Timestamp,bytes32 UserIdHash"); err != nil {
// 		panic(err)
// 	} else if pass {
// 		fmt.Println("pass !")
// 		fmt.Println(signer)
// 	}
// }

// // Example 2: BAS
// func Test_checkOffchainAttestation2(t *testing.T) {
// 	attestationStr := `{"domain":{"name":"EAS Attestation","version":"1.3.0","chainId":"97","verifyingContract":"0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"},"primaryType":"Attest","message":{"recipient":"0x0000000000000000000000000000000000000000","expirationTime":"0","time":"1704301232","revocable":true,"version":1,"nonce":"0","schema":"0xcb86ea930c2fde4952fe64237575b62903a353e4724174fd272d2fc4053165dc","refUID":"0x0000000000000000000000000000000000000000000000000000000000000000","data":"0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000037472790000000000000000000000000000000000000000000000000000000000"},"types":{"Attest":[{"name":"version","type":"uint16"},{"name":"schema","type":"bytes32"},{"name":"recipient","type":"address"},{"name":"time","type":"uint64"},{"name":"expirationTime","type":"uint64"},{"name":"revocable","type":"bool"},{"name":"refUID","type":"bytes32"},{"name":"data","type":"bytes"}]},"signature":{"v":28,"r":"0xc4c47d41380c3deadd91f4eb9db899fa517174e816a1bd1d9d2a547c76f33547","s":"0x41219a7d0776fe5b793117e91b7673281fda48fd36defa111b7a2b4578fc1f57"},"uid":"0x0c73ec7475661e62f85aa8958c0f59507b7559595930d658f79dd3d7957a1561"}`
// 	if pass, err := CheckOffchainAttestation(
// 		attestationStr,
// 		"0x471543A3bd04486008c8a38c5C00543B73F1769e",
// 		"0x0c73ec7475661e62f85aa8958c0f59507b7559595930d658f79dd3d7957a1561",
// 		"",
// 		"string ProofType,string Source,string Content,string Condition,bytes32 SourceUserIdHash,bool Result,uint64 Timestamp,bytes32 UserIdHash"); err != nil {
// 		panic(err)
// 	} else if pass {
// 		fmt.Println("pass !")
// 	}
// }

// func Test_abi(t *testing.T) {

// }
package agent

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/bnb-attestation-service/bas-go/offchain"
)

func TestCreateOffchainAttestation(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"p":     "1212",
		"tick":  "1212",
		"amt":   10,
		"nonce": 10,
		"vote":  0,
	}

	if res, err := _agent.OffchainNewAttestation(
		"0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f",
		"string p,string tick,uint256 amt,uint8 vote,uint256 nonce",
		data,
		"0x16abBD7f92CDF1703beb6D314885d2a79B0497fb",
		false,
		"",
		0,
		1703255628,
		0,
		1,
	); err != nil {
		panic(err)
	} else {
		if _b, err := json.Marshal(res); err != nil {
			panic(err)
		} else {
			fmt.Println(string(_b))
		}

	}
}

func Test_checkOffchainAttestation(t *testing.T) {
	attestationStr := `{
		"domain": {
		  "Name": "BAS Attestation",
		  "Version": "1.3.0",
		  "ChainId": "97",
		  "VerifyingContract": "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"
		},
		"message": {
		  "data": "0x00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000004313231320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000043132313200000000000000000000000000000000000000000000000000000000",
		  "expirationTime": 0,
		  "nonce": 0,
		  "recipient": "0x16abBD7f92CDF1703beb6D314885d2a79B0497fb",
		  "refUID": "0x0000000000000000000000000000000000000000000000000000000000000000",
		  "revocable": false,
		  "schema": "0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f",
		  "time": 1703255628,
		  "version": 1
		},
		"types": {
		  "Attest": [
			{
			  "name": "version",
			  "type": "uint16"
			},
			{
			  "name": "schema",
			  "type": "bytes32"
			},
			{
			  "name": "recipient",
			  "type": "address"
			},
			{
			  "name": "time",
			  "type": "uint64"
			},
			{
			  "name": "expirationTime",
			  "type": "uint64"
			},
			{
			  "name": "revocable",
			  "type": "bool"
			},
			{
			  "name": "refUID",
			  "type": "bytes32"
			},
			{
			  "name": "data",
			  "type": "bytes"
			},
			{
			  "name": "nonce",
			  "type": "uint64"
			}
		  ],
		  "EIP712Domain": [
			{
			  "name": "name",
			  "type": "string"
			},
			{
			  "name": "version",
			  "type": "string"
			},
			{
			  "name": "chainId",
			  "type": "uint256"
			},
			{
			  "name": "verifyingContract",
			  "type": "address"
			}
		  ]
		},
		"signature": {
		  "R": "0x205845584be9c05caebecc345c88ff1e3ef9b3bf2e1600d5fb8bde91a8f2e8f1",
		  "S": "0x7ce08ca0895367313062c3fece0435ace2d7d312eb02fb6ac3c58105f904305c",
		  "V": 28
		},
		"primaryType": "Attest",
		"uid": "0x461d59aa0d460c55014c10851fc166a220a2bf898f71fe77c1ff19a443ae1d8e"
	  }`
	if pass, signer, err := offchain.CheckOffchainAttestationRecSigner(
		attestationStr,
		"0x461d59aa0d460c55014c10851fc166a220a2bf898f71fe77c1ff19a443ae1d8e",
		"",
		"string p,string tick,uint256 amt,uint8 vote,uint256 nonce"); err != nil {
		panic(err)
	} else if pass {
		fmt.Println("pass !")
		fmt.Println(signer)
	}
}

func TestUploadToGF(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	_agent.ConfigBucket("bas-90498da77d3ab65e3f2589f0e7ea515266a80a40")
	data := map[string]interface{}{
		"p":     "1212",
		"tick":  "1212",
		"amt":   10,
		"nonce": 10,
		"vote":  0,
	}

	if res, err := _agent.OffchainNewAttestation(
		"0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f",
		"string p,string tick,uint256 amt,uint8 vote,uint256 nonce",
		data,
		"0x16abBD7f92CDF1703beb6D314885d2a79B0497fb",
		false,
		"0x0000000000000000000000000000000000000000000000000000000000000000",
		0,
		1703255628,
		0,
		3,
	); err != nil {
		panic(err)
	} else {
		if hash, err := _agent.OffchainUploadAttestationToGF(res); err != nil {
			panic(err)
		} else {
			fmt.Println(hash)
		}

	}
}

func TestCreateBucket(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if err := _agent.CreateBucket(); err != nil {
		panic(err)
	}

}

func TestPublicAttestation(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	_agent.ConfigBucket("bas-90498da77d3ab65e3f2589f0e7ea515266a80a40")
	if hash, err := _agent.OffchainChangeAttestationVisible("0x6fafa31aec106f515e021d367563ba2f3feb1d99cb302fe5a25723747d1be356", "0xcde3a9daa0c6d42753ce8a21ffb1764decbe2fbdd7e7bfc328acfe89a970bfc1", true); err != nil {
		panic(err)
	} else {
		t.Log(hash)
	}

}

func TestBundleToGF(t *testing.T) {

	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}

	var datas []offchain.SingleBundleObject
	for i := 0; i < 5; i++ {
		var data offchain.SingleBundleObject
		data.Name = strconv.Itoa(i)
		data.Data = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(i)}
		datas = append(datas, data)
	}

	tx, err := _agent.OffchainUploadBundleToGF(datas, "test-data", "bas-bundle")
	if err != nil {
		panic(err)
	}
	fmt.Println(tx)

}

func TestCheckWritePermission(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	fmt.Println(_agent.CheckWritePermission("bas"))
}
