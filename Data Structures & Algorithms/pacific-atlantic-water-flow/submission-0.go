func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pac := make(map[[2]int]bool) 
    alt := make(map[[2]int]bool)
    results := make([][]int, 0)

	var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)

	dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
		coord := [2]int{r,c}
		if visit[coord] || r < 0 || c < 0 || r >= rows || c >= cols || prevHeight > heights[r][c] {
			return
		}
		visit[coord] = true

		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	for r:=0; r < rows; r++ {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols-1, alt, heights[r][cols-1])
	}
	for c:=0; c < cols; c++ {
		dfs(0, c, pac, heights[0][c])
		dfs(rows-1, c, alt, heights[rows-1][c])
	}

	for r:=0; r < rows; r++ {
		for c:=0; c < cols; c++ {
			coord := [2]int{r,c}
			if pac[coord] && alt[coord] {
				results = append(results, []int{r,c})
			}
		}	
	}
	return results
}
