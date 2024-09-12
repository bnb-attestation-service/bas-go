package agent

func sliceToArray(s []byte) [32]byte {
	var a [32]byte
	copy(a[:], s[:])
	return a
}
