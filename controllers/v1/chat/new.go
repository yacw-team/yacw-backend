package chat

import (
	"context"
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

type ErrorCode struct {
	ErrorStatus string `json:"error_status"`
}

// 新建对话API，路由/v1/chat/new
func NewChat(c *gin.Context) {
	var request NewChatRequest
	var response NewChatResponse
	var err = c.BindJSON(&request)
	var errPersonality error
	var max, modelId int
	var maxMessage, personalityId int
	var chatConversation models.ChatConversation
	var systemMessage, userMessage, assistantMessage models.ChatMessage
	var systemContent string
	var assistantResponse, title openai.ChatCompletionResponse
	var gptError1, gptError2 error
	//获取modelid，转int
	if request.ModelId == "" {
		modelId = 0
	} else {
		temp, errModelId := strconv.Atoi(request.ModelId)
		modelId = temp
		if errModelId != nil {
			var errorStatus ErrorCode
			errorStatus.ErrorStatus = "1005"
			c.JSON(http.StatusBadRequest, errorStatus)
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
		if errPersonality != nil {
			var statusError ErrorCode
			statusError.ErrorStatus = "1005"
			c.JSON(http.StatusBadRequest, statusError)
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
			assistantResponse, gptError1 = ChattingWithGPT(request.ApiKey, request.Content.User, systemContent, modelId)
			if gptError1 != nil {
				var errorStatus ErrorCode
				errorStatus.ErrorStatus = "3001"
				c.JSON(http.StatusInternalServerError, errorStatus)
			} else {
				assistantMessage.Content = assistantResponse.Choices[0].Message.Content
				utils.DB.Table("chatmessage").Create(&assistantMessage)
				titleString := request.Content.User + ";帮我给这句话取个不是病句、简洁的中文标题"
				title, gptError2 = ChattingWithGPT(request.ApiKey, titleString, systemContent, modelId)
				if gptError2 != nil {
					var errorStatus ErrorCode
					errorStatus.ErrorStatus = "3001"
					c.JSON(http.StatusInternalServerError, errorStatus)
				} else {
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
			}
		}
	} else {
		var errorStatus ErrorCode
		errorStatus.ErrorStatus = "1010"
		c.JSON(http.StatusBadRequest, errorStatus)
	}
}

// 调用chatgpt的函数参数有apikey，问题，人格，模型种类
func ChattingWithGPT(apiKey string, question string, system string, modelId int) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(apiKey)
	var model string
	if modelId == 0 {
		model = openai.GPT3Dot5Turbo
	} else if modelId == 1 {
		model = openai.GPT3Dot5Turbo0301
	} else if modelId == 2 {
		model = openai.GPT4
	} else if modelId == 3 {
		model = openai.GPT432K
	} else if modelId == 4 {
		model = openai.GPT432K0314
	} else if modelId == 5 {
		model = openai.GPT40314
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
		return resp, err
	} else {
		return resp, nil
	}
}
