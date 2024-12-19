package agent

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	storageTypes "github.com/bnb-chain/greenfield/x/storage/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/umbracle/ethgo/abi"
	"math/rand"

	"strconv"
	"testing"

	"github.com/bnb-attestation-service/bas-go/offchain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
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

	tx, err := _agent.OffchainUploadBundleToGF(datas, "test-data", "bas-bundle", storageTypes.VISIBILITY_TYPE_PUBLIC_READ)
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
	fmt.Println(_agent.OffchainMultiAttestByBundle(attestations, schemaUid, bucket, storageTypes.VISIBILITY_TYPE_PUBLIC_READ))

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
	attestations, _, err := _agent.OffchainParseAttestationsFromBundle(bundle, bundleName)
	if err != nil {
		panic(err)
	}
	for k, v := range attestations {
		fmt.Println(k)
		fmt.Println(v)
	}
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
		fmt.Println(_agent.OffchainMultiAttestByBundle(attestations, schemaUid, bucket, storageTypes.VISIBILITY_TYPE_PUBLIC_READ))
	}

}

func TestName(t *testing.T) {
	rdb, err := gorm.Open(postgres.Open("postgres://postgres:Gbs1767359487@bas-instance-1.ccggmi9astti.us-east-1.rds.amazonaws.com/op_bas_mainnet"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}

	var _agent *Agent
	if _agent, err = NewAgentFromKey(privateKey, BAS, SCHEAMA, BSCRPC, BSCCHAINID, GFRPC, GFCHAINID); err != nil {
		panic(err)
	}
	schemaUID := "0x6adca6e46080fcbaae1a1b19592fc88f20845bfbd78f284a360b05b442cc3a82"
	schema := "bytes32 uHash,string source,bytes32 publicDataHash,bool followBASTwitter"
	bucket := "bas-0x0a39809058b35a5068541d892194952963516025"
	datas := getHash(schemaUID, schema)
	limit := 1181
	repeat := 9

	for i := 0; i < repeat; i++ {
		var pks []string
		sql := `select key from pks offset ? limit ?`
		if err = rdb.Raw(sql, i*1000, limit).Scan(&pks).Error; err != nil {
			panic(err)
		}

		var attestations []*offchain.OffchainAttestationParam
		for j, pk := range pks {

			randomNumber := rand.Intn(1734278400-1732983880) + 1732983880

			priKey, err := crypto.HexToECDSA(pk[2:])
			if err != nil {
				panic(err)
			}
			address := crypto.PubkeyToAddress(*(priKey.Public().(*ecdsa.PublicKey)))
			att, err := _agent.OffchainNewAttestation(schemaUID, offchain.OffchainAttestationDomain{
				Name:              BASDOMAIN.Name,
				Version:           BASDOMAIN.Version,
				ChainId:           BASDOMAIN.ChainId,
				VerifyingContract: BASDOMAIN.VerifyingContract,
			}, datas[i%len(datas)], address.String(), false, "", 0, uint64(randomNumber), 0, 0)
			if err != nil {
				panic(err)
			}
			attestations = append(attestations, att)
			t.Logf("no: %d", i*len(pks)+j)
		}
		hash, err := _agent.OffchainMultiAttestByBundle(attestations, schemaUID, bucket, storageTypes.VISIBILITY_TYPE_PUBLIC_READ)
		t.Log(hash, err)
	}
}

func TestName1(t *testing.T) {
	schemaUID := "0x6adca6e46080fcbaae1a1b19592fc88f20845bfbd78f284a360b05b442cc3a82"
	rdb, err := gorm.Open(postgres.Open("postgres://postgres:Gbs1767359487@bas-instance-1.ccggmi9astti.us-east-1.rds.amazonaws.com/bas_mainnet"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}

	sql := "select uid,timestamps from attestation where schema_uid = ?"
	var result []struct {
		Uid        string
		Timestamps int64
	}

	err = rdb.Raw(sql, schemaUID).Scan(&result).Error
	if err != nil {
		panic(err)
	}

	times := make(map[int64]int)

	for _, r := range result {
		times[r.Timestamps]++
	}
	for t, i := range times {
		fmt.Println(t, ":", i)
	}
	t.Log(len(result))
}

func getHash(schemaId, schema string) []map[string]interface{} {
	_schema := fmt.Sprintf("tuple(%s)", schema)

	typ := abi.MustNewType(_schema)

	rdb, err := gorm.Open(postgres.Open("postgres://postgres:Gbs1767359487@bas-instance-1.ccggmi9astti.us-east-1.rds.amazonaws.com/bas_mainnet"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(err)
	}

	var data [][]byte

	sql := `select raw_data from attestation where schema_uid = ?`
	err = rdb.Raw(sql, schemaId).Scan(&data).Error
	if err != nil {
		panic(data)
	}

	var result []map[string]interface{}
	for _, _data := range data {
		var offchainAtt offchain.OffchainAttestationParam
		if err = json.Unmarshal(_data, &offchainAtt); err != nil {
			panic(err)
		}

		dataStr := offchainAtt.Message["data"].(string)
		dataByte := common.Hex2Bytes(dataStr[2:])
		v, err := typ.Decode(dataByte)
		if err != nil {
			panic(err)
		}
		result = append(result, v.(map[string]interface{}))
	}
	return result
}
