package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010

var (
	n, m int
	g [N][N]byte
)

type pair struct { x, y int }

func bfs(x, y int) {
	q := make([]pair, 0)
	q = append(q, pair{x, y})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		px, py := cur.x, cur.y
		if px < 0 || py < 0 || px >= n || py >= m { continue }
		if g[px][py] != 'W' { continue }
		g[px][py] = '-'
		for i := px - 1; i <= px + 1; i++ {
			for j := py - 1; j <= py + 1; j++ {
				if i == px && j == py { continue }
				q = append(q, pair{i, j})
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	ans := 0
	for i := 0; i < n; i++ {
		var line string
		Fscan(in, &line)
		for j := 0; j < m; j++ {
			g[i][j] = line[j]
		}
	}
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		Fprint(out, g[i][j])
	// 	}
	// 	Fprintln(out)
	// }
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == 'W' {
				bfs(i, j)
				ans++
			}
		}
	}
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }