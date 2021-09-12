package substitutionciphers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAlpahbet(t *testing.T) {
	t.Log(string(alpahbet))
	t.Log(indexes)
}

func TestEncrypt(t *testing.T) {
	message := "ABCDE"
	encrypted := encrypt(3, message)
	assert.Equal(t, encrypted, "defgh")
}

func TestDecrypt(t *testing.T) {
	key := 3
	message := "ABCDE"
	encrypted := encrypt(key, message)
	assert.Equal(t, encrypted, "defgh")
	decrypted := decrypt(key, encrypted)
	assert.Equal(t, decrypted, "abcde")
}
