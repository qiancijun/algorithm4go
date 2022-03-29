package main

import (
	"bufio"
	. "fmt"
	"os"
)

var points [110][2]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		var x, y int
		Fscan(in, &x, &y)
		points[i] = [2]int{x, y}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i != j && j != k && i != k {
					if points[i][0] == points[j][0] && points[j][1] == points[k][1] {
						res = max(res, abs(points[j][0]-points[k][0]) * abs(points[j][1]-points[i][1]))
					}
				}
			}
		}
	}
	Fprintln(out, res)
}

func max(a, b int) int { if a > b { return a }; return b }
func abs(v int) int { if v > 0 { return v }; return -v }