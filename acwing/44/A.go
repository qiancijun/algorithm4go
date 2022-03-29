package main

import (
	"bufio"
	. "fmt"
	"os"
)

var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	s := make(map[int]bool)
	for i := 1; i <= n; i++ {
		var a int
		Fscan(in, &a)
		s[a] = true
	}
	if _, has := s[0]; has {
		Fprintln(out, len(s)-1)
	} else {

		Fprintln(out, len(s))
	}
}
