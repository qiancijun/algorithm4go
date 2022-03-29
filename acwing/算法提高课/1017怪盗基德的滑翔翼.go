package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const N int = 110

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	nums, dp := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &nums[i])
	}
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	for _, v := range dp {
		ans = max(ans, v)
	}
	Fprintln(out, ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	Fscan(in, &t)
	for ; t > 0; t-- {
		solve(in, out)
	}
}

func max(a, b int) int { if a < b { return b }; return a }