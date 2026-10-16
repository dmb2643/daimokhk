package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthGET(t *testing.T) {
	r := newRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status: got %d, want %d", res.Code, http.StatusOK)
	}

	if res.Body.String() != "ok" {
		t.Fatalf("body: got %q, want %q", res.Body.String(), "ok")
	}
}

func TestHealthPOST(t *testing.T) {
	r := newRouter()

	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status: got %d, want %d", res.Code, http.StatusMethodNotAllowed)
	}

	if got := res.Header().Get("Allow"); got != "GET, HEAD" {
		t.Fatalf("Allow: got %q, want %q", got, "GET, HEAD")
	}
}
