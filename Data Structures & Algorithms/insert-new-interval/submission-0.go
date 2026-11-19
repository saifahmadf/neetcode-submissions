func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	
	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			result = append(result, intervals[i])
		} else {
			newInterval = []int {
				min(newInterval[0], intervals[i][0]),
				max(newInterval[1], intervals[i][1]),
			}
		}
	}
	result = append(result, newInterval)
	return result
}
