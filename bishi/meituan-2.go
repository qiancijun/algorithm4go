package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1e5 + 10

var n int
var a [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	ans := 0x3f3f3f3f
	mei, tuan := a[1], a[n]
	for i := 1; i <= n; i++ {
		mid := a[i]
		dis := abs((mid - mei) - (tuan - mid))
		ans = min(ans, dis)
	}
	Fprintln(out, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}