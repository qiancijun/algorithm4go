package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 2010
const INF = 1e8

var n, m, k int
var s1, s2 [N]int
var a, b [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var c int
		Fscan(in, &c)
		s1[i] = s1[i-1] + c
	}
	for i := 1; i <= m; i++ {
		var c int
		Fscan(in, &c)
		s2[i] = s2[i-1] + c
	}
	Fscan(in, &k)

	for length := 1; length <= n; length++ {
		a[length] = INF
		for i := 1; i + length - 1 <= n; i++ {
			a[length] = min(a[length], s1[i+length-1] - s1[i-1])
		}
	}

	for length := 1; length <= m; length++ {
		b[length] = INF
		for i := 1; i + length - 1 <= m; i++ {
			b[length] = min(b[length], s2[i+length-1] - s2[i-1])
		}
	}

	res := 0
	for i, j := 1, m; i <= n; i++ {
		for j > 0 && a[i] > k / b[j] {
			j--
		}
		res = max(res, i*j)
	}
	Fprintln(out, res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
