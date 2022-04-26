package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 110
	INF int = 0x3f3f3f3f
)
var (
	n int
	a [N]int
	f [N][3]int
)

func init() {
	for i := range f {
		for j := range f[i] {
			f[i][j] = INF
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	f[0][0] = 0
	for i := 1; i <= n; i++ {
		f[i][0] = min(f[i-1][0], min(f[i-1][1], f[i-1][2])) + 1
		if a[i] == 1 || a[i] == 3 { f[i][1] = min(f[i-1][0], f[i-1][2]) }
		if a[i] > 1 { f[i][2] = min(f[i-1][0], f[i-1][1]) }
	}
	Fprintln(out, min(f[n][0], min(f[n][1], f[n][2])))
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }