package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1010
	MOD int = 1e9 + 7
)
var n, m int
var dp, cnt [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := range cnt {
		cnt[i] = 1
	}
	for i := 1; i <= n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		for j := m; j >= v; j-- {
			tmp := dp[j-v] + w
			if tmp > dp[j] {
				dp[j] = tmp 
				cnt[j] = cnt[j-v]
			} else if dp[j] == tmp {
				cnt[j] = (cnt[j] + cnt[j-v]) % MOD
			}
		}
	}
	Fprintln(out, dp[m])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }