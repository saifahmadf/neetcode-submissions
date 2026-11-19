type TrieNode struct {
	Children [26]*TrieNode
	Word string
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	root := &TrieNode{}
	for _, word := range words {
		node := root
		for _, c := range word {
			idx := c - 'a'
			if node.Children[idx] == nil {
				node.Children[idx] = &TrieNode{}
			}
			node = node.Children[idx]
		}
		node.Word = word
	}

	row, col := len(board), len(board[0])

	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		idx := board[r][c] - 'a'
		if node.Children[idx] == nil {
			return
		}
		node = node.Children[idx]
		
		if node.Word != "" {
			result = append(result, node.Word)
			node.Word = ""
		}

		ch := board[r][c]
		board[r][c] = '#'
		
		dirs := [4][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
		for i := 0; i < 4; i++ {
			nr, nc := r+dirs[i][0], c+dirs[i][1]
			if nr >= 0 && nc >= 0 && nr < row && nc < col && board[nr][nc] != '#' {
				dfs(nr, nc, node)
			}
		}
		board[r][c] = ch
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			dfs(r, c, root)
		}
	}
	return result
}
