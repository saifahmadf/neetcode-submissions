func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	
	for _, val := range nums {
		m[val]++
	}

	arr2 := [][]int{}
	for k, v := range m{
		arr2 = append(arr2, []int{k,v})
	}

	sort.Slice(arr2, func(i, j int) bool{
		return arr2[i][1] > arr2[j][1]
	})
	
	r := []int{}
	for i := 0; i < k; i++{
		r = append(r, arr2[i][0])
	}
	return r
}
