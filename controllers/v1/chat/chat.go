package chat

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

type ChattingRequest struct {
	ApiKey  string         `json:"apiKey"`
	ChatId  string         `json:"chatId"`
	Content RequestContent `json:"content"`
}

type RequestContent struct {
	User string `json:"user"`
}

type ChattingResponse struct {
	ChatId  string          `json:"chatId"`
	Content ResponseContent `json:"content"`
	Id      ResponseId      `json:"id"`
}

type ResponseContent struct {
	User      string `json:"user"`
	Assistant string `json:"assistant"`
}

type ResponseId struct {
	UserMsgId string `json:"usermsgid"`
	AssMsgId  string `json:"assmsgid"`
}

func SendChat(c *gin.Context) {
	var request ChattingRequest
	var response ChattingResponse
	err := c.BindJSON(&request)
	response.ChatId = request.ChatId
	response.Content.User = request.Content.User
	chatId, errChange := strconv.Atoi(request.ChatId)
	var max int
	var userMessage models.ChatMessage
	var assistantMessage models.ChatMessage
	var assistantResponse openai.ChatCompletionResponse
	if err == nil && errChange == nil {
		utils.DB.Table("chatmessage").Select("max(id) a").Find(&max)
		userMessage.Id = max + 1
		response.Id.UserMsgId = strconv.Itoa(max + 1)
		userMessage.Content = request.Content.User
		userMessage.ChatId = chatId
		userMessage.Actor = "user"
		utils.DB.Table("chatmessage").Create(&userMessage)
		assistantMessage.Id = max + 2
		assistantMessage.ChatId = chatId
		assistantMessage.Actor = "assistant"
		assistantResponse = ChattingWithGPT3Dot5Turbo(request.ApiKey, request.Content.User)
		assistantMessage.Content = assistantResponse.Choices[0].Message.Content
		response.Id.AssMsgId = strconv.Itoa(max + 2)
		response.Content.Assistant = assistantResponse.Choices[0].Message.Content
		utils.DB.Table("chatmessage").Create(&assistantMessage)
		c.JSON(http.StatusOK, response)
	} else {
		c.String(http.StatusInternalServerError, "类型转换或数据绑定失败")
	}
}

func ChattingWithGPT3Dot5Turbo(apiKey string, question string) openai.ChatCompletionResponse {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	return resp
}
