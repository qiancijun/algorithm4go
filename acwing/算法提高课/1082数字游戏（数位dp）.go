package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const N int = 15

var f [N][N]int

func init() {
	for i := 0; i <= 9; i++ { f[1][i] = 1}
	for i := 2; i < N; i++ {
		for j := 0; j <= 9; j++ {
			for k := j; k <= 9; k++ {
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
		// 不合法
		if x < last {
			break
		}

		for j := last; j < x; j++ {
			ans += f[i+1][j]
		}
		last = x
		if i == 0 {
			ans++
		}
	}

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	l, r := 0, 0
	for {
		_, err := Fscan(in, &l, &r)
		if err == io.EOF {
			break
		} else {
			Fprintln(out, dp(r) - dp(l - 1))
		}
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }