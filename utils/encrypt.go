package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// EncryptPassword 用sha256对传入的字符串进行加密
func EncryptPassword(password string) string {
	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return ""
	}
	encrypted := hex.EncodeToString(hash.Sum(nil))
	return encrypted
}
