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

type RequestCreatePrompt struct {
	ApiKey      string `json:"apiKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
}

type RequestCreatePromptWrong struct {
	ApiKey      string `json:"apiKey"`
	Nam         string `json:"nam"`
	Description string `json:"description"`
	Prompts     string `json:"prompts"`
}

func TestCreatePromptCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreatePromptMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptNameNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1007"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptDescriptionNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreatePromptPromptNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1007"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptAllNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "",
		Description: "",
		Prompts:     "",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1007"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptDataBaseNull(t *testing.T) {
	utils.InitDBNullTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3009"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptWrongQuest(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePromptWrong{
		ApiKey:      apiKey,
		Nam:         "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestCreatePromptApiKeySaltNot(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestCreatePrompt := &RequestCreatePrompt{
		ApiKey:      apiKey,
		Name:        "111",
		Description: "111",
		Prompts:     "111",
	}
	jsonStr, _ := json.Marshal(requestCreatePrompt)
	req, err := http.NewRequest("POST", "/api/v1/chat/prompts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3006"}`
	assert.Equal(t, expected, rr.Body.String())
}
