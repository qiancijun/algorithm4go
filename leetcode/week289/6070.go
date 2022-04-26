package week289

import (
	"fmt"
	"strconv"
)

func helper(s string) string {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		ans += int(s[i] - '0')
	}
	return strconv.Itoa(ans)
}

func digitSum(s string, k int) string {
	for len(s) > k {
		n := len(s)
		l := 0
		tmp := ""
		for ; l+k < n; l += k {
			r := l + k
			sub := s[l:r]
			sub = helper(sub)
			// fmt.Println(sub)
			tmp += sub
		}
		tail := s[l:n]
		tail = helper(tail)
		tmp += tail
		s = tmp
	}
	
	return s
}