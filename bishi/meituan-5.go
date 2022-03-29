package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 510
var w, din, ans [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	q := make([]int, 0)
	g := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		var l, c int
		Fscan(in, &l, &c)
		w[i] = l
		if c == 0 {
			// 入度为0
			ans[i] = l
			q = append(q, i)
		} else {
			for j := 1; j <= c; j++ {
				var pre int
				Fscan(in, &pre)
				din[i]++
				g[pre] = append(g[pre], i)
			}
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, v := range g[cur] {
			din[v]--
			ans[v] = max(ans[v], w[v] + w[cur])
			if din[v] == 0 {
				w[v] = ans[v]
				q = append(q, v)
			}
		}
	}
	for i := 1; i <= n; i++ {
		Fprintln(out, ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}