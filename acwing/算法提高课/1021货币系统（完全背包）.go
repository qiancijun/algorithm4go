package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 3010
var n, m int
var cnt [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		var a int
		Fscan(in, &a)
		cnt[a]++
		for j := a; j <= m; j++ {
			cnt[j] += cnt[j-a]
		}
	}
	Fprintln(out, cnt[m])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }