package main

import (
	"strconv"
	"strings"
)

func GetSortedWardrobe(arr []int) string {
	numOccurrences := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		numOccurrences[arr[i]] = append(numOccurrences[arr[i]], arr[i])
	}
	var sb strings.Builder
	for i := 0; i <= 2; i++ {
		tempNumStrRepr := ""
		for j := 0; j < len(numOccurrences[i]); j++ {
			if tempNumStrRepr == "" {
				tempNumStrRepr = strconv.Itoa(i)
			}
			sb.WriteString(tempNumStrRepr + " ")
		}
	}
	return strings.TrimRight(sb.String(), " ")
}
