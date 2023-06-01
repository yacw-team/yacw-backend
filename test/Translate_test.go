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

type TranslateContent struct {
	Emotion      string `json:"emotion"`
	Style        string `json:"style"`
	PreTranslate string `json:"preTranslate"`
}

type RequestTranslate struct {
	ApiKey  string           `json:"apiKey"`
	ModelId string           `json:"modelId"`
	Content TranslateContent `json:"content"`
	From    string           `json:"from"`
	To      string           `json:"to"`
}

func TestTranslateCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "1",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestTranslateMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "1",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestTranslateExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "1",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestTranslateFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "1",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestTranslateModelIdNull(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2006"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestTranslateModelIdCross(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "11",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestTranslateModelIdMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestTranslate := &RequestTranslate{
		ApiKey:  apiKey,
		ModelId: "我",
		Content: TranslateContent{
			Emotion:      "happy",
			Style:        "",
			PreTranslate: "happy",
		},
		From: "english",
		To:   "chinese",
	}
	jsonStr, _ := json.Marshal(requestTranslate)
	req, err := http.NewRequest("POST", "/api/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2006"}`
	assert.Equal(t, expected, rr.Body.String())
}
