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

	s := make(map[byte]int)
	for {
		var ch byte
		Fscanf(in, "%c", &ch)
		if ch == '{' || ch == ',' || ch == ' ' {
			continue
		} else if ch == '}' {
			break
		} else {
			s[ch]++
		}
	}
	Fprintln(out, len(s))

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }