package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N int = 1e5 + 10
)

var a [N]int
var n, m int

type node struct { l, r, val int }
type seg []node

func (t seg) pushup(u int) {
	t[u].val = (t[u<<1].val + t[u<<1|1].val)
}

func (t seg) build(u, l, r int) {
	if l == r {
		t[u] = node{l, r, a[l]}
	} else {
		t[u].l, t[u].r = l, r
		m := (l + r) >> 1
		t.build(u<<1, l, m)
		t.build(u<<1|1, m+1, r)
		t.pushup(u)
	}
}

func (t seg) query(u, l, r int) int {
	if t[u].l >= l && t[u].r <= r {
		return t[u].val
	}
	m := (t[u].l + t[u].r) >> 1
	ret := 0
	if l <= m {
		ret = t.query(u<<1, l, r)
	}
	if m < r {
		ret += t.query(u<<1|1, l, r)
	}
	return ret
}

func (t seg) update(u, x, v int) {
	if t[u].l == t[u].r {
		t[u].val += v
	} else {
		m := (t[u].l + t[u].r) >> 1
		if x <= m {
			t.update(u<<1, x, v)
		} else {
			t.update(u<<1|1, x, v)
		}
		t.pushup(u)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	t := make(seg, 4 * N)
	t.build(1, 1, n)

	for ; m > 0; m-- {
		var k, a, b int
		Fscan(in, &k, &a, &b)
		if k == 1 {
			t.update(1, a, b)
		} else {
			Fprintln(out, t.query(1, a, b))
		}
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }