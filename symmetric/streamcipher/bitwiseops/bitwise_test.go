package bitwise

import (
	"strconv"
	"testing"
)

func TestBitLen(t *testing.T) {
	var bint uint8 = 70 //bits to be XOR'd
	var kint uint8 = 1  //key

	t.Log("Bits:   " + strconv.FormatInt(int64(bint), 2))
	t.Log("Key:     " + strconv.FormatInt(int64(kint), 2))

	xor := int64(bint) ^ int64(kint)
	t.Log("XOR:    " + strconv.FormatInt(xor, 2))
	t.Log("Decrypt:" + strconv.FormatInt(xor^int64(kint), 2))

}
