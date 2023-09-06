package main

import (
	"fmt"
	"math"
	"time"
)

func CountFibonnaci(n int) int {
	if n <= 1 {
		return 1
	}
	return CountFibonnaci(n-1) + CountFibonnaci(n-2)
}

func CommitsToDo(n int) int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	return CountFibonnaci(n)
}

func LastKNumber(n int, l int) int {
	result := n % int(math.Pow(10, float64(l)))
	return result
}
