package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerNumber(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, PowerNumber(big.NewInt(2), 3), big.NewInt(8))
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, PowerNumber(big.NewInt(2), 10), big.NewInt(1024))
	})

	t.Run("Test 3", func(t *testing.T) {
		b := new(big.Int)
		b.SetString("4722366482869645213696", 10)
		assert.Equal(t, PowerNumber(big.NewInt(2), 72), b)
	})
}
