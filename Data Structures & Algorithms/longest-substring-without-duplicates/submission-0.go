func lengthOfLongestSubstring(s string) int {
	l, res := 0, 0
	m := make(map[byte]bool)

	for key, _ := range s {
		for m[s[key]]{
			delete(m, s[l])
			l++
		}
		m[s[key]] = true
		if key - l + 1 > res {
			res = key - l + 1
		}
	}

	return res
}
