package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var alpha, str string
	Fscan(in, &alpha, &str)
	cnt := 0

	i, j := 0, 0
	for i < len(str) {
		for i < len(str) && j < len(alpha) {
			if str[i] == alpha[j] {
				i++
			}
			j++
		}
		if j == len(alpha) {
			cnt++
			j = 0
		}
	}
	cnt ++ 

	Fprintln(out, cnt)
}