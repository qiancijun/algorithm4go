package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N   int = 1e4 + 10
	INF int = 1 << 4
)

var (
	n, m int
	ans  int = INF
	dis [N]int
)

func dfs(v, pos, depth int) bool {
	if pos > depth {
		return false
	}
	if v < 0 {
		return false
	}
	if v == m {
		ans = min(ans, pos)
		return true
	}
	// red button *2
	return dfs(v*2, pos+1, depth) || dfs(v-1, pos+1, depth)
	// blue button -1
}

func bfs() {
	q := make([]int, 0)
	q = append(q, n)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if dis[cur-1] == 0 && cur - 1 > 0 {
			dis[cur-1] = dis[cur] + 1
			q = append(q, cur - 1)
		}
		if cur * 2 < N && dis[cur*2] == 0 {
			dis[cur*2] = dis[cur] + 1
			q = append(q, cur * 2)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	bfs()
	Fprintln(out, dis[m])

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	}
	return a
}
func lcm(a, b int) int { return a * b / gcd(a, b) }
