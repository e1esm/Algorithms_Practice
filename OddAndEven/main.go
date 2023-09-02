package main

import (
	"fmt"
	"time"
)

var WIN = "WIN"
var DEFEAT = "FAIL"

func GetParityStatus(arr []int) string {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	isOdd := false
	isEven := false

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			isOdd = true
		} else {
			isEven = true
		}
		if isEven && isOdd {
			return DEFEAT
		}
	}
	return WIN
}
