func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	islands := 0

	var bfs func(r, c int)
	bfs = func(r, c int) {
		queue := list.New()
		queue.PushBack([]int{r, c})
		grid[r][c] = '#'

		for queue.Len() > 0 {
			front := queue.Front()
			node := front.Value.([]int)
			row, col := node[0], node[1]
			queue.Remove(front)
			directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, d := range directions {
					r := row + d[0]
					c := col + d[1]
					if r >=0 && r < rows && c>=0 && c < cols && grid[r][c] != '#' && grid[r][c] == '1' {
						queue.PushBack([]int{r, c})
						grid[r][c] = '#'
					}
			}
		}

	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '#' && grid[r][c] == '1' {
				bfs(r, c)
				islands++	
			}
		}
	}

	return islands
}
