package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGameBackgroundCorrectExample(t *testing.T) {
	utils.InitDBTest()
	req, err := http.NewRequest("GET", "/api/v1/game/story", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetGameBackgroundDataBaseNil(t *testing.T) {
	utils.InitDBNilTest()
	req, err := http.NewRequest("GET", "/api/v1/game/story", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `[]`
	assert.Equal(t, expected, rr.Body.String())
}
