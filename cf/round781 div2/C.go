package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
	"container/heap"
)

const (
	N int = 2e5 + 10
)

func solve(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	cnt := make([]int, n+1)
	
	nums := make([]int, 0)
	nums = append(nums, 1)
	for i := 0; i < n; i++ {
		var a int
		Fscan(in, &a)
		cnt[a]++
		nums = append(nums, cnt[a])
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
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

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }