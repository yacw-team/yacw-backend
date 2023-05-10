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

func TestGetChatIdCorrectExample(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"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"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getChat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"chat":[{"chatId":2222,"title":"33333"}]}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatIdMissingLength(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"sk-hISgKGQQ5cZNGHZxbQFXT3BlbkFJ8vyxitPPXM6oqfgTeNl"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getChat", bytes.NewBuffer(jsonStr))
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
	var jsonStr = []byte(`{"apiKey":"sk-hISgKGQQ5cZNGHZxbQFXT3BlbkFJ8vyxitPPXM6oqfgTeNlxl"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getChat", bytes.NewBuffer(jsonStr))
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
	var jsonStr = []byte(`{"apiKey":"sk-hISgKGQQ5cZNGHZxbQFXT3BlbkFJ8vyxitPPXM6oqfgTeNl我"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getChat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}
