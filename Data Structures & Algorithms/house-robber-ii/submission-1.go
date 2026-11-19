func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
    r1, r2 := 0, 0
	r3, r4 := 0, 0
	num1 := nums[0:len(nums)-1]
	num2 := nums[1:len(nums)]
	for _, v := range num1 {
		temp := max(v + r1, r2)
		r1 = r2
		r2 = temp
		
	}
	for _, v := range num2 {
		temp := max(v + r3, r4)
		r3 = r4
		r4 = temp
		
	}
	if r2 > r4 {
		return r2
	}
	return r4
}
