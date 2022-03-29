//go:build ignore
// +build ignore

package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 5e5 + 10

type node struct {
	l, r int
	val  int64
}

type seg []node

var a [N]int64

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (t seg) pushup(u int) {
	t[u].val = gcd(t[u<<1].val, t[u<<1|1].val)
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

var cnt int = 0
func (t seg) query(u, l, r int, v int64) {
	if t[u].l == t[u].r {
		cnt++
		return
	}
	if cnt > 1 {
		return
	}
	m := (t[u].l + t[u].r) >> 1
	if l <= m && t[u<<1].val % v != 0 {
		t.query(u<<1, l, r, v)
	}
	if m < r && t[u<<1|1].val % v != 0 {
		t.query(u<<1|1, l, r, v)
	}
	return
}

func (t seg) update(u, x int, v int64) {
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
	var n int
	Fscan(in, &n)
	t := make(seg, 4*n+1)

	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	t.build(1, 1, n)

	var q int
	Fscan(in, &q)

	for i := 1; i <= q; i++ {
		var op int
		Fscan(in, &op)
		if op == 1 {
			var l, r int
			var x int64
			Fscan(in, &l, &r, &x)
			cnt = 0
			t.query(1, l, r, x)
			if cnt <= 1 {
				Fprintln(out, "YES")
			} else {
				Fprintln(out, "NO")
			}
		} else {
			var x int
			var v int64
			Fscan(in, &x, &v)
			t.update(1, x, v)
		}
	}
}
