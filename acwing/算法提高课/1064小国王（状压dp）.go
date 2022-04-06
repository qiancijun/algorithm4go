package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 12
)

var n, k int
var cnt [1<<N]int
var state []int
var g [1<<N][]int
var dp [N][N*N][1<<N]int // dp[i][j][s] 表示第 i 行放置了 j 个国王，从 s 状态转移过来

func check(val int) bool {
	for i := 0; i < n; i++ {
		if (val >> i & 1) == 1 && (val >> (i+1) & 1) == 1 {
			return false
		}
	}
	return true
}


func count(val int) int {
	ans := 0
	for i := 0; i < n; i++ {
		if val >> i & 1 == 1 {
			ans++
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fscan(in, &n, &k)

	state = make([]int, 0)

	// 枚举所有没有连续1的情况
	for i := 0; i < (1<<n); i++ {
		if check(i) {
			state = append(state, i)
			cnt[i] = count(i)
		}
	}

	// 枚举所有相邻合法的情况
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state); j++ {
			a, b := state[i], state[j]
			if (a & b) == 0 && check(a | b) {
				// a & b 保证了上下不相邻，a | b 保证了斜对角不相邻
				g[i] = append(g[i], j)

			}
		}
	}

	dp[0][0][0] = 1 // 什么都不放一种情况
	for i := 1; i <= n + 1; i++ {
		for j := 0; j <= k; j++ {
			for a := 0; a < len(state); a++ {
				for b := 0; b < len(g[a]); b++ {
					c := cnt[state[a]]
					if j - c >= 0 {
						dp[i][j][a] += dp[i-1][j-c][g[a][b]]
					}
				}
			}
		}
	}
	Fprintln(out, dp[n+1][k][0])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }