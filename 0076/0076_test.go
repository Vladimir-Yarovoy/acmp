package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__GetMaxVisitors(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, GetMaxVisitors("0076_testdata_01"), 4)
	})

	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, GetMaxVisitors("0076_testdata_02"), 3)
	})

	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, GetMaxVisitors("0076_testdata_03"), 1)
	})

	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, GetMaxVisitors("0076_testdata_04"), 1)
	})

	t.Run("Test 5", func(t *testing.T) {
		assert.Equal(t, GetMaxVisitors("0076_testdata_05"), 1)
	})

}

func Test__getVisitorSchedule(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		comingTime, departureTime := getVisitorSchedule("09:00 10:07")
		assert.Equal(t, comingTime, "09:00")
		assert.Equal(t, departureTime, "10:07")
	})

}

func Test__convertTimeToMinutes(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, convertTimeToMinutes("10:07"), 607)
	})

}
