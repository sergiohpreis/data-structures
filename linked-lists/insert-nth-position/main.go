package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head *Node

func insert(data int, n int) {
	temp1 := &Node{data: data}

	// inserting at the beggining
	if n == 1 {
		temp1.next = head
		head = temp1
		return
	}

	//Go to the (n-1)th node
	temp2 := head
	for i := 0; i < n-2; i++ {
		temp2 = temp2.next
	}

	temp1.next = temp2.next
	temp2.next = temp1
}

func print() {
	temp := head

	for temp != nil {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
	}
}

func main() {
	insert(2, 1) // List: 2
	insert(3, 2) // List: 2, 3
	insert(4, 1) // List: 4, 2, 3
	insert(5, 2) // List: 4, 5, 2, 3
	print()      // 4 5 2 3
}
