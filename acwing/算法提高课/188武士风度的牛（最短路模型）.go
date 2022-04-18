package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 155

var (
	n, m int
	g [N]string
	st [N][N]int
	dx, dy = [8]int{1, 2, 2, 1, -1, -2, -2, -1}, [8]int{2, 1, -1, -2, -2, -1, 1, 2}
)

type pair struct { x, y int }

func bfs(begin, end pair) {
	q := make([]pair, 0)
	q = append(q, begin)
	// st[begin.x][begin.y] = 1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		x, y := cur.x, cur.y
		if x == end.x && y == end.y { 
			return
		}
		for i := 0; i < 8; i++ {
			px, py := x + dx[i], y + dy[i]
			if px < 0 || py < 0 || px >= m || py >= n || g[px][py] == '*' || st[px][py] != 0 { continue }
			st[px][py] = st[x][y] + 1
			q = append(q, pair{px, py})
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		Fscan(in, &g[i])
	}

	var begin, end pair
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 'K' {
				begin = pair{i, j}
			}
			if g[i][j] == 'H' {
				end = pair{i, j}
			}
		}
	}

	bfs(begin, end)
	Fprintln(out, st[end.x][end.y])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }