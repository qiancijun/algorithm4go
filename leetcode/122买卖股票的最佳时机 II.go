package leetcode

func maxProfit(prices []int) int {
	res := 0
	p, v := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		for ; i < len(prices)-1 && prices[i] >= prices[i+1]; i++ {
		}
		v = prices[i]
		for ; i < len(prices)-1 && prices[i] <= prices[i+1]; i++ {
		}
		p = prices[i]
		res = res + p - v
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}