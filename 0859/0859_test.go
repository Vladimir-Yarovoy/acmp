package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfNumbers(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, SumOfNumbers("5"), big.NewInt(15))
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, SumOfNumbers("3"), big.NewInt(6))
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, SumOfNumbers("10"), big.NewInt(55))
	})

}
