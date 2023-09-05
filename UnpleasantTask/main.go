package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func RemoveNode(head *ListNode, target int) int {
	if head.val == target {
		head = head.next
		return target
	}
	curr := head
	for curr.next != nil {
		if curr.next.val == target {
			curr.next = curr.next.next
			return target
		}
		curr = curr.next
	}

	return -1
}

func Iterate(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

func main() {
	head := &ListNode{val: 0, next: &ListNode{val: 1, next: &ListNode{val: 2, next: &ListNode{val: 3, next: &ListNode{val: 4, next: &ListNode{val: 5, next: &ListNode{val: 6}}}}}}}
	fmt.Println(RemoveNode(head, 1))
	Iterate(head)
}
