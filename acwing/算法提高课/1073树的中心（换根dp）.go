package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e4 + 10
	INF int = 0x3f3f3f3f
)

var n int
var he, e, ne, w, d1, d2, up, p [2*N]int
var idx = 0

func add(a, b, c int) {
	e[idx] = b 
	w[idx] = c
	ne[idx] = he[a]
	he[a] = idx
	idx++
}

// 从上往下走
func dfs_down(u, fa int) int {
	d1[u], d2[u] = -INF, -INF
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa { continue }
		// 往下走先算子节点的
		d := dfs_down(j, u) + w[i]
		if d >= d1[u] {
			d2[u] = d1[u]
			d1[u] = d
			p[u] = j
		} else if d > d2[u] {
			d2[u] = d
		}
	}
	if d1[u] == -INF {
		d1[u], d2[u] = 0, 0
	}
	return d1[u]
}

// 从下往上走
func dfs_up(u, fa int) {
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa { continue }
		// 往上走先算自己的
		if p[u] == j {
			up[j] = max(up[u], d2[u]) + w[i]
		} else {
			up[j] = max(up[u], d1[u]) + w[i]
		}
		dfs_up(j, u)
	}
}

func init() {
	for i := range he {
		he[i] = -1
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)

	// 建图
	for i := 0; i < n - 1; i++ {
		var u, v, w int
		Fscan(in, &u, &v, &w)
		add(u, v, w)
		add(v, u, w)
	}
	dfs_down(1, -1)
	dfs_up(1, -1)

	res := INF
	for i := 1; i <= n; i++ {
		res = min(res, max(up[i], d1[i]))
	}
	Fprintln(out, res)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }