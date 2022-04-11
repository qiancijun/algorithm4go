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
	
	nums := make([]int, n)
	st := make([]bool, n+1)
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			for _, v := range nums { Fprint(out, v, " ") }
			Fprintln(out)
			return
		}
		for i := 1; i <= n; i++ {
			if !st[i] {
				st[i] = true
				nums[pos] = i
				dfs(pos+1)
				nums[pos] = 0
				st[i] = false
			}
		}
	}
	dfs(0)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }