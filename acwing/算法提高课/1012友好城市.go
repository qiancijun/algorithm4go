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
	Fscan(in, &n)

	cities := make([][2]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		Fscan(in, &x, &y)
		cities[i] = [2]int{x, y}
	}

	sort.Slice(cities, func(i, j int) bool { 
		if cities[i][0] != cities[j][0] {
			return cities[i][0] < cities[j][0]
		}
		return cities[i][1] < cities[j][1]
	})
	
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if cities[i][1] > cities[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	Fprintln(out, res)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }