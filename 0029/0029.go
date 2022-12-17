package main

import (
	"fmt"
)

func main() {
	var n int
	var element int
	var heights []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&element)
		heights = append(heights, element)
	}

	fmt.Print(MinEnergy(heights))

}

func MinEnergy(heights []int) int {
	if len(heights) == 1 {
		return 0
	}
	minEnergy := make([]int, len(heights))
	minEnergy[1] = heights[1] - heights[0]
	if minEnergy[1] < 0 {
		minEnergy[1] = heights[0] - heights[1]
	}
	for i := 2; i < len(minEnergy); i++ {
		energyForOneStep := heights[i] - heights[i-1]
		energyForDoubleStep := (heights[i] - heights[i-2]) * 3
		if energyForOneStep < 0 {
			energyForOneStep = heights[i-1] - heights[i]
		}
		if energyForDoubleStep < 0 {
			energyForDoubleStep = (heights[i-2] - heights[i]) * 3
		}

		if energyForOneStep+minEnergy[i-1] > energyForDoubleStep+minEnergy[i-2] {
			minEnergy[i] = energyForDoubleStep + minEnergy[i-2]
		} else {
			minEnergy[i] = energyForOneStep + minEnergy[i-1]
		}
	}
	return minEnergy[len(minEnergy)-1]
}
