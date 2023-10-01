package main

import "fmt"

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	var remainder int
	if k > len(nums) && len(nums)%k == 0 {
		for k > 0 {
			remainder = k % len(nums)
			k /= len(nums)
		}
	} else {
		remainder = k
	}

	for remainder != 0 {
		last := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = last
		remainder--
	}
	//fmt.Println(nums)
}

func main() {
	rotate([]int{-1}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
	rotate([]int{1, 2, 3, 4, 5, 6}, 11)
}
