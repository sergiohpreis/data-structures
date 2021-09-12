package singlylinkedlist

import (
	"errors"
	"fmt"
)

func (list *List) InsertAtBeginning(data interface{}) {
	node := &Node{data: data}

	if list.head != nil {
		node.next = list.head
	}

	list.head = node
}

func (list *List) InsertAtEnd(data interface{}) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}

		current.next = node
	}
}

func (list *List) InsertAtNth(data interface{}, n int) error {
	node := &Node{data: data}

	if n == 1 {
		node.next = list.head
		return nil
	}

	current := list.head

	for i := 1; i < n-1; i++ {
		if current.next == nil {
			e := fmt.Sprintf("index %d out of range", n)
			return errors.New(e)
		}
		current = current.next
	}

	node.next = current.next
	current.next = node

	return nil
}
