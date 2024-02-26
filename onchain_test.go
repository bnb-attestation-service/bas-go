package bas_go

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/bnb-attestation-service/bas_go/eas"
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
	// //key, _ := crypto.GenerateKey()

	// keyByte := common.FromHex(privateKey)
	// key, err := x509.ParseECPrivateKey(keyByte)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(97))
	// if err != nil {
	// 	panic(err)
	// }

	// 连接到以太坊网络的客户端
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
	// 你的合约地址
	contractAddress := common.HexToAddress(EASContract)

	// 创建一个新的私钥
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 实例化一个新的 transactor 对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(97))
	if err != nil {
		panic(err)
	}

	// 设置 Gas 价格和限制
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000) // 设置你期望的 Gas 限制

	// 设置 Nonce
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 创建一个与智能合约交互的对象
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
