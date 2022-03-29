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
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		var str string
		Fscan(in, &n, &str)
		ans := 0
		i, j := 0, -1
		for i < n {
			if str[i] == '1' {
				i++
				continue
			}
			if j == -1 {
				j = i
				i++
			} else {
				if i - j == 1 {
					ans += 2
				} else if i - j == 2 {
					ans += 1
				}
				j = i
				i++
			}
		}
		Fprintln(out, ans)
	}
	defer out.Flush()
}