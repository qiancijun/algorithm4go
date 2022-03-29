package main

import (
	"bufio"
	"os"
	."fmt"
)

var n, m, k int
var a [5010]int64
var pre [5010]int64 // 前缀和
var mx int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m, &k)

	for i := 1; i <= n; i ++ {
		Fscan(in, &a[i])
		pre[i] = pre[i - 1] + a[i]
	}

	mx = 0
	dfs(0, 0, 0)

	Fprintln(out, mx)
}

func dfs(begin, pos int, sum int64) {
	if pos == k {
		mx = max(mx, sum)
		return
	}
	for i := begin + m; i <= n; i ++ {
		cur := pre[i] - pre[i - m]
		if 
		dfs(i, pos + 1, sum + cur)
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}