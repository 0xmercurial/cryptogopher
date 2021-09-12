package caeser

import "strings"

var alpahbet []byte
var indexes map[byte]int = map[byte]int{}

func init() {
	var ch byte
	var index int
	for ch = 'a'; ch <= 'z'; ch++ {
		alpahbet = append(alpahbet, ch)
		indexes[ch] = index
		index++
	}
}

func encrypt(key int, message string) string {
	var encrypted []byte
	message = strings.ToLower(message)
	for _, ch := range message {
		index := indexes[byte(ch)]
		encrypted = append(encrypted, alpahbet[(index+key)%26])
	}
	return string(encrypted)
}

func decrypt(key int, message string) string {
	var decrypted []byte
	message = strings.ToLower(message)
	for _, ch := range message {
		index := indexes[byte(ch)]
		decrypted = append(decrypted, alpahbet[(index-key)%26])
	}
	return string(decrypted)
}

/*
1. read letter
2. find index of letter in slice
3. add key to index
4. retrieve letter at new index
*/
