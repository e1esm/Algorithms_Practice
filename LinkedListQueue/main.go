package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		val:  val,
		next: nil,
	}
}

type Queue struct {
	head *ListNode
	size int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		size: 0,
	}
}

func (q *Queue) put(val int) {
	defer func() {
		q.size++
	}()
	if q.head == nil {
		q.head = &ListNode{val: val}
		return
	}
	curr := q.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = NewListNode(val)
}

func (q *Queue) get() (*ListNode, error) {
	if q.size == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}
	curr := q.head
	q.head = q.head.next
	q.size--
	return curr, nil
}

func (q *Queue) Size() int {
	return q.size
}
