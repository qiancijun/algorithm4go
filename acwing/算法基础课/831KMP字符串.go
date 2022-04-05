package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e5 + 10
	M int = 1e6 + 10
)
var n, m int
var s, p string
var next [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &p, &m, &s)
	p, s = " " + p, " " + s

	for i, j := 2, 0; i <= n; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}

	for i, j := 1, 0; i <= m; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == n {
			Fprint(out, i - n)
			j = next[j]
		}
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }