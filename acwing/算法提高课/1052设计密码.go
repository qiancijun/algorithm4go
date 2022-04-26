package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 55
	MOD = 1e9 + 7
)

var (
	n, m int
	next [N]int
	f [N][N]int
	str string
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &str)
	m = len(str)
	str = " " + str

	for i, j := 2, 0; i <= m; i++ {
		for j > 0 && str[i] != str[j+1] { j = next[j] }
		if str[i] == str[j+1] { j++ }
		next[i] = j
	}

	f[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 'a'; k <= 'z'; k++ {
				u := j
				for u > 0 && byte(k) != str[u+1] { u = next[u] }
				if byte(k) == str[u+1] { u++ }
				if u < m {
					f[i+1][u] = (f[i+1][u] + f[i][j]) % MOD
				}
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + f[n][i]) % MOD
	}
	Fprintln(out, res)
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }