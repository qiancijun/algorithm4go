package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1e5 + 10

var (
	n, k int
	st [N]int
)

func bfs() {
	for i := range st {
		st[i] = -1
	}
	q := make([]int, 0)
	q = append(q, n)
	st[n] = 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == k {
			return
		}
		if cur - 1 >= 0 && st[cur-1] == -1 {
			st[cur-1] = st[cur] + 1
			q = append(q, cur - 1)
		}
		if cur + 1 <= k && st[cur+1] == -1 {
			st[cur+1] = st[cur] + 1
			q = append(q, cur + 1)
		}
		if cur * 2 < N && st[cur*2] == -1 {
			st[cur*2] = st[cur] + 1
			q = append(q, cur * 2)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &k)
	bfs()
	Fprintln(out, st[k])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }