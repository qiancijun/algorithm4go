package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
)

const (
	N int = 105
	M int = 1 << 10
)

var (
	n, m int
	g [N]int
	state = make([]int, 0)
	dp [2][M][M]int // 对于前 i 行，第 i - 1 行的状态是 j，第 i 行的状态是 k
	cnt [M]int
)

func check(v int) bool {
	for i := 0; i < m; i++ {
		if (v >> i & 1) == 1 && ((v >> (i+1) & 1) == 1 || (v >> (i+2) & 1) == 1) {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var s string
		Fscan(in, &s)
		for j := 0; j < m; j++ {
			if s[j] == 'H' {
				g[i] += 1 << j
			}
		}
	}

	for i := 0; i < 1 << m; i++ {
		if check(i) {
			state = append(state, i)
			cnt[i] = bits.OnesCount(uint(i))
		}
	}

	for i := 1; i <= n + 2; i++ {
		for j := 0; j < len(state); j++ {
			for k := 0; k < len(state); k++ {
				for u := 0; u < len(state); u++ {
					a, b, c := state[j], state[k], state[u]
					if (a & b) | (a & c) | (b & c) != 0 { continue }
					if (g[i-1] & a) | (g[i] & b) != 0 { continue } 
					dp[i&1][j][k] = max(dp[i&1][j][k], dp[(i-1)&1][u][j] + cnt[b])
				}
			}
		}
	}
	Fprintln(out, dp[(n+2)&1][0][0])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }