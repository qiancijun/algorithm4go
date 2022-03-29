package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

const N = 1e5 + 10
var cnt [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, l int
	Fscan(in, &n, &l)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &nums[i])
		cnt[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	// Fprintln(out, nums)
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] >= i+1 {
			ans = i+1
		}
	}
	if l >= cnt[ans] {
		ans += cnt[ans]
	}
	Fprintln(out, ans)
}

func max(a, b int) int { if a > b { return a }; return b }
