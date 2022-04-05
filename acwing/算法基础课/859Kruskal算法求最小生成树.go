package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

const (
	M int = 2e5 + 10
	N int = 1e5 + 10
)

type edge struct { u, v, w int}

var n, m int
var p [N]int

func find(x int) int {
	for x != p[x] {
		p[x] = p[p[x]]
		x = p[x]
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		var u, v, w int
		Fscan(in, &u, &v, &w)
		edges[i] = edge{u, v, w}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w } )
	
	for i := range p {
		p[i] = i
	}
	cnt, ans := 0, 0
	for i := 0; i < m; i++ {
		a, b, c := edges[i].u, edges[i].v, edges[i].w
		a, b = find(a), find(b)
		if a != b {
			ans += c
			p[a] = b
			cnt++
		}
	}
	if cnt < n - 1 {
		Fprintln(out, "impossible")
	} else {
		Fprintln(out, ans)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }