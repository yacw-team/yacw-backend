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

func TestTranslateCorrectExample(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{
    "apiKey":"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",
    "modelId":"1",
    "content":{
        "emotion":"anxious",
        "style":"",
        "preTranslate":"your mother is a docter"
    },
    "from":"english",
    "to":"chinese"
}`)
	req, err := http.NewRequest("POST", "/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestTranslateMissingLength(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{
    "apiKey":"sk-ZwWkbaSbC6fdzsH3RE0DT3BlbkFJH2KzpKW9JiyTOIWpasS",
    "modelId":"1",
    "content":{
        "emotion":"anxious",
        "style":"",
        "preTranslate":"your mother is a docter"
    },
    "from":"english",
    "to":"chinese"
}`)
	req, err := http.NewRequest("POST", "/v1/translate/translate", bytes.NewBuffer(jsonStr))
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
	var jsonStr = []byte(`{
    "apiKey":"sk-ZwWkbaSbC6fdzsH3RE0DT3BlbkFJH2KzpKW9JiyTOIWpasSg1",
    "modelId":"1",
    "content":{
        "emotion":"anxious",
        "style":"",
        "preTranslate":"your mother is a docter"
    },
    "from":"english",
    "to":"chinese"
}`)
	req, err := http.NewRequest("POST", "/v1/translate/translate", bytes.NewBuffer(jsonStr))
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
	var jsonStr = []byte(`{
    "apiKey":"sk-ZwWkbaSbC6fdzsH3RE0DT3BlbkFJH2KzpKW9JiyTOIWpasS我",
    "modelId":"1",
    "content":{
        "emotion":"anxious",
        "style":"",
        "preTranslate":"your mother is a docter"
    },
    "from":"english",
    "to":"chinese"
}`)
	req, err := http.NewRequest("POST", "/v1/translate/translate", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}
