package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	Fscanf(in, "%d%d\n", &n, &k)
	bytes, _, _ := in.ReadLine()
	strs := strings.Split(string(bytes), " ")
	j := 0
	length := 0
	for j < n {
		word := strs[j]
		if length + len(word) <= k {
			Fprint(out, word, " ")
		} else {
			length = 0
			Fprintln(out)
			Fprint(out, word, " ")
		}
		length += len(word)
		j++
	}
}
