package offchain

import (
	"encoding/json"
	"fmt"
)

func CheckOffchainAttestation(attestationStr string, signer string, uid string, resolver string, schema string) (bool, error) {
	var offchainAttestation OffchainAttestationParam
	if err := json.Unmarshal([]byte(attestationStr), &offchainAttestation); err != nil {
		return false, fmt.Errorf("error attestation type" + err.Error())
	}

	return _checkOffchainAttestation(offchainAttestation, signer, uid, resolver, schema)

}

func _checkOffchainAttestation(attestation OffchainAttestationParam, signer string, uid string, resolver string, schema string) (bool, error) {
	//step 1: check uid
	var message MessageForUid
	if err := message.Decode(attestation.Message); err != nil {
		return false, fmt.Errorf("err for attestation uid: " + err.Error())
	}
	if recUid := GetOffChainAttestationUid(message); recUid != uid {
		return false, fmt.Errorf("un-matched attestation uid: " + recUid + " for " + uid)
	}

	//step 2: check the signer
	if recSigner, err := GetSigner(attestation.Signature, attestation.Domain, attestation.Type, attestation.Message); err != nil {
		return false, err
	} else {
		if recSigner != signer {
			return false, fmt.Errorf("un-matched signer: " + recSigner + " for " + signer)
		}
	}

	//step 3: resolver
	//TODO

	//step 4: check the data
	if _data, ok := attestation.Message["data"].(string); !ok {
		return false, fmt.Errorf("can not find data " + " for " + schema)
	} else if ok := checkSchema(_data, schema); !ok {
		return false, fmt.Errorf("un-matched attestation data: " + _data + " for " + schema)
	}

	return true, nil

}

func CheckOffchainAttestationRecSigner(attestationStr string, uid string, resolver string, schema string) (bool, string, error) {
	var offchainAttestation OffchainAttestationParam
	if err := json.Unmarshal([]byte(attestationStr), &offchainAttestation); err != nil {
		return false, "", fmt.Errorf("error attestation type" + err.Error())
	}

	return _checkOffchainAttestationRecSigner(offchainAttestation, uid, resolver, schema)

}

func _checkOffchainAttestationRecSigner(attestation OffchainAttestationParam, uid string, resolver string, schema string) (bool, string, error) {
	//step 1: check uid
	var message MessageForUid
	if err := message.Decode(attestation.Message); err != nil {
		return false, "", fmt.Errorf("err for attestation uid: " + err.Error())
	}

	if recUid := GetOffChainAttestationUid(message); recUid != uid {
		return false, "", fmt.Errorf("un-matched attestation uid: " + recUid + " for " + uid)
	}

	//step 2: check the signer
	var signer string
	if recSigner, err := GetSigner(attestation.Signature, attestation.Domain, attestation.Type, attestation.Message); err != nil {
		return false, "", err
	} else {
		signer = recSigner
	}

	//step 3: resolver
	//TODO

	//step 4: check the data
	if _data, ok := attestation.Message["data"].(string); !ok {
		return false, "", fmt.Errorf("can not find data " + " for " + schema)
	} else if ok := checkSchema(_data, schema); !ok {
		return false, "", fmt.Errorf("un-matched attestation data: " + _data + " for " + schema)
	}

	return true, signer, nil

}
