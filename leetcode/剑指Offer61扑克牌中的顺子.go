package leetcode

func isStraight(nums []int) bool {
	s := make(map[int]bool)
	mx, mn := 0, 14
	for _, v := range nums {
		if v == 0 {
			continue
		}
		mx = max(mx, v)
		mn = min(mn, v)
		if _, has := s[v]; has {
			return false
		}
		s[v] = true
	}
	return mx - mn < 5
}

func max(a, b int) int { if a > b { return a }; return b }
func min(a, b int) int { if a > b { return b }; return a }