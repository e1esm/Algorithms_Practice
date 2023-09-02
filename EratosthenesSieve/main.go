package main

import "fmt"

func EratosthenesSieve(limit int) {
	arr := make([]bool, limit+1)
	arr[0], arr[1] = false, false

	for i := 2; i < len(arr); i++ {
		arr[i] = true
	}

	for num := 2; num < len(arr); num++ {
		if arr[num] {
			for j := num * num; j < len(arr)+1; j += num {
				arr[j] = false
			}
		}
	}

	for i, v := range arr {
		if v {
			fmt.Println(i)
		}
	}
}

func main() {
	EratosthenesSieve(10)
}
