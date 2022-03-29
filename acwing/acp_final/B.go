package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1e5 + 10

var n, m int
var a, b [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		Fscan(in, &b[i])
	}

	ans := 0
	i, j := -1, -1
	sumA, sumB := 0, 0
	for i < n && j < m {
		if i == n-1 && j == m-1 {
			break
		}
		if sumA < sumB {
			i++
			sumA += a[i]
		} else {
			j++
			sumB += b[j]
		}
		if sumA == sumB {
			ans++
			sumA, sumB = 0, 0
		}
	}
	Fprintln(out, ans)
}
