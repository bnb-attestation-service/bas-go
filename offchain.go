package agent

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bnb-attestation-service/bas-go/offchain"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	permissionTypes "github.com/bnb-chain/greenfield/x/permission/types"
	storageTypes "github.com/bnb-chain/greenfield/x/storage/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetBASBucketName(addr string) string {
	return "bas"
}
func (a *Agent) ConfigBucket(bucket string) {
	a.gfBucket = bucket
}

// TODO: change to bas bucket name
func (a *Agent) CreateBucket() error {
	ctx := context.Background()
	// get storage providers list
	spLists, err := a.gfClient.ListStorageProviders(ctx, true)
	if err != nil {
		log.Fatalf("fail to list in service sps")
	}
	// choose the first sp to be the primary SP
	primarySP := spLists[0].GetOperatorAddress()
	addr := crypto.FromECDSA(a.privKey)
	bucketName := GetBASBucketName(hex.EncodeToString(addr))
	if hash, err := a.gfClient.CreateBucket(ctx, bucketName, primarySP, types.CreateBucketOptions{}); err != nil {
		return err
	} else {
		fmt.Println("================================================")
		fmt.Println("Create bucket: " + bucketName)
		fmt.Println("TX hash: " + "0x" + hash)
		fmt.Println("You can find your bucket at : https://dcellar.io")
		fmt.Println("================================================")
	}
	a.gfBucket = bucketName
	return nil
}

func (a *Agent) OffchainNewAttestation(schemaUid string, schema string, data map[string]interface{}, recipient string, revocable bool, refUid string, nonce uint64, time uint64, expirationTime uint64, version uint16) (*offchain.OffchainAttestationParam, error) {
	return offchain.NewBASOffchainAttestation(schemaUid, schema, data, recipient, revocable, refUid, nonce, time, expirationTime, version, a.privKey)
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
		return "", fmt.Errorf("please config or create gf bucket first")
	}
	objName := fmt.Sprintf("%s.%s", attestation.Message["schema"], attestation.Uid)
	if _data, err := json.Marshal(attestation); err != nil {
		return "", fmt.Errorf("marshal offchain attestation error: " + err.Error())
	} else {
		var buffer bytes.Buffer
		ctx := context.Background()

		buffer.WriteString(fmt.Sprintf("%s", _data))
		var txHash string
		if txHash, err = a.gfClient.CreateObject(ctx, a.gfBucket, objName, bytes.NewReader(buffer.Bytes()), types.CreateObjectOptions{Visibility: storageTypes.VISIBILITY_TYPE_PUBLIC_READ}); err != nil {
			return "", fmt.Errorf("create obj gf err: " + err.Error())
		}

		if err = a.gfClient.PutObject(ctx, a.gfBucket, objName, int64(buffer.Len()), bytes.NewReader(buffer.Bytes()), types.PutObjectOptions{TxnHash: txHash}); err != nil {
			return "", fmt.Errorf("put obj gf err: " + err.Error())
		}
		return txHash, nil
	}
}

func (a *Agent) OffchainChangeAttestationVisible(schemaUid string, attestationUid string, visible bool) (string, error) {
	if a.gfBucket == "" || a.gfBucket[:3] != "bas" {
		return "", fmt.Errorf("please config or create gf bucket first")
	}
	objName := fmt.Sprintf("%s.%s", schemaUid, attestationUid)
	ctx := context.Background()
	var v storageTypes.VisibilityType
	if visible {
		v = storageTypes.VISIBILITY_TYPE_PUBLIC_READ
	} else {
		v = storageTypes.VISIBILITY_TYPE_PRIVATE
	}
	if hash, err := a.gfClient.UpdateObjectVisibility(ctx, a.gfBucket, objName, v, types.UpdateObjectOption{}); err != nil {
		return "", err
	} else {
		return hash, nil
	}
}

func (a *Agent) OffchainUploadBundleToGF(datas []offchain.SingleBundleObject, name, bucket string) (string, error) {

	bundleData, size, err := offchain.GetBundle(datas)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	var txHash string
	if txHash, err = a.gfClient.CreateObject(ctx, bucket, name, bytes.NewReader(bundleData), types.CreateObjectOptions{Visibility: storageTypes.VISIBILITY_TYPE_PUBLIC_READ}); err != nil {
		return "", fmt.Errorf("create obj gf err: " + err.Error())
	}

	if err = a.gfClient.PutObject(ctx, bucket, name, size, bytes.NewReader(bundleData), types.PutObjectOptions{TxnHash: txHash}); err != nil {
		return "", fmt.Errorf("put obj gf err: " + err.Error())
	}
	return txHash, nil

}

func (a *Agent) CheckWritePermission(bucket string) (bool, error) {
	ctx := context.Background()
	if permission, err := a.gfClient.IsBucketPermissionAllowed(ctx, a.address, bucket, permissionTypes.ACTION_CREATE_OBJECT); err != nil {
		return false, err
	} else {
		if permission != permissionTypes.EFFECT_ALLOW {
			return false, nil
		}
	}

	return true, nil
}
