package leetcode

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	s := make([]string, 0)
	for _, v := range nums {
		t := strconv.Itoa(v)
		s = append(s, t)
	}
	sort.Slice(s, func(i, j int) bool { return s[i] + s[j] < s[j] + s[i] })
	ans := ""
	for _, v := range s {
		ans += v
	}
	return ans
}