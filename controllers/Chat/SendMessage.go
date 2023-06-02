package Chat

import (
	"context"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/controllers"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

func SendMessage(c *gin.Context) {
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKey, ok := reqBody["apiKey"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	user, ok := reqBody["content"].(map[string]interface{})["user"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	chatId, ok := reqBody["chatId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}

	slice := []string{apiKey, user, chatId}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

	apiKeyCheck := utils.IsValidApiKey(apiKey)
	if !apiKeyCheck {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	var systemMessage models.ChatMessage
	err := utils.DB.Table("chatmessage").Where("chatid = ? AND actor = ?", chatId, "system").Find(&systemMessage).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	if systemMessage.Content == "" {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1005"})
		return
	}

	var modelId int
	err = utils.DB.Table("chatconversation").Where("id = ?", chatId).Select("modelid").Scan(&modelId).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	var history []string
	err = utils.DB.Table("chatmessage").Where("chatId = ? AND (actor = ? OR actor = ?)", chatId, "user", "assistant").Select("content").Scan(&history).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	message := append(getMessage(history), openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: user,
	})

	message = append([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemMessage.Content,
		},
	}, message...)

	req := openai.ChatCompletionRequest{
		Model:    controllers.Model[modelId],
		Messages: message,
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		errCode := utils.GPTRequestErrorCode(err)
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
		return
	}
	assistant := resp.Choices[0].Message.Content

	err = utils.DB.Table("chatmessage").Create(&models.ChatMessage{
		Content: user,
		ChatId:  chatId,
		Actor:   "user",
		Show:    1,
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	err = utils.DB.Table("chatmessage").Create(&models.ChatMessage{
		Content: assistant,
		ChatId:  chatId,
		Actor:   "assistant",
		Show:    1,
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chatId": chatId,
		"content": gin.H{
			"user":      user,
			"assistant": assistant,
		},
	})
}

func getMessage(history []string) []openai.ChatCompletionMessage {
	message := make([]openai.ChatCompletionMessage, 0)
	var begin int

	if len(history) < 10 {
		begin = 0
	} else {
		begin = len(history) - 10
	}

	for i := begin; i < len(history); i++ {
		var newMessage openai.ChatCompletionMessage

		if i%2 == 0 {
			newMessage.Role = openai.ChatMessageRoleUser
			newMessage.Content = history[i]
		} else {
			newMessage.Role = openai.ChatMessageRoleAssistant
			newMessage.Content = history[i]
		}
		message = append(message, newMessage)
	}
	return message
}
