func maxProduct(nums []int) int {
   maxSum, minSum := 1, 1
   result := maxInSlice(nums)
   for _, n := range nums {
		if n == 0 {
			maxSum, minSum = 1, 1
			continue
		}

		temp := maxSum
		maxSum = max(maxSum * n, minSum * n, n)
		minSum = min(temp * n, minSum * n, n)
		result = max(maxSum, result)
   }
   return result
}

func maxInSlice(nums []int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}