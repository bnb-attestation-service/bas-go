package bas_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/bnb-attestation-service/bas_go/offchain"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	storageTypes "github.com/bnb-chain/greenfield/x/storage/types"
)

func (a *Agent) ConfigBucket(bucket string) {
	a.gfBucket = bucket
}

func (a *Agent) OffchainNewAttestation(schemaUid string, schema string, data map[string]interface{}, recipient string, revocable bool, refUid string, salt string, nonce uint64, time uint64, expirationTime uint64, version uint16) (*offchain.OffchainAttestationParam, error) {
	return offchain.NewBASOffchainAttestation(schemaUid, schema, data, recipient, revocable, refUid, salt, nonce, time, expirationTime, version, a.privKey)
}

func (a *Agent) OffchainGetAttestationJson(attestation *offchain.OffchainAttestationParam) (string, error) {

	if _b, err := json.Marshal(attestation); err != nil {
		return "", err
	} else {
		return string(_b), nil
	}
}

func (a *Agent) OffchainUploadAttestationToGF(attestation *offchain.OffchainAttestationParam) (string, error) {
	if a.gfBucket == "" || a.gfBucket[:3] != "bas" {
		return "", fmt.Errorf("please config or new gf bucket first")
	}
	objName := fmt.Sprintf("%s.%s", attestation.Message["schema"], attestation.Uid)
	if _data, err := json.Marshal(attestation); err != nil {
		return "", fmt.Errorf("marshal offchain attestation error: " + err.Error())
	} else {
		var buffer bytes.Buffer
		ctx := context.Background()

		buffer.WriteString(fmt.Sprintf("%s", _data))

		if txnHash, err := a.gfClient.CreateObject(ctx, a.gfBucket, objName, bytes.NewReader(buffer.Bytes()), types.CreateObjectOptions{Visibility: storageTypes.VISIBILITY_TYPE_PUBLIC_READ}); err != nil {
			return "", fmt.Errorf("upload to gf err: " + err.Error())
		} else {
			return txnHash, nil
		}
	}
}
