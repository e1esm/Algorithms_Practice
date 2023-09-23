package main

import (
	"fmt"
	"strings"
)

func InsertionSort(n int, arr []int) string {
	var sb strings.Builder

	for i := 1; i < n; i++ {
		toBeReplaced := arr[i]
		j := i
		for j >= 0 && arr[j-1] > toBeReplaced {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = toBeReplaced
		sb.WriteString(fmt.Sprintf("%v\n", arr))
	}

	return sb.String()[:len(sb.String())-1]
}

func main() {
	fmt.Println(InsertionSort(7, []int{1, 2, 4, 3, 7, 6, 5}))
}
