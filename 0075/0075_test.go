package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfMulOfDigitsOfNumber(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, SumOfMulOfDigitsOfNumber(big.NewInt(45), 1), big.NewInt(45))
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, SumOfMulOfDigitsOfNumber(big.NewInt(45), 3), big.NewInt(91125))
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, SumOfMulOfDigitsOfNumber(big.NewInt(45), 5), big.NewInt(184528125))
	})

	t.Run("Test 4", func(t *testing.T) {
		b := new(big.Int)
		b.SetString("107473400867440739228877320814929316405916989133564879594563305329166209446515624359452800717917853035032749176025390625", 10)
		assert.Equal(t, SumOfMulOfDigitsOfNumber(big.NewInt(45), 72), b)
	})
}
