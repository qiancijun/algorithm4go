package week288

func largestInteger(num int) int {
	nums := make([]int, 0)
	tmp := num
	for tmp > 0 {
		nums = append(nums, tmp%10)
		tmp /= 10
	}
	
	for i := len(nums)-1; i >= 0; i-- {
		f := nums[i] % 2
		mx, idx := nums[i], -1
		for j := i-1; j >= 0; j-- {
			if nums[j]%2 == f && nums[j]>mx {
				mx = nums[j]
				idx = j
			}
		}
		if idx != -1 {
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		ans = ans*10 + nums[i]
	}
	return ans
}