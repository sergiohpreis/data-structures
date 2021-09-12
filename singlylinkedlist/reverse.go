package singlylinkedlist

import "github.com/sergiohpreis/datastructures/stack"

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

func (list *List) ReverseUsingStack() {
	if list.head == nil {
		return
	}

	var stack stack.Stack
	temp := list.head

	for temp != nil {
		stack.Push(temp)
		temp = temp.next
	}

	temp = (stack.Top()).(*Node)
	list.head = temp

	stack.Pop()

	for len(stack) > 0 {
		temp.next = (stack.Top()).(*Node)
		stack.Pop()
		temp = temp.next
	}
	temp.next = nil
}
