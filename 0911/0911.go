package main

import (
	"fmt"
)

func main() {
	var n int
	var c string

	fmt.Scan(&c, &n)

	fmt.Print(WeighingWith3(c, n))

}

func WeighingWith3(side string, number int) (string, string) {
	power3 := []string{"1", "3", "9", "27", "81", "243", "729", "2187", "6561", "19683", "59049", "177147", "531441", "1594323", "4782969", "14348907", "43046721", "129140163", "387420489", "1162261467"}
	var side1, side2, left, right string
	numberTo3System := convertToTernarySystem(number)
	for i := 0; i < len(numberTo3System); i++ {
		if numberTo3System[i] == '-' {
			side1 = side1 + power3[i] + " "
		}
		if numberTo3System[i] == '+' {
			side2 = side2 + power3[i] + " "
		}

	}

	if side == "L" {
		left = "L:" + side1 + "\n"
		right = "R:" + side2
	}

	if side == "R" {
		left = "L:" + side2 + "\n"
		right = "R:" + side1
	}

	return left, right
}

func convertToTernarySystem(number int) string {
	var result, digit string
	for number > 0 {
		b := number % 3
		if b == 0 {
			digit = "0"
		}
		if b == 1 {
			digit = "+"
		}
		if b == 2 {
			digit = "-"
			number++
		}
		result = result + digit
		number = number / 3

	}
	return result
}
