package main

import (
	"bufio"
	. "fmt"
	"os"
)

const MOD int = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	Fscan(in, &T)
	for ; T > 0; T-- {
		var n int
		Fscan(in, &n)
		if n & 1 == 1 {
			Fprintln(out, 0)
			continue
		}
		var ans int64 = 1
		for i := 1; i <= n / 2; i++ {
			ans *= int64(i*i%MOD)
			ans %= int64(MOD)
		}
		Fprintln(out, ans)
	}
}
