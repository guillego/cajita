package handler

import (
	"encoding/json"
	"github.com/guillego/cajita/internal/store"
	"log"
	"net/http"
)

// Handler struct to hold the store instance.
type Handler struct {
	store *store.Store
}

// NewHandler initializes a new Handler.
func NewHandler(store *store.Store) *Handler {
	return &Handler{store: store}
}

// GetHandler handles GET requests.
func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		log.Println("GetHandler: key parameter is required")
		http.Error(w, "key parameter is required", http.StatusBadRequest)
		return
	}

	value, exists := h.store.Get(key)
	if !exists {
		log.Printf("GetHandler: key=%s not found", key)
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"value": value}); err != nil {
		log.Printf("GetHandler: failed to encode response for key=%s, error=%v", key, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// SetHandler handles POST requests.
func (h *Handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("SetHandler: invalid request body, error=%v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	key, keyExists := req["key"]
	value, valueExists := req["value"]
	if !keyExists || !valueExists {
		log.Println("SetHandler: key and value are required")
		http.Error(w, "key and value are required", http.StatusBadRequest)
		return
	}

	h.store.Set(key, value)
	log.Printf("SetHandler: key=%s set with value=%s", key, value)
	w.WriteHeader(http.StatusOK)
}

// DeleteHandler handles DELETE requests.
func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		log.Println("DeleteHandler: key parameter is required")
		http.Error(w, "key parameter is required", http.StatusBadRequest)
		return
	}

	h.store.Delete(key)
	log.Printf("DeleteHandler: key=%s deleted", key)
	w.WriteHeader(http.StatusOK)
}

