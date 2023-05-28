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

// SendMessage 发送对话
func SendMessage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2007"})
			// 进行适当的处理
		}
	}()
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	//获取数据
	apiKey := reqBody["apiKey"].(string)
	user := reqBody["content"].(map[string]interface{})["user"].(string)
	chatId := reqBody["chatId"].(string)

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

	// 创建 OpenAI 客户端
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	//查找第一句system
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

	//查找modelId
	var modelId int
	err = utils.DB.Table("chatconversation").Where("id = ?", chatId).Select("modelid").Scan(&modelId).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	//查找历史的对话
	var history []string
	err = utils.DB.Table("chatmessage").Where("chatId = ? AND (actor = ? OR actor = ?)", chatId, "user", "assistant").Select("content").Scan(&history).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	//在最后加入用户的新对话
	message := append(getMessage(history), openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: user,
	})

	//在开头加入system字段
	message = append([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemMessage.Content,
		},
	}, message...)

	//构造请求体
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
	//获取回复
	assistant := resp.Choices[0].Message.Content

	//将用户的对话写入数据库
	err = utils.DB.Table("chatmessage").Create(&models.ChatMessage{
		Content: user,
		ChatId:  chatId,
		Actor:   "user", //代表是用户
		Show:    1,      //代表要展示

	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	//将API的回复写入数据库
	err = utils.DB.Table("chatmessage").Create(&models.ChatMessage{
		Content: assistant,
		ChatId:  chatId,
		Actor:   "assistant", //代表是回复
		Show:    1,           //代表要展示

	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
		return
	}

	//将用户的对话和API的回复存入数据库
	c.JSON(http.StatusOK, gin.H{
		"chatId": chatId,
		"content": gin.H{
			"user":      user,
			"assistant": assistant,
		},
	})
}

// 包装历史信息
func getMessage(history []string) []openai.ChatCompletionMessage {
	message := make([]openai.ChatCompletionMessage, 0)
	var begin int

	if len(history) < 10 {
		begin = 0
	} else {
		begin = len(history) - 10
	}

	//包含历史的后10条对话，一问一答
	for i := begin; i < len(history); i++ {
		var newMessage openai.ChatCompletionMessage
		//偶数是用户的,奇数是AI回复
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
