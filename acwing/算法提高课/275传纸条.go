package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 55

var n, m int
var g [N][N]int
var dp [N*2][N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			Fscan(in, &g[i][j])
		}
	}
	for k := 2; k <= n+m; k++ {
		for i1 := 1; i1 <= n; i1++ {
			for i2 := 1; i2 <= n; i2++ {
				j1, j2 := k - i1, k - i2
				if j1 >= 1 && j1 <= m && j2 >= 1 && j2 <= m {
					val := g[i1][j1]
					if i1 != i2 {
						val += g[i2][j2]
					}
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1][i2] + val)
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1-1][i2] + val)
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1][i2-1] + val)
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1-1][i2-1] + val)
				}
			}
		}
	}
	Fprintln(out, dp[n+m][n][n])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }