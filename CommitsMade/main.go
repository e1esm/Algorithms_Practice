package main

import (
	"fmt"
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
