package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

type deleteChatReqBody struct {
	ApiKey string `json:"apiKey"`
	ChatId string `json:"chatId"`
}

func DeleteChat(c *gin.Context) {
	var request deleteChatReqBody
	var err = c.BindJSON(&request)
	var chatConversation models.ChatConversation
	if err == nil {
		id, errChange := strconv.Atoi(request.ChatId)
		uid := utils.Encrypt(request.ApiKey)
		if errChange == nil {
			utils.DB.Table("chatconversation").Where("id=? and uid=?", id, uid).Find(&chatConversation)
			if chatConversation.Uid == "" {
				c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1007"})
			} else {
				var temp []models.ChatMessage
				utils.DB.Table("chatmessage").Where("chatid=?", id).Find(&temp)
				utils.DB.Table("chatconversation").Delete(&chatConversation)
				utils.DB.Table("chatmessage").Delete(&temp)
				c.JSON(http.StatusOK, models.ErrCode{ErrCode: "0000"})
			}
		} else {
			c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1005"})
		}
	} else {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1010"})
	}
}
