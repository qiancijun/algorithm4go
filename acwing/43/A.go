package main

import (
	"bufio"
	"os"
	."fmt"
)

var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)

	ans := 0
	for a := 1; a <= n; a ++ {
		for b := 1; b <= a; b ++ {
			for c := 1; c <= b; c ++ {
				if a ^ b ^ c == 0 && isR(a, b, c) {
					ans ++
				}
			}
		}
	}
	Fprintln(out, ans)
}

func isR(a, b, c int) bool {
	if a + b <= c {
		return false
	}
	if a + c <= b {
		return false
	}
	if b + c <= a {
		return false
	}
	return true
}