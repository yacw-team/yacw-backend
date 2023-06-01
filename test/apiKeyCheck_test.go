package test

import (
	"fmt"
	"github.com/yacw-team/yacw/utils"
	"testing"
)

func TestApiKeyCheck(t *testing.T) {
	str1 := "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"
	if utils.IsValidApiKey(str1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
