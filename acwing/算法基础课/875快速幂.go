package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func qmi(a, b, p int) int {
	res, t := 1 % p, a
	for b > 0 {
		if b & 1 == 1 {
			res = res * t % p
		}
		
		t = t * t % p
		b >>= 1
	}
	return res
}

func solve(in io.Reader, out io.Writer) {
	var a, b, p int
	Fscan(in, &a, &b, &p)
	Fprintln(out, qmi(a, b, p))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	Fscan(in, &T)
	for ; T > 0; T-- {
		solve(in, out)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }