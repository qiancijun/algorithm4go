package main

import (
	"bufio"
	. "fmt"
	"os"
)

var n, m int
var g, dp [15][20]int
var cnt[15]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			Fscan(in, &g[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][j-k] + g[i][k])
			}
		}
	}
	Fprintln(out, dp[n][m])
	for i, j := n, m; i > 0; i-- {
		for k := 0; k <= j; k++ {
			if dp[i][j] == dp[i-1][j-k] + g[i][k] {
				cnt[i] = k;
				j -= k
				break
			}
		}
	}
	for i := 1; i <= n; i++ {
		Fprintln(out, i, cnt[i])
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }