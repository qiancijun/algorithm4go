package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010

var n, m int
var dp [N][N]int
var v, w [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		Fscan(in, &v[i], &w[i])
	}

	for i := n; i >= 1; i-- {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i+1][j]
			if j >= v[i] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-v[i]] + w[i])
			}
		}
	}

	for i, j := 1, m; i <= n; i++ {
		if j >= v[i] && dp[i][j] == dp[i+1][j-v[i]] + w[i] {
			Fprint(out, i, " ")
			j -= v[i]
		}
	}
	
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }