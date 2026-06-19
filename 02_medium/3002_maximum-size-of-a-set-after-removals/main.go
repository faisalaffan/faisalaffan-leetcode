package main

// LeetCode #3002: Maximum Size of a Set After Removals
// https://leetcode.com/problems/maximum-size-of-a-set-after-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumSetSize([]int{1, 2, 1, 2}, []int{1, 2, 1, 2}))
	fmt.Println(maximumSetSize([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}))
	fmt.Println(maximumSetSize([]int{1, 2, 1, 2}, []int{3, 4, 5, 6}))
}

func maximumSetSize(nums1 []int, nums2 []int) int {
	s1 := map[int]bool{}
	s2 := map[int]bool{}
	for _, x := range nums1 {
		s1[x] = true
	}
	for _, x := range nums2 {
		s2[x] = true
	}
	a, b, c := 0, 0, 0
	for x := range s1 {
		if !s2[x] {
			a++
		}
	}
	for x := range s2 {
		if !s1[x] {
			b++
		} else {
			c++
		}
	}
	n := len(nums1)
	if a > n/2 {
		a = n / 2
	}
	if b > n/2 {
		b = n / 2
	}
	res := a + b + c
	if res > n {
		res = n
	}
	return res
}
