package main

import (
	"fmt"
)

type LRUCache struct {
	capacity int

	//LinkedListNode holds key and value pairs
	cache     map[int]*LinkedListNode
	cacheVals LinkedList
}

func (lr *LRUCache) Get(key int) *LinkedListNode {
	if _, ok := lr.cache[key]; !ok {
		return nil
	} else {
		value := lr.cache[key].data
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		return lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Set(key, value int) {
	if _, ok := lr.cache[key]; !ok {
		if lr.cacheVals.size >= lr.capacity {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
			delete(lr.cache, lr.cacheVals.GetHead().key)
			lr.cacheVals.RemoveHead()
		} else {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
		}
	} else {
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		lr.cache[key] = lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Print() {
	curr := lr.cacheVals.head
	for curr != nil {
		fmt.Printf("(%v,%v)", curr.key, curr.data)
		curr = curr.next
	}
	fmt.Println("")
}

func main() {
	cache := &LRUCache{capacity: 3, cache: make(map[int]*LinkedListNode)}
	fmt.Println("The most recently watched titles are: (key, value)")
	cache.Set(10, 20)
	cache.Print()

	cache.Set(15, 25)
	cache.Print()

	cache.Set(20, 30)
	cache.Print()

	cache.Set(25, 35)
	cache.Print()

	cache.Set(5, 40)
	cache.Print()

	cache.Get(25)
	cache.Print()
}

type LinkedListNode struct {
	key  int
	data int
	next *LinkedListNode
	prev *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func (l *LinkedList) InsertAtHead(key, data int) {

	newNode := &LinkedListNode{key: key, data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList) InsertAtTail(key, data int) {
	newNode := &LinkedListNode{key: key, data: data}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		newNode.next = nil
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		newNode.next = nil
	}
	l.size++
}

func (l *LinkedList) GetHead() *LinkedListNode {
	return l.head
}

func (l *LinkedList) GetTail() *LinkedListNode {
	return l.tail
}

func (l *LinkedList) RemoveNode(node *LinkedListNode) *LinkedListNode {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {

		l.head = l.head.next
	}
	if node == l.tail {
		l.tail = l.tail.prev
	}
	l.size--
	return node
}

func (l *LinkedList) Remove(data int) {
	i := l.GetHead()
	for i != nil {
		if i.data == data {
			l.RemoveNode(i)
		}
		i = i.next
	}
}

func (l *LinkedList) RemoveHead() *LinkedListNode {
	return l.RemoveNode(l.head)
}

func (l *LinkedList) RemoveTail() *LinkedListNode {
	return l.RemoveNode(l.tail)
}
