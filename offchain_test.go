package agent

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/bnb-attestation-service/bas-go/offchain"
)

func TestCreateOffchainAttestation(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
		OFFCHAINBASTESTDOMAIN,
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
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
		OFFCHAINBASTESTDOMAIN,
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
		if hash, err := _agent.OffchainUploadAttestationToGF(res, true); err != nil {
			panic(err)
		} else {
			fmt.Println(hash)
		}

	}
}

func TestCreateBucket(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if _, err := _agent.CreateBucket("0xccC793c4D92f7c425Ef5C2b418b9186ad180708d"); err != nil {
		panic(err)
	}

}

func TestPublicAttestation(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	fmt.Println(_agent.CheckWritePermission("bas"))
}

func TestDownloadBundle(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	fmt.Println(_agent.OffchainDownloadBundle("bas-bundle", "bundle.0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f.0x8db66dda4b46008695f4dcab09245a3b2694da353da17ebe58ca29f79887a9dd", ""))
}

func TestOffchainMultiAttest(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}

	var attestations []*offchain.OffchainAttestationParam
	for _, _amount := range []int{10, 20, 30, 40} {
		data := map[string]interface{}{
			"p":     "1212",
			"tick":  "1212",
			"amt":   _amount,
			"nonce": _amount,
			"vote":  0,
		}

		if res, err := _agent.OffchainNewAttestation(
			"0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f",
			OFFCHAINBASTESTDOMAIN,
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
			attestations = append(attestations, res)
		}
	}
	schemaUid := "0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f"
	bucket := "bas-bundle"
	fmt.Println(_agent.OffchainMultiAttestByBundle(attestations, schemaUid, bucket))

}

func TestOffchainParseAttestationsFromBundle(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	fmt.Println(_agent.GetAddress())
	bundle := "/var/folders/nf/z40nschs2b5dkhzm7d9mrt3m0000gn/T/bundle-741682584"
	bundleName := "bundle.0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f.0x8db66dda4b46008695f4dcab09245a3b2694da353da17ebe58ca29f79887a9dd"
	attestations, data, err := _agent.OffchainParseAttestationsFromBundle(bundle, bundleName)
	if err != nil {
		panic(err)
	}
	for k, v := range attestations {
		fmt.Println(k)
		fmt.Println(v)
	}
	os.WriteFile("att.json", data[0], 0777)
}

func TestBundleCapasity(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	fmt.Println(_agent.GetAddress())

	for {
		var attestations []*offchain.OffchainAttestationParam
		bundleSize := 100
		for _amount := 0; _amount < bundleSize; _amount++ {
			data := map[string]interface{}{
				"p":     "1212",
				"tick":  "1212",
				"amt":   _amount,
				"nonce": _amount,
				"vote":  0,
			}

			if res, err := _agent.OffchainNewAttestation(
				"0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f",
				OFFCHAINBASTESTDOMAIN,
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
				attestations = append(attestations, res)
			}
		}
		schemaUid := "0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f"
		bucket := "bas-bundle"
		fmt.Println(_agent.OffchainMultiAttestByBundle(attestations, schemaUid, bucket))
	}

}
