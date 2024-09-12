package offchain

import (
	"bytes"

	"github.com/bnb-chain/greenfield-bundle-sdk/bundle"
)

func GetBundle(obj []SingleBundleObject) ([]byte, int64, error) {
	bundle, err := bundle.NewBundle()
	if err != nil {
		return nil, 0, err
	}
	defer bundle.Close()

	for _, object := range obj {
		buf := bytes.NewReader(object.Data)
		_, err := bundle.AppendObject(object.Name, buf, nil)
		if err != nil {
			return nil, 0, err
		}

	}
	bundledObject, size, err := bundle.FinalizeBundle()
	if err != nil {
		return nil, 0, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(bundledObject)
	if err != nil {
		return nil, 0, err
	}
	return buf.Bytes(), size, nil

}

func RecoverBundle(file string) (*bundle.Bundle, error) {
	bundle, err := bundle.NewBundleFromFile(file)
	if err != nil {
		return nil, err
	}
	return bundle, nil
}
