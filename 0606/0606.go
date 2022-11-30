package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Print(IsTriangle(a, b, c))

}

func IsTriangle(a, b, c int) string {
	if (a+b > c) && (a+c > b) && (c+b > a) {
		return "YES"
	}

	return "NO"
}
