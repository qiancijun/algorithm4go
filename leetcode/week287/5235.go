package week287

const N int = 1e5 + 10

func findWinners(matches [][]int) [][]int {
	var din [N]int
	for i := range din {
		din[i] = -1
	}
	for _, match := range matches {
		win, lose := match[0], match[1]
		if din[win] == -1 {
			din[win] = 0
		}
		if din[lose] == - 1 {
			din[lose] = 0
		}
		din[lose]++
	}
	ans1, ans2 := make([]int, 0), make([]int, 0)

	for i, v := range din {
		if v == 0 {
			ans1 = append(ans1, i)
		} else if v == 1 {
			ans2 = append(ans2, i)
		}
	}
	
	ans := make([][]int, 0)
	ans = append(ans, ans1, ans2)
	return ans
}
