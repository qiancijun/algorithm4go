package main

import (
	"fmt"
)

func solve(arr []int, n, m int) int {
	// 请添加具体实现
	s := make(map[int]bool)
	for _, v := range arr {
		s[v] = true
	}
	ans := 0
	for _, v := range arr {
		if a, _ := s[v]; !a { continue }
		f := m - v
		if f == v { continue } // 不重复
		if a, has := s[f]; has && a {
			ans++
			s[v] = false
			s[f] = false
		}
	}
	return ans
}

func main() {
	fmt.Println(solve([]int{1, 2, 3}, 3, 9))
}
