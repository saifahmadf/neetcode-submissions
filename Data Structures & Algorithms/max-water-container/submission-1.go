func maxArea(heights []int) int {
	var h int

	for i := 0 ; i < len(heights); i++ {
		
	}
	l, r := 0, len(heights) - 1

	for l < r { 
		tall := 0
		if heights[l] >= heights[r] {
			tall = heights[r]
		} else {
			tall = heights[l]
		}
		area := tall * (r-l)
		if h < area {
			h = area
		}
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}
	return h
}
