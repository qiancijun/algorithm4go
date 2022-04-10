package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const (
	N int = 2e5 + 10
)

func solve(in io.Reader, out io.Writer) {
	var cnt [N]int
	n, base, maxLen1, maxLen2 := 0, 0, 0, 0
	Fscan(in, &n)
	for i := 2; i <= n; i++ {
		var x int
		Fscan(in, &x)
		if cnt[x] == 0 { base++ }
		cnt[x]++
		if cnt[x] > maxLen1 {
			maxLen2 = maxLen1
			maxLen1 = cnt[x]
		} else if cnt[x] > maxLen2 {
			maxLen2 = cnt[x]
		}
	}

	ans := 0
	if base >= maxLen1 { 
		ans = base + 1 
	} else {
		step := (maxLen1 - base + 1) / 2
		if base == 1 {
			ans += step + 1 + base
			Fprintln(out, ans)
			return
		}
		if step >= maxLen2 {
			ans += step
		} else {
			ans += step
			ans += (maxLen2 - step) / 2
		}
		ans += (1 + base)
	}
	Fprintln(out, ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	Fscan(in, &t)
	for ; t > 0; t-- {
		solve(in, out)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }