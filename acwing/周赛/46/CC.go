package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 10010
)

var n, q int
var strs [N]string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	m := make(map[string]map[string]bool)

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &strs[i])
		m[strs[i]] = make(map[string]bool)
		for j := 0; j < len(strs[i]); j++ {
			for k := j; k < len(strs[i]); k++ {
				substr := strs[i][j:k+1]
				m[strs[i]][substr] = true
			}
		}
	}
	Fscan(in, &q)
	for i := 0; i < q; i++ {
		var temp string
		Fscan(in, &temp)
		ans1, ans2 := 0, "-"
		for _, str := range strs {
			if _, has := m[str][temp]; has {
				ans1++
				if ans2 == "-" { ans2 = str }
			}
		}
		Fprintln(out, ans1, ans2)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }