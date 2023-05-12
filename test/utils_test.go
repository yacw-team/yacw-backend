package test

import (
	"bytes"
	json2 "encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/controllers/v1/chat"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 测试加密函数
func TestEncryptPassword(t *testing.T) {
	input := "password"
	expectedOutput := "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
	output := utils.EncryptPassword(input)
	assert.Equal(t, expectedOutput, output)
}

func TestDeleteChat(t *testing.T) {
	utils.InitDBTest()
	r := routes.SetupRouter()
	utils.DB.Table("chatconversation").Create(&models.ChatConversation{Id: 10, Title: "123", Uid: "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3", ModelId: 0})
	utils.DB.Table("chatmessage").Create(&models.ChatMessage{Id: 50, Content: "123", ChatId: 10, Actor: "user", Show: 1})
	var temp = chat.DeleteRequest{ApiKey: "123", ChatId: "10"}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("DELETE", "/v1/chat/chat", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		print(resp.Body.String())
		assert.Equal(t, http.StatusOK, resp.Code)
	} else {
		assert.Equal(t, 1, 2)
	}

}
