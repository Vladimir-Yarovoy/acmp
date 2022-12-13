package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println(SumOfFactorialsOfNumbers(n))

}

func SumOfFactorialsOfNumbers(numbers int) *big.Int {
	factorials := factorialsOfNumbers(numbers)
	result := big.NewInt(0)
	for i := 1; numbers >= i; i++ {
		currentNumber := new(big.Int)
		currentNumber.SetString(factorials[i], 10)
		result.Add(result, currentNumber)
	}
	return result
}

func factorialsOfNumbers(numbers int) []string {
	var factorials []string
	var result = big.NewInt(1)
	factorials = append(factorials, "1")
	for i := 1; numbers >= i; i++ {
		previousNumber := new(big.Int)
		previousNumber.SetString(factorials[i-1], 10)
		result.Mul(previousNumber, big.NewInt(int64(i)))
		factorials = append(factorials, result.String())
	}
	return factorials
}
