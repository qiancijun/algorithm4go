package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010
var n, m, k int
var cnt, dmg [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m, &k)

	for i := range dmg {
		dmg[i] = m
	}

	for i := 1; i <= k; i++ {
		var c, d int
		Fscan(in, &c, &d)
		
	}
	Fprintln(out, cnt[n], dmg[m])

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }