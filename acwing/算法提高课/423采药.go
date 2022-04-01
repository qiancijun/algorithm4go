package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010
var t, m int
var w, v [N]int
var dp [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &t, &m)
	for i := 1; i <= m; i++ {
		var a, b int
		Fscan(in, &a, &b)
		w[i], v[i] = a, b
	}

	for i := 1; i <= m; i++ {
		for j := 0; j <= t; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	Fprintln(out, dp[m][t])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }