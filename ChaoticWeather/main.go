package main

import (
	"fmt"
	"time"
)

func GetChaoticLength(arr []int) int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	length := 0

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			if arr[i] > arr[i+1] {
				length++
			}
		} else if i < len(arr)-1 {
			if arr[i] > arr[i+1] && arr[i] > arr[i-1] {
				length++
			}
		} else {
			if arr[i] > arr[i-1] {
				length++
			}
		}

	}
	return length
}
