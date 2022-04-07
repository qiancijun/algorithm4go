package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1e5 + 10

var n, m, q int
var next [N]int
var S, T string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m, &q)
	Fscan(in, &S, &T)
	S, T = " " + S, " " + T

	for i, j := 2, 0; i <= m; i++ {
		for j > 0 && T[i] != T[j+1] { j = next[j] }
		if T[i] == T[j+1] { j++ }
		next[i] = j
	} 

	var sub string
	for ; q > 0; q-- {
		l, r := 0, 0
		Fscan(in, &l, &r)
		sub = " " + S[l:r+1]

		ans := 0
		for i, j := 1, 0; i <= r-l+1; i++ {
			for j > 0 && sub[i] != T[j+1] { j = next[j] }
			if sub[i] == T[j+1] { j++ }
			if j == m { 
				ans++ 
				j = next[j]
			}
		}
		Fprintln(out, ans)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }