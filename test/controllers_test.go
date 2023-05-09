package test

import (
	"bytes"
	"github.com/yacw-team/yacw/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetChatId(t *testing.T) {
	var jsonStr = []byte(`{"apiKey":"1"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getChat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
