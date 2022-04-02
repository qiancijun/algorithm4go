package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b, c, d int 
	Fscan(in, &a, &b, &c, &d)
	ans := 0
	var str string
	Fscan(in, &str)
	for _, v := range str {
		if v == '1' {
			ans += a
		} else if v == '2' {
			ans += b
		} else if v == '3' {
			ans += c
		} else if v == '4' {
			ans += d
		}
	}
	Fprintln(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }