package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e4 + 10
)

var n int
var he, e, ne, w, dis [N]int
var idx = 0

func add(a, b, c int) {
	e[idx] = b 
	w[idx] = c
	ne[a] = he[a]
	he[a] = idx
	idx++
}

func dfs1(u, fa int) {
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa { continue }
		dfs1(j, i)
		dis[u] 
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := range he {
		he[i] = -1
	}

	Fscan(in, &n)

	// 建图
	for i := 0; i < n; i++ {
		var u, v, w int
		Fscan(in, &u, &v, &w)
		add(u, v, w)
		add(v, u, w)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }