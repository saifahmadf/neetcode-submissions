type PrefixTree struct {
	Children [26]*PrefixTree
	EndOfWord bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &PrefixTree{}
		}
		node = node.Children[idx]
	}
	node.EndOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			return false
		}
		node = node.Children[idx]
	}
	return node.EndOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		idx := c - 'a'
		if node.Children[idx] == nil {
			return false
		}
		node = node.Children[idx]
	}
	return true
}
