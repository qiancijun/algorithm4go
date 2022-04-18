package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e5 + 10
)

var (
	f [N][2]int
	w [N]int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		Fscan(in, &n)
		for i := 1; i <= n; i++ {
			Fscan(in, &w[i])
		}

		for i := 1; i <= n; i++ {
			f[i][0] = max(f[i-1][0], f[i-1][1])
			f[i][1] = f[i-1][0] + w[i]
		}
		Fprintln(out, max(f[n][1], f[n][0]))
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }