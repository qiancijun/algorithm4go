package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 6010
var n, m int
var dp [N]int

type pii struct { x, y int }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)

	list := make([]pii, 0)
	for i := 0; i < n; i++ {
		var v, w, s int
		Fscan(in, &v, &w, &s)
		for k := 1; k <= s; k++ {
			s -= k
			list = append(list, pii{k * v, k * w})
		}
		if s > 0 {
			list = append(list, pii{s * v, s * w})
		}
	}

	for _, v := range list {
		for i := m; i >= v.x; i-- {
			dp[i] = max(dp[i], dp[i-v.x] + v.y)
		}
	}

	Fprintln(out, dp[m])

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }