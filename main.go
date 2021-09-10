package main

import (
	"github.com/sergiohpreis/datastructures/doublylinkedlist"
	"github.com/sergiohpreis/datastructures/linkedlist"
	"github.com/sergiohpreis/datastructures/stack"
)

func createSinglyLinkedList() *linkedlist.Node {
	var list *linkedlist.Node

	list = linkedlist.InsertAtEnd(list, 1)
	list = linkedlist.InsertAtEnd(list, 2)
	list = linkedlist.InsertAtEnd(list, 3)
	list = linkedlist.InsertAtEnd(list, 4)
	list = linkedlist.InsertAtEnd(list, 5)
	list = linkedlist.InsertAtEnd(list, 6)
	list = linkedlist.InsertAtEnd(list, 7)
	list, _ = linkedlist.InsertAtNth(list, 100, 4)
	list = linkedlist.InsertAtBeginning(list, 0)

	list = linkedlist.DeleteFirstNode(list)
	list = linkedlist.DeleteLastNode(list)
	list, _ = linkedlist.DeleteNthNode(list, 2)

	list = linkedlist.Reverse(list)

	return list
}

func createDoublyLinkedList() *doublylinkedlist.Node {
	var head *doublylinkedlist.Node

	head = doublylinkedlist.InsertAtHead(head, 1)
	head = doublylinkedlist.InsertAtHead(head, 2)
	head = doublylinkedlist.InsertAtTail(head, 3)
	head = doublylinkedlist.InsertAtTail(head, 4)
	return head
}

func createStack() *stack.Stack {
	var stack stack.Stack

	stack.Create()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	return &stack

}

func main() {
	singly := createSinglyLinkedList()
	linkedlist.Print(singly)
	linkedlist.PrintReverse(singly)

	doubly := createDoublyLinkedList()
	doublylinkedlist.Print(doubly)
	doublylinkedlist.ReversePrint(doubly)

	stack := createStack()
	stack.Print()
}
