package singlylinkedlist

import (
	"errors"
	"fmt"
)

func (list *List) DeleteFirstNode() {
	list.head = list.head.next
}

func (list *List) DeleteLastNode() {
	current := list.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}

func (list *List) DeleteNthNode(n int) error {
	if n == 1 {
		list.head = list.head.next
	} else {
		current := list.head

		for i := 1; i < n-1; i++ {
			if current.next == nil {
				e := fmt.Sprintf("index %d out of range", n)
				return errors.New(e)
			}
			current = current.next
		}

		current.next = current.next.next
	}

	return nil
}
