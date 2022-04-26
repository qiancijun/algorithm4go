package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1010
)

var (
	n, m int
	g [N][N]int
	pre [N][N]pair
	dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
)

type pair struct { x, y int }

func bfs(a, b int) {
	q := make([]pair, 0)
	q = append(q, pair{a, b})
	g[a][b] = 1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		x, y := cur.x, cur.y
		for i := 0; i < 4; i++ {
			px, py := x + dx[i], y + dy[i]
			if px < 0 || py < 0 || px >= n || py >= n || g[px][py] == 1 { continue }
			q = append(q, pair{px, py})
			pre[px][py] = cur
			g[px][py] = 1
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &g[i][j])
		}
	}
	bfs(n - 1, n - 1)
	t := pair{0, 0}
	for {
		if t.x == n - 1 && t.y == n - 1 {
			Fprintln(out, t.x, t.y)
			break
		} else {
			Fprintln(out, t.x, t.y)
			next := pre[t.x][t.y]
			t = next
		}	
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }