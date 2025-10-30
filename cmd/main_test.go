package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETIndexPage(t *testing.T) {
	t.Run("returns index html page", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		IndexPageHandler(resp, req)
		got := resp.Body.String()
		want := `<html></html>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
