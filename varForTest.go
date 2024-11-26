package agent

import (
	"github.com/bnb-attestation-service/bas-go/offchain"
	"github.com/bnb-attestation-service/bas-go/onchain"
)

const (
	// 'Test' for Testnet (bsc, opbnb)
	TESTBAS      = "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD"
	OPBNBTESTBAS = "0x5e905F77f59491F03eBB78c204986aaDEB0C6bDa"

	TESTSCHEMA      = "0x08C8b8417313fF130526862f90cd822B55002D72"
	OPBNBTESTSCHEMA = "" // could be set to nil if not used in agent

)

const (
	BNBTESTRPC     = "https://data-seed-prebsc-1-s1.bnbchain.org:8545"
	BNBTESTCHAINID = 97

	OPBNBRPC     = "https://opbnb-mainnet-rpc.bnbchain.org"
	OPBNBCHAINID = 204

	OPBNBTESTRPC     = "https://opbnb-testnet-rpc.bnbchain.org"
	OPBNBTESTCHAINID = 5611

	GFTESTRPC     = "https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org:443"
	GFTESTCHAINID = "greenfield_5600-1"
)

const (
	privateKey = ""
)

var BASDOMAIN = onchain.OnchainAttestationDomain{
	Name:              "BAS Attestation",
	Version:           "1.3.0",d
	ChainId:           "56",
	VerifyingContract: "0x01dAc45529a070Cb67Fc5B328a7eBE394644355B",
}

var BASTESTDOMAIN = onchain.OnchainAttestationDomain{
	Name:              "BAS Attestation",
	Version:           "1.3.0",
	ChainId:           "97",
	VerifyingContract: "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD",
}

var OPBNBTESTDOMAIN = onchain.OnchainAttestationDomain{
	Name:              "OPBNB ATTESTATION",
	Version:           "1.3.0",
	ChainId:           "5611",
	VerifyingContract: "0x50D7b7DB694715Ec62F999736FCBF872438A01f1",
}

var OPBNBDOAIN = onchain.OnchainAttestationDomain{
	Name:              "OPBNB ATTESTATION",
	Version:           "1.3.0",
	ChainId:           "204",
	VerifyingContract: "0x6f9397703f9911Ec39C52D344431e81FE5a6710b",
}

var OFFCHAINBASTESTDOMAIN = offchain.OffchainAttestationDomain{
	Name:              "BAS Attestation",
	Version:           "1.3.0",
	ChainId:           "97",
	VerifyingContract: "0x6c2270298b1e6046898a322acB3Cbad6F99f7CBD",
}
