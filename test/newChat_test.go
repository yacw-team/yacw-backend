package test

import (
	"bytes"
	json2 "encoding/json"
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
	Content NewChatContent `json:"content"`
}

type NewChatContent struct {
	PersonalityId string `json:"personalityId"`
	User          string `json:"user"`
}

func TestNewChatCorrect(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	r := routes.SetupRouter()
	var tempContent = NewChatContent{PersonalityId: "1", User: "怎么成为一名肌肉男"}
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "0", Content: tempContent}
	json, err := json2.Marshal(&temp)
	reader := bytes.NewReader(json)
	if err == nil {
		req, _ := http.NewRequest("POST", "/api/v1/chat/new", reader)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	} else {
		t.Fatal(err)
	}
}

func TestNewChatModelIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	r := routes.SetupRouter()
	var tempContent = NewChatContent{PersonalityId: "1", User: "怎么成为一名肌肉男"}
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "", Content: tempContent}
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
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "10", Content: tempContent}
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
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "1", Content: tempContent}
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

func TestNewChatPersonalityIdWrong(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	r := routes.SetupRouter()
	var tempContent = NewChatContent{PersonalityId: "100", User: "怎么成为一名肌肉男"}
	var temp = NewChatRequest{ApiKey: apiKey, ModelId: "1", Content: tempContent}
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
