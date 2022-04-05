package main

import (
	"bufio"
	"os"
	."fmt"
)

const N = 2e5 + 10

var n int
var t int64
var a [N]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &t)

	for i := 1; i <= n; i ++ {
		var tmp int64
		Fscan(in, &tmp)
		a[i] = a[i - 1] + tmp
	}
	ans := 0

	for i := 1; i <= n; i ++ {
		Fprintf(out, "%d ", a[i])
	}

	Fprintln(out, ans)
}
