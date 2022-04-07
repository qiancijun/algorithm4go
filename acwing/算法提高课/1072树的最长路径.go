package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e4 + 10
	M int = 2 * N
)

var he [N]int
var e, ne, w [M]int
var idx int = 0
var n int
var ans int = 0

func add(a, b, c int) {
	e[idx] = b; w[idx] = c; ne[idx] = he[a]; he[a] = idx; idx++
}

func init() {
	for i := range he {
		he[i] = -1
	}
}

func dfs(u, fa int) int {
	dis, d1, d2 := 0, 0, 0
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa {
			continue
		}
		d := dfs(j, u) + w[i]
		dis = max(dis, d)
		if d >= d1 {
			d2, d1 = d1, d
		} else if d > d2 {
			d2 = d
		}
	}
	ans = max(ans, d1 + d2)
	return dis
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 0; i < n - 1; i++ {
		var u, v, w int
		Fscan(in, &u, &v, &w)
		add(u, v, w)
		add(v, u, w)
	}

	dfs(1, -1)
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }