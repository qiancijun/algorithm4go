package week290

const N int = 1010

func intersection(nums [][]int) []int {
	n := len(nums)
	cnt := make([]int, N)
	for i := range nums {
		for j := range nums[i] {
			v := nums[i][j]
			cnt[v]++ 
		}
	}
	res := make([]int, 0)
	for i, v := range cnt {
		if v == n {
			res = append(res, i)
		}
	}
	return res
}