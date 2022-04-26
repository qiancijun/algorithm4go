package main

import (
	"bufio"
	. "fmt"
	"os"
)

type pair struct {
	pre string
	opt byte
}

var (
	g [2][4]rune
	begin, end string
	pre = make(map[string]pair)
	dist = make(map[string]int)
)

func setState(state string) [2][4]string {
	var ret [2][4]string
	idx := 0
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		ret[i][j] = string(state[idx])
	// 		idx++
	// 	}
	// }
	for i := 0; i < 4; i++ { ret[0][i] = string(state[idx]); idx++ }
	for i := 3; i >= 0; i-- { ret[1][i] = string(state[idx]); idx++ }
	return ret
}

func moveToA(state string) string {
	tmp := setState(state)
	for i := 0; i < 4; i++ {
		tmp[0][i], tmp[1][i] = tmp[1][i], tmp[0][i]
	}
	str := ""
	for i := 0; i <= 3; i++ { str += tmp[0][i] }
	for i := 3; i >= 0; i-- { str += tmp[1][i] }
	return str
}

func moveToB(state string) string {
	tmp := setState(state)
	a, b := tmp[0][3], tmp[1][3]
	for i := 3; i > 0; i-- {
		tmp[0][i] = tmp[0][i-1]
		tmp[1][i] = tmp[1][i-1]
	}
	tmp[0][0], tmp[1][0] = a, b
	str := ""
	for i := 0; i <= 3; i++ { str += tmp[0][i] }
	for i := 3; i >= 0; i-- { str += tmp[1][i] }
	return str
}

func moveToC(state string) string {
	tmp := setState(state)
	tmp[0][2], tmp[1][2], tmp[1][1], tmp[0][1] = tmp[0][1], tmp[0][2], tmp[1][2], tmp[1][1]
	str := ""
	for i := 0; i <= 3; i++ { str += tmp[0][i] }
	for i := 3; i >= 0; i-- { str += tmp[1][i] }
	return str
}

func bfs(begin, end string) {
	if begin == end {
		return
	}
	q := make([]string, 0)
	q = append(q, begin)
	dist[begin] = 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == end {
			break
		}
		a, b, c := moveToA(cur), moveToB(cur), moveToC(cur)
		if _, has := dist[a]; !has {
			pre[a] = pair{cur, 'A'}
			dist[a] = dist[cur] + 1
			q = append(q, a)
		}
		if _, has := dist[b]; !has {
			pre[b] = pair{cur, 'B'}
			dist[b] = dist[cur] + 1
			q = append(q, b)
		}
		if _, has := dist[c]; !has {
			pre[c] = pair{cur, 'C'}
			dist[c] = dist[cur] + 1
			q = append(q, c)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var b string
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			Fscan(in, &b)
			end += b
		}
	}
	begin = "12345678"
	bfs(begin, end)

	ans := end
	res := make([]byte, 0)
	for ans != begin {
		l := pre[ans]
		res = append(res, l.opt)
		ans = l.pre
	}
	Fprintln(out, dist[end])
	for i := len(res) - 1; i >= 0; i-- {
		Fprint(out, string(res[i]))
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }