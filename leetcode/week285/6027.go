package main

func countHillValley(nums []int) int {
	nums = removeDup(nums)
	n := len(nums)
	s := make(map[int]bool)
	for i := 1; i < n - 1; i ++ {
		if nums[i - 1] < nums[i] && nums[i] > nums[i + 1] {
			s[i] = true
		}
		if nums[i - 1] > nums[i] && nums[i + 1] > nums[i] {
			s[i] = true
		}
	}
	return len(s)
}

func removeDup(a []int)[]int {
    i := 0
    for j := 1; j < len(a); j ++ {
        if a[i] != a[j] {
            i ++
            a[i] = a[j]
        }
    }
    return a[:i+1]
}