package main

import "fmt"

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func Reverse(head *ListNode) {
	curr := head

	for curr != nil {
		temp := curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}
}
