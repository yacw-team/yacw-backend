package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type RequestChatMessage struct {
	ApiKey string `json:"apiKey"`
	ChatId string `json:"chatId"`
}

func TestGetChatMessageCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChatMessage := &RequestChatMessage{
		ApiKey: apiKey,
		ChatId: "1",
	}
	jsonStr, err := json.Marshal(requestChatMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/getmessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetChatMessageMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING_MISSING")
	requestChatMessage := &RequestChatMessage{
		ApiKey: apiKey,
		ChatId: "1",
	}
	jsonStr, err := json.Marshal(requestChatMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/getmessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestChatMessage := &RequestChatMessage{
		ApiKey: apiKey,
		ChatId: "1",
	}
	jsonStr, err := json.Marshal(requestChatMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/getmessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestChatMessage := &RequestChatMessage{
		ApiKey: apiKey,
		ChatId: "1",
	}
	jsonStr, err := json.Marshal(requestChatMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/getmessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageApiKeyNull(t *testing.T) {
	utils.InitDBTest()
	apiKey := ""
	requestChatMessage := &RequestChatMessage{
		ApiKey: apiKey,
		ChatId: "1",
	}
	jsonStr, err := json.Marshal(requestChatMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/getmessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageChatIdMiss(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChatMessage := &RequestChatMessage{
		ApiKey: apiKey,
		ChatId: "9999",
	}
	jsonStr, err := json.Marshal(requestChatMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/getmessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3009"}`
	assert.Equal(t, expected, rr.Body.String())
}
