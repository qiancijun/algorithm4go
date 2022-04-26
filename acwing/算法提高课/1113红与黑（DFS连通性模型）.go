package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 25

type pair struct { x, y int }

var (
	n, m int
	g [N]string
	st [N][N]bool
)

func dfs(pos pair) int {
	x, y := pos.x, pos.y
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for {
		Fscan(in, &n, &m)
		if n == 0 && m == 0 {
			break
		}
		var begin pair
		for i := 0; i < m; i++ {
			Fscan(in, &g[i])
			for j := 0; j < n; j++ {
				if g[i][j] == '@' {
					begin = pair{i, j}
				}
			}
		}
		cnt := dfs(begin)
		Fprintln(out, cnt)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }