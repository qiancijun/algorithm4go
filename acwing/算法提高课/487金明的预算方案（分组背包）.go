package main

import (
	"bufio"
	. "fmt"
	"os"
)

type pii struct {
	v, w int
}

const N int = 32010
var n, m int
var dp [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	goods := make([][]pii, m+1)
	for i := 1; i <= m; i++ {
		var v, p, q int
		Fscan(in, &v, &p, &q)
		if q == 0 {
			// 主件
			goods[i] = append(goods[i], pii{v, v * p})
		} else {
			goods[q] = append(goods[q], pii{v, v * p})
		}
	}
	for _, good := range goods {
		list := make([]pii, 0)
		for i := range good {
			if i == 0 {
				list = append(list, pii{good[i].v, good[i].w})
			} else if i == 1 {
				list = append(list, pii{good[i].v + good[0].v, good[i].w + good[0].w})
			} else if i == 2 {
				list = append(list, pii{good[i].v + good[0].v, good[i].w + good[0].w})
				list = append(list, pii{good[i].v + good[1].v + good[0].v, good[i].w + good[1].w + good[0].w})
			}
		}
		for j := n; j >= 0; j-- {
			for _, l := range list {
				v, w := l.v, l.w
				if j >= v {
					dp[j] = max(dp[j], dp[j-v] + w)
				}
			}
		}
	}
	Fprintln(out, dp[n])
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }