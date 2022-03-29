package week284

import (
	"container/heap"
	"math"
)

type edge struct {
	to, w int
}

type hp []edge

func (h hp) Len() int              { return len(h) }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool    { return h[i].w < h[j].w || h[i].to < h[j].to }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(edge)) }

func dijkstra(g [][]edge, n, st int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64 / 3
	}
	dist[st] = 0
	h := hp{{st, 0}}
	for len(h) > 0 {
		cur := heap.Pop(&h).(edge)
		v, d := cur.to, cur.w
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newDis := d + e.w; newDis < dist[w] {
				dist[w] = newDis
				heap.Push(&h, edge{w, newDis})
			}
		}
	}
	return dist
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	g := make([][]edge, n)
	rg := make([][]edge, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g[u] = append(g[u], edge{v, w})
		rg[v] = append(rg[v], edge{u, w})
	}
	
	d1 := dijkstra(g, n, src1)
	d2 := dijkstra(g, n, src2)
	d3 := dijkstra(rg, n, dest)

	var ans = math.MaxInt64 / 3
	for x := 0; x < n; x++ {
		tmp := d1[x] + d2[x] + d3[x]
		ans = min(ans, tmp)
	}
	if ans < math.MaxInt64 / 3 {
		return int64(ans)
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}