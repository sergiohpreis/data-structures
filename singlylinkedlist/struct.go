package singlylinkedlist

type Node struct {
	next *Node
	data interface{}
}

type List struct {
	head *Node
}
