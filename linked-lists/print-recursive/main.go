package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func insert(head *Node, data int) *Node {
	temp := &Node{data: data}

	if head == nil {
		head = temp
	} else {
		temp1 := head

		for temp1.next != nil {
			temp1 = temp1.next
		}
		temp1.next = temp
	}

	return head
}

// For the normal traveste, iterative approach is more efficient
// We would need just one temp var, while in the recursive, we need a var to each call on the stack
func print(p *Node) {
	if p == nil {
		fmt.Println()
		return // Exit condition
	}
	fmt.Printf("%d ", p.data) // first print the value in the node
	print(p.next)             // recursive call
}

func reversePrint(p *Node) {
	if p == nil {
		fmt.Println()
		return // Exit condition
	}
	reversePrint(p.next)      // first recursive call
	fmt.Printf("%d ", p.data) // print the value in the node
}

func main() {
	var head *Node
	head = insert(head, 2)
	head = insert(head, 4)
	head = insert(head, 6)
	head = insert(head, 5)
	print(head)
	reversePrint(head)
}
