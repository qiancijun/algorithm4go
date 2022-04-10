package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 2010
	MOD int = 1e9 + 7
)
var f [N][N]int
var n int

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 { 
				f[i][j] = 1 
			} else { 
				f[i][j] = (f[i-1][j-1] + f[i-1][j]) % MOD
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fscan(in, &n)
	for ; n > 0; n-- {
		var a, b int
		Fscan(in, &a, &b)
		Fprintln(out, f[a][b])
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }