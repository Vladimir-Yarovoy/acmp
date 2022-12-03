package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, IsDivided("121"), "YES")
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, IsDivided("1211"), "NO")
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, IsDivided("1771561"), "YES")
	})

	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, IsDivided("1719191"), "NO")
	})

	t.Run("Test 5", func(t *testing.T) {
		assert.Equal(t, IsDivided("7171717"), "NO")
	})
}
