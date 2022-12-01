package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, LastDigit(2, 2), 4)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, LastDigit(3, 7), 7)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, LastDigit(24, 9), 4)
	})

}
