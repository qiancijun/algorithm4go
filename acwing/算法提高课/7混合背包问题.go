package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010
var n, m int
var dp [N]int

type pii struct {
	kind, v, w int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	list := make([]pii, 0)
	for i := 0; i < n; i++ {
		var v, w, s int
		Fscan(in, &v, &w, &s)
		if s == -1 {
			list = append(list, pii{-1, v, w})
		} else if s == 0 {
			list = append(list, pii{0, v, w})
		} else {
			for k := 1; k <= s; k *= 2 {
				s -= k
				list = append(list, pii{-1, k * v, k * w})
			}
			if s > 0 {
				list = append(list, pii{-1, s * v, s * w})
			}
		}
	}

	for _, v := range list {
		if v.kind == -1 {
			// 01 背包
			for j := m; j >= v.v; j-- {
				dp[j] = max(dp[j], dp[j-v.v] + v.w)
			}
		} else {
			// 完全背包
			for j := v.v; j <= m; j++ {
				dp[j] = max(dp[j], dp[j-v.v] + v.w)
			}
		}
	}
	Fprintln(out, dp[m])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }