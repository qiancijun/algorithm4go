package week286

func maxValueOfCoins(piles [][]int, k int) int {
	dp := make([]int, k+1)

	for _, pile := range piles {
		n := len(pile)
		for i := 1; i < len(pile); i++ {
			pile[i] += pile[i-1]
		}
		for j := k; j > 0; j-- {
			for m := n-1; m >= 0; m-- {
				if j - m - 1 >= 0 {
					dp[j] = max(dp[j], dp[j-m-1]+pile[m])
				}
			}
		}
	}
	return dp[k]
}

func max(a, b int) int { if a > b { return a }; return b }