package main

import (
	"fmt"
	"strings"
	"time"
)

func GetLongestWord(sentence string) (string, int) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	var word strings.Builder
	var maxWord string
	var maxLength int
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			if word.Len() > maxLength {
				maxLength = word.Len()
				maxWord = word.String()
				word.Reset()
			}
		}
		word.WriteByte(sentence[i])
	}
	return maxWord, maxLength
}
