package store

import (
	"log"
	"sync"
)

// Store holds the data and provides thread-safe operations.
type Store struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewStore initializes a new Store.
func NewStore() *Store {
	return &Store{
		store: make(map[string]string),
	}
}

// Get retrieves a value from the store.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.store[key]
	if exists {
		log.Printf("Get: key=%s, value=%s", key, value)
	} else {
		log.Printf("Get: key=%s not found", key)
	}
	return value, exists
}

// Set sets a value in the store.
func (s *Store) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
	log.Printf("Set: key=%s, value=%s", key, value)
}

// Delete removes a key from the store.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
	log.Printf("Delete: key=%s", key)
}
