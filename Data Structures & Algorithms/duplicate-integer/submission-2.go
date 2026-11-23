func hasDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, found := m[num]; !found {
			m[num] = true
		} else {
			return true
		}
	}
	return false
}
