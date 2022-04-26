package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N int = 110

var (
	n, k int
	a [N]int
	root *Node = new(Node)
)

type Node struct {
	idx int
	next *Node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n, &k)
	for i := 1; i <= k; i++ {
		Fscan(in, &a[i])
	}

	head := root
	for i := 1; i <= n; i++ {
		tmp := &Node{i, nil}
		head.next = tmp
		head = head.next
	}
	root = root.next
	head.next = root

	for i := 1; i <= k; i++ {
		t := a[i] % n
		if t == 0 {
			t = n
		}
		// Fprintln(out, n, t)
		for j := 1; j < t; j++ {
			root = root.next
		}
		// Fprint(out, root.idx, " ")
		Fprint(out, root.next.idx, " ")
		root.next = root.next.next
		root = root.next
		n--
		// Fprintln(out, root.idx)
	}

}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
func abs(v int) int { if v > 0 { return v }; return -v }
func gcd(a, b int) int { if b > 0 { return gcd(b, a % b) }; return a }
func lcm(a, b int) int { return a * b / gcd(a, b) }