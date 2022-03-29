package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

var s = map[string]int{
	"Ox":      0,
	"Tiger":   1,
	"Rabbit":  2,
	"Dragon":  3,
	"Snake":   4,
	"Horse":   5,
	"Goat":    6,
	"Monkey":  7,
	"Rooster": 8,
	"Dog":     9,
	"Pig":     10,
	"Rat":     11,
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscanf(in, "%d\n", &n)

	m := make(map[string]int)
	m["Bessie"] = 0
	for ; n > 0; n-- {
		str, _, _ := in.ReadLine()
		lines := strings.Split(string(str), " ")
		// Println(len(lines))
		target, opt, year, from := lines[0], lines[3], lines[4], lines[7]
		if opt == "previous" {
			ans := ((m[from]-s[year])%12 + 12) % 12
			if ans == 0 {
				ans = 12
			}
			m[target] = m[from] - ans
		} else {
			ans := ((s[year]-m[from])%12 + 12) % 12
			if ans == 0 {
				ans = 12
			}
			m[target] = m[from] + ans
		}
	}
	Fprintln(out, abs(m["Elsie"]))
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
