package main

import (
	"fmt"
	"strconv"
)

func add(a, b string) string {
	n, _ := strconv.ParseInt(a, 36, 0)
	m, _ := strconv.ParseInt(b, 36, 0)
	num := n + m
	res := strconv.FormatInt(num, 36)
	return res
}

func main() {
	t := add("abbbb", "1")
	fmt.Println(t)
}