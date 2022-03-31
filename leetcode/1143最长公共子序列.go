package leetcode

func longestCommonSubsequence(a string, b string) int {
	n, m := len(a), len(b)
	a = " " + a
	b = " " + b
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func max(a, b int) int { if a > b { return a }; return b }