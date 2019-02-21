package main

type Node struct {
	Next  map[uint8]Node
	Value int
}

type Trie struct {
	root Node
}

func NewNode() Node {
	return Node{Value: 0}
}

// leetcode add sum
func (t *Trie) insert(word string, value int) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.Next[c]; !ok {
			cur.Next[c] = NewNode()
		}
		cur = cur.Next[c]
	}

	cur.Value = value
}

func (t *Trie) Sum(prefix string) int {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := cur.Next[c]; !ok {
			return 0
		}
		cur = cur.Next[c]
	}
	return sum(cur)
}

func sum(node Node) int {
	res := node.Value
	for _, value := range node.Next {
		res += sum(value)
	}
	return res
}
