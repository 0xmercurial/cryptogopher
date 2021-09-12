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
	log.Println(key.A)
	log.Println(gcd(key.A, 26))
	if gcd(key.A, 26) != 1 {
		log.Println("Invalid key. A and 26 have a GCD that is greater than 1.")
		return "none"
	}
	for _, ch := range message {
		index := indexes[byte(ch)]
		encrypted = append(encrypted, alpahbet[((key.A*index)+key.B)%26])
	}
	return string(encrypted)
}

func decrypt(key affineKey, message string) string {
	var decrypted []byte
	message = strings.ToLower(message)
	aInverse := modInverse(key.A, 26)
	for _, ch := range message {
		index := indexes[byte(ch)]
		absIdx := absMod(aInverse*(index-key.B), 26)
		log.Println(absIdx)
		decrypted = append(decrypted, alpahbet[absIdx])
	}
	return string(decrypted)
}

func gcd(a int, b int) int {

	var divisor int
	for d := a; d > 0; d-- {
		divisor = d
		if a%d == 0 && b%d == 0 {
			return divisor
		}
	}
	return divisor
}

func absMod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func modInverse(a int, m int) int {
	for i := 0; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return 0
}

/*
1. assume int a is max possible value, d
2. divide a and b by d
 - if a mod d AND b mod d == 0, you have found gcd
 - else, decrease d by 1
3.
*/
