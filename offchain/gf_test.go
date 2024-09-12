package offchain

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestBundle(t *testing.T) {

	var datas []SingleBundleObject
	for i := 0; i < 10; i++ {
		var data SingleBundleObject
		data.Name = strconv.Itoa(i)
		data.Data = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(i)}
		datas = append(datas, data)
	}

	bundleData, size, err := GetBundle(datas)
	if err != nil {
		panic(err)
	}
	fmt.Println(bundleData, len(bundleData), size)

}

func TestRecoverBundle(t *testing.T) {

	path := ""

	b, err := RecoverBundle(path)
	if err != nil {
		panic(err)
	}
	res, _, err := b.GetObject("2")
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())
}
