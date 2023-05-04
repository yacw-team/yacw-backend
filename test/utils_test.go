package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/utils"
	"testing"
)

// 测试加密函数
func TestEncryptPassword(t *testing.T) {
	input := "password"
	expectedOutput := "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
	output := utils.EncryptPassword(input)
	assert.Equal(t, expectedOutput, output)
}
