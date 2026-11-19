func characterReplacement(s string, k int) int {
	l, result := 0, 0
	m := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		maxV := mapMaxVal(m)
		if r-l+1-maxV <= k {
			if r-l+1 > result {
				result = r - l + 1
			}
		} else {
			m[s[l]]--
			l++
		}
	}
	return result
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
