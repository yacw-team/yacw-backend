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

type Content struct {
	User string `json:"user"`
}

type RequestSendMessage struct {
	ApiKey  string  `json:"apiKey"`
	ChatId  string  `json:"chatId"`
	Content Content `json:"content"`
}

func TestSendMessageCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "1",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestSendMessageMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "2",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendMessageExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "2",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendMessageFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "2",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendMessageApiKeyNull(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{
    "apiKey": "",
    "chatId": "2",
    "content": {
        "user": "再多说一些"
    }
}`)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2006"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendMessageChatIdNull(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendMessageChatIdNoExist(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "0",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendMessageChatIdMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendMessage := &RequestSendMessage{
		ApiKey: apiKey,
		ChatId: "我",
		Content: Content{
			User: "再多说一些",
		},
	}
	jsonStr, err := json.Marshal(requestSendMessage)
	req, err := http.NewRequest("POST", "/api/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}
