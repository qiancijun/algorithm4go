package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 110

var (
	n int
	g [N]string
	st [N][N]bool
	ax, ay, bx ,by int
)

func dfs(x, y int) bool {
	if x < 0 || y < 0 || x >= n || y >= n || g[x][y] == '#' || st[x][y] { return false }
	if x == bx && y == by { return true }
	st[x][y] = true
	return dfs(x + 1, y) || dfs(x - 1, y) || dfs(x, y + 1) || dfs(x, y - 1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	Fscan(in, &t)
	for ; t > 0; t-- {
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &g[i])
		}

		for i := range st {
			for j := range st[i] {
				st[i][j] = false
			}
		}
		Fscan(in, &ax, &ay, &bx, &by)
		if dfs(ax, ay) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }