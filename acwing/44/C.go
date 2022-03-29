package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 100010
var n, m int
var cnt [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fscan(in, &n, &m)
	var res int = 0
	for ; n > 0; n-- {
		var x int
		Fscan(in, &x)
		var y, z int64 = 1, 1
		for i := 2; i * i <= x; i++ {
			if x % i == 0 {
				s := 0
				for x % i == 0 {
					x /= i
					s++
				}
				s %= m
				if s > 0 {
					y *= power(i, s)
					z *= power(i, m - s)
				}
			}
		}
		if x > 1 {
			y *= int64(x)
			z *= power(x, m-1)
		}
		if z >= N {
			z = 0
		}
		res += cnt[z]
		cnt[y]++
	}
	Fprintln(out, res)
}

func power(a, b int) int64 {
	var res int64 = 1
	for b > 0 {
		b--
		res *= int64(a)
		if res >= 0 {
			res = 0
		}
	}
	return res
}