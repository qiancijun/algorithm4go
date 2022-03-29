package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 2e5 + 10

var n int
var st [N][N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)

	mx := 0
	for i := 1; i < n; i++ {
		var u, v int
		Fscan(in, &u, &v)
		st[u][v] = true
	}

}
