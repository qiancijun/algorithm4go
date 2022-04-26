package week289

const N int = 1e5 + 10

var (
	n, m int
	col, row [N]int
)

func maxTrailingZeros(grid [][]int) int {
	n, m = len(grid), len(grid[0])
	idx, mxRow := 0, 0
	for i := range grid {
		cnt := 0
		for j := range grid[i] {
			if grid[i][j] % 5 == 0 { cnt++ }
		}
		if mxRow < cnt {
			idx, mxRow = i, cnt
		}
	}
	ans := 0
	t := mxRow
	for j := 0; j < m; j++ {
		cnt := 0
		for i := 0; i < idx; i++ {
			if grid[i][j] % 5 == 0 { cnt++ }
		}
		ans = max(ans, cnt + t)
		if grid[idx][j] % 5 == 0 { t-- }
	}

	return ans

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }