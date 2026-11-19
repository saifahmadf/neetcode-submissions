func missingNumber(nums []int) int {
	res := len(nums)
	for indx:=0; indx < len(nums); indx++ {
		res += indx - nums[indx]
	}
	return res
}
