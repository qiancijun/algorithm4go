package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 12

var f [N][N]int

func init() {
	for i := 0; i <= 9; i++ { f[1][i] = 1 } // 一位数处理，去掉4
	f[1][4] = 0
	for i := 2; i < N; i++ {
		for j := 0; j <= 9; j++ {
			if j == 4 { continue }
			for k := 0; k <= 9; k++ {
				if k == 4 { continue }
				if j == 6 && k == 2 { continue }
				f[i][j] += f[i-1][k]
			}
		}
	}
}

func dp(n int) int {
	if n == 0 {
		return 1
	}
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n % 10)
		n /= 10
	}
	
	ans, last := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		for j := 0; j < x; j++ {
			if j == 4 { continue }
			if last == 6 && j == 2 { continue }
			ans += f[i+1][j]
		}
		if x == 4 || (last == 6 && x == 2){ break }
		last = x
		if i == 0 { ans++ }
	}

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	for {
		Fscan(in, &n, &m)
		if n + m == 0 {
			break
		}
		Fprintln(out, dp(m) - dp(n - 1))
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }