package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-attestation-service/bas-go/offchain"
	bundletypes "github.com/bnb-chain/greenfield-bundle-sdk/types"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	permissionTypes "github.com/bnb-chain/greenfield/x/permission/types"
	storageTypes "github.com/bnb-chain/greenfield/x/storage/types"
)

func GetBASBucketName(addr string) string {
	return "bas"
}
func (a *Agent) ConfigBucket(bucket string) {
	a.gfBucket = bucket
}

// TODO: change to bas bucket name
func (a *Agent) CreateBucket(address string) (string, error) {
	if !common.IsHexAddress(address) {
		return "", errors.New("invalid address")
	}

	var bucketPrefix string
	switch a.evmChainId {
	case 56, 97:
		bucketPrefix = "bas"
	case 204, 5611:
		bucketPrefix = "obas"
	default:
		return "", errors.New("invalid chain id")
	}

	name := strings.ToLower(fmt.Sprintf("%s-%s", bucketPrefix, address))
	ctx := context.Background()
	// get storage providers list
	spLists, err := a.gfClient.ListStorageProviders(ctx, true)
	if err != nil {
		return "", fmt.Errorf("fail to list in service sps")
	}
	// choose the first sp to be the primary SP
	primarySP := spLists[0].GetOperatorAddress()
	// addr := crypto.FromECDSA(a.privKey)
	// bucketName := GetBASBucketName(hex.EncodeToString(addr))
	bucketName := name
	if hash, err := a.gfClient.CreateBucket(ctx, bucketName, primarySP, types.CreateBucketOptions{}); err != nil {
		return "", err
	} else {
		fmt.Println("================================================")
		fmt.Println("Create bucket: " + bucketName)
		fmt.Println("TX hash: " + "0x" + hash)
		fmt.Println("You can find your bucket at : https://dcellar.io")
		fmt.Println("================================================")
	}
	a.gfBucket = bucketName
	return bucketName, nil
}

func (a *Agent) OffchainNewAttestation(schemaUid string, domain offchain.OffchainAttestationDomain, data map[string]interface{}, recipient string, revocable bool, refUid string, nonce uint64, time uint64, expirationTime uint64, version uint16) (*offchain.OffchainAttestationParam, error) {
	bSchemaUid := common.HexToHash(schemaUid)
	schemaRecord, err := a.schemaContract.GetSchema(new(bind.CallOpts), bSchemaUid)
	if err != nil {
		return nil, err
	}
	schema := schemaRecord.Schema
	return offchain.NewBASOffchainAttestation(schemaUid, schema, domain, data, recipient, revocable, refUid, nonce, time, expirationTime, version, a.privKey)
}

func (a *Agent) OffchainNewAttestation2(schemaUid string, domain offchain.OffchainAttestationDomain, data []byte, recipient string, revocable bool, refUid string, nonce uint64, time uint64, expirationTime uint64, version uint16) (*offchain.OffchainAttestationParam, error) {
	return offchain.NewBASOffChainAttestation2(schemaUid, domain, data, recipient, revocable, refUid, nonce, time, expirationTime, version, a.privKey)
}

func (a *Agent) OffchainGetAttestationJson(attestation *offchain.OffchainAttestationParam) (string, error) {

	if _b, err := json.Marshal(attestation); err != nil {
		return "", err
	} else {
		return string(_b), nil
	}
}

func (a *Agent) OffchainUploadAttestationToGF(attestation *offchain.OffchainAttestationParam, visible bool) (string, error) {
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
		var visibility storageTypes.VisibilityType
		if visible {
			visibility = storageTypes.VISIBILITY_TYPE_PUBLIC_READ
		} else {
			visibility = storageTypes.VISIBILITY_TYPE_PRIVATE
		}
		var txHash string
		if txHash, err = a.gfClient.CreateObject(ctx, a.gfBucket, objName, bytes.NewReader(buffer.Bytes()), types.CreateObjectOptions{Visibility: visibility}); err != nil {
			return "", fmt.Errorf("create obj gf err: " + err.Error())
		}

		if err = a.gfClient.PutObject(ctx, a.gfBucket, objName, int64(buffer.Len()), bytes.NewReader(buffer.Bytes()), types.PutObjectOptions{TxnHash: txHash}); err != nil {
			return "", fmt.Errorf("put obj gf err: " + err.Error())
		}
		return txHash, nil
	}
}

