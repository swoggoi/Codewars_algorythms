package kata

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Crack(hash string) string {
	for i := 0; i < 100000; i++ {
		pin := fmt.Sprintf("%05d", i)
		sum := md5.Sum([]byte(pin))
		hashed := hex.EncodeToString(sum[:])
		if hashed == hash {
			return pin
		}
	}
	return ""
}
