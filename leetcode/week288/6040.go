package week288

import "sort"

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	sort.Ints(flowers)
	n := len(flowers)

	// 如果最少的花园都满足要求了，那么可以直接出答案
	if flowers[0] >= target { return int64(n * full) }

	need := int(newFlowers) - target * n
	for i, flower := range flowers {
		flowers[i] = min(flower, target)
		need += flowers[i]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}