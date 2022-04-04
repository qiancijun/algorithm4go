package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 30010
var n, m int
var dp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		w *= v
		for j := n; j >= v; j-- {
			dp[j] = max(dp[j], dp[j-v] + w)
		}
	}
	Fprintln(out, dp[n])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }