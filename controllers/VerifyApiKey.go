package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
)

func VerifyApiKey(c *gin.Context) {

	apiKey := c.PostForm("apiKey")

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "",
			},
		},
	}

	_, err := client.CreateChatCompletion(ctx, req)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.JSON(http.StatusOK, "OK")
}
