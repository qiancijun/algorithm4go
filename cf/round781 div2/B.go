package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	cnt := make(map[int]int)
	mx := 0
	for i := 0; i < n; i++ {
		var x int
		Fscan(in, &x)
		cnt[x]++
		mx = max(mx, cnt[x])
	}
	if n == 1 { Fprintln(out, 0); return }

	step, ans := 0, mx
	for ans < n {
		if ans * 2 <= n {
			step = step + 1 + ans
			ans *= 2
		} else {
			step = step + 1 + (n - ans)
			break
		}
	}
	Fprintln(out, step)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int 
	Fscan(in, &t)
	for; t > 0; t-- {
		solve(in, out)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }