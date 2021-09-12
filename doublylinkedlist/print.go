package doublylinkedlist

import "fmt"

func (list *List) Print() {
	current := list.head

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (list *List) ReversePrint() {
	if list.head == nil {
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}

	fmt.Println()
}
