package linkedlist

func reverse(head **Node, p *Node) {
	if p.Next == nil {
		*head = p
		return
	}

	reverse(head, p.Next)
	prev := p // n-1
	current := p.Next
	prev.Next = nil
	current.Next = prev
}

func Reverse(p *Node) *Node {
	var head *Node
	reverse(&head, p)
	return head
}
