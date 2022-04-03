package main

import (
	"bufio"
	. "fmt"
	"os"
)

var N, V, M int
var dp [110][110]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &N, &V, &M)
	for i := 0; i < N; i++ {
		var v, m, w int
		Fscan(in, &v, &m, &w)
		for j := V; j >= v; j-- {
			for k := M; k >= m; k-- {
				dp[j][k] = max(dp[j][k], dp[j-v][k-m] + w)
			}
		}
	}
	Fprintln(out, dp[V][M])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }