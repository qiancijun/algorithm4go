package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

var T int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Fscan(in, &T)
	for t := 0; t < T; t++ {
		var x, y int
		Fscan(in, &x, &y)
		if x == 0 && y == 0 {
			Fprintln(out, 0)
			continue
		}
		dis := x*x + y*y
		tmp := int(math.Sqrt(float64(dis)))
		if tmp * tmp == dis {
			Fprintln(out, 1)
		} else {
			Fprintln(out, 2)
		}
	}
}