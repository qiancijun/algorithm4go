package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

const (
	N int = 110
	M int = 10010
)

type stone struct { s, e, l int }
var dp [N][M]int

func solve(in io.Reader, out io.Writer) int {
	var n int
	Fscan(in, &n)
	stones := make([]stone, n+1)
	m := 0
	for i := 1; i <= n; i++ {
		var s, e, l int
		Fscan(in, &s, &e, &l)
		stones[i] = stone{s, e, l}
		m += s
	}
	sort.Slice(stones, func(i, j int) bool { return stones[i].s * stones[j].l < stones[i].l * stones[j].s})
	// for i := range dp {
	// 	dp[i] = -0x3f3f3f3f
	// }
	// dp[0] = 0
	for i := 1; i <= n; i++ {
		s, e, l := stones[i].s, stones[i].e, stones[i].l
		// for j := m; j >= s; j-- {
		// 	dp[j] = max(dp[j], dp[j-s] + max(0, e - l * (j - s)))
		// }
		for j := 0; j <= m; j++ {
			if j >= s {
				dp[i][j] = max(dp[i][j], dp[i-1][j-s] + max(0, e - l * (j - s)))
			}
		}
	}

	ans := 0
	for i := 0; i <= m; i++ {
		ans = max(ans, dp[n][i])
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	Fscan(in, &T)
	for i := 1; i <= T; i++ {
		ans := solve(in, out)
		Fprintf(out, "Case #%d: %d\n", i, ans)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }