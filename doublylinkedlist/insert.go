package doublylinkedlist

func InsertAtHead(head *Node, data int) *Node {
	node := &Node{Data: data}

	if head != nil {
		head.Prev = node
		node.Next = head
	}

	head = node
	return head
}

func InsertAtTail(head *Node, data int) *Node {
	node := &Node{Data: data}

	current := head

	for current.Next != nil {
		current = current.Next
	}

	node.Prev = current
	current.Next = node
	return head
}
