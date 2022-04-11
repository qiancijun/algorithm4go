package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var m, s int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &m, &s)

	// n 位数，最小的情况为1000，最坏的情况9999
	if m * 9 < s {
		Fprintln(out, -1, -1)
		return
	}
	if m != 1 && s == 0 {
		Fprintln(out, -1, -1)
		return
	}
	smallest, largest := "", ""
	for s > 9 {
		s -= 9
		smallest = "9" + smallest
		largest += "9"
		m--
	}
	last := strconv.Itoa(s)
	if m == 1 {
		smallest = last + smallest
		largest += last
	} else {
		largest += last
		remain := s
		for m > 1 {
			if remain > 1 {
				smallest = strconv.Itoa(remain-1) + smallest
				remain = 1
			} else {
				smallest = "0" + smallest
			}
			largest += "0"
			m--
		}
		smallest = "1" + smallest
	}

	Fprintln(out, smallest, largest)

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }