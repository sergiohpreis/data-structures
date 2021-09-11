package linkedlist

func reverse(head **Node, p *Node) {
	if p.Next == nil {
		*head = p
		return
	}

	reverse(head, p.Next)
	current := p.Next
	current.Next = p
	p.Next = nil
}

func Reverse(p *Node) *Node {
	var head *Node
	reverse(&head, p)
	return head
}

// TODO: Remove from here and add to stack
func top(stack []*Node) *Node {
	return stack[len(stack)-1]
}
func pop(stack []*Node) []*Node {
	return stack[:len(stack)-1]
}

func ReverseUsingStack(head *Node) *Node {
	if head == nil {
		return head
	}

	var stack []*Node
	temp := head

	for temp != nil {
		stack = append(stack, temp)
		temp = temp.Next
	}

	temp = top(stack)
	head = temp

	stack = pop(stack)

	for len(stack) > 0 {
		temp.Next = top(stack)
		stack = pop(stack)
		temp = temp.Next
	}
	temp.Next = nil

	return head
}
