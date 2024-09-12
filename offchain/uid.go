package offchain

import (
	"encoding/hex"

	solsha3 "github.com/bnb-attestation-service/go-solidity-sha3"
)

func getOffChainAttestationUid(param MessageForUid) string {

	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint16", "bytes", "address", "address", "uint64", "uint64", "bool", "bytes32", "bytes", "uint32"},

		// values
		[]interface{}{
			param.Version,
			param.Schema,
			param.Recipient,
			ZEROADDRESS,
			param.Time,
			param.ExpirationTime,
			param.Revocable,
			param.RefUID,
			param.Data,
			0,
		},
	)

	recUid := hex.EncodeToString(hash)
	return "0x" + recUid

}
