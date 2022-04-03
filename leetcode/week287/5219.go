package week287

import (
	"sort"
)

func maximumCandies(candies []int, k int64) int {
	var sum int64 = 0
	mx := 0
	for _, v := range candies {
		sum += int64(v)
		mx = max(mx, v)
	}
	if sum < k {
		return 0
	}
	n := len(candies)
	sort.Ints(candies)

	var check func(int) bool
	check = func(m int) bool {
		var ret int64 = 0
		for i := 0; i < n; i++ {
			cnt := candies[i] / m
			ret += int64(cnt)
		}
		return ret >= k
	}

	l, r := 0, mx
	for l < r {
		m := (l + r + 1) >> 1
		// tmp := candies[m]
		if check(m) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}

func max(a, b int) int { if a < b { return b }; return a }