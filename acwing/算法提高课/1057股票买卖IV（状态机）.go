package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e5 + 10
	INF = 0x3f3f3f3f
)

var (
	n, k int
	w [N]int
	f [N][110][2] int
)

func init() {
	for i := range f {
		for j := 0; j < 110; j++ {
			f[i][j][1] = -INF
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		Fscan(in, &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			f[i][j][0] = max(f[i-1][j][0], f[i-1][j][1] + w[i])
			f[i][j][1] = max(f[i-1][j][1], f[i-1][j-1][0] - w[i])
		}
	}
	ans := 0
	for i := 0; i <= k; i++ {
		ans = max(ans, f[n][i][0])
	}
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }