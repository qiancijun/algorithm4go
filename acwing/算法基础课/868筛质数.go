package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1000010

var cnt = 0
var prime [N]int
var st [N]bool

// 垃圾版
func is_prime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i <= x / i; i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
}

// nlogn 做法，容易理解
func get_prime1(x int) {
	for i := 2; i <= x; i++ {
		if !st[i] {
			prime[cnt] = i
			cnt++
			// 优化一下，只有是质数的时候才筛掉它的倍数
			for j := i + i; j <= x; j += i {
				st[j] = true
			}
		}
	}
}

// 线性筛
func get_prime2(x int) {
	for i := 2; i <= x; i++ {
		if !st[i] {
			prime[cnt] = i
			cnt++	
		}
		for j := 0; prime[j] <= x / i; j++ {
			st[prime[j]*i] = true
			if i % prime[j] == 0 {
				break
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	
	// n, ans := 0, 0
	// Fscan(in, &n)
	// for i := 1; i <= n; i++ {
	// 	if is_prime(i) {
	// 		ans++
	// 	}
	// }
	// Fprintln(out, ans)

	n := 0
	Fscan(in, &n)
	get_prime2(n)
	Fprintln(out, cnt)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }