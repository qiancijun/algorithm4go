package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nums [7]int
	for i := 0; i < 7; i++ {
		Fscan(in, &nums[i])
	}
	sort.Ints(nums[:])
	a, b, abc := nums[0], nums[1], nums[6]
	Fprintln(out, a, b, abc-a-b)
}