package Chat

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"strconv"
)

type RequestGetMessage struct {
	Apikey  string `json:"apiKey"`
	ChatStr string `json:"chatId"`
}

type ResponseGetMessage struct {
	ChatId   string    `gorm:"column:chatid" json:"chatId"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func GetChatMessage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2007"})
		}
	}()
	var requestGetMessage RequestGetMessage
	var responseGetMessage ResponseGetMessage
	var errRequestGetMessage = c.ShouldBindJSON(&requestGetMessage)
	var chatMessages []models.ChatMessage
	apiKeyCheck := utils.IsValidApiKey(requestGetMessage.Apikey)
	if apiKeyCheck {
		if errRequestGetMessage == nil {
			slice := []string{requestGetMessage.ChatStr, requestGetMessage.Apikey}
			if !utils.Utf8Check(slice) {
				c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
				return
			}
			chatId, err := strconv.Atoi(requestGetMessage.ChatStr)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
				return
			}
			if err = utils.DB.Table("chatmessage").Where("chatid = ? AND show = ?", chatId, 1).Order("id ASC").Find(&chatMessages).Error; err == nil {
				if len(chatMessages) > 0 {
					responseGetMessage.ChatId = chatMessages[0].ChatId
					for i := 0; i < len(chatMessages); i++ {
						responseGetMessage.Messages = append(responseGetMessage.Messages, Message{
							Type:    chatMessages[i].Actor,
							Content: chatMessages[i].Content,
						})

					}
				}
			} else {
				c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3009"})
				return
			}
		}
		c.JSON(http.StatusOK, responseGetMessage)
	} else {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
	}
}
