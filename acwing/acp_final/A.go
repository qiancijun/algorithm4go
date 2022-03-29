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
	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 1000; j++ {
			if b + i*a == d + j*c {
				Fprintln(out, b+i*a)
				return
			}
		}
	}
	Fprintln(out, -1)
}
