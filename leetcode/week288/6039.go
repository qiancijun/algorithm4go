package week288

import "container/heap"


const MOD int = 1e9 + 7
type hp []int

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i] < h[j] }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(int)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func maximumProduct(nums []int, k int) int {
	pq := hp{}
	for _, v := range nums {
		heap.Push(&pq, v)
	}
	for k > 0 {
		cur := heap.Pop(&pq).(int)
		cur++
		heap.Push(&pq, cur)
		k--
	}

	ans := 1
	for len(pq) > 0 {
		cur := heap.Pop(&pq).(int)
		ans = (ans * cur) % MOD
	}
	return ans
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