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
	} else if strings.Contains(errStr, "An existing connection was forcibly closed by the remote host") {
		errCode = "3026"
	} else if strings.Contains(errStr, "A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond") {
		errCode = "3027"
	}
	return errCode
}
