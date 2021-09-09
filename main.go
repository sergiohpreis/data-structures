package main

import (
	"github.com/sergiohpreis/datastructures/linkedlist"
)

func createLinkedList() {
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
	linkedlist.Print(list)

	list = linkedlist.DeleteFirstNode(list)
	linkedlist.Print(list)

	list = linkedlist.DeleteLastNode(list)
	linkedlist.Print(list)

	list, _ = linkedlist.DeleteNthNode(list, 2)
	linkedlist.Print(list)
	linkedlist.PrintReverse(list)
	list = linkedlist.Reverse(list)
	linkedlist.Print(list)
}

func main() {
	createLinkedList()
}
