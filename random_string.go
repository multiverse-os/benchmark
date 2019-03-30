package benchmark

import (
	"math/rand"
	"strings"
	//"time"
)

//func init() {
//	rand.Seed(time.Now().UTC().UnixNano())
//}

const (
	ALPHABETIC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMERIC    = "0123456789"
	SYMBOLS    = "!#$%&()*+-;<=>?@^_`{|}~"
)

var (
	ALPHANUMERIC = ALPHABETIC + NUMERIC
	// rfc1924
	BASE85 = ALPHANUMERIC + strings.ToLower(ALPHABETIC) + SYMBOLS

	// rfc4648 alphabets
	BASE16    = BASE85[:16]
	BASE32    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	BASE32HEX = BASE85[:32]
	BASE64    = ALPHABETIC + strings.ToLower(ALPHABETIC) + NUMERIC + "+/"
	BASE64URL = BASE64[:62] + "-_"

	// Only contain alphanumeric characters
	BASE62 = BASE85[:62]
)

func RandomString(minLength, maxLength int) string {
	return RandomFilteredString(minLength, maxLength, BASE62)
}

// StringA returns a random string with length from minLength to maxLength
// containing only characters from provided alphabet.
// Will return empty string if maxLength < 0
func RandomFilteredString(minLength, maxLength int, alphabet string) string {
	var length int
	if minLength < 0 {
		minLength = 0
	}
	if maxLength < 0 {
		return ""
	}
	length = Int(minLength, maxLength)
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(buf)
}
