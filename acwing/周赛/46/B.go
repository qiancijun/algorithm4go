package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"os"
)

const (
	N int = 2e5 + 10
)
var n, k int

type pair struct {
	idx, val int
}

type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].val > h[j].val }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	val := make([]int, n)
	pq := hp{}
	for i := 0; i < n; i++ { Fscan(in, &a[i]) }
	for i := 0; i < n; i++ { Fscan(in, &b[i]) }
	for i := 0; i < n; i++ { 
		val[i] = a[i] - b[i]
		if val[i] > 0 {
			heap.Push(&pq, pair{i, val[i]})
		}
	}
	k = n - k // 可以翻 k 张
	for len(pq) > 0 && k > 0 {
		cur := heap.Pop(&pq).(pair)
		idx, _ := cur.idx, cur.val
		a[idx] = b[idx]
		k--
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i]
	}
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }