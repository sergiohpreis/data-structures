package linkedlists

import "fmt"

func Print(head *Node) {
	current := head

	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
