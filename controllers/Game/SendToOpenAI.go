package Game

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/controllers"
)

func SendToOpenAI(message []openai.ChatCompletionMessage, modelId int, apiKey string) (string, error) {
	// 创建 OpenAI 客户端
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	//构造请求体
	req := openai.ChatCompletionRequest{
		Model:    controllers.Model[modelId],
		Messages: message,
	}

	//获取回复
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	//获取回复
	assistant := resp.Choices[0].Message.Content
	return assistant, nil
}
