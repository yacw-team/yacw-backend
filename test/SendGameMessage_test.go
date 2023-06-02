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

type RequestSendGameMessage struct {
	ApiKey   string `json:"apiKey"`
	ChoiceId string `json:"choiceID"`
	ModelId  string `json:"modelId"`
}

func TestSendGameMessageMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "A",
		ModelId:  "1",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendGameMessageExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "A",
		ModelId:  "1",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendGameMessageFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "A",
		ModelId:  "1",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendGameMessageModelIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "A",
		ModelId:  "",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendGameMessageModelIdNoExist(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "A",
		ModelId:  "100",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendGameMessageChoiceIdNoExist(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "q",
		ModelId:  "1",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestSendGameMessageChoiceIdSmallLetter(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "a",
		ModelId:  "1",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	b := rr.Body.String()
	assert.Equal(t, expected, b)
}

func TestSendGameMessageChoiceIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestSendGameMessage := &RequestSendGameMessage{
		ApiKey:   apiKey,
		ChoiceId: "",
		ModelId:  "1",
	}
	jsonStr, _ := json.Marshal(requestSendGameMessage)
	req, err := http.NewRequest("POST", "/api/v1/game/chat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}
