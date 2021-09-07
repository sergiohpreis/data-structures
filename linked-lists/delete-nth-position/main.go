package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var head *Node

// inserting in the end of the list
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

func print() {
	temp := head

	for temp != nil {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func delete(n int) {
	temp1 := head

	// wants to remove the first element
	if n == 1 {
		head = temp1.next // head now points to the second Node
		return
	}

	for i := 0; i < n-2; i++ {
		temp1 = temp1.next // points to (n-1)th Node
	}

	temp2 := temp1.next     // nth Node
	temp1.next = temp2.next // (n+1)th Node
}

func main() {
	insert(2)
	insert(4)
	insert(6)
	insert(5) // List: 2 4 6 5

	print()

	var n int
	fmt.Println("Enter a position")
	fmt.Scanf("%d", &n)

	delete(n)
	print()
}
