package main

import (
	"bufio"
	. "fmt"
	"os"
)

var T, n int
var B, x, y int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &T)
	for t := 0; t < T; t++ {
		Fscan(in, &n, &B, &x, &y)
		var last, ans int64
		last, ans = 0, 0
		for i := 1; i <= n; i++ {
			if last + x <= B {
				last = last + x 
			} else {
				last = last - y
			}
			ans += last
		}
		Fprintln(out, ans)
	}
}