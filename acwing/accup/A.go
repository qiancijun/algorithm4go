package main

import (
	"bufio"
	"os"
	."fmt"
)

var n int
var a = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &n)

	if n == 0 {
		Fprintln(out, 1)
		return
	}
	ans := 0
	for n > 0 {
		tmp := n % 16
		n /= 16
		// res = a[tmp] + res
		if tmp == 0 || tmp == 4 || tmp == 6 || tmp == 9 || tmp == 10 || tmp == 13 {
			ans ++
		} else if tmp == 8 || tmp == 11 {
			ans += 2
		}
	}

	Fprintln(out,ans)
}