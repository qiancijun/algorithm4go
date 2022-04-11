package week288

import (
	"strconv"
	"strings"
)

var ans int = 0x3f3f3f3f
var l, r int
var n, m int

func qmi(p, q int) int {
	res, t := 1, p
	for q > 0 {
		if q & 1 == 1 {
			res = res * t
		}
		t = t * t
		q >>= 1
	}
	return res
}

func dfs(n1, n2, p1, p2 int) {
	if p1 == 0 || p2 == m { return }
	n1_left, n1_right := n1 / qmi(10, p1), n1 % qmi(10, p1) 
	n2_left, n2_right := n2 / qmi(10, p2), n2 % qmi(10, p2)
	if n1_left == 0 { n1_left = 1 }
	if n2_right == 0 { n2_right = 1 }
	tmp := (n1_right + n2_left) * n1_left * n2_right
	if ans > tmp {
		ans = tmp
		l, r = p1, p2
	}
	dfs(n1, n2, p1-1, p2)
	dfs(n1, n2, p1, p2+1)
}

func minimizeResult(expression string) string {
	ans = 0x3f3f3f3f
	nums := strings.Split(expression, "+")
	n, m = len(nums[0]), len(nums[1])
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	dfs(n1, n2, n, 0)
	// fmt.Println(l, r)
	ans1 := nums[0][:n-l] + "(" + nums[0][n-l:n]
	ans2 := nums[1][:m-r] + ")" + nums[1][m-r:m]
	return ans1 + "+" + ans2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}