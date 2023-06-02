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

type RequestChooseGameBackground struct {
	ApiKey  string `json:"apiKey"`
	GameId  string `json:"gameId"`
	ModelId string `json:"modelId"`
}

func TestChooseGameBackgroundCorrectExample(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "1",
		ModelId: "1",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestChooseGameBackgroundMissingLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MISSING")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "1",
		ModelId: "1",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestChooseGameBackgroundExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_EXCESSIVE")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "1",
		ModelId: "1",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestChooseGameBackgroundFormatMixing(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY_MIXING")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "1",
		ModelId: "1",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestChooseGameBackgroundGameIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "",
		ModelId: "1",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1008"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestChooseGameBackgroundModelIdNil(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "1",
		ModelId: "",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"2005"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestChooseGameBackgroundGameIdNoExist(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "100",
		ModelId: "1",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1008"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestChooseGameBackgroundModelIdNoExist(t *testing.T) {
	utils.InitDBTest()
	apiKey := os.Getenv("TEST_OPENAI_KEY")
	requestChooseGameBackground := &RequestChooseGameBackground{
		ApiKey:  apiKey,
		GameId:  "1",
		ModelId: "100",
	}
	jsonStr, _ := json.Marshal(requestChooseGameBackground)
	req, err := http.NewRequest("POST", "/api/v1/game/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"1005"}`
	assert.Equal(t, expected, rr.Body.String())
}
