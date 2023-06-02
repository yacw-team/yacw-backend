package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/models"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

type RequestDeletePrompt struct {
	Apikey    string `json:"apiKey"`
	PromptsId string `json:"promptsId"`
}

func TestDeletePromptCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	uid, _ := utils.Encrypt(apiKey)
	var prompts []models.Prompt
	err := utils.DB.Table("prompt").Where("uid=?", uid).Find(&prompts).Error
	if err == nil {
		requestDeletePrompt := &RequestDeletePrompt{
			Apikey:    apiKey,
			PromptsId: strconv.Itoa(prompts[0].Id),
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
	routes.SetupRouter().ServeHTTP(rr, req)
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
		PromptsId: "10000",
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
