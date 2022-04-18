package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	Fscan(in, &str)
	n := len(str)
	ans1, ans2 := false, false
	for i := 0; i + 2 <= n; i++ {
		j := i + 2
		substr := str[i:j]
		if substr == "AB" && !ans1 { 
			ans1 = true 
			i = j - 1
		}
		if ans1 && substr == "BA" && !ans2 {
			ans2 = true
			i = j - 1
		}
	}

	if ans1 && ans2 {
		Fprintln(out, "YES")
		return
	}

	ans1, ans2 = false, false
	for i := 0; i + 2 <= n; i++ {
		j := i + 2
		substr := str[i:j]
		if substr == "BA" && !ans1 { 
			ans1 = true 
			i = j - 1
		}
		if ans1 && substr == "AB" && !ans2 {
			ans2 = true
			i = j - 1
		}
	}

	if ans1 && ans2 {
		Fprintln(out, "YES")
		return
	}

	Fprintln(out, "NO")
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }