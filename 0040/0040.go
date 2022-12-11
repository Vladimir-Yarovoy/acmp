package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int

	fmt.Scan(&n)

	a := big.NewInt(2)

	fmt.Print(PowerNumber(a, n))

}

func PowerNumber(number *big.Int, power int) *big.Int {
	power2System := convertPowerTo2System(power)
	result := big.NewInt(1)
	for i := len(power2System) - 1; i >= 0; i-- {
		if power2System[i] == 0 {
			result.Mul(result, result)
		} else {
			result.Mul(result, result)
			result.Mul(result, number)
		}
	}
	return result
}

func convertPowerTo2System(power int) []int {
	var power2System []int
	for i := 0; power > 0; i++ {
		power2System = append(power2System, power%2)
		power = power / 2
	}
	return power2System
}
