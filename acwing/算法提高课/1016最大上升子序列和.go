package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010
var n int
var nums, dp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &nums[i])
	}

	for i := 1; i <= n; i++ {
		dp[i] = nums[i]
		for j := 1; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+nums[i])
			}
		}
	}
	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	Fprint(out, ans)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }