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
	if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
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
	if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}
	if attest, err := _agent.OnchainRevoke(schema, uid); err != nil {
		panic(err)
	} else {
		fmt.Println(attest)
	}

}

func TestOnchainDelegate(t *testing.T) {
	var _agent *Agent
	var err error
	if _agent, err = NewAgentFromKey(privateKey, BNBTESTRPC, BNBTESTCHAINID, GFTESTRPC, GFTESTCHAINID); err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"p":     "1212",
		"tick":  "1212",
		"amt":   10,
		"nonce": 10,
		"vote":  0,
	}

	var param onchain.OnchainDelegateAttestationParam
	param.Attestor = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.Data = data
	param.Deadline = 1721552480
	param.ExpirationTime = 0
	param.Recipient = "0x471543A3bd04486008c8a38c5C00543B73F1769e"
	param.RefUid = "0x0000000000000000000000000000000000000000000000000000000000000000"
	param.Revocable = false
	param.Schema = "string p,string tick,uint256 amt,uint8 vote,uint256 nonce"
	param.SchemaUid = "0x5bb3334a97088f7c018fafb6cdd5f06d17c6734ba10fe3944115b815b8b89d2f"
	param.Value = ""
	if res, err := _agent.OnchainSignDelegateAttestation(
		param,
	); err != nil {
		panic(err)
	} else {
		fmt.Println(res)

	}

}
