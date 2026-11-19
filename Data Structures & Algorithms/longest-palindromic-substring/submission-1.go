
func longestPalindrome(s string) string {
    pal := ""
	result := 0
	for i:=0; i < len(s); i++ {
		l, r := i, i
		for l >=0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > result {
				result = r - l + 1
				pal = s[l:r+1]
			}
			l -=1
			r +=1
		}
		
		l,r = i, i+1
		for l >=0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > result {
				result = r - l + 1
				pal = s[l:r+1]
			}
			l -=1
			r +=1
		}
		
	}
	return pal
}
