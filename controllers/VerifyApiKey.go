package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

func VerifyApiKey(c *gin.Context) {

	apiKey := c.PostForm("apiKey")

	slice := []string{apiKey}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

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
		errCode := utils.GPTRequestErrorCode(err)
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
		return
	}

	c.JSON(http.StatusOK, models.ErrCode{ErrCode: "1002"})
}
