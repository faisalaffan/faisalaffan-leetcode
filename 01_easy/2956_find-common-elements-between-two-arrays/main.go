package main

// LeetCode #2956: Find Common Elements Between Two Arrays
// https://leetcode.com/problems/find-common-elements-between-two-arrays/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findIntersectionValues
	fmt.Println(FindCommonElementsBetweenTwoArrays([]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6})) // [3, 4]
	fmt.Println(FindCommonElementsBetweenTwoArrays([]int{3, 4, 2, 3}, []int{1, 5})) // [0, 0]
}

// Time: O(n + m) | Space: O(n + m)
// LeetCode submission name: findIntersectionValues
func FindCommonElementsBetweenTwoArrays(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	count1 := 0
	for _, v := range nums1 {
		if set2[v] {
			count1++
		}
	}
	count2 := 0
	for _, v := range nums2 {
		if set1[v] {
			count2++
		}
	}
	return []int{count1, count2}
}
