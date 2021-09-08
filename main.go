package main

import (
	"github.com/sergiohpreis/datastructures/linkedlists"
)

func linkedlist() {
	var list *linkedlists.Node

	list = linkedlists.InsertAtEnd(list, 1)
	list = linkedlists.InsertAtEnd(list, 2)
	list = linkedlists.InsertAtEnd(list, 3)
	list = linkedlists.InsertAtEnd(list, 4)
	list = linkedlists.InsertAtEnd(list, 5)
	list = linkedlists.InsertAtEnd(list, 6)
	list = linkedlists.InsertAtEnd(list, 7)
	list, _ = linkedlists.InsertAtNth(list, 100, 4)
	list = linkedlists.InsertAtBeginning(list, 0)

	linkedlists.Print(list)
}

func main() {
	linkedlist()
}