func (a *Agent) OffchainUploadAttestationToGFFromRaw(attestationBytes []byte, bucket string, visible bool) (string, *offchain.OffchainAttestationParam, error) {
	var attestation offchain.OffchainAttestationParam
	if err := json.Unmarshal(attestationBytes, &attestation); err != nil {
		return "", nil, fmt.Errorf("error attestation type" + err.Error())
	}

	objName := fmt.Sprintf("%s.%s", attestation.Message["schema"], attestation.Uid)
	if _data, err := json.Marshal(attestation); err != nil {
		return "", nil, fmt.Errorf("marshal offchain attestation error: " + err.Error())
	} else {
		var buffer bytes.Buffer
		ctx := context.Background()

		buffer.WriteString(fmt.Sprintf("%s", _data))
		var visibility storageTypes.VisibilityType
		if visible {
			visibility = storageTypes.VISIBILITY_TYPE_PUBLIC_READ
		} else {
			visibility = storageTypes.VISIBILITY_TYPE_PRIVATE
		}
		var txHash string
		if txHash, err = a.gfClient.CreateObject(ctx, bucket, objName, bytes.NewReader(buffer.Bytes()), types.CreateObjectOptions{Visibility: visibility}); err != nil {
			return "", nil, fmt.Errorf("create obj gf err: " + err.Error())
		}

		if err = a.gfClient.PutObject(ctx, bucket, objName, int64(buffer.Len()), bytes.NewReader(buffer.Bytes()), types.PutObjectOptions{TxnHash: txHash}); err != nil {
			return "", nil, fmt.Errorf("put obj gf err: " + err.Error())
		}
		return txHash, &attestation, nil
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

func (a *Agent) OffchainUploadBundleToGF(datas []offchain.SingleBundleObject, name, bucket string, visibility storageTypes.VisibilityType) (string, error) {

	bundleData, size, err := offchain.GetBundle(datas)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	var txHash string
	if txHash, err = a.gfClient.CreateObject(ctx, bucket, name, bytes.NewReader(bundleData), types.CreateObjectOptions{Visibility: visibility}); err != nil {
		return "", fmt.Errorf("create obj gf err: " + err.Error())
	}

	if err = a.gfClient.PutObject(ctx, bucket, name, size, bytes.NewReader(bundleData), types.PutObjectOptions{TxnHash: txHash}); err != nil {
		return "", fmt.Errorf("put obj gf err: " + err.Error())
	}
	return txHash, nil

}

func (a *Agent) OffchainDownloadBundle(bucketName, objName, savePath string) (string, error) {
	ctx := context.Background()
	// Get bundle object from Greenfield
	bundledObject, _, err := a.gfClient.GetObject(ctx, bucketName, objName, types.GetObjectOptions{})
	if err != nil {
		return "", err
	}
	// Write bundle object into temp file
	bundleFile, err := os.CreateTemp(savePath, bundletypes.TempBundleFilePrefix)
	if err != nil {
		return "", err
	}
	defer bundleFile.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(bundledObject)

	if err != nil {
		return "", err
	}
	_, err = bundleFile.Write(buf.Bytes())
	if err != nil {
		return "", err
	}
	return bundleFile.Name(), nil
}

func (a *Agent) OffchainMultiAttestByBundle(attestations []*offchain.OffchainAttestationParam, schemaUid, bucket string, visibility storageTypes.VisibilityType) (string, error) {
	var objs []offchain.SingleBundleObject
	var attestationUids []string
	for _, attestation := range attestations {
		attestationUids = append(attestationUids, attestation.Uid)
		if _b, err := json.Marshal(attestation); err != nil {
			return "", fmt.Errorf("offchain multi attest error: %v", err)
		} else {
			var obj offchain.SingleBundleObject
			obj.Data = _b
			obj.Name = attestation.Uid
			objs = append(objs, obj)
		}
	}
	bundleUid, err := offchain.GetBundleUid(attestationUids)
	if err != nil {
		return "", err
	}
	objName := fmt.Sprintf("bundle.%s.%s", schemaUid, bundleUid)
	return a.OffchainUploadBundleToGF(objs, objName, bucket, visibility)

}

func (a *Agent) OffchainParseAttestationsFromBundle(bundleFile string, bundleName string) (map[string]offchain.OffChainAttestation, map[string][]byte, error) {
	// we have to check schema ID
	re := regexp.MustCompile(`bundle\.(0x[a-fA-F0-9]{64})\.(0x[a-fA-F0-9]{64})`)

	matches := re.FindStringSubmatch(bundleName)
	var attestationsRawData = make(map[string][]byte)
	var schemaId string
	var bundleUid string
	if len(matches) > 2 {
		schemaId = matches[1]
		bundleUid = matches[2]
	} else {
		return nil, nil, fmt.Errorf("invalid schema Id in bundle")
	}

	data, err := offchain.RecoverBundle(bundleFile)
	if err != nil {
		return nil, nil, err
	}
	sources := data.GetBundleObjectsMeta()

	results := map[string]offchain.OffChainAttestation{}

	var attestationUids []string
	for _, source := range sources {
		obj, _, err := data.GetObject(source.Name)
		if err != nil || obj == nil {
			return nil, nil, fmt.Errorf("parse object in bundled object failed: %v", err)
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(obj)
		if err != nil {
			return nil, nil, fmt.Errorf("parse object in bundled object failed: %v", err)
		}
		attest := buf.Bytes()
		var offchainAttestation offchain.OffchainAttestationParam
		if err := json.Unmarshal(attest, &offchainAttestation); err != nil {
			return nil, nil, fmt.Errorf("parse object in bundled object failed: %v", err)
		}
		var message offchain.MessageForUid
		if err := message.Decode(offchainAttestation.Message); err != nil {
			return nil, nil, fmt.Errorf("parse object in bundled object failed: %v", err)
		}
		if message.Schema != schemaId {
			return nil, nil, fmt.Errorf("parse object in bundled object failed: get an invalid schemaId")
		}

		uid := offchain.GetOffChainAttestationUid(message)
		if uid != source.Name {
			return nil, nil, fmt.Errorf("parse object in bundled object failed: {%s} has unmatched uid", source.Name)
		}
		attestationUids = append(attestationUids, uid)
		attestationsRawData[uid] = attest
		attestor, err := offchain.GetSigner(offchainAttestation.Signature, offchainAttestation.Domain, offchainAttestation.Type, offchainAttestation.Message)
		if err != nil {
			return nil, nil, fmt.Errorf("get signer from attestation Param failed: %v", err)
		}
		results[source.Name] = offchain.OffChainAttestation{
			Attestor:      attestor,
			MessageForUid: message,
		}
	}

	bundleUidRec, err := offchain.GetBundleUid(attestationUids)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid attestation uids")
	}
	if bundleUid != bundleUidRec {
		return nil, nil, fmt.Errorf("invalid bundle uid")
	}
	return results, attestationsRawData, nil
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
