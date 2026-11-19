func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	list := make([]int, len(nums))
	for i := range list {
		list[i] = 1
	}

	for i := len(nums); i >= 0; i-- {
		for j := i+1; j < len(nums); j++ {
			if nums[i] < nums[j] {
					if 1 + list[j] > list[i] {
					list[i] = 1 + list[j]
				}
			}
		}
	}

	maxL := 1
	for _, v := range list{
		if v > maxL {
			maxL = v
		}
	}
	return maxL
}
