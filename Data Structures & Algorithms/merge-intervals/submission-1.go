func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
    newInterval := [][]int{}
	newInterval = append(newInterval, intervals[0])

	for i := 1; i < len(intervals); i++{
		start, end := intervals[i][0], intervals[i][1]
		lastEnd := newInterval[len(newInterval)-1][1]

		if lastEnd < start {
			newInterval = append(newInterval, []int{start, end})
		} else {
			newInterval[len(newInterval)-1][1] = max(lastEnd, end)
		}
	}	
	return newInterval
}
