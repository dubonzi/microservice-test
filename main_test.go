package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("[FATAL] error creating the request: ", err)
	}
	rec := httptest.NewRecorder()

	hand := http.HandlerFunc(RootHandler)

	hand.ServeHTTP(rec, req)

	expected := "Hello from AWS!"
	got := rec.Body.String()
	if got != expected {
		t.Errorf("Expected response to be '%s' but got '%s'", expected, got)
	}
}

func TestThanksRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/thanks", nil)
	if err != nil {
		t.Fatal("[FATAL] error creating the request: ", err)
	}
	rec := httptest.NewRecorder()

	hand := http.HandlerFunc(ThanksHandler)

	hand.ServeHTTP(rec, req)

	expected := "Thank you Clara! :)"
	got := rec.Body.String()
	if got != expected {
		t.Errorf("Expected response to be '%s' but got '%s'", expected, got)
	}
}
