package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	s := NewStore()

	s.Set("foo", "bar")
	value, exists := s.Get("foo")
	if !exists || value != "bar" {
		t.Errorf("expected value 'bar', got %v", value)
	}

	s.Delete("foo")
	_, exists = s.Get("foo")
	if exists {
		t.Error("expected key 'foo' to be deleted")
	}
}

