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
	i := 0
	if errRequestGetMessage == nil {
		if err := utils.DB.Table("chatmessage").Where("chatid = ?", requestGetMessage.ChatId).Order("id ASC").Find(&chatMessages).Error; err == nil {
			if len(chatMessages) > 0 {
				responseGetMessage.ChatId = chatMessages[0].ChatId
				for ; i < len(chatMessages); i++ {
					if chatMessages[i].Show == 0 {
						responseGetMessage.Messages = append(responseGetMessage.Messages, Message{
							Type:    chatMessages[i].Actor,
							Content: chatMessages[i].Content,
						})
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK, responseGetMessage)
}
