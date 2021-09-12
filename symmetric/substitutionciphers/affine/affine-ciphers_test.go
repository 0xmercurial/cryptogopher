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
	assert.Equal(t, gcd(3, 26), 1)
	assert.Equal(t, gcd(5, 26), 1)
}

func TestValidEncrypt(t *testing.T) {
	key := affineKey{A: 9, B: 13}
	assert.Equal(t, encrypt(key, "ATTACK"), "nccnfz")
}

func TestInvalidEncrypt(t *testing.T) {
	key := affineKey{A: 2, B: 13}
	assert.Equal(t, encrypt(key, "ATTACK"), "none")
}

func TestAbsMod(t *testing.T) {
	t.Log(absMod(-33, 26))
}
func TestModInverse(t *testing.T) {
	assert.Equal(t, modInverse(3, 26), 9)
}

func TestValidDecrypt(t *testing.T) {
	key := affineKey{A: 9, B: 13}
	assert.Equal(t, encrypt(key, "ATTACK"), "nccnfz")
	t.Log(decrypt(key, "nccnfz"))
	assert.Equal(t, decrypt(key, "nccnfz"), "attack")
}
