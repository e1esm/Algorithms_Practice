package main

import (
	"fmt"
	"strings"
)

var answer []string

func generateCombinations(digits string, buffer []string, i int) {
	if i == len(digits) {
		answer = append(answer, strings.Join(buffer, ""))
		return
	}
	for _, symbol := range keyboard[rune(digits[i])] {
		buffer[i] = symbol
		generateCombinations(digits, buffer, i+1)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	answer = make([]string, 0)
	buffer := strings.Split(digits, "")

	generateCombinations(digits, buffer, 0)

	return answer
}

var keyboard = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func main() {
	fmt.Println(letterCombinations("23"))
}
