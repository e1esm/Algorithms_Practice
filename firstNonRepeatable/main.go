package main

import "strings"

func FirstNonRepeating(str string) string {
	str = strings.ToLower(str)
	occurrences := make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		occurrences[str[i]]++
	}

	for i := 0; i < len(str); i++ {
		if occurrences[str[i]] == 1 {
			return string(str[i])
		}
	}
	return ""
}

func main() {

}
