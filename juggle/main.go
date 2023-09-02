package main

import (
	"fmt"
	"time"
)

func getPoints(tries int, matrix [][]rune) int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	points := 0
	matrixLength := len(matrix)

	currentTime := 48

	for i := 0; i < matrixLength; i++ {
		runeRepr := rune(currentTime)
		currTries := tries
		amountFound := 0
		for j := 0; j < matrixLength; j++ {
			if matrix[i][j] == runeRepr {
				currTries--
				amountFound++
			}
		}
		if currentTime+1 < 58 {
			currentTime++
		}
		if currTries >= 0 && amountFound > 0 {
			points++
		}
	}

	return points
}
