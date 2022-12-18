package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQtyCombinations(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, QtyCombinations(1, 3), big.NewInt(1))
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, QtyCombinations(2, 7), big.NewInt(21))
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, QtyCombinations(3, 10), big.NewInt(274))
	})

}
