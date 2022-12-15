package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimalEdgeWeight(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, OptimalEdgeWeight([][]int{{1, 1, 1, 1}, {5, 2, 2, 100}, {9, 4, 2, 1}}), 8)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, OptimalEdgeWeight([][]int{{1, 1, 1, 1, 1}, {3, 100, 100, 100, 100}, {1, 1, 1, 1, 1}, {2, 2, 2, 2, 1}, {1, 1, 1, 1, 1}}), 11)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, OptimalEdgeWeight([][]int{{1}}), 1)
	})

}
