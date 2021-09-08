package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head *Node

func insert(data int) {
	node := &Node{data: data}

	if head == nil {
		head = node
		return
	}

	current := head

	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func print() {
	current := head

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func reverse(p *Node) {
	if p.next == nil { // last node reached
		head = p
		return
	}

	reverse(p.next)

	q := p.next
	q.next = p
	p.next = nil
}

func main() {
	insert(1)
	insert(2)
	insert(3)
	insert(4)
	print()
	reverse(head)
	print()
}
