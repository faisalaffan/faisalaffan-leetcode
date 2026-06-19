package main

// LeetCode #2918: Minimum Equal Sum of Two Arrays After Replacing Zeros
// https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/
// Difficulty: Medium
// Time: O(n+m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minSum([]int{3, 2, 0, 1, 0}, []int{6, 5, 0}))
	fmt.Println(minSum([]int{2, 0, 2, 0}, []int{1, 4}))
	fmt.Println(minSum([]int{1, 2}, []int{3, 4}))
}

func minSum(nums1 []int, nums2 []int) int64 {
	s1, s2 := int64(0), int64(0)
	z1, z2 := 0, 0
	for _, x := range nums1 {
		if x == 0 {
			z1++
		} else {
			s1 += int64(x)
		}
	}
	for _, x := range nums2 {
		if x == 0 {
			z2++
		} else {
			s2 += int64(x)
		}
	}
	min1, min2 := s1+int64(z1), s2+int64(z2)
	if min1 < min2 && z1 == 0 {
		return -1
	}
	if min2 < min1 && z2 == 0 {
		return -1
	}
	if min1 > min2 {
		return min1
	}
	return min2
}
