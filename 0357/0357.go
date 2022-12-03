package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Print(DividerN(n))

}

func DividerN(n int) string {
	var digitsOfNumberN []int
	for i := 0; n > 0; i++ {
		if n <= 9 {
			digitsOfNumberN = append(digitsOfNumberN, n)
			break
		}
		digitsOfNumberN = append(digitsOfNumberN, n%10)
		n = n / 10
	}

	var sumUnevenDigits, sumEvenDigits int
	for i := 0; i < len(digitsOfNumberN); i++ {
		if i%2 == 0 {
			sumUnevenDigits += digitsOfNumberN[i]
		}
		if i%2 != 0 {
			sumEvenDigits += digitsOfNumberN[i]
		}
	}

	result := (sumEvenDigits - sumUnevenDigits) % 11

	if result == 0 {
		return "YES"
	}
	return "NO"

}
