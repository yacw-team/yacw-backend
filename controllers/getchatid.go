package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type RequestGetChatId struct {
	ApiKey string `json:"apiKey"`
}

type Chat struct {
	ChatId int    `json:"chatId"`
	Title  string `json:"title"`
}

type ResponseGetChatId struct {
	Chat []Chat `json:"chat"`
}

func GetChatId(c *gin.Context) {
	var requestGetChatId RequestGetChatId
	var responseGetChatId ResponseGetChatId
	var errRequestGetChatId = c.ShouldBindJSON(&requestGetChatId)
	var chatConservations []models.ChatConversation
	var chatTemps []Chat
	var i = 0
	Uid := utils.EncryptPassword(requestGetChatId.ApiKey)
	if errRequestGetChatId == nil {
		if err := utils.DB.Table("chatconversation").Where("uid=?", Uid).Find(&chatConservations).Error; err == nil {
			if len(chatConservations) > 0 {
				for ; i < len(chatConservations); i++ {
					var temp Chat
					temp.Title = chatConservations[i].Title
					temp.ChatId = chatConservations[i].Id
					chatTemps = append(chatTemps, temp)
				}
			}
		}
	}
	responseGetChatId.Chat = chatTemps
	c.JSON(http.StatusOK, responseGetChatId)
}
