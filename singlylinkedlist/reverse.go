package singlylinkedlist

func reverse(head **Node, p *Node) {
	if p.next == nil {
		*head = p
		return
	}

	reverse(head, p.next)
	current := p.next
	current.next = p
	p.next = nil
}

func (list *List) Reverse() {
	reverse(&list.head, list.head)
}

// TODO: Remove from here and add to stack
func top(stack []*Node) *Node {
	return stack[len(stack)-1]
}
func pop(stack []*Node) []*Node {
	return stack[:len(stack)-1]
}

func (list *List) ReverseUsingStack() {
	if list.head == nil {
		return
	}

	var stack []*Node
	temp := list.head

	for temp != nil {
		stack = append(stack, temp)
		temp = temp.next
	}

	temp = top(stack)
	list.head = temp

	stack = pop(stack)

	for len(stack) > 0 {
		temp.next = top(stack)
		stack = pop(stack)
		temp = temp.next
	}
	temp.next = nil
}
