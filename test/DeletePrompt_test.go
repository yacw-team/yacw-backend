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

type RequestDeletePrompt struct {
	Apikey    string `json:"apiKey"`
	PromptsId string `json:"promptsId"`
}

func TestDeletePromptCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestDeletePrompt := &RequestDeletePrompt{
		Apikey:    apiKey,
		PromptsId: "65",
	}
	jsonStr, _ := json.Marshal(requestDeletePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/deleteprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeletePromptMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestDeletePrompt := &RequestDeletePrompt{
		Apikey:    apiKey,
		PromptsId: "53",
	}
	jsonStr, _ := json.Marshal(requestDeletePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/deleteprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestDeletePromptExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestDeletePrompt := &RequestDeletePrompt{
		Apikey:    apiKey,
		PromptsId: "53",
	}
	jsonStr, _ := json.Marshal(requestDeletePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/deleteprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req) //11
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestDeletePromptFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestDeletePrompt := &RequestDeletePrompt{
		Apikey:    apiKey,
		PromptsId: "53",
	}
	jsonStr, _ := json.Marshal(requestDeletePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/deleteprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestDeletePromptPromptsIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestDeletePrompt := &RequestDeletePrompt{
		Apikey:    apiKey,
		PromptsId: "",
	}
	jsonStr, _ := json.Marshal(requestDeletePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/deleteprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1008"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestDeletePromptPromptsIdNoExist(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestDeletePrompt := &RequestDeletePrompt{
		Apikey:    apiKey,
		PromptsId: "100",
	}
	jsonStr, _ := json.Marshal(requestDeletePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/deleteprompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1008"}`
	assert.Equal(t, expected, rr.Body.String())
}
