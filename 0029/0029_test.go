package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimalEdgeWeight(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{1, 5, 10}), 9)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{1, 5, 2}), 3)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{5, 1}), 4)
	})

	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{10, 15, 8, 18, 7, 18, 7, 19}), 17)
	})

	t.Run("Test 5", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{1, 99, 2, 99, 3, 99, 4, 99, 5, 99}), 98)
	})

	t.Run("Test 6", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{3, 4, 8, 10, 1, 7, 1, 10, 12, 15}), 30)
	})

	t.Run("Test 7", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{1, 5, 2, 5}), 4)
	})

	t.Run("Test 8", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{100, 1}), 99)
	})

	t.Run("Test 9", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{30000, 3000, 300}), 29700)
	})

	t.Run("Test 10", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{1, 1, 100, 1, 100, 1, 100, 100, 1, 1, 1}), 198)
	})

	t.Run("Test 11", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{1, 100, 100, 1}), 198)
	})

	t.Run("Test 12", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{5, 10, 5, 0}), 5)
	})

	t.Run("Test 13", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{0, 5, 10, 5}), 5)
	})

	t.Run("Test 14", func(t *testing.T) {
		assert.Equal(t, MinEnergy([]int{18, 8, 15, 10}), 14)
	})
}
