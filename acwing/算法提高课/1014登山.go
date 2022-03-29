package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010

var nums [N]int
var dp [N][2]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &nums[i])
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i][0] = max(dp[i][0], dp[j][0]+1)
			}
		}
	}
	for i := n-1; i >= 0; i-- {
		dp[i][1] = 1
		for j := n-1; j > i; j-- {
			if nums[i] > nums[j] {
				dp[i][1] = max(dp[i][1], dp[j][1]+1)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dp[i][0]+dp[i][1]-1)
	}
	Fprintln(out, ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
