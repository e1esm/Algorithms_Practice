package main

import (
	"fmt"
	"strings"
)

func BubbleSort(n int, arr []int) string {
	var sb strings.Builder

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				sb.WriteString(fmt.Sprintf("%v\n", arr))
			}
		}
	}

	return sb.String()[:len(sb.String())-1]
}

func main() {
	fmt.Println(BubbleSort(5, []int{1, 3, 4, 2, 1, 9}))
}
