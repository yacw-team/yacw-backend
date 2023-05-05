package chat

import (
	"github.com/gin-gonic/gin"
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
	ChatId  string  `json:"chatId"`
	ModelId string  `json:"modelId"`
	Content Content `json:"content"`
}

type Content struct {
	PersonalityId string `json:"personalityId"`
	PromptsId     string `json:"promptsId"`
	System        string `json:"system"`
}

func NewChat(c *gin.Context) {
	var request NewChatRequest
	var response NewChatResponse
	var err = c.BindJSON(&request)
	var errChange, errChange1 error
	var max int
	if err == nil {
		var chatConversation models.ChatConversation
		uid := utils.EncryptPassword(request.ApiKey)
		utils.DB.Table("chatconversation").Select("max(id) a").Find(&max)
		chatConversation.Id = max + 1
		chatConversation.SystemPrompt = request.Content.System
		chatConversation.Uid = uid
		chatConversation.ModelId, errChange = strconv.Atoi(request.ModelId)
		chatConversation.PromptId, errChange1 = strconv.Atoi(request.Content.PromptsId)
		if errChange != nil || errChange1 != nil {
			c.JSON(http.StatusInternalServerError, "转换失败")
		} else {
			utils.DB.Table("chatconversation").Create(&chatConversation)
			chatConversation.SystemPrompt = request.Content.System
			response.Content.System = request.Content.System
			response.Content.PromptsId = request.Content.PromptsId
			response.Content.PersonalityId = request.Content.PersonalityId
			response.ModelId = request.ModelId
			response.ChatId = strconv.Itoa(max + 1)
			c.JSON(http.StatusOK, response)
		}
	} else {
		c.String(http.StatusInternalServerError, "数据绑定失败")
	}
}
