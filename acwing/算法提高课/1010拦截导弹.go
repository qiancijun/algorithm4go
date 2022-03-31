package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1010
var n int = 0
var nums, dp, f [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for {
		var val int
		if c, _ := Fscan(in, &val); c != 0 {
			nums[n] = val
			n++
		} else {
			break
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	Fprintln(out, res)

	res = 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		res = max(res, f[i])
	}
	Fprint(out, res)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }