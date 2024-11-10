package agent

import (
	"encoding/hex"
	"fmt"

	"github.com/bnb-attestation-service/bas-go/schemaRegistry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/ethgo/abi"
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

func (a *Agent) SetSchemaName(uid string, name string, schemaNameUid string) (string, error) {
	if uid[:2] == "0x" {
		uid = uid[2:]
	}
	_uid, err := hex.DecodeString(uid)
	if err != nil || len(_uid) != 32 {
		return "", fmt.Errorf("can not parse uid: " + uid)
	}

	if nData, err := encodeSchemaName([32]byte(_uid), name); err != nil {
		return "", err
	} else {
		return a.OnchainAttest(schemaNameUid, ZEROADDRESS, nData, true, 0)
	}
}

func encodeSchemaName(uid [32]byte, name string) ([]byte, error) {

	n := map[string]interface{}{}
	n["name"] = name
	n["schemaId"] = uid

	typ := abi.MustNewType("tuple(bytes32 schemaId,string name)")

	res, err := typ.Encode(&n)
	if err != nil {
		return nil, fmt.Errorf("fail to encode schema attestors data with error %v", err)
	}

	return res, nil
}

func (a *Agent) SetSchemaDescription(uid string, description string, schemaDescriptionUid string) (string, error) {
	if uid[:2] == "0x" {
		uid = uid[2:]
	}
	_uid, err := hex.DecodeString(uid)
	if err != nil || len(_uid) != 32 {
		return "", fmt.Errorf("can not parse uid: " + uid)
	}

	if nData, err := encodeSchemDescription([32]byte(_uid), description); err != nil {
		return "", err
	} else {
		return a.OnchainAttest(schemaDescriptionUid, ZEROADDRESS, nData, true, 0)
	}
}

func encodeSchemDescription(uid [32]byte, description string) ([]byte, error) {

	n := map[string]interface{}{}
	n["description"] = description
	n["schemaId"] = uid

	typ := abi.MustNewType("tuple(bytes32 schemaId,string description)")

	res, err := typ.Encode(&n)
	if err != nil {
		return nil, fmt.Errorf("fail to encode schema attestors data with error %v", err)
	}

	return res, nil
}
