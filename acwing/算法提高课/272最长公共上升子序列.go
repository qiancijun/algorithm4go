package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 3010

var n int
var a, b [N]int
var dp [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)

	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
	}

	for i := 1; i <= n; i++ {
		mx := 1 // j 从1到n的过程中最优的答案
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if a[i] == b[j] {
				dp[i][j] = max(dp[i][j], mx)
			}
			if b[j] < a[i] {
				mx = max(mx, dp[i-1][j]+1)
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dp[n][i])
	}
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }