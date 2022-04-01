package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 20010
var v, n int
var w [40]int
var dp [40][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &v, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= v; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]] + w[i])
			}
		}
	}
	Fprintln(out, v - dp[n][v])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }