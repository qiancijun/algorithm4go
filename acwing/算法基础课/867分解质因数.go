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
	var solve func(v int) = func(v int) {
		for i := 2; i <= v / i; i++ {
			if v % i == 0 {
				cnt := 0
				for v % i == 0 {
					v /= i
					cnt++
				}
				Fprintln(out, i, cnt)
			}
		}
		if v > 1 {
			Fprintln(out, v, 1)
		}
	}

	Fscan(in, &n)
	for ; n > 0; n-- {
		var x int
		Fscan(in, &x)
		solve(x)
		Fprintln(out)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }