package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int

	var matrix []string

	fmt.Scan(&n, &m)

	var row string
	for i := 0; i < n; i++ {
		fmt.Scan(&row)
		matrix = append(matrix, row)
	}

	fmt.Print(Distance(matrix))

}

func Distance(matrix []string) int {
	min := 98
	for i := 0; i < len(matrix); i++ {
		count := strings.Count(matrix[i], ".")
		if count < min {
			min = count
		}
	}

	return min
}
