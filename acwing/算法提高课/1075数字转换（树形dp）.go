package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 50010
)

var n int
var he, ne, e, st [N]int
var idx = 0
var ans = 0

func init() {
	for i := range he {
		he[i] = -1
	}
}

func add(a, b int) {
	e[idx] = b; ne[idx] = he[a]; he[a] = idx; idx++
}

func dfs(u, fa int) int {
	dis, d1, d2 := 0, 0, 0
	for i := he[u]; i != -1; i = ne[i] {
		j := e[i]
		if j == fa { continue }
		d := dfs(j, u) + 1
		dis = max(dis, d)
		if d >= d1 {
			d2, d1 = d1, d
		} else if d > d2 {
			d2 = d
		}
	}
	ans = max(ans, d1 + d2)
	return dis
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)

	// 试除法
	// for i := 2; i <= n; i++ {
	// 	sum := 0
	// 	for j := 1; j <= i / j; j++ {
	// 		if i % j == 0 {
	// 			sum += j
	// 			if j != i / j {
	// 				sum += i / j
	// 			}
	// 		}
	// 	}
	// 	sum -= i
	// 	if sum < i {
	// 		add(sum, i)
	// 	}
	// }

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 2; j * i <= n; j++ {
			sum[i*j] += i
		}
	}

	for i := 2; i <= n; i++ {
		if sum[i] < i {
			add(sum[i], i)
		}
	}
	dfs(1, -1)
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }