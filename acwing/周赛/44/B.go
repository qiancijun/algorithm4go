package main

import (
	"bufio"
	. "fmt"
	"os"
)

var str string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &str)
	if len(str) == 1 {
		Fprintln(out, "YES")
		return
	}
	s := make(map[[2]int]bool)
	cur := [2]int{0, 0}
	s[cur] = true
	l, r, d, u := 0, 0, 0, 0
	for i := 0; i < len(str); i++ {
		var node [2]int
		switch str[i] {
		case 'D':
			d++
			node = [2]int{cur[0]-1, cur[1]}
		case 'U':
			u++
			node = [2]int{cur[0]+1, cur[1]}
		case 'R':
			r++
			node = [2]int{cur[0], cur[1]+1}
		case 'L':
			l++
			node = [2]int{cur[0], cur[1]-1}
		}
		cur = node
		// Fprintln(out, node)
		if _, has := s[node]; has {
			Fprintln(out, "NO")
			return
		}
		s[node] = true
	}
	if l == r || u == d {
		Fprintln(out, "NO")
		return
	}
	Fprintln(out, "YES")
}
