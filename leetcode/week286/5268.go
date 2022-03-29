package week286

import "sort"

func findDifference(nums1 []int, nums2 []int) [][]int {
	sort.Slice(nums1, func(i, j int) bool {return nums1[i] < nums1[j]})
	sort.Slice(nums2, func(i, j int) bool {return nums2[i] < nums2[j]})
	n1 := removeDup(nums1)
	n2 := removeDup(nums2)
	s1, s2 := make(map[int]bool), make(map[int]bool)
	for _, v := range n1 {
		s1[v] = true
	}
	for _, v := range n2 {
		s2[v] = true
	}
	ans := make([][]int, 0)
	a1, a2 := []int{}, []int{}
	for _, v := range n1 {
		if _, has := s2[v]; !has {
			a1 = append(a1, v)
		}
	}
	for _, v := range n2 {
		if _, has := s1[v]; !has {
			a2 = append(a2, v)
		}
	}
	ans = append(ans, a1, a2)
	return ans
}

func removeDup(a []int) []int {
	i := 0
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i++
			a[i] = a[j]
		}
	}
	return a[:i+1]
}