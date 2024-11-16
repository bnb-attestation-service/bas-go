package agent

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

type Agent struct {
	contract       *eas.EAS
	schemaContract *schemaRegistry.SchemaRegistry
	txOp           *bind.TransactOpts
	callOp         *bind.CallOpts

	evmClient *ethclient.Client
	gfClient  gf.IClient
	gfBucket  string

	privKey *ecdsa.PrivateKey
	address string

	BasContract    string
	SchemaContract string
}

func NewAgentFromKey(privKey string, basAddress string, schemaAddress string, evmRPC string, evmChainId uint64, gfRPC string, gfChainId string) (*Agent, error) {
	client, err := ethclient.Dial(evmRPC)
	if err != nil {
		return nil, err
	}

	// Contract Address
	contractAddress := common.HexToAddress(basAddress)
	schemaContractAddress := common.HexToAddress(schemaAddress)

	// Create a new privKey
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(evmChainId)))
	if err != nil {
		return nil, err
	}

	// // Gas Limit
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(300000)

	// nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	// if err != nil {
	// 	return nil, err
	// }
	// auth.Nonce = big.NewInt(int64(nonce))

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

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	return &Agent{
		contract:       contract,
		schemaContract: schemaContract,
		txOp:           auth,
		callOp:         &caller,

		evmClient: client,
		gfClient:  cli,
		privKey:   privateKey,
		address:   address,
	}, nil
}

// return agent's address
func (a *Agent) GetAddress() string {
	return a.address
}

func (a *Agent) SetGas() error {
	// Gas Limit
	gasPrice, err := a.evmClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	a.txOp.GasPrice = gasPrice
	a.txOp.GasLimit = uint64(300000)

	nonce, err := a.evmClient.PendingNonceAt(context.Background(), a.txOp.From)
	if err != nil {
		return err
	}
	a.txOp.Nonce = big.NewInt(int64(nonce))
	return nil
}
