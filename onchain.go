package agent

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/bnb-attestation-service/bas-go/eas"
	"github.com/bnb-attestation-service/bas-go/onchain"
)

func (a *Agent) OnchainAttest(schemaUid string, data []byte, revocable bool, expirationTime uint64) (string, error) {
	if schemaUid[:2] == "0x" {
		schemaUid = schemaUid[2:]
	}
	_schema, err := hex.DecodeString(schemaUid)
	if err != nil || len(_schema) != 32 {
		return "", fmt.Errorf("can not parse schema uid: " + schemaUid)
	}

	req := eas.AttestationRequest{
		Schema: [32]byte(_schema),
		Data: eas.AttestationRequestData{
			ExpirationTime: expirationTime,
			Revocable:      revocable,
			Data:           data,
			Value:          big.NewInt(0),
		},
	}
	if tx, err := a.contract.Attest(a.txOp, req); err != nil {
		return "", fmt.Errorf("create attestation onchain error: " + err.Error())
	} else {
		return tx.Hash().Hex(), nil
	}
}

func (a *Agent) OnchainGetAttestation(uid string) (*eas.Attestation, error) {
	if uid[:2] == "0x" {
		uid = uid[2:]
	}
	_uid, err := hex.DecodeString(uid)
	if err != nil || len(_uid) != 32 {
		return nil, fmt.Errorf("can not parse uid: " + uid)
	}

	if attest, err := a.contract.GetAttestation(a.callOp, [32]byte(_uid)); err != nil {
		return nil, fmt.Errorf("get attestation onchain error: " + err.Error())
	} else {
		return &attest, nil
	}
}

func (a *Agent) OnchainRevoke(schema string, uid string) (string, error) {
	if schema[:2] == "0x" {
		schema = schema[2:]
	}
	_schema, err := hex.DecodeString(schema)
	if err != nil || len(_schema) != 32 {
		return "", fmt.Errorf("can not parse schema uid: " + schema)
	}

	if uid[:2] == "0x" {
		uid = uid[2:]
	}
	_uid, err := hex.DecodeString(uid)
	if err != nil || len(_uid) != 32 {
		return "", fmt.Errorf("can not parse uid: " + uid)
	}

	req := eas.RevocationRequest{
		Schema: [32]byte(_schema),
		Data: eas.RevocationRequestData{
			Uid:   [32]byte(_uid),
			Value: big.NewInt(0),
		},
	}
	if tx, err := a.contract.Revoke(a.txOp, req); err != nil {
		return "", fmt.Errorf("revoke attestation onchain error: " + err.Error())
	} else {
		return tx.Hash().Hex(), nil
	}
}

func (a *Agent) OnchainRevokeOffchain(uid string) (string, error) {

	if uid[:2] == "0x" {
		uid = uid[2:]
	}
	_uid, err := hex.DecodeString(uid)
	if err != nil || len(_uid) != 32 {
		return "", fmt.Errorf("can not parse uid: " + uid)
	}

	if tx, err := a.contract.RevokeOffchain(a.txOp, sliceToArray(_uid)); err != nil {
		return "", fmt.Errorf("revoke offchain attestation onchain error: " + err.Error())
	} else {
		return tx.Hash().Hex(), nil
	}
}

func (a *Agent) OnchainSignDelegateAttestation(attest onchain.OnchainDelegateAttestationParam, domain onchain.OnchainAttestationDomain) (*onchain.DelegatedProxyAttestation, error) {
	if sig, err := onchain.NewBASOnchainDelegateAttestation(
		attest,
		domain,
		a.privKey,
	); err != nil {
		return nil, err
	} else {
		return sig, nil
	}

}
