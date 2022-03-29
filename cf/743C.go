// +build ignore
package main

import (
	."fmt"
	"bufio"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	Fprintln(out, n, 2*n, 2*n)
}