package offchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	solsha3 "github.com/bnb-attestation-service/go-solidity-sha3"
)

func GetOffChainAttestationUid(param MessageForUid) string {

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

func GetBundleUid(attestationUids []string) (string, error) {
	if len(attestationUids) < 1 {
		return "", fmt.Errorf("empty attestations")
	}
	for _, uid := range attestationUids {
		if len(uid) != 66 || uid[:2] != "0x" {
			return "", fmt.Errorf("invalid attestation uid")
		}
		if _, err := hex.DecodeString(uid[2:]); err != nil {
			return "", fmt.Errorf("invalid attestation uid")
		}
	}
	sort.Strings(attestationUids)
	sortedAddresses := strings.Join(attestationUids, "")
	hash := sha256.New()
	hash.Write([]byte(sortedAddresses))
	uid := hex.EncodeToString(hash.Sum(nil))
	return "0x" + uid, nil
}
