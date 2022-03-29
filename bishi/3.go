package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 5e4 + 10
const INF = 0x3f3f3f3f

var n int
var a [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	ans := -INF

	var dfs func(int, int)
	dfs = func(pos, sum int) {
		if pos == n {
			if sum % 7 == 0 {
				ans = max(ans, sum)
			}
			return
		}
		// choice
		dfs(pos+1, sum+a[pos])
		// dont choice
		dfs(pos+1, sum)
	}
	dfs(0, 0)
	Fprintln(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}