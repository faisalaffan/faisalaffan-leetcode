package main

// LeetCode #349: Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n)
func IntersectionOfTwoArrays(nums1, nums2 []int) []int {
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}
	var result []int
	for _, v := range nums2 {
		if set[v] {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

func main() {
	fmt.Println(IntersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(IntersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
