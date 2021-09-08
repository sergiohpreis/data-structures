package linkedlist

import "fmt"

func Print(head *Node) {
	current := head

	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func reverse(p *Node) {
	if p == nil {
		return
	}

	reverse(p.Next)
	fmt.Printf("%d ", p.Data)
}
func PrintReverse(p *Node) {
	reverse(p)
	fmt.Println()
}
