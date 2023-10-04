package main

const bucketsCount int = 8

type Node struct {
	key  int
	val  int
	next *Node
}

type MyHashMap struct {
	buckets [bucketsCount]*Node
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: [bucketsCount]*Node{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.Remove(key)
	h := key % bucketsCount

	if this.buckets[h] == nil {
		this.buckets[h] = &Node{key: key, val: value}
		return
	}

	curr := this.buckets[h]

	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &Node{key: key, val: value}

}

func (this *MyHashMap) Get(key int) int {
	h := key % bucketsCount

	curr := this.buckets[h]

	for curr != nil {
		if curr.key == key {
			return curr.val
		}
		curr = curr.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := key % bucketsCount

	curr := this.buckets[h]

	if curr == nil {
		return
	}

	if curr.key == key {
		this.buckets[h] = curr.next
	}

	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
		}
		curr = curr.next
	}

}
