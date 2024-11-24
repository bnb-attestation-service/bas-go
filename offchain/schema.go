package offchain

import (
	"encoding/hex"
	"fmt"

	"github.com/umbracle/ethgo/abi"
)

func checkSchema(dataStr string, schema string) bool {
	_schema := fmt.Sprintf("tuple(%s)", schema)

	typ := abi.MustNewType(_schema)
	if dataStr[:2] == "0x" {
		dataStr = dataStr[2:]
	}
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return false
	}

	_, err = typ.Decode(data)

	// fmt.Println(v, err)
	if err != nil {
		return false
	}
	return true

}
