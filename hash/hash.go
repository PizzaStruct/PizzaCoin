package hash

import (
	"crypto/sha256"
	"fmt"
)

func GetSHA256(in []byte) string {
	hashsum := sha256.Sum256(in)
	return fmt.Sprintf("%x", hashsum)
}
