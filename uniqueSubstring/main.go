package main

import (
	"fmt"
	"strings"
)

func GetMaxLengthOfUniqueSubstring(str string) int {
	occurrences := make(map[uint8]int)
	var sb strings.Builder
	var maxLen int

	for i := 0; i < len(str); i++ {
		if _, ok := occurrences[str[i]]; ok {
			if sb.Len() > maxLen {
				maxLen = sb.Len()
			}
			sb.Reset()
			occurrences = make(map[uint8]int)
		}
		occurrences[str[i]]++
		sb.WriteByte(str[i])
	}

	return maxLen
}

func main() {
	fmt.Println(GetMaxLengthOfUniqueSubstring("abcabcbb"))
	fmt.Println(GetMaxLengthOfUniqueSubstring("bbbbb"))

}
