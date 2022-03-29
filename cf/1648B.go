// +build ignore
package main

import (
	"bufio"
	. "fmt"
	"os"
)

var T int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &T)
	for t := 0; t < T; t++ {
		var n, c int
		Fscan(in, &n, &c)
		a := make([]int, n+1)
		cnt := make([]int, 2*c+1)
		pre := make([]int, 2*c+1)
		for j := 1; j <= n; j++ {
			Fscan(in, &a[j])
			cnt[a[j]]++
		}

		for i := 1; i <= 2*c; i++ {
			pre[i] += pre[i-1] + cnt[i]
		}
		
		ok := true
		for i := 1; i <= c; i++ {
			if cnt[i] > 0 {
				for j := 1; i*j <= c; j++ {
					if cnt[j] == 0 && pre[i*j-1] != pre[i*(j+1)-1] {
						ok = false
						break
					}
				}
				if !ok {
					break
				}
			}
		}
		if ok {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}
