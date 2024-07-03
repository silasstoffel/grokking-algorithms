package bfs_test

import (
	"testing"

	bfs "github.com/silasstoffel/grokking-algorithms/internal/bfs"
)

var sellers = map[string][]string{
	"me":     {"bob", "claire", "alice"},
	"bob":    {"anju", "peggy"},
	"claire": {"thom", "jonny"},
	"alice":  {"peggy"},
	"anju":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

func TestBfs(t *testing.T) {
	t.Parallel()

	t.Run("Should find the mango seller", func(t *testing.T) {
		name, err := bfs.Bfs("me", sellers)
		if err != nil {
			t.Error(err)
		}
		if name != "claire" {
			t.Error("Expected claire, got", name)
		}
	})
}
