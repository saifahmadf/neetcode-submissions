func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	
	for _, str := range strs {
		var count [26]int
		for _, s := range str {
			count[s - 'a']++
		}
		m[count] = append(m[count], str)
	}

	var result [][]string

	for _, v := range m {
		result = append(result, v)
	}
	return result
}
