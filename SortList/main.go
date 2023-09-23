package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	arr := make([]int, 0)
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return buildList(sortArr(arr))
}

func sortArr(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func buildList(arr []int) *ListNode {
	if len(arr) == 0{
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++{
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func main() {

}