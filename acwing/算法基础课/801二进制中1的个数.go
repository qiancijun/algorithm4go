package main

import (
	"bufio"
	. "fmt"
	"os"
)

func lowbit(x int) int {
	return x & (-x)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int 
	Fscan(in, &n)
	for ; n > 0; n-- {
		var x int
		Fscan(in, &x)
		ans := 0
		for i := x; i > 0; i -= lowbit(i) { ans++ }
		Fprint(out, ans, " ")
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }