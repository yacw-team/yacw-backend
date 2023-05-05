package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

type DeleteRequest struct {
	ApiKey string `json:"apiKey"`
	ChatId string `json:"chatId"`
}

func DeleteChat(c *gin.Context) {
	var request DeleteRequest
	var err = c.BindJSON(&request)
	var chatConversation models.ChatConversation
	if err == nil {
		id, errChange := strconv.Atoi(request.ChatId)
		uid := utils.EncryptPassword(request.ApiKey)
		if errChange == nil {
			utils.DB.Table("chatconversation").Where("id=? and uid=?", id, uid).Find(&chatConversation)
			if chatConversation.Uid == "" {
				c.String(http.StatusBadRequest, "查无此对话")
			} else {
				var temp []models.ChatMessage
				utils.DB.Table("chatmessage").Where("chatid=?", id).Find(&temp)
				utils.DB.Table("chatconversation").Delete(&chatConversation)
				utils.DB.Table("chatmessage").Delete(&temp)
				c.String(http.StatusOK, "200 OK")
			}
		} else {
			c.String(http.StatusInternalServerError, "类型转换失败")
		}
	} else {
		c.String(http.StatusInternalServerError, "数据绑定失败")
	}
}
