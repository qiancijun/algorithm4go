package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"os"
)

const (
	N int = 150010
	INF int = 0x3f3f3f3f
)

type edge struct { to, w int }
type hp []edge

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].w < h[j].w }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(edge)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

var n, m int
var g [N][]edge

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var u, v, w int
		Fscan(in, &u, &v, &w)
		g[u] = append(g[u], edge{v, w})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = INF
	}
	dist[1] = 0
	hp := hp{{1, 0}}

	for len(hp) > 0 {
		cur := heap.Pop(&hp).(edge)
		v, d := cur.to, cur.w
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newDis := d + e.w; newDis < dist[w] {
				dist[w] = newDis
				heap.Push(&hp, edge{w, newDis})
			}
		}
	}
	if dist[n] == INF {
		Fprintln(out, -1)
	} else {
		Fprintln(out, dist[n])
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }