package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e5 + 10
)

var he, ne, e, d [2*N]int // 无向图，空间开两倍
var idx int = 0
var n int
var ans int = 0x3f3f3f

func add(a, b int) {
	e[idx] = b
	ne[idx] = he[a]
	he[a] = idx
	idx++
}

func dfs(u, fa int) {
	d[u] = 1
	mx := 0
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa {
			continue
		}
		dfs(j, u)
		d[u] += d[j]
		mx = max(mx, d[j]) // 所有子节点最大连通块个数
	}
	mx = max(mx, n - d[u]) // 如果该节点是子节点，那么其父节点的连通块数量可以直接计算而来
	ans = min(ans, mx)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := range he {
		he[i] = -1
	}

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)
		add(u, v)
		add(v, u)
	}

	dfs(1, -1)
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }