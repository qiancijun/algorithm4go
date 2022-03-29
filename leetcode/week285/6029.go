package main

func maximumBobPoints(n int, aA []int) (ans []int) {
	maxScore := -1
	for i := 0; i < 1 << 12; i ++ {
		score, arrow, bB := 0, 0, [12]int{}
		for j, v := range aA {
			if (i >> j & 1) == 1 {
				arrow += v + 1
				score += j
				bB[j] = v + 1
			}
		}
		if arrow > n {
			continue
		}
		if score > maxScore {
			maxScore = score
			bB[0] += n - arrow
			ans = bB[:]
		}
	}
	return
}