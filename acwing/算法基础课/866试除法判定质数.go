package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for ; n > 0; n-- {
		var x int
		Fscan(in, &x)
		var check func(int) bool
		check = func(v int) bool {
			if v < 2 {
				return false
			}
			for i := 2; i <= v / i; i++ {
				if v % i == 0 {
					return false
				}
			}
			return true
		}
		if check(x) {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }