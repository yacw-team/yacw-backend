package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// 加盐的长度
var saltLength int64 = 8

// 截断的长度
var hashLength int64 = 16

// HashAndSalt 对apikey进行加盐 哈希 截断
func HashAndSalt(input string) string {
	// 生成随机盐
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	// 将盐和输入字符串拼接起来
	saltedInput := input + string(salt)

	// 计算哈希值
	hash := sha256.Sum256([]byte(saltedInput))

	// 截取哈希值的前 hashLength 个字节
	truncatedHash := hash[:hashLength]

	// 将盐和截断后的哈希值进行 Base64 编码，并拼接在一起作为最终结果
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	encodedHash := base64.StdEncoding.EncodeToString(truncatedHash)
	result := encodedSalt + encodedHash

	return result
}
