package main

import (
	"fmt"
	"strings"
	"time"
)

func IsPalindrome(str string) bool {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	str = strings.ToLower(str)
	isPalindrome := true
	var alphabetStart uint8 = 97
	var alphabetEnd uint8 = 122

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if IsInRange(str[i], alphabetStart, alphabetEnd) && IsInRange(str[j], alphabetStart, alphabetEnd) {
			if str[i] != str[j] {
				isPalindrome = false
				break
			}
		}
		if !IsInRange(str[i], alphabetStart, alphabetEnd) {
			j++
		}
		if !IsInRange(str[j], alphabetStart, alphabetEnd) {
			i--
		}
	}

	return isPalindrome
}

func IsInRange(target, alphabetStart, alphabetEnd uint8) bool {
	if target >= alphabetStart && target <= alphabetEnd {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
}
