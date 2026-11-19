func countSubstrings(s string) int {
    result := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >=0 && r < len(s) && s[l] == s[r]{
			result += 1
			l -= 1
			r += 1
		}
		l, r = i, i+1
		for l >=0 && r < len(s) && s[l] == s[r]{
			result += 1
			l -= 1
			r += 1
		}
	}
	return result
}
