package test

import (
	"bytes"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type NewChatRequest struct {
	ApiKey  string         `json:"apiKey"`
	ModelId string         `json:"modelId"`
	ChatId  string         `json:"chatId"`
	Content NewChatContent `json:"content"`
}

type NewChatContent struct {
	PersonalityId string `json:"personalityId"`
	User          string `json:"user"`
}

func TestNewChatModelIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	r := routes.SetupRouter()
	var tempContent = NewChatContent{PersonalityId: "1", User: "怎么成为一名肌肉男"}
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "", ChatId: "123", Content: tempContent}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("POST", "/api/v1/chat/new", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	} else {
		t.Fatal(err)
	}
}

func TestNewChatModelIdWrong(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	r := routes.SetupRouter()
	var tempContent = NewChatContent{PersonalityId: "1", User: "怎么成为一名肌肉男"}
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "10", ChatId: "123", Content: tempContent}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("POST", "/api/v1/chat/new", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	} else {
		t.Fatal(err)
	}
}

func TestNewChatAPIKeyWrong(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	r := routes.SetupRouter()
	var tempContent = NewChatContent{PersonalityId: "1", User: "怎么成为一名肌肉男"}
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "1", ChatId: "123", Content: tempContent}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("POST", "/api/v1/chat/new", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	} else {
		t.Fatal(err)
	}
}

func TestNewChatCatchPanic(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	r := routes.SetupRouter()
	var temp = gin.H{
		"apiKey": apiKey,
	}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("POST", "/api/v1/chat/new", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	} else {
		t.Fatal(err)
	}
}
