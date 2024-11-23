package agent

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/bnb-attestation-service/bas-go/eip712"

	"github.com/bnb-attestation-service/bas-go/eas"
	"github.com/bnb-attestation-service/bas-go/onchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestAttest(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(TESTBAS)

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(97))
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	contract, err := eas.NewEAS(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := hex.DecodeString("85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61")
	if err != nil {
		panic(err)
	}
	req := eas.AttestationRequest{
		Schema: [32]byte(schema),
		Data: eas.AttestationRequestData{
			ExpirationTime: 0,
			Revocable:      true,
			Data:           nil,
			Value:          big.NewInt(0),
		},
	}
	if tx, err := contract.Attest(auth, req); err != nil {
		panic(err)
	} else {
		fmt.Println(tx.Hash().Hex())
	}

}

func TestOnchainAttest(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	_agent.OnchainAttest(
		"0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61",
		"0x471543A3bd04486008c8a38c5C00543B73F1769e",
		"0x0000000000000000000000000000000000000000000000000000000000000000",
		nil,
		false,
		0,
		0,
		0,
	)
}
func TestGetAttestation(t *testing.T) {
	uid := "0x02a902caa703b0af896614c443265844d50b82ba2c321bb01232e7b34d8a1d19"
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if attest, err := _agent.OnchainGetAttestation(uid); err != nil {
		panic(err)
	} else {
		fmt.Println(attest)
	}
}

func TestOnchainRevokeAttestation(t *testing.T) {
	schema := "0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61"
	uid := "0x10e25590023c3fcdb0aaa1429712a67399328c53e2cfb8658348658e9f07d694"
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, TESTBAS, TESTSCHEMA, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if attest, err := _agent.OnchainRevoke(schema, uid); err != nil {
		panic(err)
	} else {
		fmt.Println(attest)
	}
}

func waitTx(conn *ethclient.Client, tx *types.Transaction) (string, error) {
	receipt, err := bind.WaitMined(context.Background(), conn, tx)
	if err != nil {
		return "", err
	}
	return receipt.TxHash.String(), nil
}
func TestOnchainDelegate(t *testing.T) {

	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, OPBNBTESTBAS, OPBNBTESTSCHEMA, OPBNBTESTRPC, OPBNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	fmt.Println(_agent.contract.Eip712Domain(_agent.callOp))
	d, _ := _agent.contract.GetDomainSeparator(_agent.callOp)
	fmt.Println(hex.EncodeToString(d[:]))
	data := map[string]interface{}{
		"gm": true,
	}

	var param onchain.OnchainDelegateAttestationParam
	param.Attestor = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.Data = data
	param.Deadline = 1721552480
	param.ExpirationTime = 1721552480
	param.Recipient = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.RefUid = "0x0000000000000000000000000000000000000000000000000000000000000000"
	param.Revocable = true
	param.Schema = "bool gm"
	param.SchemaUid = "0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61"
	param.Value = "0"

	if res, err := _agent.OnchainSignDelegateAttestation(
		param,
		OPBNBTESTDOMAIN,
	); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
		fmt.Println("nonce")
		fmt.Println(_agent.contract.GetNonce(_agent.callOp, common.HexToAddress("471543A3bd04486008c8a38c5C00543B73F1769e")))
		// s := "0xfeb2925a02bae3dae48d424a0437a2b6ac939aa9230ddc55a1a76f065d988076"
		// fmt.Println(hex.DecodeString(s[2:]))
		// fmt.Println(res)
		// fmt.Println(json.Marshal(res))

		// var _input eas.AttestationRequest
		// _input.Data = eas.AttestationRequestData(res.Data)
		// _input.Schema = res.Schema
		// if r, err := _agent.contract.Attest(_agent.txOp, _input); err != nil {
		// 	panic(err)
		// } else {
		// 	fmt.Println(waitTx(_agent.evmClient, r))
		// }
		var _value eas.DelegatedAttestationRequest
		_value.Attester = res.Attester
		_value.Data = eas.AttestationRequestData(res.Data)
		_value.Deadline = res.Deadline
		_value.Schema = res.Schema
		_value.Signature = eas.Signature{
			V: res.Signature.V,
			R: common.HexToHash(res.Signature.R),
			S: common.HexToHash(res.Signature.S),
		}

		fmt.Println(_value)
		if tx, err := _agent.contract.AttestByDelegation(_agent.txOp, eas.DelegatedAttestationRequest{
			Schema:    res.Schema,
			Data:      eas.AttestationRequestData(res.Data),
			Signature: _value.Signature,
			Attester:  res.Attester,
			Deadline:  res.Deadline,
		}); err != nil {
			panic(err)
		} else {
			fmt.Println(waitTx(_agent.evmClient, tx))
		}

		//req := eas.AttestationRequest{
		//	Schema: res.Schema,
		//	Data:   eas.AttestationRequestData(res.Data),
		//}
		//
		//if tx, err := _agent.contract.Attest(_agent.txOp, req); err != nil {
		//	panic(err)
		//} else {
		//	fmt.Println(waitTx(_agent.evmClient, tx))
		//}
	}
}

