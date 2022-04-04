package main

import (
	"bufio"
	. "fmt"
	"os"
)

var N, V int
var v, w [110]int
var dp [110]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &N, &V)
	for i := 0; i < N; i++ {
		var s int
		Fscan(in, &s)
		for j := 1; j <= s; j++ { Fscan(in, &v[j], &w[j]) }
		for j := V; j >= 0; j-- {
			for k := 1; k <= s; k++ {
				if j - v[k] >= 0 {
					dp[j] = max(dp[j], dp[j-v[k]] + w[k])
				}
			}
		}
	}
	Fprintln(out, dp[V])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }