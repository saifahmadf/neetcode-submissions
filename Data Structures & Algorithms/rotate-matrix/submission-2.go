func rotate(matrix [][]int)  {
    
	row, col := len(matrix), len(matrix[0])
	result := make([][]int, row)
	for i := range result {
		result[i] = make([]int, col)
	}
	
	
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[j][col-1-i] = matrix[i][j]
		}
	}
	for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            matrix[i][j] = result[i][j]
        }
    }
}
