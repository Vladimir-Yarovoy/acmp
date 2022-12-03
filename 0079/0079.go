package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	fmt.Print(LastDigitOfPower(a, b))
}

func LastDigitOfPower(base, power int) int {

	digitPeriods := [][]int{
		{0},
		{1},
		{2, 4, 8, 6},
		{3, 9, 7, 1},
		{4, 6},
		{5},
		{6},
		{7, 9, 3, 1},
		{8, 4, 2, 6},
		{9, 1}}

	lastDigitOfBase := base % 10
	lastDigitOfBasePeriod := len(digitPeriods[lastDigitOfBase])

	return digitPeriods[lastDigitOfBase][calcDigitIndex(power, lastDigitOfBasePeriod)]
}

func calcDigitIndex(power, digitPeriod int) int {

	powerModPeriod := power % digitPeriod
	if powerModPeriod == 0 {
		return digitPeriod - 1
	}

	return powerModPeriod - 1
}
