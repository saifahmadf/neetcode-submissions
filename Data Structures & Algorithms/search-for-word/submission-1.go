func exist(board [][]byte, word string) bool {
	for r := range board {
		for c := range board[0] {
			if dfs(board, word, r, c, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, r, c, i int) bool {
	if i == len(word){
		return true
	}
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != word[i] {
		return false
	}

	temp := board[r][c]
	board[r][c] = '#'
	
	ok := dfs(board, word, r+1, c, i+1) || dfs(board, word, r-1, c, i+1) || dfs(board, word, r, c+1, i+1) || dfs(board, word, r, c-1, i+1)
	board[r][c] = temp
	return ok
}