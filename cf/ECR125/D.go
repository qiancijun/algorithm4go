package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 3e5 + 5
const M = 1e6 + 5

var a [N]int64
var dp [M]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, c int
	Fscan(in, &n, &c)
	for i := 1; i <= n; i++ {
		var c, h, d int64
		Fscan(in, &c, &h, &d)
		dp[c] = max(dp[c], h*d)
	}

	for i := 1; i <= c; i++ {
		for j := i + i; j <= c; j += i {
			dp[j] = max(dp[j], dp[i]*int64((j/i)))
		}
	}

	var tmp int64 = 0
	for i := 1; i <= c; i++ {
		tmp = max(tmp, dp[i])
		dp[i] = tmp
		// Fprint(out, dp[i], " ")
	}
	var m int
	Fscan(in, &m)
	for ; m > 0; m-- {
		var h, d int64
		Fscan(in, &h, &d)
		ans := search(1, c, h*d+1)
		if dp[ans] < h*d+1 {
			Fprint(out, -1, " ")
		} else {
			Fprint(out, ans, " ")
		}
	}
}

func search(l, r int, v int64) int {
	for l < r {
		m := (l + r) >> 1
		if dp[m] >= v {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
