package main

import (
	"bufio"
	. "fmt"
	"os"
)

var n, x int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &x)

	if x == 1 {
		if n % 2 == 1 {
			n %= 6
			if n == 1 || n == 2 {
				Fprintln(out, 0)
			} else if n == 4 || n == 5 {
				Fprintln(out, 2)
			} else {
				Fprintln(out, 1)
			}
		} else {
			n %= 6
			if n == 1 || n == 2 {
				Fprintln(out, 2)
			} else if n == 4 || n == 5 {
				Fprintln(out, 0)
			} else {
				Fprintln(out, 1)
			}
		}
	} else if x == 2 {
		n %= 6
		if n == 0 || n == 5 {
			Fprintln(out, 2)
		} else if n == 1 || n == 4 {
			Fprintln(out, 1)
		} else {
			Fprintln(out, 0)
		}
	} else {
		n %= 6
		if n == 0 || n == 5 {
			Fprintln(out, 0)
		} else if n == 1 || n == 4 {
			Fprintln(out, 1)
		} else {
			Fprintln(out, 2)
		}
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }