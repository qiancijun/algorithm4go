package main

import (
	"bufio"
	. "fmt"
	"os"
	"io"
)

const N int = 110

var dp[N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	Fscan(in, &T)
	for ; T > 0; T-- {
		solve(in, out)
	}

}

func solve(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &dp[i][j])
		}
	}

	for i := 1; i < m; i++ {
		dp[0][i] += dp[0][i-1]
	}
	for i := 1; i < n; i++ {
		dp[i][0] += dp[i-1][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j]+dp[i][j], dp[i][j-1]+dp[i][j])
		}
	}
	Fprintln(out, dp[n-1][m-1])
}

func max(a, b int) int { if a > b { return a }; return b }