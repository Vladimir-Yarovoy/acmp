package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfFactorialsOfNumbers(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, SumOfFactorialsOfNumbers(1), 1)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, SumOfFactorialsOfNumbers(2), 3)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, SumOfFactorialsOfNumbers(3), 9)
	})

}
