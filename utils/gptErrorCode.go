package utils

import "strings"

func GPTRequestErrorCode(err error) string {
	errStr := err.Error()
	errCode := "3001"
	if strings.Contains(errStr, "401") {
		errCode = "3021"
	} else if strings.Contains(errStr, "403") {
		errCode = "3022"
	} else if strings.Contains(errStr, "404") {
		errCode = "3023"
	} else if strings.Contains(errStr, "429") {
		errCode = "3024"
	} else if strings.Contains(errStr, "500") {
		errCode = "3025"
	}
	return errCode
}
