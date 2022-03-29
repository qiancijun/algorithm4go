package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 2010

var n int
var a [N]int

type node struct {
	l, r int
	val float32
}

type seg []node

func (t seg) pushup(u int) {
	lo, ro := t[u<<1], t[u<<1|1]
	t[u].val = (float32(lo.r-lo.l+1)*lo.val + float32(ro.r-ro.l+1)*ro.val) / float32(t[u].r - t[u].l + 1)
}

func (t seg) build(u, l, r int) {
	if l == r {
		t[u] = node{l, r, float32(a[r])}
	} else {
		t[u].l, t[u].r = l, r
		mid := (l + r) >> 1
		t.build(u<<1, l, mid)
		t.build(u<<1|1, mid+1, r)
		t.pushup(u)
	}
}

func (t seg) query(u, l, r int) float32 {
	if t[u].l >= l && t[u].r <= r {
		return t[u].val
	}
	mid := (t[u].l + t[u].r) >> 1
	var ret float32
	if l <= mid {
		ret = t.query(u<<1, l, r)
	}
	if mid < r {
		ret = t.query(u<<1|1, l, r)
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)
	sum := 0
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	t := make(seg, 4*n)
	t.build(1, 1, n)
	Fprintln(out, t.query(1, 1, 2))
	for i := 0; i < n; i += 2 {
		for j := 1; j+i <= n; j++ {
			Fprintln(out, j, " ", j+i, " ", t.query(1, j, j+i))
			sum += int(t.query(1, j, j+i))
		}
	}
	Fprintln(out, sum)
}
