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
