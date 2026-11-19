func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[string]int)
	tm := make(map[string]int)

	for k, _ := range s{
		sc := string(s[k])
		if _, ok := sm[sc]; ok {
			sm[sc] += 1
 		} else {
			sm[sc] = 1
		}

		tc := string(t[k])
		if _, ok := tm[tc]; ok {
			tm[tc] += 1
 		} else {
			tm[tc] = 1
		}
	}

	for k, v := range sm{
		if _, ok := tm[k]; ok {
			if tm[k] != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
