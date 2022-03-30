package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	res, mn := 0, prices[0]
	for _, v := range prices {
		mn = min(mn, v)
		res = max(res, v - mn)
	}
	return res
}

func max (a, b int) int { if a > b { return a }; return b }
func min (a, b int) int { if a < b { return a }; return b }