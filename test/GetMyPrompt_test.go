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

type RequestGetMyPrompt struct {
	ApiKey string `json:"apiKey"`
}

func TestGetMyPromptCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestGetMyPrompt := &RequestGetMyPrompt{
		ApiKey: apiKey,
	}
	jsonStr, _ := json.Marshal(requestGetMyPrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/myprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetMyPromptMissingLength(t *testing.T) { //1
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestGetMyPrompt := &RequestGetMyPrompt{
		ApiKey: apiKey,
	}
	jsonStr, _ := json.Marshal(requestGetMyPrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/myprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetMyPromptExcessiveLength(t *testing.T) { //1
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestGetMyPrompt := &RequestGetMyPrompt{
		ApiKey: apiKey,
	}
	jsonStr, _ := json.Marshal(requestGetMyPrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/myprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetMyPromptFormatMixing(t *testing.T) { //1
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestGetMyPrompt := &RequestGetMyPrompt{
		ApiKey: apiKey,
	}
	jsonStr, _ := json.Marshal(requestGetMyPrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/myprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetMyPromptDataBaseNull(t *testing.T) {
	utils.InitDBNullTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestGetMyPrompt := &RequestGetMyPrompt{
		ApiKey: apiKey,
	}
	jsonStr, _ := json.Marshal(requestGetMyPrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/myprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3009"}`
	assert.Equal(t, expected, rr.Body.String())
}
