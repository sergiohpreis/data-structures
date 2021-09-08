package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head *Node

func insert(data int) {
	temp1 := &Node{data: data}

	if head == nil {
		head = temp1
		return
	}

	temp2 := head

	for temp2.next != nil {
		temp2 = temp2.next
	}

	temp2.next = temp1
}

func reverse() {
	var prev *Node

	current := head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	head = prev
}

func print() {
	temp := head

	for temp != nil {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	insert(2)
	insert(4)
	insert(6)
	insert(8)
	print()
	reverse()
	print()
}
