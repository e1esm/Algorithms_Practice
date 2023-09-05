package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
}

func Iterate(head *ListNode) string {
	var sb strings.Builder
	curr := head

	for curr != nil {
		sb.WriteString(fmt.Sprintf("%d ", curr.val))
		curr = curr.next
	}

	return sb.String()
}

func main() {
	head := &ListNode{val: 0, next: &ListNode{val: 1, next: &ListNode{val: 2, next: &ListNode{val: 3, next: &ListNode{val: 4}}}}}
	fmt.Println(Iterate(head))
}
