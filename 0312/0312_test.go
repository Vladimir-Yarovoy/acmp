package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, SequenceItem(1, 5, 3), 9)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, SequenceItem(3, 9, 7), 39)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, SequenceItem(1, 1, 5), 1)
	})

	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, SequenceItem(10, 5, 4), -5)
	})

}
