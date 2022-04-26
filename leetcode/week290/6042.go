package week290

import (
	"fmt"
	"math"
)

type pair struct{ x, y int }

func countLatticePoints(circles [][]int) int {
	m := make(map[pair]bool)
	for i := range circles {
		circle := circles[i]
		x, y, r := circle[0], circle[1], circle[2]
		leftTop, rightBottom := pair{x - r, y + r}, pair{x + r, y - r}
		for p := leftTop.x; p <= rightBottom.x; p++ {
			for q := leftTop.y; q >= rightBottom.y; q-- {
				point := pair{p, q}
				if check(point, x, y, r) {
					m[point] = true
				}
			}
		}
	}
	return len(m)
}

func check(point pair, x, y, r int) bool {
	px, py := point.x, point.y
	dis := math.Sqrt(math.Pow(float64(px - x), 2) + math.Pow(float64(py - y), 2))
	return dis <= float64(r)
}
func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }