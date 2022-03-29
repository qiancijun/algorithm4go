package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 12
var n int
var g [N][N]int
var dp [N*2][N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for {
		var r, c, v int
		Fscan(in, &r, &c, &v)
		if r + c + v == 0 {
			break
		}
		g[r][c] = v
	}

	for k := 2; k <= n+n; k++ {
		for i1 := 1; i1 <= n; i1++ {
			for i2 := 1; i2 <= n; i2++ {
				j1, j2 := k - i1, k - i2
				if j1 >= 1 && j1 <= n && j2 >= 1 && j2 <= n {
					val := g[i1][j1]
					if i1 != i2 {
						val += g[i2][j2]
					}
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1-1][i2-1]+val)
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1][i2-1]+val)
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1-1][i2]+val)
					dp[k][i1][i2] = max(dp[k][i1][i2], dp[k-1][i1][i2]+val)
				}
			}
		}
	}
	Fprintln(out, dp[n+n][n][n])
}

func max(a, b int) int { if a > b { return a }; return b }
