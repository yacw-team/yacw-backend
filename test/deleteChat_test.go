package test

import (
	"bytes"
	json2 "encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/controllers"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteChat(t *testing.T) {
	utils.InitDBTest()
	r := routes.SetupRouter()
	utils.DB.Table("chatconversation").Create(&models.ChatConversation{Id: 10, Title: "123", Uid: "eWFjdzEyMw==/W2oNvwBBerfVCz4rsyXpA==", ModelId: 0})
	utils.DB.Table("chatmessage").Create(&models.ChatMessage{Id: 50, Content: "123", ChatId: 10, Actor: "user", Show: 1})
	var temp = controllers.DeleteChatReqBody{ApiKey: "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", ChatId: "10"}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("DELETE", "/api/v1/chat/chat", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		if resp.Code != http.StatusOK {
			utils.DB.Table("chatconversation").Delete(&models.ChatConversation{}, "id=?", 10)
			utils.DB.Table("chatmessage").Delete(&models.ChatMessage{}, "id=?", 50)
		}
	} else {
		utils.DB.Table("chatconversation").Delete(&models.ChatConversation{}, "id=?", 10)
		utils.DB.Table("chatmessage").Delete(&models.ChatMessage{}, "id=?", 50)
		assert.Equal(t, 1, 2)
	}

}
