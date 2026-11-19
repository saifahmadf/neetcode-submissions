func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for k, v := range nums {
		m[v] = k
	}

	for k, v := range nums {
		complement := target - v
		if i, ok := m[complement]; ok && k != i {
			return []int{k,i}
		}
	}

	return nil
}
