func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, num := range nums {
		m[num] = index
	}

	for k, v := range nums {
		complement := target - v
		if val, found := m[complement]; found && k != val {
			return []int{k, val}
		}
	}
	return nil
}
