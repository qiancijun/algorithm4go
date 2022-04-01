package leetcode

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum & 1 == 1 {
		return false
	}

	sum /= 2
	dp := make([]int, sum+1)
	for _, v := range nums {
		for j := sum; j >= v; j-- {
			dp[j] = max(dp[j], dp[j-v]+v)
		}
	}
	return dp[sum] == sum
}

func max(a, b int) int { if a > b { return a }; return b }