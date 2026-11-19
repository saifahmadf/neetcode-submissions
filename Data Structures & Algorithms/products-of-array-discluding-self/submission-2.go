func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	temp := 1
	for k, v := range nums {
		temp = temp * v
		prefix[k] = temp
	}

	temp = 1
	for i := len(nums)-1 ; i >= 0; i-- {
		temp = temp * nums[i]
		postfix[i] = temp
	}

	for k,_ := range nums {
		if k == 0 {
			result[k] = postfix[k+1]
		} else if k == len(nums) - 1 {
			result[k] = prefix[k-1]
		} else {
			result[k] = postfix[k+1] * prefix[k-1]
		}
	
	}
	return result
}
