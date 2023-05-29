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

		// 将字符数据解析为map[string]interface{}类型
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		// 检查解析后的JSON数据是否符合预期格式
		if IsValidResult(result) {
			// 返回JSON响应
			break
		}
		time.Sleep(3 * time.Second)
	}

	return result, nil
}
