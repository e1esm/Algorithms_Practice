package main

import (
	"fmt"
	"math"
)

const (
	base    = 1000
	modulus = 1000009
)

func hashString(str string) int {
	var hashedValue int
	for i := 0; i < len(str); i++ {
		hashedValue += int(str[i]) * int(math.Pow(float64(base), float64(len(str)-i-1)))
	}
	return hashedValue
}

func isAlreadyIn(str string, arr *[]int) (int, bool) {
	hash := hashString(str)
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == hash {
			return -1, true
		}
	}
	return hash, false
}

func GetUniqueActivityList(arr []string) []string {
	hashedValues := make([]int, len(arr))
	uniqueStrings := make([]string, 0)

	for i := 0; i < len(arr); i++ {
		if hash, isOK := isAlreadyIn(arr[i], &hashedValues); !isOK {
			hashedValues = append(hashedValues, hash)
			uniqueStrings = append(uniqueStrings, arr[i])
		}
	}
	return uniqueStrings
}

func main() {
	fmt.Println(GetUniqueActivityList([]string{
		"вышивание крестиком",
		"рисование мелками на парте",
		"настольный керлинг",
		"настольный керлинг",
		"кухня африканского племени ужасмай",
		"тяжелая атлетика",
		"таракановедение",
		"таракановедение",
	}))
}
