package linkedlist

import (
	"errors"
	"fmt"
)

func DeleteFirstNode(head *Node) *Node {
	head = head.Next
	return head
}

// 1 -> 2 -> 3 -> 4
func DeleteLastNode(head *Node) *Node {
	current := head

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil

	return head
}

func DeleteNthNode(head *Node, n int) (*Node, error) {
	if n == 1 {
		head = head.Next
	} else {
		current := head

		for i := 1; i < n-1; i++ {
			if current.Next == nil {
				e := fmt.Sprintf("index %d out of range", n)
				return head, errors.New(e)
			}
			current = current.Next
		}

		current.Next = current.Next.Next
	}

	return head, nil
}
