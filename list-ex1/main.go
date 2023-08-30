package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[any]) length() int {
	if l.next == nil {
		return 1
	}
	return 1 + l.next.length()
}

func main() {
	// Create a List of integers
	var head *List[int] = &List[int]{val: 0}
	var tail *List[int] = head

	for i := 1; i < 10; i++ {
		newNode := &List[int]{val: i * 10}
		tail.next = newNode
		tail = tail.next
	}

	// Traverse and print the List
	current := head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

	fmt.Println("l", head.length())
}
