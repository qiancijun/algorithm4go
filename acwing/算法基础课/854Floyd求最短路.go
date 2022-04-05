package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 210
	INF int = 1e9
)

var n, m, k int
var g [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m, &k)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				g[i][j] = 0
			} else {
				g[i][j] = INF
			}
		}
	}

	for ; m > 0; m-- {
		var u, v, w int
		Fscan(in, &u, &v, &w)
		g[u][v] = min(g[u][v], w)
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				g[i][j] = min(g[i][j], g[i][k] + g[k][j])
			}
		}
	}

	for ; k > 0; k-- {
		var u, v int
		Fscan(in, &u, &v)
		if g[u][v] >= INF / 2 {
			Fprintln(out, "impossible")
		} else {
			Fprintln(out, g[u][v])
		}
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }