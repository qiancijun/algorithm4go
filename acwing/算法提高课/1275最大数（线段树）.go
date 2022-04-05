package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 2e5 + 10

var m, p int

type node struct { l, r, val int}
type seg []node

func (t seg) pushup(u int) {
	lo, ro := t[u<<1], t[u<<1|1]
	t[u].val = max(lo.val, ro.val)
}

func (t seg) build(u, l, r int) {
	t[u].l, t[u].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(u<<1, l, m)
	t.build(u<<1|1, m+1, r)
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
		ret = max(ret, t.query(u<<1|1, l, r))
	}
	return ret
}

func (t seg) update(u, x, v int) {
	if t[u].l == t[u].r {
		t[u].val = v
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

	Fscan(in, &m, &p)

	n, last := 0, 0
	t := make(seg, 4 * N)
	t.build(1, 1, m)

	for i := 0; i < m; i++ {
		var op string
		var v int
		Fscan(in, &op, &v)
		if op == "Q" {
			last = t.query(1, n-v+1, n)
			Fprintln(out, last)
		} else {
			t.update(1, n+1, (last+v)%p)
			n++
		}
	}
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }