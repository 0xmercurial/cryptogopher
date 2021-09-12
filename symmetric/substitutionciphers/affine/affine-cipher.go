package affine

import (
	"log"
	"strings"
)

var alpahbet []byte
var indexes map[byte]int = map[byte]int{}

type affineKey struct {
	A int
	B int
}

func init() {
	var ch byte
	var index int
	for ch = 'a'; ch <= 'z'; ch++ {
		alpahbet = append(alpahbet, ch)
		indexes[ch] = index
		index++
	}
}

func encrypt(key affineKey, message string) string {
	var encrypted []byte
	message = strings.ToLower(message)
	if gcd(key.A, 26) != 1 {
		log.Println("Invalid key. A and 26 have a GCD that is greater than 1 ")
		return "none"
	}
	for _, ch := range message {
		index := indexes[byte(ch)]
		log.Print(((key.A * index) + key.B) % 26)
		encrypted = append(encrypted, alpahbet[((key.A*index)+key.B)%26])
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

func gcd(a int, b int) int {

	for d := a; d > 0; d-- {
		if a%d == 0 && b%d == 0 {
			return d
		}
	}
	return 1
}

/*
1. assume int a is max possible value, d
2. divide a and b by d
 - if a mod d AND b mod d == 0, you have found gcd
 - else, decrease d by 1
3.
*/
