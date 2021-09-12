package singlylinkedlist

import "fmt"

func (list *List) Print() {
	current := list.head

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func printReverse(p *Node) {
	if p == nil {
		return
	}

	printReverse(p.next)
	fmt.Printf("%d ", p.data)
}
func (list *List) PrintReverse() {
	printReverse(list.head)
	fmt.Println()
}
