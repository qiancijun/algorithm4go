package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 310

var (
	n int
	s [N]int
	dp [N][N]int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var x int
		Fscan(in, &x)
		s[i] = s[i-1] + x
	}

	for i := 2; i <= n; i++ {
		for l := 1; l + i - 1 <= n; l++ {
			r := l + i - 1
			dp[l][r] = 0x3f3f3f3f
			for k := l; k < r; k++ {
				dp[l][r] = min(dp[l][r], dp[l][k] + dp[k+1][r] + s[r] - s[l-1])
			}
		}
	}
	Fprintln(out, dp[1][n])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }