package controllers

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

// NewChat 新建对话API，路由/v1/chat/new
func NewChat(c *gin.Context) {
	var response NewChatResponse
	var max int
	var maxMessage int
	var chatConversation models.ChatConversation
	var systemMessage, userMessage, assistantMessage models.ChatMessage
	var systemContent string
	var assistantResponse, title openai.ChatCompletionResponse
	var err error

	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if ok == false {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKey := reqBody["apiKey"].(string)
	modelStr := reqBody["modelId"].(string)
	modelId, err := strconv.Atoi(modelStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}
	personalityId := reqBody["content"].(map[string]interface{})["personalityId"].(string)
	user := reqBody["content"].(map[string]interface{})["user"].(string)

	uid := utils.Encrypt(apiKey)
	//获取最大的对话id
	err = utils.DB.Table("chatconversation").Select("COALESCE(MAX(id), 0)").Row().Scan(&max)
	if err != nil {
		max = 0
	}
	//获取最大的消息id
	err = utils.DB.Table("chatmessage").Select("COALESCE(MAX(id), 0)").Find(&maxMessage).Error
	if err != nil {
		maxMessage = 0
	}
	chatConversation.Id = max + 1
	chatConversation.Uid = uid
	chatConversation.ModelId = modelId
	//人格设置，默认你是个帮手
	if personalityId == "" {
		systemContent = "You are a helper."
	} else {
		var personality models.Personality
		utils.DB.Table("personality").Where("id=?", personalityId).Find(&personality)
		systemContent = personality.Prompts
	}

	systemMessage.Id = maxMessage + 1
	systemMessage.Content = systemContent
	systemMessage.ChatId = chatConversation.Id
	systemMessage.Actor = "system"
	systemMessage.Show = 0
	userMessage.Id = maxMessage + 2
	userMessage.Content = user
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
	assistantResponse, err = ChattingWithGPT(apiKey, user, systemContent, modelId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3001"})
		return
	}

	assistantMessage.Content = assistantResponse.Choices[0].Message.Content
	utils.DB.Table("chatmessage").Create(&assistantMessage)
	titleString := "帮我根据以下的文本想一个标题（注意直接返回一个标题，我想直接使用，正式一些）：" + user
	title, err = ChattingWithGPT(apiKey, titleString, systemContent, modelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3001"})
		return
	}

	chatConversation.Title = title.Choices[0].Message.Content
	//插入对话
	err = utils.DB.Table("chatconversation").Create(&chatConversation).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	//返回请求体
	response.Content.Assistant = assistantResponse.Choices[0].Message.Content
	response.Content.Title = title.Choices[0].Message.Content
	response.Content.User = user
	response.Content.PersonalityId = personalityId
	response.ModelId = strconv.Itoa(modelId)
	response.ChatId = strconv.Itoa(max + 1)
	c.JSON(http.StatusOK, response)

}

// ChattingWithGPT 调用chatgpt的函数参数有apikey，问题，人格，模型种类
func ChattingWithGPT(apiKey string, question string, system string, modelId int) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model[modelId],
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
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