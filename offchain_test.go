package bas_go

import (
	"fmt"
	"testing"
)

// NOTE!!!!!!!!!
// POINT ONE: MODIFY go/src/github.com/go-ethereum/signer/core/apitypes/types.go to remove line 347/348 like:
// *************************
// Verify extra data
//
//	if exp, got := len(typedData.Types[primaryType]), len(data); exp < got {
//		return nil, fmt.Errorf("there is extra data provided in the message (%d < %d)", exp, got)
//	}
//
// *************************

// POINT TWO: MODIFY go/pkg/mod/github.com/miguelmota/go-solidity-sha3@v0.1.1/utils.go to add at line 43 like:
// *************************
// switch typ {
// case "address":
// 	if _isArray {
// 		return padZeros(Address(value), 32)
// 	}

// 	return Address(value)
// case "bytes":
// 	bytes, err := hex.DecodeString(value.(string)[2:])
// 	if err != nil {
// 		panic(err)
// 	}
// 	return bytes

// case "string":
// 	return String(value)
// case "bool":
// 	if _isArray {
// 		return padZeros(Bool(value), 32)
// 	}

//		return Bool(value)
//	}
//
// *************************
// Example 1: PADO
func Test_checkOffchainAttestation(t *testing.T) {
	attestationStr := `{"types":{"Attest":[{"name":"schema","type":"bytes32"},{"name":"recipient","type":"address"},{"name":"expirationTime","type":"uint64"},{"name":"revocable","type":"bool"},{"name":"refUID","type":"bytes32"},{"name":"data","type":"bytes"},{"name":"deadline","type":"uint64"}],"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}]},"primaryType":"Attest","message":{"schema":"0x5f868b117fd34565f3626396ba91ef0c9a607a0e406972655c5137c6d4291af9","recipient":"0x024e45d7f868c41f3723b13fd7ae03aa5a181362","expirationTime":0,"revocable":true,"data":"0x00000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c025deb5f92761aa19e8f8b872f11a9e08b6b8d9bf58aef0365a0ccde9bb92eab900000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000018cd756d28c8fc677e1cd95db805d6784c88c20ba812d254657777ee80828f49bf8670d15ae00000000000000000000000000000000000000000000000000000000000000084964656e746974790000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006474f4f474c45000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000114163636f756e74204f776e65727368697000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000085665726966696564000000000000000000000000000000000000000000000000","refUID":"0x0000000000000000000000000000000000000000000000000000000000000000","deadline":0,"time":"1704419840652","version":1},"domain":{"name":"PermissionedEIP712Proxy","version":"0.1","chainId":"97","verifyingContract":"0x620e84546d71A775A82491e1e527292e94a7165A","salt":null},"uid":"0x0a63960e61305942570547f9ad8def43283a7dcff082eb8dee083008b83a5a7f","signature":{"v":27,"r":"0x022206869ae40e84e0234bd73c655c1fe0a634defd740be9098bde59dcafd11d","s":"0x1182c9eaa8669d2e078fcf9a19cc1d70f26dce450ea2f86b40a3462b15af623d"}}`
	if pass, signer, err := CheckOffchainAttestationRecSigner(
		attestationStr,
		"0x0a63960e61305942570547f9ad8def43283a7dcff082eb8dee083008b83a5a7f",
		"",
		"string ProofType,string Source,string Content,string Condition,bytes32 SourceUserIdHash,bool Result,uint64 Timestamp,bytes32 UserIdHash"); err != nil {
		panic(err)
	} else if pass {
		fmt.Println("pass !")
		fmt.Println(signer)
	}
}

// Example 2: BAS
func Test_checkOffchainAttestation2(t *testing.T) {
	attestationStr := `{"domain":{"name":"EAS Attestation","version":"1.3.0","chainId":"97","verifyingContract":"0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"},"primaryType":"Attest","message":{"recipient":"0x0000000000000000000000000000000000000000","expirationTime":"0","time":"1704301232","revocable":true,"version":1,"nonce":"0","schema":"0xcb86ea930c2fde4952fe64237575b62903a353e4724174fd272d2fc4053165dc","refUID":"0x0000000000000000000000000000000000000000000000000000000000000000","data":"0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000037472790000000000000000000000000000000000000000000000000000000000"},"types":{"Attest":[{"name":"version","type":"uint16"},{"name":"schema","type":"bytes32"},{"name":"recipient","type":"address"},{"name":"time","type":"uint64"},{"name":"expirationTime","type":"uint64"},{"name":"revocable","type":"bool"},{"name":"refUID","type":"bytes32"},{"name":"data","type":"bytes"}]},"signature":{"v":28,"r":"0xc4c47d41380c3deadd91f4eb9db899fa517174e816a1bd1d9d2a547c76f33547","s":"0x41219a7d0776fe5b793117e91b7673281fda48fd36defa111b7a2b4578fc1f57"},"uid":"0x0c73ec7475661e62f85aa8958c0f59507b7559595930d658f79dd3d7957a1561"}`
	if pass, err := CheckOffchainAttestation(
		attestationStr,
		"0x471543A3bd04486008c8a38c5C00543B73F1769e",
		"0x0c73ec7475661e62f85aa8958c0f59507b7559595930d658f79dd3d7957a1561",
		"",
		"string ProofType,string Source,string Content,string Condition,bytes32 SourceUserIdHash,bool Result,uint64 Timestamp,bytes32 UserIdHash"); err != nil {
		panic(err)
	} else if pass {
		fmt.Println("pass !")
	}
}
