package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yacw-team/yacw/models"
	"net/http"
	"strconv"
)

// Translate 翻译
func Translate(c *gin.Context) {
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if ok == false {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	// 创建 OpenAI 客户端
	client := openai.NewClient(reqBody["apiKey"].(string))
	ctx := context.Background()

	//原语言
	origin := reqBody["from"].(string)
	//目标语言
	goal := reqBody["to"].(string)
	//情感
	emotion := reqBody["content"].(map[string]interface{})["emotion"].(string)
	//模型的id
	modelId, err := strconv.Atoi(reqBody["modelId"].(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}

	if emotion == "" {
		emotion = "normal"
	}

	style := reqBody["content"].(map[string]interface{})["style"].(string)

	if style == "" {
		style = "normal"
	}

	//设置翻译的身份
	system := "You are a translator who can translate sentences with a given emotion and style. You can't change the meaning of the original sentence because of emotion and style."
	prompt := "Translate this text into " + goal + " with a " + emotion + " tone: "
	//翻译的语句
	user := reqBody["content"].(map[string]interface{})["preTranslate"].(string)

	prompt += user

	req := openai.ChatCompletionRequest{
		Model:     model[modelId],
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
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3001"})
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
