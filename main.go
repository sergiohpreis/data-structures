package main

import (
	"fmt"

	"github.com/sergiohpreis/datastructures/doublylinkedlist"
	"github.com/sergiohpreis/datastructures/singlylinkedlist"
	"github.com/sergiohpreis/datastructures/stack"
)

func createSinglyLinkedList() singlylinkedlist.List {
	fmt.Println("Singly Linked List")
	var list singlylinkedlist.List

	list.InsertAtEnd(1)
	list.InsertAtEnd(2)
	list.InsertAtEnd(3)
	list.InsertAtEnd(4)
	list.InsertAtEnd(5)
	list.InsertAtEnd(6)
	list.InsertAtEnd(7)
	list.InsertAtNth(100, 4)
	list.InsertAtBeginning(0)

	list.DeleteFirstNode()
	list.DeleteLastNode()
	list.DeleteNthNode(2)

	list.Reverse()

	return list
}

func createDoublyLinkedList() doublylinkedlist.List {
	fmt.Println("Doubly Linked List")
	var list doublylinkedlist.List

	list.InsertAtHead(1)
	list.InsertAtHead(2)
	list.InsertAtTail(3)
	list.InsertAtTail(4)

	return list
}

func createStack() stack.Stack {
	fmt.Println("Stack")
	var stack stack.Stack

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	return stack

}

func main() {
	singly := createSinglyLinkedList()
	singly.Print()
	singly.PrintReverse()
	singly.ReverseUsingStack()
	singly.Print()

	doubly := createDoublyLinkedList()
	doubly.Print()
	doubly.ReversePrint()

	s := createStack()
	s.Print()

	fmt.Println(stack.ReverseString("Teste"))
	fmt.Println(stack.ReverseStringUsingSwap("Teste"))
}
