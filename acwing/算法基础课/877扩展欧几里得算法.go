package main

import (
	"bufio"
	. "fmt"
	"os"
)

func exgcd(a, b int, x, y *int) int {
	if b == 0 {
		*x, *y = 1, 0
		return a
	}
	d := exgcd(b, a % b, y, x)
	*y -= (a / b) * *x
	return d
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	for ; n > 0; n-- {
		var a, b, x, y int
		Fscan(in, &a, &b)
		exgcd(a, b, &x, &y)
		Fprintln(out, x, y)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }