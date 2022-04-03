package week287

import (
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
	ss := strings.Split(current, ":")
	c1, c2 := ss[0], ss[1]
	s1, _ := strconv.Atoi(c1)
	s2, _ := strconv.Atoi(c2)

	ss = strings.Split(correct, ":")
	c1, c2 = ss[0], ss[1]
	v1, _ := strconv.Atoi(c1)
	v2, _ := strconv.Atoi(c2)


	ans := 0

	if v2 < s2 {
		v2 += 60
		v1 --
	}

	ans += (v1 - s1)

	dis := v2 - s2

	tmp := dis / 15
	ans += tmp
	dis = dis - tmp * 15

	tmp = dis / 5
	ans += tmp
	dis = dis - tmp * 5

	ans += dis
	
	return ans
}