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

// Translate 翻译
func Translate(c *gin.Context) {
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKeyCheck := utils.IsValidApiKey(reqBody["apiKey"].(string))
	if !apiKeyCheck {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
		return
	}

	// 创建 OpenAI 客户端
	client := openai.NewClient(reqBody["apiKey"].(string))
	ctx := context.Background()

	//原语言
	origin, ok := reqBody["from"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	//目标语言
	goal, ok := reqBody["to"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	//情感
	emotion, ok := reqBody["content"].(map[string]interface{})["emotion"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	//模型的id
	modelStr, ok := reqBody["modelId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}
	//文体
	style, ok := reqBody["content"].(map[string]interface{})["style"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}

	slice := []string{origin, goal, emotion, modelStr, style}

	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}

	modelId, err := strconv.Atoi(modelStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}

	if modelId < 0 || modelId > 6 {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1005"})
		return
	}

	if emotion == "" {
		emotion = "normal"
	}

	if style == "" {
		style = "normal"
	}

	if origin == "" {
		origin = "depended by yourself"
	}

	//设置翻译的身份
	system := "You are a translator and talker who can translate sentences with a given emotion and style. You can't change the meaning of the original sentence because of emotion and style."
	prompt := "Translate this text into " + goal + " with " + emotion + " tone," + style + " literary form and its original language which is " + origin + ":"
	//翻译的语句
	user, ok := reqBody["content"].(map[string]interface{})["preTranslate"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "1010"})
		return
	}

	prompt += user

	req := openai.ChatCompletionRequest{
		Model:     Model[modelId],
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: system,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Translate this text into English with positive tone,normal literary form and its original language which is chinese:我很高兴见到你。",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "I am delighted to meet you.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Translate this text into English with negative tone,normal literary form and its original language which is chinese:我很高兴见到你。",
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
		errCode := utils.GPTRequestErrorCode(err)
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: errCode})
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
