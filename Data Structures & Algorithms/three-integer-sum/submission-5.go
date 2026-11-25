func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i:=0; i < len(nums); i++ {
		l, r := i+1, len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		for l < r {
			s := nums[i] + nums[l] + nums[r]

			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l<r && nums[l] == nums[l-1] {
					l++
				}
				for l<r && nums[r] == nums[r+1] {
					r--
				}
			} else if s > 0 {
				r--
			} else {
				l++
			}
		} 
	}
	return res
}
