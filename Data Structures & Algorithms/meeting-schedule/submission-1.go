/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0{
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})
	prevEnd := intervals[0].end

	for _, interval := range intervals[1:]{
		start := interval.start
		end := interval.end
		if start < prevEnd {
			return false
		}
		prevEnd = end
	}
	return true
}
