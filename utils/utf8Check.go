package utils

import "unicode/utf8" //nolint:all

func Utf8Check(str []string) bool {
	for i := 0; i < len(str); i++ {
		if !utf8.ValidString(str[i]) && str[i] != "" {
			return false
		}
	}
	return true
}
