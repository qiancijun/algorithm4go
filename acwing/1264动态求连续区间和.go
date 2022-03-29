package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1e5 + 10
const INF = 2<<31 - 1

var n, m int
var a [N]int

type node struct {
	l, r, val int
}

type seg []node

func (seg) op(a, b int) int {
	return a + b
}

func (t seg) pushup(u int) {
	lo, ro := t[u<<1], t[u<<1|1]
	t[u].val = t.op(lo.val, ro.val)
}

func (t seg) build(u, l, r int) {
	if l == r {
		t[u] = node{l, r, a[r]}
	} else {
		t[u].l, t[u].r = l, r
		mid := (l + r) >> 1
		t.build(u<<1, l, mid)
		t.build(u<<1|1, mid+1, r)
		t.pushup(u)
	}
}

func (t seg) query(u, l, r int) int {
	if t[u].l >= l && t[u].r <= r {
		return t[u].val
	}
	mid := (t[u].l + t[u].r) >> 1
	ret := 0
	if l <= mid {
		ret = t.query(u<<1, l, r)
	}
	if mid < r {
		ret = t.op(ret, t.query(u<<1|1, l, r))
	}
	return ret
}

func (t seg) update(u, x, v int) {
	if t[u].l == t[u].r {
		t[u].val = t.op(t[u].val, v)
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

// EXTRA: 查询整个区间小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 需要先判断 t[1].min < v
func (t seg) queryFirstLessPos(u, v int) int {
	if t[u].l == t[u].r {
		return t[u].l
	}
	if t[u<<1].val < v {
		return t.queryFirstLessPos(u<<1, v)
	}
	return t.queryFirstLessPos(u<<1|1, v)
}


func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	t := make(seg, 4*n)

	t.build(1, 1, n)

	for i := 0; i < m; i++ {
		var k, a, b int
		Fscan(in, &k, &a, &b)
		if k == 0 {
			Fprintln(out, t.query(1, a, b))
		} else {
			t.update(1, a, b)
		}
	}
}
