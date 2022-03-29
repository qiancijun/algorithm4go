package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1e5 + 10
var n, t int
var a, pre [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fscan(in, &n, &t)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		pre[i] = pre[i-1] + a[i]
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if a[i] > t {
			continue
		}
		l, r := i, n 
		for l < r {
			m := (l + r + 1) >> 1
			if pre[m] - pre[i-1] <= t {
				l = m
			} else {
				r = m - 1
			}
		}
		ans = max(ans, l-i+1)
	}
	Fprintln(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}