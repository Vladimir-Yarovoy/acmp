package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	fmt.Print(LastDigit(a, b))

}

func LastDigit(a, b int) int {

	matrix := [][]int{
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

	lastDigitA := a % 10

	bModLastDigitA := b % len(matrix[lastDigitA])

	if bModLastDigitA == 0 {
		return matrix[lastDigitA][len(matrix[lastDigitA])-1]
	}

	return matrix[lastDigitA][bModLastDigitA-1]
}
