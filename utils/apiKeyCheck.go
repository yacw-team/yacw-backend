package utils

import "regexp"

func IsValidApiKey(apiKey string) bool {
	// OpenAI API key由以"sk-"开头的51个字符组成，包含数字和字母（大小写敏感）
	regex := regexp.MustCompile(`^sk-[a-zA-Z0-9]{48}$`)
	return regex.MatchString(apiKey)
}
