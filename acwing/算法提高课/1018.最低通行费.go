package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 110

var n int
var dp [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &dp[i][j])
		}
	}

	for i := 1; i < n; i++ {
		dp[0][i] += dp[0][i-1]
		dp[i][0] += dp[i-1][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			val := dp[i][j]
			dp[i][j] = min(dp[i-1][j]+val, dp[i][j-1]+val)
		}
	}
	Fprintln(out, dp[n-1][n-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
