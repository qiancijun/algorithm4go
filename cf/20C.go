//go:build ignore

package main

import (
	"bufio"
	. "fmt"
	"os"
	."container/heap"
)

const INF int64 = 1e18

type edge struct {
	to int
	w  int64
}

type hp []edge

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].w < h[j].w }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(edge)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

var n, m int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	g := make([][]edge, n+1)
	for ; m > 0; m-- {
		var u, v int
		var w int64
		Fscan(in, &u, &v, &w)
		g[u] = append(g[u], edge{v, w})
		g[v] = append(g[v], edge{u, w})
	}

	dist := make([]int64, n+1)
	for i := range dist {
		dist[i] = INF
	}
	dist[1] = 0
	fa := make([]int, n+1)
	h := hp{{1, 0}}

	for len(h) > 0 {
		cur := Pop(&h).(edge)
		v, d := cur.to, cur.w
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newDis := d + e.w; newDis < dist[w] {
				dist[w] = newDis
				fa[w] = v
				Push(&h, edge{w, newDis})
			}
		}
	}

	if dist[n] == INF {
		Fprintln(out, -1)
		return
	}
	path := []int{}
	for x := n; x > 0; x = fa[x] {
		path = append(path, x)
	}
	for i := len(path) - 1; i >= 0; i-- {
		Fprint(out, path[i], " ")
	}

}
