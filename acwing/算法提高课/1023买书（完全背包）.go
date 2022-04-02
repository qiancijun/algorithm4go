package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010
var dp [N]int
var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	dp[0] = 1
	list := []int{10, 20, 50, 100}
	for _, v := range list {
		for j := v; j <= n; j++ {
			dp[j] += dp[j-v]
		}
	}
	Fprintln(out, dp[n])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }