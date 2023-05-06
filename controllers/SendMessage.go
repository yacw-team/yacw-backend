package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

// 接受json的格式
type reqMessage struct {
	ApiKey  string `json:"apiKey"`
	ChatId  string `json:"chatId"`
	Content struct {
		User string `json:"user"`
	}
}

// 写入数据库的message的结构
type chatMessage struct {
	content string `gorm:"content"`
	chatId  string `gorm:"chatid"`
	actor   int    `gorm:"actor"`
	show    int    `gorm:"show"`
}

// SendMessage 发送对话
func SendMessage(c *gin.Context) {
	var reqMessage reqMessage
	err := c.BindJSON(reqMessage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据绑定错误")
		return
	}

	//获取数据
	apiKey := reqMessage.ApiKey
	chatId := reqMessage.ChatId
	user := reqMessage.Content.User

	// 创建 OpenAI 客户端
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	//查找system字段
	var system string
	err = utils.DB.Table("chatmessage").Where("chatId = ? AND actor = ?", chatId, 0).Select("content").Scan(&system).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据库查询错误")
		return
	}

	//查找历史的对话
	var history []string
	err = utils.DB.Table("chatmessage").Where("chatId = ? AND (actor = ? OR actor = ?)", chatId, 1, 2).Select("content").Scan(&history).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据库查询错误")
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
			Content: system,
		},
	}, message...)

	//构造请求体
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 100,
		Messages:  message,
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	//获取回复
	assistant := resp.Choices[0].Message.Content

	//将用户的对话写入数据库
	err = utils.DB.Table("chatmessage").Create(chatMessage{
		content: user,
		chatId:  chatId,
		actor:   1, //代表是用户
		show:    1, //代表要展示

	}).Error

	//将API的回复写入数据库
	err = utils.DB.Table("chatmessage").Create(chatMessage{
		content: assistant,
		chatId:  chatId,
		actor:   2, //代表是用户
		show:    1, //代表要展示

	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据库存储失败")
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
	var message []openai.ChatCompletionMessage
	var begin int

	if len(history) < 10 {
		begin = 0
	} else {
		begin = len(history) - 10
	}

	//包含历史的后10条对话，一问一答
	for i := begin; i < len(history); i++ {
		//偶数是用户的,奇数是AI回复
		if i%2 == 0 {
			message[i].Role = openai.ChatMessageRoleUser
			message[i].Content = history[i]
		} else {
			message[i].Role = openai.ChatMessageRoleAssistant
			message[i].Content = history[i]
		}
	}
	return message
}
