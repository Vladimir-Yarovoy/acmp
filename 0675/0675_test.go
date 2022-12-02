package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, Distance([]string{"AA.BB"}), 1)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, Distance([]string{"AA..B", "A...B"}), 2)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, Distance([]string{"AA..B", "A...B", "AAABB"}), 0)
	})

}
