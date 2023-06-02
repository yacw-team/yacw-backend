package Game

import (
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"time"
)

func CheckAndReSend(message []openai.ChatCompletionMessage, modelId int, apiKey string) (map[string]interface{}, error) {
	var result map[string]interface{}

	for {
		data, err := SendToOpenAI(message, modelId, apiKey)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		if IsValidResult(result) {
			break
		}
		time.Sleep(3 * time.Second)
	}

	return result, nil
}
