package main

// LeetCode #3131: Find the Integer Added to Array I
// https://leetcode.com/problems/find-the-integer-added-to-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: addedInteger
	fmt.Println(FindTheIntegerAddedToArrayI([]int{2, 6, 4}, []int{9, 7, 5})) // 3
	fmt.Println(FindTheIntegerAddedToArrayI([]int{10}, []int{5}))             // -5
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: addedInteger
func FindTheIntegerAddedToArrayI(nums1 []int, nums2 []int) int {
	min1, min2 := nums1[0], nums2[0]
	for _, v := range nums1 {
		if v < min1 {
			min1 = v
		}
	}
	for _, v := range nums2 {
		if v < min2 {
			min2 = v
		}
	}
	return min2 - min1
}
