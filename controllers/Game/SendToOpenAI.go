package Game

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/controllers"
)

func SendToOpenAI(message []openai.ChatCompletionMessage, modelId int, apiKey string) (string, error) {

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:    controllers.Model[modelId],
		Messages: message,
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	assistant := resp.Choices[0].Message.Content
	return assistant, nil
}
