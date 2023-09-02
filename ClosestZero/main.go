package main

import (
	"fmt"
	"time"
)

func getDistanceArray(houseNumbers []int) []int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	i := 0
	length := len(houseNumbers)
	zeroIndexes := make([]int, 0)
	distances := make([]int, length)

	for zeroIndex := 0; zeroIndex < length; zeroIndex++ {
		if houseNumbers[zeroIndex] == 0 {
			zeroIndexes = append(zeroIndexes, zeroIndex)
		}
	}

	for currIndex := 0; currIndex < length; currIndex++ {
		minimumLength := length + 1
		for _, zIndex := range zeroIndexes {
			diff := getAbsResult(zIndex, currIndex)
			if diff < minimumLength {
				minimumLength = diff
			}
		}
		distances[i] = minimumLength
		i++
	}

	return distances
}

func getAbsResult(first, second int) int {
	if second > first {
		return second - first
	}
	return first - second
}

func main() {

}
