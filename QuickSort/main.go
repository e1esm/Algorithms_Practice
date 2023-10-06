package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr) // 1

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	fmt.Println(quickSort([]int{5, 9, 6, 3, 1}))
}
