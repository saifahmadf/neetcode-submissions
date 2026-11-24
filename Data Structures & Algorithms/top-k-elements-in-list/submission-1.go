func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	var temp [][]int
	for k, v := range m {
		temp = append(temp, []int{k, v})
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i][1] > temp[j][1]
	})

	var result []int
	for i :=0; i < k; i++ {
		result = append(result, temp[i][0])
	}
	fmt.Println("result: ", result)


	return result
}
