package main

import (
	"bufio"
	. "fmt"
	"os"
)

<<<<<<< HEAD
const (
	N int = 1e5 + 10
)

var (
	s string
)

=======
>>>>>>> 8c5717383ac412c04a5951cc62ff58eceed3bc72
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

<<<<<<< HEAD
	Fscan(in, &s)
	n := 2
	ans1, ans2 := false, false
	for i := 0; i + n <= len(s); i++ {
		r := i + n
		sub := s[i:r]
		if sub == "AB" && !ans1 {
			ans1 = true
			i = r-1
		}
		if sub == "BA" && !ans2 {
			ans2 = true
			i = r-1
		}
	}
	if ans1 && ans2 {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}

=======
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
>>>>>>> 8c5717383ac412c04a5951cc62ff58eceed3bc72
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }