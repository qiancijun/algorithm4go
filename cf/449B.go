//go:build ignore

package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"os"
)

const INF int64 = 1 << 32

var n, m, k int

type edge struct {
	to int
	w  int64
}

type hp []edge

func (h hp) Len() int              { return len(h) }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool    { return h[i].w < h[j].w || h[i].to < h[j].to }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(edge)) }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m, &k)
	g := make([][]edge, n+1)

	var u, v int
	var w int64
	for ; m > 0; m-- {
		Fscan(in, &u, &v, &w)
		g[u] = append(g[u], edge{v, w})
		g[v] = append(g[v], edge{u, w})
	}

	var s int
	var y int64
	si := make([]int, k+1)
	yi := make([]int64, k+1)
	for i := 1; i <= k; i++ {
		Fscan(in, &s, &y)
		si[i], yi[i] = s, y
		g[1] = append(g[1], edge{s, y})
		g[s] = append(g[s], edge{1, y})
	}

	dist := make([]int64, n+1)
	for i := range dist {
		dist[i] = INF
	}
	// dist[1] = 0
	// for i := range dist {
	// 	Fprint(out, dist[i], " ")
	// }
	// Fprintln(out)

	// vi := make([]bool, n+1)
	cnt := make([]int, n+1)

	h := hp{{1, 0}}
	for len(h) > 0 {
		cur := Pop(&h).(edge)
		v, d := cur.to, cur.w
		if dist[v] < d {
			continue
		}
		// if vi[v] {
		// 	continue
		// }
		// vi[v] = true
		for _, e := range g[v] {
			w := e.to
			newDis := d + e.w
			if newDis < dist[w] {
				dist[w] = newDis
				cnt[w] = 1
				Push(&h, edge{e.to, newDis})
			} else if newDis == dist[w] {
				cnt[w]++
			}

		}
	}

	// for i := range dist {
	// 	Fprint(out, dist[i], " ")
	// }
	// Fprintln(out)

	ans := 0
	for i := 1; i <= k; i++ {
		if dist[si[i]] < yi[i] {
			ans++
		} else if dist[si[i]] == yi[i] && cnt[si[i]] > 1 {
			ans++
			cnt[si[i]]--
		}
	}
	Fprintln(out, ans)
}
