package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 14
	MOD int = 1e8
)

var m, n int
var g [N]int
var state = make([]int, 0)
var head [1<<13][]int
var dp [N][1<<13]int

func check(v int) bool {
	for i := 0; i + 1 < m; i++ {
		if (v >> i & 1) == 1 && (v >> (i+1) & 1) == 1 { return false }
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ { 
			var t int
			Fscan(in, &t)
			g[i] += (1 - t) * (1 << j)
		}
	}
	for i := 0; i < 1 << m; i++ {
		if check(i) {
			state = append(state, i)
		}
	}

	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state); j++ {
			a, b := state[i], state[j]
			if (a & b) == 0 { head[i] = append(head[i], j) }
		}
	}

	dp[0][0] = 1
	for i := 1; i <= n + 1; i++ {
		for a := 0; a < len(state); a++ {
			if g[i] & state[a] == 0 {
				for _, b := range head[a] {
					dp[i][a] = (dp[i][a] + dp[i-1][b]) % MOD
				}
			}
		}
	}
	Fprintln(out, dp[n+1][0])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }