package test

import (
	"fmt"
	"github.com/yacw-team/yacw/utils"
	"testing"
)

func TestEncrypt(t *testing.T) {
	input := "password"
	output := utils.Encrypt(input)
	fmt.Println(output)
}
