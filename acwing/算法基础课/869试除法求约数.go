package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var solve func(v int) = func(v int) {
		nums := make([]int, 0)
		for i := 1; i <= v / i; i++ {
			if v % i == 0 {
				nums = append(nums, i)
				if i != v / i {
					nums = append(nums, v / i)
				}
			}
		}
		sort.Ints(nums)
		for _, v := range nums {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}

	Fscan(in, &n)
	for ; n > 0; n-- {
		var x int
		Fscan(in, &x)
		solve(x)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }