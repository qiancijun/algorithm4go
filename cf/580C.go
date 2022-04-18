package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1e5 + 10

var (
	n, m int
	he, ne, e, w, dis [2*N]int
	leaf [N]bool
	idx = 0
)

func init() {
	for i := range he {
		he[i] = -1
	}
}

func add(a, b int) {
	e[idx], ne[idx], he[a] = b, he[a], idx; idx++
}

func dfs(u, fa int) {
	if ne[he[u]] == -1 {
		// 叶节点
		leaf[u] = true
	}
	dis[u] = w[u]
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa { continue }
		if w[u] >= 1 && w[j] == 1 {
			w[j] += w[u]
		}
		dis[j] = max(dis[j], w[u])
		
		dfs(j, u)
	}
}

func dfs2(u, fa int) {
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa { continue }
		w[j] = max(w[u], w[j])
		dfs2(j, u)
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ { Fscan(in, &w[i]) }
	for i := 1; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)
		add(u, v)
		add(v, u)
	}
	dfs(1, -1)
	dfs2(1, -1)
	ans := 0
	leaf[1] = false
	for i := 1; i <= n; i++ {
		if leaf[i] && w[i] <= m {
			ans++
		}
	}
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }