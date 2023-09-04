package main

import (
	"fmt"
	"time"
)

func ResultOf(sl1, sl2 []int) []int {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	result := getNumFrom(sl1) + getNumFrom(sl2)
	arr := make([]int, 0)

	for result != 0 {
		arr = append(arr, result%10)
		result = result / 10
	}

	return reverse(arr)
}

func getNumFrom(sl []int) int {
	var num = 0
	for i := 0; i < len(sl); i++ {
		num = num*10 + sl[i]
	}
	return num
}

func reverse(sl []int) []int {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return sl
}
