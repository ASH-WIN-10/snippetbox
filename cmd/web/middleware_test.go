package main

import (
	"net/http"
	"testing"

	"github.com/ASH-WIN-10/snippetbox/internal/assert"
)

func TestCommonHeaders(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	ts := newTestServer(t, commonHeaders(next))
	defer ts.Close()

	statusCode, headers, body := ts.get(t, "/")

	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, headers.Get("Content-Security-Policy"), expectedValue)

	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, headers.Get("Referrer-Policy"), expectedValue)

	expectedValue = "nosniff"
	assert.Equal(t, headers.Get("X-Content-Type-Options"), expectedValue)

	expectedValue = "deny"
	assert.Equal(t, headers.Get("X-Frame-Options"), expectedValue)

	expectedValue = "0"
	assert.Equal(t, headers.Get("X-XSS-Protection"), expectedValue)

	expectedValue = "Go"
	assert.Equal(t, headers.Get("Server"), expectedValue)

	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
