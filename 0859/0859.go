package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n string

	fmt.Scan(&n)

	fmt.Print(SumOfNumbers(n))

}

func SumOfNumbers(number string) *big.Int {
	nextNumber, currentNumber, sum := new(big.Int), new(big.Int), new(big.Int)
	currentNumber.SetString(number, 10)
	nextNumber.Add(currentNumber, big.NewInt(1))
	sum.Mul(currentNumber, nextNumber)
	sum.Div(sum, big.NewInt(2))
	return sum

}
