package main

import (
	"fmt"
	"math/rand"
)

func merge2SortedLists(l1, l2 *LinkedListNode) *LinkedListNode {
	dummy := &LinkedListNode{data: -1}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			prev.next = l1
			l1 = l1.next
		} else {
			prev.next = l2
			l2 = l2.next
		}
		prev = prev.next
	}

	if l1 == nil {
		prev.next = l2
	} else {
		prev.next = l1
	}

	return dummy.next
}

func mergeKSortedLists(lists []*LinkedListNode) *LinkedListNode {
	if len(lists) > 0 {
		res := lists[0]

		for i := 1; i < len(lists); i++ {
			res = merge2SortedLists(res, lists[i])
		}
		return res
	}

	return &LinkedListNode{data: -1}
}

func main() {
	a := createLinkedList([]int{11, 41, 51})
	b := createLinkedList([]int{21, 23, 42})
	c := createLinkedList([]int{25, 56, 66, 72})
	list1 := []*LinkedListNode{a, b, c}
	display(mergeKSortedLists(list1))
}

type LinkedListNode struct {
	key              int
	data             int
	next             *LinkedListNode
	arbitrartPointer *LinkedListNode
}

func createLinkedList(lst []int) *LinkedListNode {
	var head *LinkedListNode
	var tail *LinkedListNode
	for _, x := range lst {
		newNode := &LinkedListNode{data: x}
		if head == nil {
			head = newNode
		} else {
			tail.next = newNode
		}
		tail = newNode
	}
	return head
}

func insertAtHead(head *LinkedListNode, data int) *LinkedListNode {
	newNode := &LinkedListNode{data: data}
	newNode.next = head
	return newNode
}

func insertAtTail(head *LinkedListNode, data int) *LinkedListNode {
	newNode := &LinkedListNode{data: data}
	if head == nil {
		return newNode
	}

	temp := head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	return head
}

func createRandomList(length int) *LinkedListNode {
	var listHead *LinkedListNode
	for i := 0; i < length; i++ {
		listHead = insertAtHead(listHead, rand.Intn(100))
	}
	return listHead
}

func toList(head *LinkedListNode) []int {
	var lst []int
	temp := head
	for temp != nil {
		lst = append(lst, temp.data)
		temp = temp.next
	}
	return lst
}

func display(head *LinkedListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d", temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}

func mergeAlternating(list1, list2 *LinkedListNode) *LinkedListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := list1

	for list1.next != nil && list2 != nil {
		temp := list2
		list2 = list2.next

		temp.next = list1.next
		list1.next = temp
		list1 = temp.next
	}

	if list1.next == nil {
		list1.next = list2
	}

	return head
}

func isEqual(list1, list2 *LinkedListNode) bool {
	if list1 == list2 {
		return true
	}

	for list1 != nil && list2 != nil {
		if list1.data != list2.data {
			return false
		}

		list1 = list1.next
		list2 = list2.next
	}

	return list1 == list2
}
