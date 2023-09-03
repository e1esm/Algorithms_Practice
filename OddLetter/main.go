package main

import (
	"fmt"
	"time"
)

type Pair struct {
	firstAmount  int
	secondAmount int
}

func GetOddLetter(first, second string) byte {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	pairs := make(map[byte]Pair, 0)
	maxLen := Max(len(first), len(second))

	for i := 0; i < maxLen; i++ {
		if i < len(second) {
			v, ok := pairs[second[i]]
			if ok {
				v.secondAmount++
				pairs[second[i]] = v
			} else {
				pairs[second[i]] = Pair{secondAmount: 1}
			}
		}
		if i < len(first) {
			v, ok := pairs[first[i]]
			if ok {
				v.firstAmount++
				pairs[first[i]] = v
			} else {
				pairs[first[i]] = Pair{firstAmount: 1}
			}
		}
	}

	for k, v := range pairs {
		if v.firstAmount != v.secondAmount {
			return k
		}
	}
	return 0
}

func Max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func main() {
	fmt.Println(GetOddLetter("hey", "he"))
}
