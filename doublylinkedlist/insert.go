package doublylinkedlist

func (list *List) InsertAtHead(data interface{}) {
	node := &Node{data: data}

	if list.head != nil {
		list.head.prev = node
		node.next = list.head
	}

	list.head = node
}

func (list *List) InsertAtTail(data interface{}) {
	node := &Node{data: data}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	node.prev = current
	current.next = node
}
