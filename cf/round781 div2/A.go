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

	var t int
	Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		Fscan(in, &n)
		Fprintln(out, 1, n - 3, 1, 1)
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }