package leetcode

func lengthOfLIS(nums []int) int {
    n := len(nums)

    g := make([]int, n+1)
    cnt := 0
    for i := 0; i < n; i++ {
        k := 0
        for k < cnt && g[k] < nums[i] {
            k++
        }
        if k == cnt {
            g[k] = nums[i]
            cnt++
        } else {
            g[k] = nums[i]
        }
    }
    return cnt
}