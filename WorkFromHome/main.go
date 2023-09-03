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
	return sb.String()
}
