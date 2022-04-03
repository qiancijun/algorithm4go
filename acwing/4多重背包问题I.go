package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 110
var n, m int
var dp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		var v, w, s int
		Fscan(in, &v, &w, &s)
		for j := m; j >= 0; j-- {
			for k := 1; k <= s && k * v <= j; k++ {
				dp[j] = max(dp[j], dp[j-k*v]+k*w)
			}
		}
	}
	Fprintln(out, dp[m])

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }