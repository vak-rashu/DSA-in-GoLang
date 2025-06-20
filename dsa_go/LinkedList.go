package dsa

import (
	"fmt"
)

// implementing a Linked List
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// appending to the end
func (l *LinkedList) Append(data int) {
	NewNode := &Node{data: data}

	if l.head == nil {
		l.head = NewNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = NewNode
	}
}

// displaying by traversing the list
func (l *LinkedList) Display() {
	current := l.head
	if current == nil {
		fmt.Println("Empty Linked List")
	}
	fmt.Println("Linked List:")
	for current != nil {
		fmt.Println(current, "->", current.data)
		current = current.next
	}
}

// inserting at a given position(index)
func (l *LinkedList) InsertAtPos(data, pos int) {
	NewNode := &Node{data: data}
	current := l.head
	for range pos - 1 {
		/*if current == nil || current.next == nil {
			return l.head // Position exceeds list length
		}*/
		current = current.next
	}
	NewNode.next = current.next
	current.next = NewNode
}

// deleting the given data
func (l *LinkedList) Delete(data int) int {
	current := l.head
	if current.data == data {
		l.head = current.next
		return data
	}
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return data
		}
		current = current.next
	}
	return 0
}
func ImplementOps() {
	list := LinkedList{}
	list.Append(10)
	list.Append(100)
	list.Append(10000)
	list.InsertAtPos(1000, 2)
	list.Display()
	fmt.Println("Deleted Value is:", list.Delete(10))
	list.Display()
}
