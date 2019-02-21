package common

type Node struct {
	Key    string
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

type AVL struct {
	root *Node
	size int
}

func (a *AVL) add(node *Node, key string, value int) {
	if  node == nil {
		a.size++

	}
}
