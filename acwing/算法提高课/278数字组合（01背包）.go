package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 10010
var n, m int
var dp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)

	dp[0] = 1
	for i := 0; i < n; i++ {
		var a int
		Fscan(in, &a)
		for j := m; j >= a; j-- {
			dp[j] += dp[j-a]
		}
	}
	Fprintln(out, dp[m])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }