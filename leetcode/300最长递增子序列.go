package leetcode

// dp 版本
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	return ans
}

// 贪心版本
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	g := make([]int, n)

	cnt := 0
	for i := 0; i < n; i++ {
		k := 0
		for k < cnt && g[k] < nums[i] {
			k++
		}
		g[k] = nums[i]
		if k >= cnt {
			cnt++
		}
	}
	return cnt
}

func max(a, b int) int { if a > b { return a }; return b }