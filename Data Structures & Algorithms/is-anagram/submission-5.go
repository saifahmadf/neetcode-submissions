func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sm := make(map[byte]int)
	tm := make(map[byte]int)

	for indx, _ := range s {
		sm[s[indx]] += 1
		tm[t[indx]] += 1
	}

	for k, v := range sm {
		if _, found := tm[k]; !found || tm[k] != v {
			return false
		}
	}

	for k, v := range tm {
		if _, found := sm[k]; !found  || sm[k] != v {
			return false
		}
	}

	
	return true
}
