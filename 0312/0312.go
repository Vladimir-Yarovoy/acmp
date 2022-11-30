package main

import "fmt"

func main() {
	var a1, a2, n int

	fmt.Scan(&a1, &a2, &n)

	fmt.Print(SequenceItem(a1, a2, n))

}

func SequenceItem(a1, a2, n int) int {

	d := a2 - a1
	return d*(n-1) + a1

}
