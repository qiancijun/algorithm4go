package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 5e5 + 10

var n, m int
var nums [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &nums[i])
	}

	l, r := 0, 0
	s := make(map[int]int)
	i, j := 0, 0
	for j < n {
		
		s[nums[j]]++
		for i < j && len(s) > m {
			s[nums[i]]--
			if s[nums[i]] <= 0 {
				delete(s, nums[i])
			}
			i++
		}
		if j - i > r - l {
			l, r = i, j
		}
		
		j++
	}
	Fprintln(out, l+1, r+1)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }