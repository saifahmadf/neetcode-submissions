func longestCommonSubsequence(text1 string, text2 string) int {
   r, c := len(text1), len(text2)
   dp := make([][]int, r+1)
   
   for i := range dp {
	dp[i] = make([]int, c+1)
   }

	for i := r-1; i >=0; i-- {
		for j := c-1; j >=0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

   return dp[0][0]
}
