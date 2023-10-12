package main

import (
	"fmt"
	"math"
)

func GeneratePolynomialHash(r, q int, str string) int {
	hashedValue := 0

	for i, v := range str {
		fmt.Println(v)
		fmt.Println(q)
		fmt.Println(float64(len(str) - i))
		hashedValue += int(v) * int(math.Pow(float64(q), float64(len(str)-i)-1))
	}
	hashedValue %= r

	return hashedValue
}

func main() {
	fmt.Println(GeneratePolynomialHash(100003, 123, "hash"))
}
