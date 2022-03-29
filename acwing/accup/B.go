package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

const N = 210

var T int
var n, k int
var a [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &T)
	for T > 0 {
		T--
		Fscan(in, &n, &k)
		for i := 0; i < k; i ++ {
			Fscan(in, &a[i])
		}
		// solution

		ans := 0
		l, r := 1, n
		var time float64
		time = float64(a[0] - l + 1)
		ans = max(ans, int(math.Ceil(time)))
		// Fprintln(out, ans)
		for j := 1; j < k; j ++ {
			time = float64(a[j] - l - 1) / 2 + 1
			
			ans = max(ans, int(math.Ceil(time)))
			// Fprintln(out, ans)
			l = a[j]
		}
		time = float64(r - a[k - 1] + 1)
		ans = max(ans, int(math.Ceil(time)))
		Fprintln(out, ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}