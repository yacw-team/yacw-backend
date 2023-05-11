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

type RequestChatId struct {
	ApiKey string `json:"apiKey"`
}

func TestGetChatIdCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChatId := &RequestChatId{
		ApiKey: apiKey,
	}
	jsonStr, err := json.Marshal(requestChatId)
	req, err := http.NewRequest("POST", "/api/v1/chat/getchat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetChatIdMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestChatId := &RequestChatId{
		ApiKey: apiKey,
	}
	jsonStr, err := json.Marshal(requestChatId)
	req, err := http.NewRequest("POST", "/api/v1/chat/getchat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatIdExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestChatId := &RequestChatId{
		ApiKey: apiKey,
	}
	jsonStr, err := json.Marshal(requestChatId)
	req, err := http.NewRequest("POST", "/api/v1/chat/getchat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatIdFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestChatId := &RequestChatId{
		ApiKey: apiKey,
	}
	jsonStr, err := json.Marshal(requestChatId)
	req, err := http.NewRequest("POST", "/api/v1/chat/getchat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatIdApiKeyNull(t *testing.T) {
	utils.InitDBTest()
	apiKey := ""
	requestChatId := &RequestChatId{
		ApiKey: apiKey,
	}
	jsonStr, err := json.Marshal(requestChatId)
	req, err := http.NewRequest("POST", "/api/v1/chat/getchat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}
