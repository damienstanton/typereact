package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticServe(t *testing.T) {
	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("error running GET request on the root path: %v", err)
	}

	fileServer().ServeHTTP(respRec, req)

	statusOK := 200
	if respRec.Code != http.StatusOK {
		t.Fatalf("test server returned a non-200 response. Expected %v, got %v",
			statusOK, respRec.Code)
	}
}
