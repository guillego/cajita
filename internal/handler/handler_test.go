package handler

import (
	"bytes"
	"github.com/guillego/cajita/internal/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	// Test SetHandler
	reqBody := `{"key": "foo", "value": "bar"}`
	req := httptest.NewRequest(http.MethodPost, "/set", bytes.NewBufferString(reqBody))
	w := httptest.NewRecorder()
	h.SetHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", w.Code)
	}

	// Test GetHandler
	req = httptest.NewRequest(http.MethodGet, "/get?key=foo", nil)
	w = httptest.NewRecorder()
	h.GetHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", w.Code)
	}

	expected := `{"value":"bar"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("expected body %v, got %v", expected, w.Body.String())
	}

	// Test DeleteHandler
	req = httptest.NewRequest(http.MethodDelete, "/delete?key=foo", nil)
	w = httptest.NewRecorder()
	h.DeleteHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", w.Code)
	}

	// Verify deletion
	req = httptest.NewRequest(http.MethodGet, "/get?key=foo", nil)
	w = httptest.NewRecorder()
	h.GetHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %v", w.Code)
	}
}
