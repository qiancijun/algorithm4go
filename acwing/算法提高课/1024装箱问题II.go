package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 20010
var v, n int
var dp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &v, &n)
	for i := 1; i <= n; i++ {
		var a int
		Fscan(in, &a)
		for j := v; j >= a; j-- {
			dp[j] = max(dp[j], dp[j-a] + a)
		}
	}
	Fprintln(out, v - dp[v])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }