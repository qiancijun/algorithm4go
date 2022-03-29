package main

import (
	"bufio"
	. "fmt"
	"os"
)

var T, n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &T)
	for t := 0; t < T; t++ {
		Fscan(in, &n)
		var seq string
		Fscan(in, &seq)
		ans, i := 0, 0
		for i < len(seq) {
			if seq[i] == '(' && i != n-1 {
				i += 2
				ans++
			} else {
				j := i + 1
				for j < n && seq[j] == '(' {
					j ++ 
				}
				if j == n {
					break
				} else {
					ans ++ 
					i = j + 1
				}
			}
		}
		Fprintln(out, ans, n - i)
	}
}
