package main

import (
	"fmt"
	"strings"
	"time"
)

func BinaryOf(num int) string {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	var sb strings.Builder

	for num != 0 {
		sb.WriteString(fmt.Sprintf("%d", num%2))
		num = num / 2
	}
	result := []rune(sb.String())
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
