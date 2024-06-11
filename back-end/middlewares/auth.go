package middlewares

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	res := fmt.Sprintf("%x", h.Sum(nil))
	return res
}
