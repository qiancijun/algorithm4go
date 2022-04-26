package week290

import (
	"fmt"
	"sort"
)

func countRectangles(rectangles [][]int, points [][]int) []int {
	s := make([]int, 5, 10)
	return s
	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][0] != rectangles[j][0] {
			return rectangles[i][0] < rectangles[j][0]
		}
		return rectangles[i][1] < rectangles[j][1]
	})
	n := len(rectangles)
	res := make([]int, 0)
	for _, point := range points {
		x, y := point[0], point[1]
		// 二分找到横坐标
		l, r := 0, n - 1
		for l < r {
			m := (l + r) >> 1
			if rectangles[m][0] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		cnt := 0
		if x > rectangles[l][0] {
			res = append(res, 0)
			continue
		}
		for i := l; i < n; i++ {
			if rectangles[i][1] >= y { cnt++ }
		}
		res = append(res, cnt)
	}
	return res
}