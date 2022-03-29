package main

var lastS = -1

func countCollisions(directions string) int {
	ans := 0
	rr:= 0
	lastS = -1
	for i, v := range directions {
		switch v {
		case 'R':
			rr ++ 
		case 'L':
			if rr != 0 {
				ans += 2
				ans += (rr - 1)
				rr = 0
				lastS = i
			} else if lastS != -1 {
				ans ++ 
				lastS = i
			}
		case 'S':
			ans += rr
			lastS = i
			rr = 0
		}
	}
	return ans
}