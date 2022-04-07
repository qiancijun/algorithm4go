package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 15

var f [N][N]int

func init() {
	for i := 0; i <= 9; i++ { f[1][i] = 1 } // 必须从 0 开始枚举
	for i := 2; i < N; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if (abs(j - k) >= 2) {
					f[i][j] += f[i-1][k]
				}
			}
		}
	}
}

func dp(n int) int {
	if n == 0 {
		return 0
	}

	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n % 10)
		n /= 10
	}

	ans, last := 0, -2
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i]
		j := 0
		if i == len(nums) - 1 {
			j = 1
		}
		for ; j < x; j++ {
			if abs(last - j) >= 2 {
				ans += f[i+1][j]
			}
		}
		
		if abs(x - last) >= 2 {
			last = x 
		} else {
			break
		}
		if i == 0 {
			ans++
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= 9; j++ {
			ans += f[i][j]
		}
	}

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	a, b := 0, 0
	Fscan(in, &a, &b)
	Fprintln(out, dp(b) - dp(a - 1))

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }