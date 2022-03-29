package week286

import (
	"fmt"
)

func helper(k int) (int64, int64) {
	var begin, end int64 = 1, 9
	t := k
	for ; t > 1; t-- {
		begin *= 10
		end *= 10
		end += 9
	}
	return begin, end
}

func check(v int64, k int) bool {
	var n int64 = 0
	tmp := v
	for tmp > 0 {
		a := tmp % 10
		n *= 10
		n += a
		tmp /= 10
	}
	return n == v
}

func kthPalindrome(queries []int, intLength int) []int64 {
	b1, e1 := helper(9)
	fmt.Println(e1 - b1 + 1)
	ans := make([]int64, 0)
	for _, v := range queries {
		idx := 0
		begin, end := helper(intLength)
		// fmt.Println(begin, " ", end, " ", v)
		if end - begin + 1 < int64(v) {
			ans = append(ans, -1)
			continue
		}
		flag := false
		for j := begin; j <= end; j++ {
			if check(j, intLength) {
				idx++
				if idx == v {
					ans = append(ans, j)
					flag = true
					break
				}
			}
		}
		if !flag {
			ans = append(ans, -1)
		}
	}
	return ans
}