package main

import (
	"bufio"
	. "fmt"
	"os"
)

func qmi(p, q, m int) int {
	res, t := 1 % m, p
	for q > 0 {
		if q & 1 == 1 {
			res = res * t % m
		}
		t = t * t % m
		q >>= 1
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for ; n > 0; n-- {
		var p, q, m int
		Fscan(in, &p, &q, &m)
		Fprintln(out, qmi(p, q, m))
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }