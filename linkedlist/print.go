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

func printReverse(p *Node) {
	if p == nil {
		return
	}

	printReverse(p.Next)
	fmt.Printf("%d ", p.Data)
}
func PrintReverse(p *Node) {
	printReverse(p)
	fmt.Println()
}
