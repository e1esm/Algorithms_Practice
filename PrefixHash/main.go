package main

import (
	"fmt"
	"math"
)

const (
	modulus = 1000009
	base    = 1000
)

func hashStr(str string) int {
	var hashValue int

	for i := 0; i < len(str); i++ {
		hashValue += int(str[i]) * int(math.Pow(base, float64(len(str)-i-1)))
	}

	return hashValue % modulus
}

func GenerateHashFor(str string, start, end int) int {
	return hashStr(str[start-1 : end])
}

func main() {
	fmt.Println(GenerateHashFor("abcdefgh", 1, 1))
}
