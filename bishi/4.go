package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	var sum int64 = 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j]})
	pre := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + int64(a[i-1])
	}
	l, r := 1, n
	for l < r {
		sum += (pre[r] - pre[l-1])
		l++
		r--
	}
	Fprintln(out, sum)
}
