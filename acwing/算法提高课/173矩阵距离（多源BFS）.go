package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010

var (
	n, m int
	g [N]string
	st [N][N]int
	dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
	q = make([]pair, 0)
)

type pair struct { x, y int }

func bfs() {
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			px, py := cur.x + dx[i], cur.y + dy[i]
			if px < 0 || py < 0 || px >= n || py >= m || st[px][py] != -1 || g[px][py] == '1' { continue }
			st[px][py] = st[cur.x][cur.y] + 1
			q = append(q, pair{px, py})
		}
	}
}

func init() {
	for i := range st {
		for j := range st[i] {
			st[i][j] = -1
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &g[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == '1' {
				st[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}
	bfs()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fprint(out, st[i][j], " ")
		}
		Fprintln(out)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }