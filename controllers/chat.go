package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/context"
	"net/http"
	"os"
)

func Chat(c *gin.Context) {
	userComment, isCommentOK := c.GetQuery("comment")
	userKey, isKeyOK := c.GetQuery("key")
	token := os.Getenv(userKey)
	if isCommentOK && isKeyOK {
		client := openai.NewClient(token)
		response, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: userComment,
					},
				},
			},
		)
		if err != nil {
			return
		}
		c.String(http.StatusOK, response.Choices[0].Message.Content)
	}

}
