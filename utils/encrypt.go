package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"flag"
	"os"
)

// 截断的长度
const hashLength = 16

// Encrypt 对apikey进行加盐 哈希 截断
func Encrypt(input string) (string, error) {

	salt, err := getSalt()

	if err != nil {
		return "", errors.New(err.Error())
	}

	// 将盐和输入字符串拼接起来
	saltedInput := input + salt

	// 计算哈希值
	hash := sha256.Sum256([]byte(saltedInput))

	// 截取哈希值的前 hashLength 个字节
	truncatedHash := hash[:hashLength]

	// 将盐和截断后的哈希值进行 Base64 编码，并拼接在一起作为最终结果
	encodedSalt := base64.StdEncoding.EncodeToString([]byte(salt))
	encodedHash := base64.StdEncoding.EncodeToString(truncatedHash)
	result := encodedSalt + encodedHash

	return result, nil
}

func getSalt() (string, error) {

	//从环境变量获取盐
	salt := os.Getenv("SALT")
	if salt != "" {
		return salt, nil
	}
	//从命令行参数获取盐
	//变量指针、命令行参数的名称、默认值和命令行参数的描述
	flag.StringVar(&salt, "salt", "", "the salt value")
	if salt == "" {
		return "", errors.New("3006")
	}
	return salt, nil
}
