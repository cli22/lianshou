package datastruct

type Node struct {
	IsWord bool
	Next   map[uint8]Node
}

func NewNode() Node {
	return Node{IsWord: false}
}

type Trie struct {
	root Node
	size int
}

func (t *Trie) GetSize() int {
	return t.size
}

// 添加单词
func (t *Trie) Add(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.Next[c]; !ok {
			cur.Next[c] = NewNode()
		}
		cur = cur.Next[c]
	}

	if !cur.IsWord {
		cur.IsWord = true
		t.size++
	}
}

// 查询单词
func (t *Trie) Find(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.Next[c]; !ok {
			return false
		}
		cur = cur.Next[c]
	}
	return cur.IsWord
}

// 是否存在以prefix为前缀的单词
func (t *Trie) IsPrefix(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.Next[c]; !ok {
			return false
		}
		cur = cur.Next[c]
	}
	return true
}

//是否包含"."模式匹配的查找
func (t *Trie) Match(node Node, word string, index int) bool {
	if index == len(word) {
		return node.IsWord
	}

	c := word[index]
	if c != '.' {
		if _, ok := node.Next[c]; !ok {
			return false
		}
		return t.Match(node.Next[c], word, index+1)
	} else {
		for _, value := range node.Next {
			if t.Match(value, word, index+1) {
				return true
			}
		}
		return false
	}
}
