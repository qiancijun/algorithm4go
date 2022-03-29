package week286

func minDeletion(nums []int) int {
	ans := 0
	for i := 0; ; {
		if i >= len(nums)-1 {
			break
		}
		if i%2 == 0 && nums[i] == nums[i+1] {
			ans++
			nums = nums[i+1:]
			i = 0
		} else {
			i++
		}
	}
	if len(nums)%2 == 1 {
		ans++
	}
	return ans
}