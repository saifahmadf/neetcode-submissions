type WordDictionary struct {
    Children [26]*WordDictionary
	IsEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string)  {
	node := this
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &WordDictionary{}
		}
		node = node.Children[idx]
	}
	node.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.dfs(this, word, 0)
}

func (this *WordDictionary) dfs(node *WordDictionary, word string, i int) bool {
	if i == len(word) {
		return node.IsEnd
	}

	c := word[i]

	if c == '.' {
		for j := 0; j < 26; j++ {
			if node.Children[j] != nil && this.dfs(node.Children[j], word, i+1){
				return true
			}
		}
		return false
	}

	idx := c - 'a'
	if node.Children[idx] == nil {
		return false
	}
	return this.dfs(node.Children[idx], word, i+1)
}
