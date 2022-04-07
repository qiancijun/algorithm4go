package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 32

var f [N][N]int
var x, y, k, b int

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = f[i-1][j] + f[i-1][j-1]
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
		nums = append(nums, n % b)
		n /= b
	}
	ans, last := 0, 0
	for i := len(nums) - 1; i >= 0 ; i-- {
		x := nums[i]
		if x > 0 {
			ans += f[i][k-last] // 如果x取0，就从后面的几位中取1
			// 只有大于1的时候才有选择
			// 选择大于1的数都可以直接算出来
			if x > 1 {
				if k - last - 1 >= 0 {
					ans += f[i][k-last-1]
					break
				} 
			} else {
				// 选的是1
				last++
				if last > k {
					break
				}
			}
		}
		if i == 0 && last == k {
			ans++ // 数本身是不是合法
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &x, &y, &k, &b)
	Fprintln(out, dp(y) - dp(x - 1))

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }