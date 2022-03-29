package main

import (
	"bufio"
	. "fmt"
	"os"
)

// acbcca
var cnt [26]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var str string
	tmp := "acbcca"
	Fscan(in, &str)
	for _, v := range str {
		cnt[v-'a']++
	}

	ans := 0
	n := len(str)
	idx := 0
	for i := 0; i < n; i++ {
		if idx == len(tmp) {
			ans++
			idx = 0
		}
		if idx == 0 && ans > 0 {
			idx++
			continue 
		}
		if cnt[tmp[idx]-'a'] != 0 {
			cnt[tmp[idx]-'a']--
			idx++
		}
	}
	Fprintln(out, ans)
}
