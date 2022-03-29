package main

import (
	"bufio"
	. "fmt"
	"os"
)

var k, n int
var nums [30][30]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &k, &n)
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &nums[i][j])
		}
	}

	ans := 0
	vi := make(map[[2]int]bool)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if _, has := vi[[2]int{i, j}]; has || i == j {
				continue
			}
			flag := true
			for t := 0; t < k; t++ {
				x := find(t, i)
				y := find(t, j)
				if x > y {
					flag = false
					break
				}
			}
			if flag {
				vi[[2]int{i, j}] = true
				vi[[2]int{j, i}] = true
				ans++
			}
		}
	} 
	Fprintln(out, ans)
}

func find(level, val int) int {
	for i, v := range nums[level] {
		if v == val {
			return i
		}
	}
	return -1
}