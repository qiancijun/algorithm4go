// +build ignore
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 5e5 + 10

var n int
var a, pre [N]int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		pre[i] = pre[i-1] + a[i]
	}

	if pre[n] % 3 != 0 {
		Fprintln(out, 0)
		return
	}

	k := pre[n] / 3
	var cnti, cnt int64
	cnti, cnt = 0, 0
	for i := 1; i < n; i++ {
		if pre[i] == 2 * k {
			cnt += cnti
		}
		if pre[i] == k {
			cnti++
		}
	}
	Fprintln(out, cnt)
}
