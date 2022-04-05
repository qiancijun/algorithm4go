package main

import (
	"bufio"
	"os"
	."fmt"
)

const N = 2e5 + 10

var n int
var s int64
var a [N]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &s)

	var sum int64
	sum = 0
	for i := 0; i < n; i ++ {
		Fscan(in, &a[i])
		sum += a[i]
	}
	
	if n == 1 {
		Fprint(out, a[0] - 1)
		return
	}

	if sum == s {
		for i := 0; i < n; i ++ {
			Fprintf(out, "%d ", a[i] - 1)
		}
	} else if sum > s {
		for i := 0; i < n - 1; i ++ {
			Fprintf(out, "%d ", 0)
		}
		t := sum - s
		Fprintf(out, "%d ", a[n - 1] - t)
	}
}
