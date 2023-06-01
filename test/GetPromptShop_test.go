package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPromptShopCorrectExample(t *testing.T) {
	utils.InitDBTest()
	req, err := http.NewRequest("GET", "/api/v1/chat/prompts?type=工具", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetPromptShopDataBaseNil(t *testing.T) {
	utils.InitDBNilTest()
	req, err := http.NewRequest("GET", "/api/v1/chat/prompts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"Prompts":[]}`
	assert.Equal(t, expected, rr.Body.String())
}
