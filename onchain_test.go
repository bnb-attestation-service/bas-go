package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/bnb-attestation-service/bas-go/eas"
	"github.com/bnb-attestation-service/bas-go/onchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	privateKey  = "19dfa30d6165181386c6a706f065bd841bf41fb9457a7f4a7a9b8c5df5d4de89"
	EASContract = "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"
)

func TestAttest(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")

	contractAddress := common.HexToAddress(EASContract)

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
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	_agent.OnchainAttest(
		"0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61",
		nil,
		false,
		0,
	)
}
func TestGetAttestation(t *testing.T) {
	uid := "0x02a902caa703b0af896614c443265844d50b82ba2c321bb01232e7b34d8a1d19"
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	if _agent, err = NewAgentFromKey(privateKey, BAS, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	if _agent, err = NewAgentFromKey(privateKey, OPBNBTESTBAS, OPBNBTESTRPC, OPBNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	param.ExpirationTime = 0
	param.Recipient = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.RefUid = "0x0000000000000000000000000000000000000000000000000000000000000000"
	param.Revocable = true
	param.Schema = "bool gm"
	param.SchemaUid = "0x85500e806cf1e74844d51a20a6d893fe1ed6f6b0738b50e43d774827d08eca61"
	param.Value = "100000"
	param.Nonce = "0"
	if res, err := _agent.OnchainSignDelegateAttestation(
		param,
	); err != nil {
		panic(err)
	} else {
		// s := "0xfeb2925a02bae3dae48d424a0437a2b6ac939aa9230ddc55a1a76f065d988076"
		// fmt.Println(hex.DecodeString(s[2:]))
		// fmt.Println(res)
		// fmt.Println(json.Marshal(res))

		var _value eas.DelegatedAttestationRequest
		_value.Attester = res.Attester
		_value.Data = eas.AttestationRequestData(res.Data)
		_value.Deadline = res.Deadline
		_value.Schema = res.Schema
		if r, err := hex.DecodeString(res.Signature.R[2:]); err != nil || len(r) != 32 {
			panic(err)
		} else {
			_value.Signature.R = [32]byte(r)
		}

		if s, err := hex.DecodeString(res.Signature.S[2:]); err != nil || len(s) != 32 {
			panic(err)
		} else {
			_value.Signature.S = [32]byte(s)
		}

		_value.Signature.V = res.Signature.V

		fmt.Println(_value)
		if tx, err := _agent.contract.AttestByDelegation(_agent.txOp, _value); err != nil {
			panic(err)
		} else {
			fmt.Println(waitTx(_agent.evmClient, tx))

		}

	}

}
