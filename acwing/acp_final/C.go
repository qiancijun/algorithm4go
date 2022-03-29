package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 2e5 + 10
	M     = N * 2
)

var ne, e, w [M]int
var he, down, up [N]int
var n, idx = 0, 0

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = he[a]
	he[a] = idx
	idx++
}

func dfs_down(u, from int) {
	for i := he[u]; i != -1; i = ne[i] {
		if i == (from ^ 1) {
			continue
		}
		j := e[i]
		dfs_down(j, i)
		down[u] += down[j] + w[i]
	}
}

func dfs_up(u, from int) {
	if from != -1 {
		fa := e[from^1]
		up[u] = up[fa] + down[fa] - down[u] - w[from] + w[from^1]
	}
	for i := he[u]; i != -1; i = ne[i] {
		if i == (from ^ 1) {
			continue
		}
		j := e[i]
		dfs_up(j, i)
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
	for i := 0; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)
		add(u, v, 0)
		add(v, u, 1)
	}

	dfs_down(1, -1)
	dfs_up(1, -1)

	res := N
	for i := 1; i <= n; i++ {
		res = min(res, down[i]+up[i])
	}
	Fprintln(out, res)

	for i := 1; i <= n; i++ {
		if down[i]+up[i] == res {
			Fprint(out, i, " ")
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
