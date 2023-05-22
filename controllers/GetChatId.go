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
	ChatId string `json:"chatId"`
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
	uid, err := utils.Encrypt(requestGetChatId.ApiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
		return
	}
	apiKeyCheck := utils.IsValidApiKey(requestGetChatId.ApiKey)
	if apiKeyCheck {
		if errRequestGetChatId == nil {
			slice := []string{requestGetChatId.ApiKey}
			if !utils.Utf8Check(slice) {
				c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
				return
			}
			if err := utils.DB.Table("chatconversation").Where("uid=?", uid).Find(&chatConservations).Error; err == nil {
				if len(chatConservations) > 0 {
					for ; i < len(chatConservations); i++ {
						var temp Chat
						temp.Title = chatConservations[i].Title
						temp.ChatId = chatConservations[i].Id
						chatTemps = append(chatTemps, temp)
					}
				}
			} else {
				c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
				return
			}
		}
		responseGetChatId.Chat = chatTemps
		c.JSON(http.StatusOK, responseGetChatId)
	} else {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
	}
}
