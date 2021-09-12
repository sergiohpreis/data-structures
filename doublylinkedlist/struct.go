package doublylinkedlist

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

type List struct {
	head *Node
}
