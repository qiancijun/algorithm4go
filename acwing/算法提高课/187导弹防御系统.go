package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 55
var nums, down, up [N]int
var n, ans int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		Fscan(in, &n)
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			Fscan(in, &nums[i])
		}

		ans = n
		dfs(0, 0, 0)
		Fprintln(out, ans)
	}
}

func dfs(u, su, sd int) {
	if su + sd >= ans {
		return
	}
	if u == n {
		ans = su + sd // 更新答案
		return
	}
	// 1. 将当前数放入上升子序列中
	k := 0
	for k < su && up[k] >= nums[u] {
		k++
	}
	t := up[k]
	up[k] = nums[u]
	if k < su {
		dfs(u+1, su, sd)
	} else {
		dfs(u+1, su+1, sd)
	}
	up[k] = t // 回溯

	// 2. 将当前数放入下降子序列中
	k = 0
	for k < sd && down[k] <= nums[u] {
		k++
	}
	t = down[k]
	down[k] = nums[u]
	if k < sd {
		dfs(u+1, su, sd)
	} else {
		dfs(u+1, su, sd+1)
	}
	down[k] = t // 回溯
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }