package agent

import (
	"encoding/hex"
	"fmt"

	"github.com/bnb-attestation-service/bas-go/schemaRegistry"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: return uid
func (a *Agent) CreateSchema(schema string, revocable bool, resolver string) (string, error) {
	_resolver := common.HexToAddress(resolver)
	uid, err := a.schemaContract.Register(a.txOp, schema, _resolver, revocable)
	if err != nil {
		return "", fmt.Errorf("create schema error: " + err.Error())
	} else {
		return uid.Hash().Hex(), nil
	}
}

func (a *Agent) GetSchema(uid string) (*schemaRegistry.SchemaRecord, error) {
	if uid[:2] == "0x" {
		uid = uid[2:]
	}
	_uid, err := hex.DecodeString(uid)
	if err != nil || len(_uid) != 32 {
		return nil, fmt.Errorf("can not parse uid: " + uid)
	}

	if schema, err := a.schemaContract.GetSchema(a.callOp, [32]byte(_uid)); err != nil {
		return nil, fmt.Errorf("get schema error: " + err.Error())
	} else {
		return &schema, nil
	}
}

func (a *Agent) SetSchemaName(uid string, name string) (string, error) {
	bUId := common.HexToHash(uid)
	nData := make(map[string]interface{})
	nData["schemaId"] = bUId
	nData["name"] = name
	return a.OnchainAttest(NameSchemaUid, ZEROADDRESS, ZeroRef, nData, true, 0, 0, 0)
}

func (a *Agent) SetSchemaDescription(uid string, description string) (string, error) {
	bUId := common.HexToHash(uid)
	nData := make(map[string]interface{})
	nData["schemaId"] = bUId
	nData["description"] = description
	return a.OnchainAttest(DescriptionSchemaUId, ZEROADDRESS, ZeroRef, nData, true, 0, 0, 0)

}
