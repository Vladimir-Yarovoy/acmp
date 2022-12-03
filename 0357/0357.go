package main

import (
	"fmt"
)

func main() {
	var n string

	fmt.Scan(&n)

	fmt.Print(IsDivided(n))

}

func IsDivided(number string) string {
	var diff int
	for i := 0; i < len(number); i++ {
		if i%2 == 0 {
			diff += int(number[i]) - '0'
		} else {
			diff -= int(number[i]) - '0'
		}
	}

	if diff%11 == 0 {
		return "YES"
	}

	return "NO"

}

/* МАТЕМАТИКА
5 % 3 = 2
-5 % 3 = -2*3 + 1


N - число, Q - делитель, R - остаток
N = a*Q + R, 0 <= R < Q

ПРОГРАММИРОВАНИЕ
Есть стандарт ассоциации IETF, которым определена операция взятия остатка в программировании и она отличается от математической


*/

// func DividerN(n int) string {
// 	var digitsOfNumberN []int
// 	for i := 0; n > 0; i++ {
// 		if n <= 9 {
// 			digitsOfNumberN = append(digitsOfNumberN, n)
// 			break
// 		}
// 		digitsOfNumberN = append(digitsOfNumberN, n%10)
// 		n = n / 10
// 	}

// 	var sumUnevenDigits, sumEvenDigits int
// 	for i := 0; i < len(digitsOfNumberN); i++ {
// 		if i%2 == 0 {
// 			sumUnevenDigits += digitsOfNumberN[i]
// 		}
// 		if i%2 != 0 {
// 			sumEvenDigits += digitsOfNumberN[i]
// 		}
// 	}

// 	result := (sumEvenDigits - sumUnevenDigits) % 11

// 	if result == 0 {
// 		return "YES"
// 	}
// 	return "NO"

// }
