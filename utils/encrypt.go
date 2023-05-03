package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptPassword(password string) string {
	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return ""
	}
	encrypted := hex.EncodeToString(hash.Sum(nil))
	return encrypted
}
