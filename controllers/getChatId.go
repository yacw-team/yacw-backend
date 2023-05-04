package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

type RequireGetChatId struct {
	ApiKey string `json:"apiKey"`
}

type Chat struct {
	ChatId string
	System string
}

type RespondGetChatId struct {
	chat Chat
}

func GetChatId(c *gin.Context) {
	var id int
	var requireGetChatId RequireGetChatId
	var respondGetChatId RespondGetChatId
	var errRequireGetChatId = c.BindJSON(&requireGetChatId)
	var chatConservation []models.ChatConversation
	var chatMessage []models.ChatMessage
	utils.DB.Table("chatconversation").Find(&chatConservation)
	utils.DB.Table("chatmessage").Find(&chatMessage)
	var i = 0
	var j = 0
	if errRequireGetChatId == nil {
		for ; i < len(chatConservation); i++ {
			if requireGetChatId.ApiKey == chatConservation[i].Uid {
				respondGetChatId.chat.System = chatConservation[i].SystemPrompt
				id = chatConservation[i].Id
			}
		}
		for ; j < len(chatMessage); j++ {
			if id == chatMessage[j].Id {
				respondGetChatId.chat.ChatId = strconv.Itoa(chatMessage[j].ChatId)
			}
		}
	}
	c.JSON(http.StatusOK, respondGetChatId)
}
