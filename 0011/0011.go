package main

import (
	"fmt"
	"math/big"
)

func main() {
	var k, n int

	fmt.Scan(&k, &n)

	fmt.Println(QtyCombinations(k, n))

}

func QtyCombinations(k, n int) *big.Int {
	var combinationsNWithStepK []string
	combinationsNWithStepK = append(combinationsNWithStepK, "1", "1")
	result := big.NewInt(0)
	iterationsByK := k
	for i := 2; i <= n; i++ {
		combinationsNWithStepK = append(combinationsNWithStepK, "0")
		for j := i - 1; iterationsByK > 0; j-- {
			if j == -1 {
				break
			}
			iterationsByK--
			currentNumber := new(big.Int)
			previousNumber := new(big.Int)
			currentNumber.SetString(combinationsNWithStepK[i], 10)
			previousNumber.SetString(combinationsNWithStepK[j], 10)
			currentNumber.Add(currentNumber, previousNumber)
			combinationsNWithStepK[i] = currentNumber.String()
		}
		iterationsByK = k
	}
	result.SetString(combinationsNWithStepK[len(combinationsNWithStepK)-1], 10)
	return result
}
