package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, IsTriangle(3, 4, 5), "YES")
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, IsTriangle(1, 1, 5), "NO")
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, IsTriangle(2, 3, 5), "NO")
	})

}
