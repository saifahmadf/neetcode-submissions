func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, val := range strs {
		var count [26]int
		val = strings.ToLower(val)
		for _, c := range val {
			count[c - 'a']++
		}
		m[count] = append(m[count], val)
	}
	var r [][]string
	for _, v := range m{
		r = append(r, v)
	}
	return r
}
