package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
)

// 绑定的结构体
type formData struct {
	APIKey  string `json:"apiKey"`
	Content struct {
		Emotion      string `json:"emotion"`
		Style        string `json:"style"`
		PreTranslate string `json:"preTranslate"`
	} `json:"content"`
	From string `json:"from"`
	To   string `json:"to"`
}

// Translate 翻译
func Translate(c *gin.Context) {
	var formData formData
	err := c.BindJSON(&formData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "数据绑定错误")
		return
	}
	fmt.Println(formData.APIKey + "知识")
	// 创建 OpenAI 客户端
	client := openai.NewClient(formData.APIKey)
	ctx := context.Background()

	//原语言
	origin := formData.From
	//目标语言
	goal := formData.To

	emotion := formData.Content.Emotion

	if emotion == "" {
		emotion = "normal"
	}

	style := formData.Content.Style

	if style == "" {
		style = "normal"
	}

	//设置翻译的身份
	system := "You are a translator who can translate sentences with a given emotion and style. You can't change the meaning of the original sentence because of emotion and style."
	prompt := "Translate this text into " + goal + " with a " + emotion + " tone: "
	//翻译的语句
	user := formData.Content.PreTranslate

	prompt += user

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: system,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Translate this text into English with a positive tone:我很高兴见到你。",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "I am delighted to meet you.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Translate this text into English with a negative tone:我很高兴见到你。",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "I guess it's nice to meet you.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": gin.H{
			"emotion":      emotion,
			"style":        style,
			"preTranslate": user,
			"translated":   resp.Choices[0].Message.Content,
		},
		"from": origin,
		"to":   goal,
	})
}
