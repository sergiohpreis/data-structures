package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

// Inserting in the beginning of the list
// head is the address (*) of the head, that points to the address(*) of a Node
func Insert(head **Node, x int) {
	// temp is the address of a node that doesn`t point to any node
	temp := &Node{data: x, next: nil}

	// if head already point to some node (list not empty)
	if *head != nil {
		temp.next = *head // temp next is the node that the head points
	}
	*head = temp // now head points to the new node (inserted)
}

func Print(head *Node) {
	fmt.Println("List is: ")

	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func main() {
	// head is a variable that points to a Node
	// head points to nil, the list is empty
	// In Go, the default value of a pointer type is nil
	var head *Node
	var n, x int

	fmt.Println("How many numbers?")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter the number")
		fmt.Scanf("%d", &x)

		// &head is the address of the head variable
		// We can change head value inside the function Insert
		Insert(&head, x)
		Print(head)
	}
}
