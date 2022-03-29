package main

import (
	"bufio"
	. "fmt"
	"os"
)

// acbcca
var cnt [26]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var str string
	Fscan(in, &str)
	for _, v := range str {
		cnt[v-'a']++
	}
	// for _, v := range cnt {
	// 	Fprint(out, v, " ")
	// }
	// Fprintln(out)
	a, b, c := cnt[0]-1, cnt[1], cnt[2]/3
	ans := min(b, c)
	Fprintln(out, min(ans, a))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}