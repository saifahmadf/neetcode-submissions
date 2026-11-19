func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	
	var result int
	for _, v := range nums {
		set[v] = true
	}
	
	for _, num := range nums {
		if set[num-1] {
			continue
		}
		seq := 1
		temp := num + 1
		for set[temp] {
			seq++
			temp++
		}
		if seq > result {
			result = seq
		}

	}

	return result
}
