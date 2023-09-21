package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GetBiggestNumOf(nums []int) int {
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}

	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	var sb strings.Builder
	for _, numStr := range numStrs {
		sb.WriteString(numStr)
	}
	res, _ := strconv.Atoi(sb.String())
	return res
}

func main() {
	fmt.Println(GetBiggestNumOf([]int{2, 4, 5, 2, 10}))
	fmt.Println(GetBiggestNumOf([]int{15, 56, 2}))
	fmt.Println(GetBiggestNumOf([]int{1, 783, 2}))
}
