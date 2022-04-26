package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e5 + 10
	INF int = 0x3f3f3f3f
)

var (
	n int
	w [N]int
	f [N][3]int
)

func init() {
	f[0][0], f[0][1] = -INF, -INF
	f[0][2] = 0
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &w[i])
	}

	for i := 1; i <= n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2] - w[i])
		f[i][1] = f[i-1][0] + w[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}

	Fprintln(out, max(f[n][1], f[n][2]))

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }