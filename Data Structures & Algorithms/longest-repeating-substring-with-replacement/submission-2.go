func characterReplacement(s string, k int) int {
	l, res := 0, 0
	m := make(map[byte]int)

	for r, _ := range s {
		m[s[r]]++
		
		if r - l + 1 - mapMaxVal(m) <= k {
			if r - l + 1 > res {
				res = r - l + 1
			}
		} else {
			m[s[l]]--
			l++
		}
	}
	return res
}

func mapMaxVal(m map[byte]int) int {
	var m_v int
	for _, v := range m {
		if v > m_v {
			m_v = v
		}
	}
	return m_v
}
