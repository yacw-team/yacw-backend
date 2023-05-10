package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessageCorrectExample(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{
    "apiKey": "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",
    "chatId": "2",
    "content": {
        "user": "再多说一些"
    }
}`)
	req, err := http.NewRequest("POST", "/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestSendMessageMissingLength(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{
    "apiKey": "sk-ZwWkbaSbC6fdzsH3RE0DT3BlbkFJH2KzpKW9JiyTOIWpasS",
    "chatId": "2",
    "content": {
        "user": "再多说一些"
    }
}`)
	req, err := http.NewRequest("POST", "/v1/chat/chat", bytes.NewBuffer(jsonStr))
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
	var jsonStr = []byte(`{
    "apiKey": "sk-ZwWkbaSbC6fdzsH3RE0DT3BlbkFJH2KzpKW9JiyTOIWpasSg1",
    "chatId": "2",
    "content": {
        "user": "再多说一些"
    }
}`)
	req, err := http.NewRequest("POST", "/v1/chat/chat", bytes.NewBuffer(jsonStr))
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
	var jsonStr = []byte(`{
    "apiKey": "sk-ZwWkbaSbC6fdzsH3RE0DT3BlbkFJH2KzpKW9JiyTOIWpasS我",
    "chatId": "2",
    "content": {
        "user": "再多说一些"
    }
}`)
	req, err := http.NewRequest("POST", "/v1/chat/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}
