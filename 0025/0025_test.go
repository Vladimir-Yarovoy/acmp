package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, CompareSymbol(5, 7), '<')
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, CompareSymbol(-7, -12), '>')
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, CompareSymbol(13, 13), '=')
	})
}
