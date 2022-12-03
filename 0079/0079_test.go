package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__LastDigitOfPower(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, LastDigitOfPower(2, 2), 4)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, LastDigitOfPower(3, 7), 7)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, LastDigitOfPower(24, 9), 4)
	})

	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, LastDigitOfPower(2, 6), 4)
	})

	t.Run("Test 5", func(t *testing.T) {
		assert.Equal(t, LastDigitOfPower(1, 99), 1)
	})

}

func Test__calcDigitIndex(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(1, 4), 0)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(10, 4), 1)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(3, 4), 2)
	})

	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(4, 4), 3)
	})

	t.Run("Test 5", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(5, 4), 0)
	})

	t.Run("Test 6", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(10, 1), 0)
	})

	t.Run("Test 7", func(t *testing.T) {
		assert.Equal(t, calcDigitIndex(10, 2), 1)
	})

}
