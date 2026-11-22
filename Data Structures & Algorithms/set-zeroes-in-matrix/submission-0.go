func setZeroes(matrix [][]int) {
    R, C := len(matrix), len(matrix[0])
	rows := make([]bool, R)
	cols := make([]bool, C)

	for r := 0; r < R; r++{
		for c := 0; c < C; c++{
			if matrix[r][c] == 0 {
				rows[r] = true
				cols[c] = true
			}
		}
	}

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if rows[r] || cols[c] {
				matrix[r][c] = 0
			}
		}
	}
}
