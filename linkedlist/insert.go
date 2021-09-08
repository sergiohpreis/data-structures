package linkedlist

import (
	"errors"
	"fmt"
)

func InsertAtBeginning(head *Node, data int) *Node {
	node := &Node{Data: data}

	if head == nil {
		head = node
	} else {
		node.Next = head
		head = node
	}

	return head
}

func InsertAtEnd(head *Node, data int) *Node {
	node := &Node{Data: data}

	if head == nil {
		head = node
	} else {
		current := head

		for current.Next != nil {
			current = current.Next
		}

		current.Next = node
	}

	return head
}

func InsertAtNth(head *Node, data int, n int) (*Node, error) {
	node := &Node{Data: data}

	if n == 1 {
		node.Next = head
		return node, nil
	}

	current := head

	for i := 1; i < n-1; i++ {
		if current.Next == nil {
			e := fmt.Sprintf("index %d out of range", n)
			return head, errors.New(e)
		}
		current = current.Next
	}

	node.Next = current.Next
	current.Next = node

	return head, nil
}
