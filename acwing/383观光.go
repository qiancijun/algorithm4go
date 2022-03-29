//go:build ignore

package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"os"
)

const INF int64 = 1e18

var T int

type hp []edge

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].w < h[j].w }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(edge)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

type edge struct {
	to int
	w  int64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &T)
	for ; T > 0; T-- {
		var n, m int
		Fscan(in, &n, &m)
		g := make([][]edge, n+1)
		for ; m > 0; m-- {
			var u, v int
			var w int64
			Fscan(in, &u, &v, &w)
			g[u] = append(g[u], edge{v, w})
		}
		var s, f int
		Fscan(in, &s, &f)

		dist := make([]int64, n+1)
		for i := range dist {
			dist[i] = INF
		}
		dist[s] = 0
		dist2 := make([]int64, n+1)
		for i := range dist {
			dist2[i] = INF
		}
		h := hp{{s, 0}}
		for len(h) > 0 {
			cur := Pop(&h).(edge)
			v, d := cur.to, cur.w
			if dist2[v] < d {
				continue
			}
			for _, e := range g[v] {
				w := e.to
				newDis := d + e.w
				if newDis < dist[w] {
					h.Push(edge{w, newDis})
					dist[w], newDis = newDis, dist[w]
				}
				if dist[w] < newDis && newDis < dist2[w] {
					h.Push(edge{w, newDis})
					dist2[w] = newDis
				}
			}
		}
		ans := 1
		if dist2[f] + 1 == dist[f] {
			ans ++ 
		}
		Fp
	}
}
