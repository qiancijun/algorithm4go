package main

import "fmt"

func minimumRounds(tasks []int) int {
	m := make(map[int]int)
	for _, v := range tasks {
		m[v]++
	}
	ans := 0
	for _, v := range tasks {
		t := m[v]
		if t == 1 {
			return -1
		} else if t == 2 {
			ans++
		} else {
			ans += t / 3
			if t%3 > 0 {
				ans++
			}
		}
		m[v] = 0
	}
	return ans
}

func main() {
	t := minimumRounds([]int{2, 3, 3})
	fmt.Println(t)
}