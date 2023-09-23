package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var preToBeDeleted *ListNode
	var size int
	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}

	i := 0
	curr = head

	if size - n == 0{
		return head.Next
	}

	for i < size-n {
		i++
		preToBeDeleted = curr
		curr = curr.Next
	}

	preToBeDeleted.Next = preToBeDeleted.Next.Next
	return head
}

func IterateOverNodes(head *ListNode){
	curr := head
	if curr == nil{
		fmt.Println("Empty")
	}
	for curr != nil{
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	//n := 2
	//head := &ListNode{Val: 1}
	//n := 1

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	n := 2

	head = removeNthFromEnd(head, n)
	IterateOverNodes(head)
}