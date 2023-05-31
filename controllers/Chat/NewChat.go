package Chat

import (
	"context"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/controllers"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

type NewChatResponse struct {
	ChatId  string             `json:"chatId"`
	ModelId string             `json:"modelId"`
	Content ResponseNewContent `json:"content"`
}

type Content struct {
	PersonalityId int    `json:"personalityId"`
	User          string `json:"user"`
}

type ResponseNewContent struct {
	PersonalityId int    `json:"personalityId"`
	User          string `json:"user"`
	Assistant     string `json:"assistant"`
	Title         string `json:"title"`
}

// NewChat 新建对话API，路由/v1/chat/new
func NewChat(c *gin.Context) {
	var response NewChatResponse
	var maxMessage int
	var chatConversation models.ChatConversation
	var systemMessage, userMessage, assistantMessage models.ChatMessage
	var systemContent string
	var assistantResponse, title openai.ChatCompletionResponse
	var err error

	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKey := reqBody["apiKey"].(string)
	modelStr := reqBody["modelId"].(string)
	chatId := reqBody["chatId"].(string)
	personalityId_s := reqBody["content"].(map[string]interface{})["personalityId"].(string)
	user := reqBody["content"].(map[string]interface{})["user"].(string)

	slice := []string{apiKey, modelStr, user}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

	modelId, err := strconv.Atoi(modelStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2005"})
		return
	}

	personalityId, _ := strconv.Atoi(personalityId_s)

	if modelId < 0 || modelId > 6 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1005"})
		return
	}

	apiKeyCheck := utils.IsValidApiKey(apiKey)
	if !apiKeyCheck {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	uid, err := utils.Encrypt(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}
	//获取最大的消息id
	err = utils.DB.Table("chatmessage").Select("COALESCE(MAX(id), 0)").Find(&maxMessage).Error
	if err != nil {
		maxMessage = 0
	}
	chatConversation.Id = chatId
	chatConversation.Uid = uid
	chatConversation.ModelId = modelId
	//人格设置，默认你是个帮手
	if personalityId == 0 {
		systemContent = "You are a helper."
	} else {
		var personality models.Personality
		err = utils.DB.Table("personality").Where("id=?", personalityId).Find(&personality).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
			return
		}
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

	assistantResponse, openAIerr_title := ChattingWithGPT(apiKey, user, systemContent, modelId)
	titleString := "帮我根据以下的文本想一个标题（注意直接返回一个标题，我想直接使用，正式一些，字数在4-6个字）：" + user
	title, openAIerr_ass := ChattingWithGPT(apiKey, titleString, systemContent, modelId)

	err = utils.DB.Table("chatmessage").Create(&systemMessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}
	err = utils.DB.Table("chatmessage").Create(&userMessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	if openAIerr_title != nil || openAIerr_ass != nil {
		var errCode string
		chatConversation.Title = ""
		//插入对话
		err = utils.DB.Table("chatconversation").Create(&chatConversation).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
			return
		}
		if openAIerr_title != nil {
			errCode = utils.GPTRequestErrorCode(openAIerr_title)
		} else {
			errCode = utils.GPTRequestErrorCode(openAIerr_ass)
		}

		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
		return
	}

	assistantMessage.Content = assistantResponse.Choices[0].Message.Content
	utils.DB.Table("chatmessage").Create(&assistantMessage)

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
	response.ChatId = chatId
	c.JSON(http.StatusOK, response)

}

// ChattingWithGPT 调用chatgpt的函数参数有apikey，问题，人格，模型种类
func ChattingWithGPT(apiKey string, question string, system string, modelId int) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: controllers.Model[modelId],
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
