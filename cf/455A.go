package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1e5 + 10
var a, cnt [N]int
var dp [N][2]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	
	dp[1][1] = int64(1 * cnt[1])
	dp[1][0] = 0
	for i := 2; i <= 1e5; i++ {
		dp[i][1] = dp[i-1][0] + int64(i*cnt[i])
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
	}
	Fprintln(out, max(dp[1e5][0], dp[1e5][1]))
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}