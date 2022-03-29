package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 5e4 + 10
const INF = 0x3f3f3f3f

var n int
var a [N]int
var dp [8]int

// 4
// 1 3 6 6

// 5
// -2 -6 15 4 5

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		var tmp int
		if a[i] > 0 {
			tmp = (a[i] % 7)
		} else {
			tmp = (7 - a[i]) % 7
		}
		if dp[tmp] != 0 {
			dp[7] = max(dp[7], dp[7 - tmp] + a[i])
		}
		dp[tmp] += a[i]
	}
	Fprintln(out, dp[7])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}