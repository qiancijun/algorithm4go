package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 1e4 + 10

var (
	n int
	a [N]int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := n-1; i >= 0; i-- {
		a[i+1] -= a[i]
	}

	ans := make([]int, 0)
	for i := 1; i <= n; i++ {
		v := a[i]
		for j := i + 1; j <= n; j++ {
			if a[j] == v {
				t := j - i
				if len(ans) == 0 {
					ans = append(ans, t)
				} else {
					last := ans[len(ans)-1]
					if t > last {
						ans = append(ans, t)
					}
				}
			}
		}
	}
	ans = append(ans, n)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}


func uniqueArr(arr []int) []int {
	newArr := make([]int, 0)
	tempArr := make(map[int]bool, len(newArr))
	for _, v := range arr {
		if tempArr[v] == false {
			tempArr[v] = true
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }