package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 35

var a, b, n int
var f[N][N][110] int

func init() {
	for i := 1; i <= 9; i++ { 
		if i % N == 0 {
			f[1][i][i] = 1
		}
	}

	// i 位数，最高位是 j，前 i 位和是 s
	for i := 2; i < N; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if n - k >= 0 {
					f[i][j][n] += f[i-1][j][n - k] 
				}
			}
		}
	}
}

func dp(v int) int {
	if v == 0 {
		return 0
	}

	nums := make([]int, 0)
	for v > 0 {
		nums = append(nums, v % 10)
		v /= 10
	}

	ans, last := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		for j := 0; j < x; j++ {
			ans += f[i][j][n]
		}
		last += x

		if i == 0 && last % n == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &a, &b, &n)
	Fprintln(out, dp(b) - dp(a - 1))
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }