package affine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAlpahbet(t *testing.T) {
	t.Log(string(alpahbet))
	t.Log(indexes)
}
func TestGCD(t *testing.T) {
	assert.Equal(t, gcd(2, 9), 1)
	assert.Equal(t, gcd(3, 9), 3)
}

func TestEncrypt(t *testing.T) {
	key := affineKey{A: 9, B: 13}
	assert.Equal(t, encrypt(key, "ATTACK"), "nccnfz")
}
