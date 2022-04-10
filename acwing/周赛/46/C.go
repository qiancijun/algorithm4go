package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const (
	N int = 10010
)

var n, q int
var strs [N]string


func solve(temp string, out io.Writer) {
	// 求next数组
	var next [10]int
	m := len(temp)
	temp = " " + temp
	// Fprintln(out, temp)
	for i, j := 2, 0; i <= m; i++ {
		for j > 0 && temp[i] != temp[j+1] { j = next[j] }
		if temp[i] == temp[j+1] { j ++ }
		next[i] = j
	}
	// 模式匹配
	ans1, ans2 := 0, "-"
	for t := 0; t < n; t++ {
		str := strs[t]
		k := len(str)
		str = " " + str
		// Fprintln(out, k, str)
		for i, j := 1, 0; i <= k; i++ {
			for j > 0 && str[i] != temp[j+1] { j = next[j] }
			if str[i] == temp[j+1] { j++ }
			if j == m {
				ans1++
				if ans2 == "-" { ans2 = strs[t] }
				// j = next[j]
				break
			}
		}
	}
	Fprintln(out, ans1, ans2)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &strs[i])
	}
	Fscan(in, &q)
	for i := 0; i < q; i++ {
		var temp string
		Fscan(in, &temp)
		solve(temp, out)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }