package Chat

import (
	"github.com/gin-gonic/gin"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/utils"
	"net/http"
)

type DeleteChatReqBody struct {
	ApiKey string `json:"apiKey"`
	ChatId string `json:"chatId"`
}

func DeleteChat(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2007"})
			// 进行适当的处理
		}
	}()
	var reqBody map[string]interface{}
	reqTemp, ok := c.Get("reqBody")
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "2006"})
		return
	}
	reqBody = reqTemp.(map[string]interface{})

	apiKey := reqBody["apiKey"].(string)
	chatId := reqBody["chatId"].(string)

	slice := []string{apiKey, chatId}
	if !utils.Utf8Check(slice) {
		c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1011"})
		return
	}
	var chatConversation models.ChatConversation
	apiKeyCheck := utils.IsValidApiKey(apiKey)
	if apiKeyCheck {
		uid, err := utils.Encrypt(apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrCode{ErrCode: "3006"})
			return
		}

		utils.DB.Table("chatconversation").Where("id=? and uid=?", chatId, uid).Find(&chatConversation)
		if chatConversation.Uid == "" {
			c.JSON(http.StatusBadRequest, models.ErrCode{ErrCode: "1007"})
		} else {
			var temp []models.ChatMessage
			utils.DB.Table("chatmessage").Where("chatid=?", chatId).Find(&temp)
			utils.DB.Table("chatconversation").Delete(&chatConversation)
			utils.DB.Table("chatmessage").Delete(&temp)
			c.JSON(http.StatusOK, models.ErrCode{ErrCode: "0000"})
		}

	} else {
		var errCode models.ErrCode
		errCode.ErrCode = "3004"
		c.JSON(http.StatusBadRequest, errCode)
	}
}
