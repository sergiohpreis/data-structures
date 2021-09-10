package doublylinkedlist

import "fmt"

func Print(head *Node) {
	fmt.Printf("Forward: ")
	current := head

	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func ReversePrint(head *Node) {
	if head == nil {
		return
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	fmt.Printf("Backward: ")
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Prev
	}

	fmt.Println()
}