func TestEIP712HashData(t *testing.T) {
	client, err := ethclient.Dial(OPBNBTESTRPC)
	if err != nil {
		panic(err)
	}

	_privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(_privateKey, big.NewInt(int64(OPBNBTESTCHAINID)))
	if err != nil {
		panic(err)
	}

	_eip712, err := eip712.NewEIP712(common.HexToAddress("0x50D7b7DB694715Ec62F999736FCBF872438A01f1"), client)
	fmt.Println(_eip712.Eip712Domain(&bind.CallOpts{From: auth.From}))
	d, _ := _eip712.GetDomainSeparator(&bind.CallOpts{From: auth.From})
	fmt.Println("domain", hex.EncodeToString(d[:]))

	var _agent *Agent
	if _agent, err = NewAgentFromKey(privateKey, OPBNBTESTBAS, OPBNBTESTSCHEMA, OPBNBTESTRPC, OPBNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"gm": true,
	}

	var param onchain.OnchainDelegateAttestationParam
	param.Attestor = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.Data = data
	param.Deadline = 1721552480
	param.ExpirationTime = 0
	param.Recipient = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.RefUid = "0x0000000000000000000000000000000000000000000000000000000000000001"
	param.Revocable = true
	param.Schema = "bool gm"
	param.SchemaUid = "0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61"
	param.Value = "0"
	//param.Nonce = "1"
	if res, err := _agent.OnchainSignDelegateAttestation(
		param,
		OPBNBTESTDOMAIN,
	); err != nil {
		panic(err)
	} else {

		// fmt.Println(res)
		// fmt.Println(json.Marshal(res))

		// var _value eip712.DelegatedProxyAttestationRequest
		// _value.Attester = res.Attester
		// _value.Data = eip712.AttestationRequestData(res.Data)
		// _value.Deadline = res.Deadline
		// _value.Schema = res.Schema
		// _value.Signature = eip712.Signature{
		// 	V: res.Signature.V,
		// 	R: common.HexToHash(res.Signature.R),
		// 	S: common.HexToHash(res.Signature.S),
		// }

		// fmt.Println(_value)
		// hash, err := _eip712.GetAttestationDataHash(&bind.CallOpts{}, _value)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// fmt.Println("============")
		// fmt.Println(hash)
		var _value eip712.DelegatedProxyAttestationRequest
		_value.Attester = res.Attester
		_value.Data = eip712.AttestationRequestData(res.Data)
		_value.Deadline = res.Deadline
		_value.Schema = res.Schema
		_value.Signature = eip712.Signature{
			V: res.Signature.V,
			R: common.HexToHash(res.Signature.R),
			S: common.HexToHash(res.Signature.S),
		}
		tx, _ := _eip712.AttestByDelegation(_agent.txOp, _value)

		fmt.Println(tx)

	}
}
