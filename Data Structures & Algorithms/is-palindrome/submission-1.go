func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var s1 string
	for _,v := range s{
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9')  {
			s1 += string(v)
		}
	}

	j := len(s1) - 1
	for i:=0; i<len(s1)/2; i++ {
		if s1[i] != s1[j] {
			return false
		}
		j--
	}

	return true
}
