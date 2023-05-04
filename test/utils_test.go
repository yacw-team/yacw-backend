package test

import (
	"bytes"
	json2 "encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/controllers/v1/chat"
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
func TestNewChat(t *testing.T) {
	utils.InitDB()
	r := routes.SetupRouter()
	var temp = chat.NewChatRequest{ApiKey: "123", ModelId: "123", Content: chat.Content{System: "123", PromptsId: "123", PersonalityId: "123"}}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("POST", "/v1/chat/new", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	} else {
		assert.Equal(t, 1, 2)
	}

}
