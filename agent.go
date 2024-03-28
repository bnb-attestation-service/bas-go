package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/bnb-attestation-service/bas-go/eas"
	"github.com/bnb-attestation-service/bas-go/schemaRegistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	gf "github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
)

const (
	BAS               = "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"
	OPBNBTESTBAS      = "0x5e905F77f59491F03eBB78c204986aaDEB0C6bDa"
	SCHEMA            = "0x08C8b8417313fF130526862f90cd822B55002D72"
	SCHEMANAME        = "0x44d562ac1d7cd77e232978687fea027ace48f719cf1d58c7888e509663bb87fc"
	SCHEMADESCRIPTION = "0x21cbc60aac46ba22125ff85dd01882ebe6e87eb4fc46628589931ccbef9b8c94"
)

const (
	BNBTESTRPC     = "https://data-seed-prebsc-1-s1.bnbchain.org:8545"
	BNBTESTCHAINID = 97

	OPBNBTESTRPC     = "https://opbnb-testnet-rpc.bnbchain.org"
	OPBNBTESTCHAINID = 5611

	GFTESTRPC     = "https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org:443"
	GFTESTCHAINID = "greenfield_5600-1"
)

type Agent struct {
	contract       *eas.EAS
	schemaContract *schemaRegistry.SchemaRegistry
	txOp           *bind.TransactOpts
	callOp         *bind.CallOpts

	evmClient *ethclient.Client
	gfClient  gf.IClient
	gfBucket  string

	privKey *ecdsa.PrivateKey
}

func NewAgentFromKey(privKey string, bas string, evmRPC string, evmChainId uint64, gfRPC string, gfChainId string) (*Agent, error) {
	client, err := ethclient.Dial(evmRPC)
	if err != nil {
		return nil, err
	}

	// 你的合约地址
	contractAddress := common.HexToAddress(bas)
	schemaContractAddress := common.HexToAddress(SCHEMA)

	// 创建一个新的私钥
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	// 实例化一个新的 transactor 对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(evmChainId)))
	if err != nil {
		return nil, err
	}

	// 设置 Gas 价格和限制
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000) // 设置你期望的 Gas 限制

	// 设置 Nonce
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	// 创建一个与智能合约交互的对象
	contract, err := eas.NewEAS(contractAddress, client)
	if err != nil {
		return nil, err
	}

	schemaContract, err := schemaRegistry.NewSchemaRegistry(schemaContractAddress, client)
	if err != nil {
		return nil, err
	}

	caller := bind.CallOpts{Context: context.Background()}

	account, err := types.NewAccountFromPrivateKey("test", privKey)
	if err != nil {
		log.Fatalf("New account from private key error, %v", err)
	}
	cli, err := gf.New(gfChainId, gfRPC, gf.Option{DefaultAccount: account})

	if err != nil {
		log.Fatalf("unable to new greenfield client, %v", err)
	}

	return &Agent{
		contract:       contract,
		schemaContract: schemaContract,
		txOp:           auth,
		callOp:         &caller,

		evmClient: client,
		gfClient:  cli,
		privKey:   privateKey,
	}, nil
}
