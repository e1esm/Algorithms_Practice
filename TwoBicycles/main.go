package main

import "fmt"

/*
1 2 4 4 6 8
 2 4
*/

func indexFrom(arr []int, currIndex, price int) int {
	if currIndex < 0 || currIndex >= len(arr) {
		return -1
	}

	if price > arr[currIndex] {
		return indexFrom(arr, currIndex+1, price)
	} else if price <= arr[currIndex] {
		if currIndex-1 >= 0 && price <= arr[currIndex-1] {
			return indexFrom(arr, currIndex-1, price)
		}
	}
	return currIndex + 1
}

func GetDaysIndexes(arr []int, price int) (int, int) {
	return indexFrom(arr, len(arr)/2, price), indexFrom(arr, len(arr)/2, price*2)
}

func main() {
	fmt.Println(GetDaysIndexes([]int{1, 2, 4, 4, 6, 8}, 3))
}
