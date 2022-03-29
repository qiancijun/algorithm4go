// +build ignore
package main

import (
	"bufio"
	. "fmt"
	"os"
)

var n, k, d int

const MOD = 1e9 + 7
const N = 1e5 + 10

var dp [N][2]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &k, &d)
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i && j <= k; j++ {
			if j >= d {
				dp[i][1] += dp[i-j][0] + dp[i-j][1]
			} else {
				dp[i][1] += dp[i-j][1]
				dp[i][0] += dp[i-j][0]
			}
			dp[i][0] %= MOD
			dp[i][1] %= MOD
		}
	}
	Fprintln(out, dp[n][1])
}
