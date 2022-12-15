package main

import (
	"fmt"
)

func main() {
	var n, m, element int
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&element)
			matrix[i][j] = element
		}
	}

	fmt.Print(OptimalEdgeWeight(matrix))

}

func OptimalEdgeWeight(matrix [][]int) int {
	optimalPath := make([][]int, len(matrix))
	for i := range optimalPath {
		optimalPath[i] = make([]int, len(matrix[0]))
	}

	optimalPath[0][0] = matrix[0][0]

	for j := 1; j < len(matrix[0]); j++ {
		optimalPath[0][j] = matrix[0][j] + optimalPath[0][j-1]
	}

	for i := 1; i < len(matrix); i++ {
		optimalPath[i][0] = matrix[i][0] + optimalPath[i-1][0]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if optimalPath[i][j-1] > optimalPath[i-1][j] {
				optimalPath[i][j] = optimalPath[i-1][j] + matrix[i][j]
			} else {
				optimalPath[i][j] = optimalPath[i][j-1] + matrix[i][j]
			}
		}
	}

	return optimalPath[len(matrix)-1][len(matrix[0])-1]
}
