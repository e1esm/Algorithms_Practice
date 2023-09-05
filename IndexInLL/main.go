package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func GetIndexOf(head *ListNode, target, index int) int {
	if head != nil && head.val != target {
		return GetIndexOf(head.next, target, index+1)
	}
	if head != nil && head.val == target {
		return index
	}
	return -1
}

func main() {
	head := &ListNode{val: 10,
		next: &ListNode{val: 15,
			next: &ListNode{val: 22,
				next: &ListNode{val: 5,
					next: &ListNode{val: 55,
						next: &ListNode{val: 30}}}}}}
	curr := head
	fmt.Println(GetIndexOf(curr, 22, 0))
}
