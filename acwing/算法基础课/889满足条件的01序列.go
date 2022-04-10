package main

import (
	"bufio"
	. "fmt"
	"os"
)

const MOD = 1e9 + 7

func catalan(n int) int {
	if n == 1 { return 1 }
	if n == 2 { return 1 }
	res := 0
	for i := 1; i <= n - 1; i++ {
		res += catalan(i) * catalan(n - i) % MOD
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	Fprintln(out, catalan(2 * n) % MOD)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }