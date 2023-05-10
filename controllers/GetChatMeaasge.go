package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type RequestGetMessage struct {
	Apikey string `json:"apiKey"`
	ChatId int    `json:"chatId"`
}

type ResponseGetMessage struct {
	ChatId   int       `gorm:"column:chatid" json:"chatId"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func GetChatMessage(c *gin.Context) {
	var requestGetMessage RequestGetMessage
	var responseGetMessage ResponseGetMessage
	var errRequestGetMessage = c.ShouldBindJSON(&requestGetMessage)
	var chatMessages []models.ChatMessage

	if errRequestGetMessage == nil {
		if err := utils.DB.Table("chatmessage").Where("chatid = ? AND show = ?", requestGetMessage.ChatId, 1).Order("id ASC").Find(&chatMessages).Error; err == nil {
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
}
