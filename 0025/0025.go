package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Print(string(CompareSymbol(a, b)))

}

func CompareSymbol(a, b int) rune {
	if a < b {
		return '<'
	}

	if a > b {
		return '>'
	}

	return '='
}
