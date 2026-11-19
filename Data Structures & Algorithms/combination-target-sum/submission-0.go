func combinationSum(nums []int, target int) [][]int {
    result := [][]int{}

	var dfs func (i int, cur []int, total int)

	dfs = func(i int, cur []int, total int) {
		if total == target {
			cur_copy := make([]int, len(cur))
			copy(cur_copy, cur)
			result = append(result, cur_copy)
			return
		}
		if total > target || i >= len(nums) {
			return 
		}

		cur = append(cur, nums[i])
		dfs(i, cur, total + nums[i])
		cur = cur[:len(cur)-1]
		dfs(i+1, cur, total)
	}

	c := []int{}
	dfs(0, c, 0)
	return result
}
