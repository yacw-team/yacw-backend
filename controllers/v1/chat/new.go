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

type NewChatRequest struct {
	ApiKey  string
	ModelId string
	Content Content
}

type NewChatResponse struct {
	ChatId  string             `json:"chatId"`
	ModelId string             `json:"modelId"`
	Content ResponseNewContent `json:"content"`
}

type Content struct {
	PersonalityId string `json:"personalityId"`
	User          string `json:"user"`
}

type ResponseNewContent struct {
	PersonalityId string `json:"personalityId"`
	User          string `json:"user"`
	Assistant     string `json:"assistant"`
	Title         string `json:"title"`
}

// 新建对话API，路由/v1/chat/new
func NewChat(c *gin.Context) {
	var request NewChatRequest
	var response NewChatResponse
	var err = c.BindJSON(&request)
	var errChange1, errPersonality error
	var max, modelId int
	var maxMessage, personalityId int
	var chatConversation models.ChatConversation
	var systemMessage, userMessage, assistantMessage models.ChatMessage
	var systemContent string
	var assistantResponse, title openai.ChatCompletionResponse
	//获取modelid，转int
	if request.ModelId == "" {
		modelId = 0
	} else {
		temp, errModelId := strconv.Atoi(request.ModelId)
		modelId = temp
		if errModelId != nil {
			c.String(http.StatusInternalServerError, "转换失败")
		}
	}
	if err == nil {
		uid := utils.EncryptPassword(request.ApiKey)
		//获取最大的对话id
		utils.DB.Table("chatconversation").Select("max(id) a").Find(&max)
		//获取最大的消息id
		utils.DB.Table("chatmessage").Select("max(id) a").Find(&maxMessage)
		chatConversation.Id = max + 1
		chatConversation.Uid = uid
		chatConversation.ModelId = modelId
		//人格设置，默认你是个帮手
		if request.Content.PersonalityId == "" {
			systemContent = "You are a helper."
		} else {
			personalityId, errPersonality = strconv.Atoi(request.Content.PersonalityId)
			var personality models.Personality
			utils.DB.Table("personality").Where("id=?", personalityId).Find(&personality)
			systemContent = personality.Prompts
		}
		if errChange1 != nil || errPersonality != nil {
			c.JSON(http.StatusInternalServerError, "转换失败")
		} else {
			systemMessage.Id = maxMessage + 1
			systemMessage.Content = systemContent
			systemMessage.ChatId = chatConversation.Id
			systemMessage.Actor = "system"
			systemMessage.Show = 0
			userMessage.Id = maxMessage + 2
			userMessage.Content = request.Content.User
			userMessage.ChatId = chatConversation.Id
			userMessage.Actor = "user"
			userMessage.Show = 1
			assistantMessage.Id = maxMessage + 3
			assistantMessage.ChatId = chatConversation.Id
			assistantMessage.Actor = "assistant"
			assistantMessage.Show = 1
			//插入系统消息，用户消息，回答的消息
			utils.DB.Table("chatmessage").Create(&systemMessage)
			utils.DB.Table("chatmessage").Create(&userMessage)
			assistantResponse = ChattingWithGPT(request.ApiKey, request.Content.User, systemContent, modelId)
			assistantMessage.Content = assistantResponse.Choices[0].Message.Content
			utils.DB.Table("chatmessage").Create(&assistantMessage)
			titleString := request.Content.User + ";帮我给这句话取个不是病句、简洁的中文标题"
			title = ChattingWithGPT(request.ApiKey, titleString, systemContent, modelId)
			chatConversation.Title = title.Choices[0].Message.Content
			//插入对话
			utils.DB.Table("chatconversation").Create(&chatConversation)
			//返回请求体
			response.Content.Assistant = assistantResponse.Choices[0].Message.Content
			response.Content.Title = title.Choices[0].Message.Content
			response.Content.User = request.Content.User
			response.Content.PersonalityId = request.Content.PersonalityId
			response.ModelId = request.ModelId
			response.ChatId = strconv.Itoa(max + 1)
			c.JSON(http.StatusOK, response)
		}

	} else {
		c.String(http.StatusInternalServerError, "数据绑定失败")
	}
}

// 调用chatgpt的函数参数有apikey，问题，人格，模型种类
func ChattingWithGPT(apiKey string, question string, system string, modelId int) openai.ChatCompletionResponse {
	client := openai.NewClient(apiKey)
	var model string
	if modelId == 0 {
		model = openai.GPT3Dot5Turbo
	} else if modelId == 1 {
		model = openai.GPT432K0314
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	return resp
}
